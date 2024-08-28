package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ip struct {
	Addr string `json:"addr"`
}

func main() {
	port := ":443"

	http.HandleFunc("/ip", ipHandler)
	log.Println("server running on port", "https://aws.mugund10.top"+port+"/ip")
	http.ListenAndServeTLS(port, "/home/ubuntu/LetsEncryptAcmeClient/cmd/acmeclient/fullchain.pem", "/home/ubuntu/LetsEncryptAcmeClient/cmd/acmeclient/pk[aws.mugund10.top].pem", nil)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ipa := extractsIP(r.RemoteAddr)
	ipaddr := ip{Addr: ipa}
	//converts to json format
	byt, err := json.Marshal(ipaddr)
	if err != nil {
		log.Println(err)
		return
	}
	//makes headers
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s", byt)
}

// removes the ports ex: 1.1.1.1:30 -> 1.1.1.1
func extractsIP(ip string) string {
	for k, v := range ip {
		if ":" == string(v) {
			ip = ip[:k]
		}
	}
	return ip
}
