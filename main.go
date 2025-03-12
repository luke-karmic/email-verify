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
	configuration, _ := truemail.NewConfiguration(truemail.ConfigurationAttr{
		VerifierEmail:         "dummyEmail",
		VerifierDomain:        "hotmail.com",
		EmailPattern:          `\A.+@(.+)\z`,
		SmtpErrorBodyPattern:  `.*(user|account).*`,
		ConnectionTimeout:     5, // Increased timeout
		ResponseTimeout:       5, // Increased timeout
		ConnectionAttempts:    2,
		ValidationTypeDefault: "smtp",
		WhitelistedDomains:    []string{"somedomain1.com", "somedomain2.com"},
		BlacklistedDomains:    []string{"somedomain3.com", "somedomain4.com"},
		ValidationTypeByDomain: map[string]string{
			"somedomain.com":  "regex",
			"otherdomain.com": "mx",
		},
	})
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
		Email   string
		IsValid bool
		Err     bool
	}

	validatedEmailAddresses := []ValidatedEmail{}

	totalEmailAddr := len(records) - 1
	fmt.Printf("Beginning email validation. You are analysing a total of %v email address records\n", totalEmailAddr)

	// Read all the records from the CSV file
	var countSuccessfulEmails int32 = 0
	var countUnsuccessfulEmails int32 = 0

	for _, record := range records {
		isValid, err := truemail.Validate(record[0], configuration)

		validEmail := false
		if err == nil && isValid.Success {
			countSuccessfulEmails++
			validEmail = true
		} else {
			fmt.Printf("ERR VALIDATING %s: %v\n", record[0], err)
			countUnsuccessfulEmails++
		}

		validatedEmailAddresses = append(validatedEmailAddresses, ValidatedEmail{
			Email:   record[0],
			IsValid: validEmail,
			Err:     err != nil,
		})
	}

	totalSuccessRate := float64(countSuccessfulEmails) / float64(totalEmailAddr) * 100
	fmt.Printf("You have successfully checked %v email addresses, %v are valid, %v are invalid, total success rate %v%%", totalEmailAddr, countSuccessfulEmails, countUnsuccessfulEmails, totalSuccessRate)

	// Custom format using fmt.Printf
	fmt.Printf("%v\n", validatedEmailAddresses)
}
