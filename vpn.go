package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/shirou/gopsutil/v3/process"
)

func startOpenConnect(w http.ResponseWriter, r *http.Request, vpnLogin, vpnPassword, vpnServer string) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	isRunning, err := isProcessRunning("openconnect")
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isRunning {
		w.Write([]byte("0"))
	} else {
		err := startVPN(vpnLogin, vpnPassword, vpnServer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("VPN started")
		w.Write([]byte("1"))
	}
}

func isProcessRunning(processName string) (bool, error) {
	processes, err := process.Processes()
	if err != nil {
		return false, err
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}

		if name == processName {
			status, _ := p.Status()
			for _, s := range status {
				if s != "Z" { // "Z" indicates a Zombie process
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func startVPN(vpnLogin, vpnPassword, vpnServer string) error {
	cmd := exec.Command("./openconnect-expect.sh")
	return cmd.Start()
}
