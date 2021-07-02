package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/smartystreets/cproxy/v2"
)

var opts options

func init() {
	flag.StringVar(&opts.addr, "addr", "localhost:9090", "addr")
	flag.StringVar(&opts.addr, "a", "localhost:9090", "addr")
	flag.Parse()
}

func main() {
	handler := cproxy.New()
	l, err := net.Listen("tcp", opts.addr)
	if err != nil {
		log.Fatal(err)
	}
	printAddr(l)
	http.Serve(l, handler)
}

func printAddr(l net.Listener) {
	fmt.Print("Listening on: ")
	la := l.Addr().(*net.TCPAddr)
	if la.IP.IsLoopback() {
		fmt.Printf("http://127.0.0.1:%d\n", la.Port)
		return
	}

	fmt.Println()

	// Print all ipv4 addrs
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalln(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP.To4()
			case *net.IPAddr:
				ip = v.IP.To4()
			}
			if ip != nil {
				fmt.Printf("  - http://%v:%d\n", ip, la.Port)
			}
		}
	}
}

type options struct {
	addr string
}
