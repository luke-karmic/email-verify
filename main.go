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
		VerifierEmail:         "luke.taaffe@hotmail.com",
		ConnectionTimeout:     3, // Increased timeout
		ResponseTimeout:       3, // Increased timeout
		ConnectionAttempts:    2,
		ValidationTypeDefault: "smtp",
	})
	if err != nil {
		log.Fatal("Failed to create configuration:", err)
	}

	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "leads.csv")

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
		Email string
		res   bool
		Err   bool
	}

	validatedEmailAddresses := []ValidatedEmail{}

	totalEmailAddr := func() int {
		if len(records)-1 >= 0 {
			return len(records) - 1
		} else {
			return 0
		}
	}()

	fmt.Printf("Beginning email validation. You are analysing a total of %v email address records\n", totalEmailAddr)

	// Read all the records from the CSV file
	var countSuccessfulEmails int32 = 0
	var countUnsuccessfulEmails int32 = 0

	for i, record := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		// Skip empty email addresses
		if record[0] == "" {
			continue
		}

		res, err := truemail.Validate(record[0], configuration)

		validEmail := false
		if err == nil && res != nil && res.Success {
			fmt.Print(res.Errors, res.SmtpDebug)
			countSuccessfulEmails++
			validEmail = true
		} else {
			fmt.Printf("ERR VALIDATING %s: %v\n", record[0], err)
			if res != nil {
				fmt.Print(res.Errors, res.SmtpDebug)
			}
			countUnsuccessfulEmails++
		}

		validatedEmailAddresses = append(validatedEmailAddresses, ValidatedEmail{
			Email: record[0],
			res:   validEmail,
			Err:   err != nil,
		})
	}

	totalSuccessRate := float64(countSuccessfulEmails) / float64(totalEmailAddr) * 100
	fmt.Printf("You have successfully checked %v email addresses, %v are valid, %v are invalid, total success rate %v%%", totalEmailAddr, countSuccessfulEmails, countUnsuccessfulEmails, totalSuccessRate)

	// Custom format using fmt.Printf
	fmt.Printf("%v\n", validatedEmailAddresses)
}
