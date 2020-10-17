package main

import (
	"flag"
	"fmt"

	msalgo "github.com/AzureAD/microsoft-authentication-library-for-go/src/msal"
)

const port = "3000"

var config = createConfig("config.json")
var publicClientApp *msalgo.PublicClientApplication
var err error
var authCodeParams *msalgo.AcquireTokenAuthCodeParameters
var cacheAccessor = &SampleCacheAccessor{"serialized_cache.json"}

func main() {
	var flowNumber string
	flag.StringVar(&flowNumber, "flow", "1", "Specify flow. Default is 1 (DeviceCode).")

	flag.Usage = func() {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("./go-project -flow flowNumber\n")
		fmt.Printf("1 : Public Client - Devide Code\n")
		fmt.Printf("2 : Public Client - Authorization Code\n")
		fmt.Printf("3 : Public Client - Username/Password\n")
		fmt.Printf("4 : Confidential Client - Auth Code\n")
		fmt.Printf("5 : Confidential Client - Client Secret\n")
		fmt.Printf("6 : Confidential Client - Certificate\n")
		//flag.PrintDefaults()
	}
	flag.Parse()

	if flowNumber == "1" {
		acquireTokenDeviceCode()
	} else if flowNumber == "2" {
		acquireByAuthorizationCodePublic()
	} else if flowNumber == "3" {
		acquireByUsernamePasswordPublic()
	} else if flowNumber == "4" {
		acquireByAuthorizationCodeConfidential()
	} else if flowNumber == "5" {
		acquireTokenClientSecret()
	} else if flowNumber == "6" {
		acquireTokenClientCertificate()
	}
}
