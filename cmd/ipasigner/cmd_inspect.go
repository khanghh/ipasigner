package main

import (
	"fmt"
	"os"

	"github.com/khanghh/ipa-signer/provisioning"
	"github.com/urfave/cli/v2"
)

var (
	inspectCommand = &cli.Command{
		Action:    inpsectProvision,
		Name:      "inspect",
		Usage:     "Inspect .mobileprovision file",
		ArgsUsage: "<file>",
		Flags:     []cli.Flag{},
		Description: `The "inspect" command allows you to inspect a .mobileprovision file. 
By running this command and specifying the path to a .mobileprovision file, you can 
view its contents and gather information such as the included certificates, entitlements, 
and provisioning profile details.`,
	}
)

// initGenesis will initialise the given JSON format genesis file and writes it as
// the zero'd block (i.e. genesis) or will fail hard if it can't succeed.
func inpsectProvision(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return fmt.Errorf("invalid argument must provide path to .mobileprovision file")
	}
	filePath := ctx.Args().First()
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read .mobileprovision file: %v", err)
	}

	provision, err := provisioning.ParseMobileProvision(data)
	if err != nil {
		return fmt.Errorf("could not parse .mobileprovision file: %v", err)
	}

	devCerts, err := provision.GetDeveloperCertificates()
	if err != nil {
		return err
	}
	fmt.Println("==== Mobile Provisioning Information ====")
	fmt.Printf("UUID:\t\t%s\n", provision.UUID)
	fmt.Printf("Name:\t\t%s\n", provision.Name)
	fmt.Printf("AppIdName:\t\t%s\n", provision.AppIDName)
	fmt.Printf("Version:\t\t%d\n", provision.Version)
	fmt.Printf("TeamName:\t\t%s\n", provision.TeamName)
	fmt.Printf("TeamIdentifier:\t\t%s\n", provision.TeamIdentifier)
	fmt.Printf("TimeToLive:\t\t%v\n", provision.TimeToLive)
	fmt.Printf("CreationDate:\t\t%s\n", provision.CreationDate)
	fmt.Printf("ExpirationDate:\t\t%v\n", provision.ExpirationDate)
	fmt.Printf("ApplicationIdentifierPrefix:\t\t%s\n", provision.ApplicationIdentifierPrefix)
	fmt.Printf("DeveloperCertificates:\t\t%v\n", devCerts)
	fmt.Println("Entitlements:")
	fmt.Printf("  KeychainAccessGroups:\t\t%v\n", provision.Entitlements.KeychainAccessGroups)
	fmt.Printf("  GetTaskAllow:\t\t%v\n", provision.Entitlements.GetTaskAllow)
	fmt.Printf("  ApplicationIDentifier:\t\t%v\n", provision.Entitlements.ApplicationIDentifier)
	fmt.Printf("  TeamIdentifier:\t\t%v\n", provision.Entitlements.TeamIdentifier)
	return nil
}
