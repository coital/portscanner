package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func single(proto, host string) bool {
	address := host
	con, err := net.DialTimeout(proto, address, 60*time.Second)

	if err != nil {
		return false
	}

	defer con.Close()

	return true
}

func main() {
	var host = os.Args[1] + ":" + os.Args[2]
	var protocol = os.Args[3]
	scan := single(string(protocol), host)
	if scan == true {
		fmt.Printf("%s %s open", strings.ToUpper(protocol), host)
	} else {
		fmt.Printf("%s %s not open", strings.ToUpper(protocol), host)
	}
}
