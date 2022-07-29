package util

import (
	"bufio"
	"os"
	"strings"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/tools/wasp-cli/log"
)

func SDAdjustmentPrompt(output iotago.Output) {
	minStorageDeposit := parameters.L1.Protocol.RentStructure.MinRent(output)
	if output.Deposit() < minStorageDeposit {
		// don't prompt if running in a script // https://stackoverflow.com/a/43947435/6749639
		fi, _ := os.Stdin.Stat()
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			log.Fatalf("transaction not sent.")
		}

		xxx

		// query the user if they want to send the Tx with adjusted storage deposit
		log.Printf(`
		The amount of base tokens to be sent are not enough to cover the Storage Deposit for the new output.\n
		need:%d, have:%d  \n
		Do you wish to continue and issue this transaction sending %d base tokens? [Y/n]\n 
		(you can automatically accept this prompt by using the following flags: --adjust-storage-deposit or -a)
		`, minStorageDeposit, output.Deposit(), minStorageDeposit)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if strings.ToLower(text) != "y" {
			log.Fatalf("transaction not sent.")
		}
	}
}
