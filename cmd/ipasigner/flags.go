package main

import "github.com/urfave/cli/v2"

var (
	pkcs12FileFlag = &cli.StringFlag{
		Name:  "pkcs12",
		Usage: "Path to PKCS#12 signing key file",
	}
	pkcs12PasswdFlag = &cli.StringFlag{
		Name:  "password",
		Usage: "Password for the provided PKCS#12 signing key",
	}
)
