package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	vpnLogin, vpnPassword, vpnServer, jProxy, target := checkEnvironmentVariables()

	http.HandleFunc("/vpn", func(w http.ResponseWriter, r *http.Request) {
		startOpenConnect(w, r, vpnLogin, vpnPassword, vpnServer)
	})
	http.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		startProxy(w, r, jProxy)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingTarget(w, r, target)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkEnvironmentVariables() (string, string, string, string, string) {
	requiredEnvVars := []string{"VPN_LOGIN", "VPN_PASSWORD", "VPN_SERVER", "JPROXY", "TARGET"}
	allEnvVarsFound := true

	vpnLogin, _ := os.LookupEnv("VPN_LOGIN")       // example: VPN_LOGIN=vpnuser
	vpnPassword, _ := os.LookupEnv("VPN_PASSWORD") // example: VPN_PASSWORD=vpnpassword
	vpnServer, _ := os.LookupEnv("VPN_SERVER")     // example: VPN_SERVER=vpnserver.com
	jProxy, _ := os.LookupEnv("JPROXY")            // example: JPROXY=:8081
	target, _ := os.LookupEnv("TARGET")            // example: TARGET=192.168.0.1

	for _, envVar := range requiredEnvVars {
		if _, exists := os.LookupEnv(envVar); !exists {
			fmt.Printf("%s environment variable not found\n", envVar)
			allEnvVarsFound = false
		}
	}

	if !allEnvVarsFound {
		os.Exit(1)
	}

	return vpnLogin, vpnPassword, vpnServer, jProxy, target
}
