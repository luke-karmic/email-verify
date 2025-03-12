package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/truemail-rb/truemail-go"
)

func main() {
	configuration, err := truemail.NewConfiguration(truemail.ConfigurationAttr{
		VerifierEmail:         "luke.taaffe@hotmail.com", // some random fake email
		ConnectionTimeout:     3,                         // Increased timeout
		ResponseTimeout:       3,                         // Increased timeout
		ConnectionAttempts:    2,
		ValidationTypeDefault: "smtp",
	})
	if err != nil {
		log.Fatal("Failed to create configuration:", err)
	}

	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "leads/leads.csv")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f := csv.NewReader(file)
	records, err := f.ReadAll()
	if err != nil {
		fmt.Println("ERROR READING", err)
	}

	type ValidatedEmail struct {
		Email  string
		Err    bool
		Record []string
	}

	successEmails := []ValidatedEmail{}
	failEmails := []ValidatedEmail{}

	totalEmailAddr := func() int {
		if len(records)-1 >= 0 {
			return len(records) - 1
		} else {
			return 0
		}
	}()

	fmt.Printf("Beginning email validation. You are analysing a total of %v email address records\n", totalEmailAddr)

	var emailFieldKey int = -1
	var headerRow []string
	for i, record := range records {
		if i == 0 {
			headerRow = record
			for j := range record {
				if record[j] == "email" {
					emailFieldKey = j
				}
			}
			continue
		}

		// Check field key is set
		if emailFieldKey == -1 {
			log.Println("Email field not found in record, exiting...")
			os.Exit(1)
		}

		res, err := truemail.Validate(record[emailFieldKey], configuration)
		if err == nil && res != nil && res.Success {
			successEmails = append(successEmails, ValidatedEmail{
				Email:  record[emailFieldKey],
				Err:    err != nil,
				Record: record,
			})
		} else {
			failEmails = append(failEmails, ValidatedEmail{
				Email:  record[emailFieldKey],
				Err:    err != nil,
				Record: record,
			})
		}
		fmt.Printf("Finished Validating email: %s: %v\n", record[emailFieldKey], func() string {
			if res.Success {
				return "valid"
			} else {
				return "invalid"
			}
		}())
	}

	totalSuccessRate := float64(len(successEmails)) / float64(totalEmailAddr) * 100
	fmt.Printf("You have successfully checked %v email addresses, %v are valid, %v are invalid, total success rate %.2f%%\n", totalEmailAddr, len(successEmails), len(failEmails), totalSuccessRate)

	fmt.Printf("Successfully Validate Email addresses: \n%v\n", successEmails)
	fmt.Printf("Un-successfully Validated Email addresses: \n%v\n", failEmails)

	// Write successEmails to leads-successful.csv
	successFile, err := os.Create(filepath.Join(wd, "leads/leads-successful.csv"))
	if err != nil {
		log.Fatal("Failed to create leads-successful.csv:", err)
	}
	defer successFile.Close()

	successWriter := csv.NewWriter(successFile)
	defer successWriter.Flush()
	successWriter.Write(headerRow)

	for _, email := range successEmails {
		err := successWriter.Write(email.Record)
		if err != nil {
			log.Fatal("Failed to write to leads-successful.csv:", err)
		}
	}
	fmt.Println("Successfully wrote leads-successful.csv")

	// Write failEmails to leads-failure.csv
	failFile, err := os.Create(filepath.Join(wd, "leads/leads-failure.csv"))
	if err != nil {
		log.Fatal("Failed to create leads-failure.csv:", err)
	}
	defer failFile.Close()

	failWriter := csv.NewWriter(failFile)
	defer failWriter.Flush()
	failWriter.Write(headerRow)

	for _, email := range failEmails {
		err := failWriter.Write(email.Record)
		if err != nil {
			log.Fatal("Failed to write to leads-failure.csv:", err)
		}
	}
	fmt.Println("Successfully wrote leads-failure.csv")
}
