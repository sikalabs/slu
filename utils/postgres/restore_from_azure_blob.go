package postgres

import "fmt"

func RestoreFromAzureBlob(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	containerName string,
	accountName string,
	accountKey string,
	sourcePrefix string,
	sourceSuffix string,
) {
	fmt.Printf(
		"host=%s\nport=%d\nuser=%s\npassword=%s\ndbName=%s\ncontainerName=%s\naccountName=%s\naccountKey=%s\nsourcePrefix=%s\nsourceSuffix=%s\n",
		host,
		port,
		user,
		password,
		dbName,
		containerName,
		accountName,
		accountKey,
		sourcePrefix,
		sourceSuffix,
	)
}
