package main

import (
	"github.com/secmohammed/port-scanner-go/port"
)

func main() {
    port.GetOpenPorts("www.stackoverflow.com", port.PortRange{
        Start: 80,
        End: 1000,
    })

}
