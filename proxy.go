package main

import (
	"net"
	"net/http"

	"github.com/win2key/jproxy"
)

func startProxy(w http.ResponseWriter, r *http.Request, jProxy string) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	address := jProxy
	_, err := isAddressAvailable(address)
	if err != nil {
		w.Write([]byte("0"))
	} else {
		go jproxy.RunProxy(address)
		w.Write([]byte("1"))
	}
}

func isAddressAvailable(address string) (bool, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false, err
	}
	defer listener.Close()
	return true, nil
}
