package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/crypto/ssh"
)

func main() {
	var hostKey ssh.PublicKey

	pemBytes, err := ioutil.ReadFile("E:/Keys-Codes/DevOps-SRE.pem")
	if err != nil {
		return nil, fmt.Errorf("Reading private key file failed %v", err)
	}

	// create signer
	signer, err := signerFromPem(pemBytes, []byte(privateKeyPassword))
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// Dial your ssh server.
	conn, err := ssh.Dial("tcp", "3.237.173.153:22", config)
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer conn.Close()

	// Request the remote side to open port 8080 on all interfaces.
	l, err := conn.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("unable to register tcp forward: ", err)
	}
	defer l.Close()

	// Serve HTTP with your SSH server acting as a reverse proxy.
	http.Serve(l, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Hello world!\n")
	}))
}