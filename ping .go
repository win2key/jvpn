package main

import (
	"net/http"
	"os/exec"
)

func pingTarget(w http.ResponseWriter, r *http.Request, target string) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	cmd := exec.Command("ping", "-c", "1", "-W", "3", target)
	err := cmd.Run()

	if err != nil {
		w.Write([]byte("0")) // Target IP is not reachable
	} else {
		w.Write([]byte("1")) // Target IP is reachable
	}
}
