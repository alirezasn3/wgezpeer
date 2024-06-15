package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"os"
	"strings"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please Provide CIDR and Endpoint Arguemnts\nEx: ./wgezpeer 10.0.0.1/24 12.23.34.45")
		return
	}

	_, _, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		fmt.Println("Invalid CIDR: " + os.Args[1])
		return
	}

	serverIP := &IPAddress{}
	err = serverIP.Parse(strings.Split(os.Args[1], "/")[0])
	if err != nil {
		fmt.Println("Failed to parse CIDR: " + os.Args[1])
		return
	}

	clientIP := &IPAddress{}
	err = clientIP.Parse(strings.Split(os.Args[1], "/")[0])
	if err != nil {
		panic(err)
	}
	clientIP.Increment()

	wgc, err := wgctrl.New()
	if err != nil {
		panic(err)
	}
	defer wgc.Close()

	listenPort := rand.IntN(65535-1000) + 1000

	serverPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		panic(err)
	}
	serverPublicKey := serverPrivateKey.PublicKey()
	clientPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		panic(err)
	}
	clientPublicKey := clientPrivateKey.PublicKey()

	serverConfig := fmt.Sprintf("[Interface]\nPrivateKey = %s\nAddress = %s\n[Peer]\nPublicKey = %s\nAllowedIPs = %s/32\nEndpoint = %s:%d", serverPrivateKey.String(), os.Args[1], clientPublicKey.String(), clientIP.ToString(), os.Args[2], listenPort)

	clientConfig := fmt.Sprintf("[Interface]\nPrivateKey = %s\nAddress = %s/32\nListenPort = %d\nDNS = 1.1.1.1,8.8.8.8\n[Peer]\nPublicKey = %s\nAllowedIPs = 0.0.0.0/0", clientPrivateKey, clientIP.ToString(), listenPort, serverPublicKey.String())

	fmt.Println(serverConfig)
	fmt.Println()
	fmt.Println(clientConfig)
}
