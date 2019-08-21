package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	app "github.com/azohra/periscope-proxy/internal/app/periscope_proxy"
)

func main() {

	var strPort = flag.String("port", "", "Proxy port number")
	var strAddr = flag.String("addr", "", "Cluster address")
	flag.Parse()

	port := strings.Trim(*strPort, "")
	addr := strings.Trim(*strAddr, "")
	if port != "" && addr != "" {
		if portInt, err := strconv.Atoi(port); err == nil {
			app.Proxy(portInt, addr)
		} else {
			fmt.Println("Invalid parameters")
		}
	} else {
		fmt.Println("Incorrect usage. Run -h for details.")
	}
}
