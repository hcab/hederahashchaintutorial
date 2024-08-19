package main

import (
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func main() {

	//Loads the .env file and throws an error if it cannot load the variables from that file correctly
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables from .env file. Error:\n%v\n", err))
	}

	//Grab your testnet account ID and private key from the .env file
	myAccountId, err := hedera.AccountIDFromString(os.Getenv("MY_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	myPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("MY_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	//Print your testnet account ID and private key to the console to make sure there was no error
	fmt.Printf("The account ID is = %v\n", myAccountId)
	fmt.Printf("The private key is = %v\n", myPrivateKey)

	//that's for the test net
	//client := hedera.ClientForTestnet()
	//client for the local net
	node := make(map[string]hedera.AccountID, 1)
	node["127.0.0.1:50211"] = hedera.AccountID{Account: 3}

	mirrorNode := []string{"127.0.0.1:5600"}

	client := hedera.ClientForNetwork(node)
	client.SetMirrorNetwork(mirrorNode)

	accountId, err := hedera.AccountIDFromString("0.0.2")
	privateKey, err := hedera.PrivateKeyFromString("302e020100300506032b65700422042091132178e72057a1d7528025956fe39b0b847f200ab59b2fdd367017f3087137")
	client.SetOperator(accountId, privateKey)

	//Submit a transaction to your local node
	newAccount, err := hedera.NewAccountCreateTransaction().
		SetKey(privateKey).
		SetInitialBalance(hedera.HbarFromTinybar(1000)).
		Execute(client)

	if err != nil {
		println(err.Error(), ": error getting balance")
		return
	}

	//Get receipt
	receipt, err := newAccount.GetReceipt(client)

	//Get the account ID
	newAccountId := receipt.AccountID
	fmt.Print(newAccountId)

	// the previous tutorial for the test network
	/*
		client.SetOperator(myAccountId, myPrivateKey)
		// Set default max transaction fee & max query payment
		client.SetDefaultMaxTransactionFee(hedera.HbarFrom(100, hedera.HbarUnits.Hbar))
		client.SetDefaultMaxQueryPayment(hedera.HbarFrom(50, hedera.HbarUnits.Hbar))

		fmt.Println("Client setup complete.")

		newAccountPrivateKey, err := hedera.PrivateKeyGenerateEd25519()
		if err != nil {
			panic(err)
		}
		newAccountPublicKey := newAccountPrivateKey.PublicKey()

		newAccount, err := hedera.NewAccountCreateTransaction().SetKey(newAccountPublicKey).SetInitialBalance(hedera.HbarFrom(1000, hedera.HbarUnits.Tinybar)).Execute(client)

		receipt, err := newAccount.GetReceipt(client)
		if err != nil {
			panic(err)
		}

		newAccountId := *receipt.AccountID
		fmt.Printf("The new account ID is %v\n", newAccountId)

		query := hedera.NewAccountBalanceQuery().SetAccountID(newAccountId)

		accountBalance, err := query.Execute(client)
		if err != nil {
			panic(err)
		}

		fmt.Println("The account balance for the new account is ", accountBalance.Hbars.AsTinybar())

		transaction := hedera.NewTransferTransaction().AddHbarTransfer(myAccountId, hedera.HbarFrom(-1000, hedera.HbarUnits.Tinybar)).AddHbarTransfer(newAccountId, hedera.HbarFrom(1000, hedera.HbarUnits.Tinybar))
		txResponse, err := transaction.Execute(client)

		if err != nil {
			panic(err)
		}

		//Request the recepit of the transaction
		transferReceipt, err := txResponse.GetReceipt(client)

		if err != nil {
			panic(err)
		}

		transactionStatus := transferReceipt.Status
		fmt.Printf("The transaction consensus status is %v\n", transactionStatus)
	*/
}
