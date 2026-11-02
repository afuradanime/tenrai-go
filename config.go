package jikan

import "os"

// Endpoint points to Tenrai API
var Endpoint = endpointFromEnvironment()

func endpointFromEnvironment() string {
	if endpoint := os.Getenv("TENRAI_ENDPOINT"); endpoint != "" {
		return endpoint
	}
	return "https://api.tenrai.org/v1"
}

// ClientTimeout specifies the default timeout value in seconds for the http client
const ClientTimeout = 60
