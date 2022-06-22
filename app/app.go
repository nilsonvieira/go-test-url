package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/urfave/cli"
)

// Return the Aplication
func Tstnet() *cli.App {
	app := cli.NewApp()
	app.Name = "Application for Network Tests"
	app.Usage = "Search Internet IP and Server Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "nilsonvieira.com.br",
		},
		cli.StringFlag{
			Name:  "url",
			Value: "http://nilsonvieira.com.br",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "8080",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP Internet address",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "ns",
			Usage:  "Search Internet Server Names",
			Flags:  flags,
			Action: searchServers,
		},
		{
			Name:   "sc",
			Usage:  "Get URL Status Code",
			Flags:  flags,
			Action: statusCode,
		},
		{
			Name:   "nc",
			Usage:  "Test IP and Port",
			Flags:  flags,
			Action: netCat,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host) //NameServer
	if error != nil {
		log.Fatal(error)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func statusCode(c *cli.Context) {
	url := c.String("url")

	res, error := http.Get(url)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Status Code:", res.StatusCode)
}

func netCat(c *cli.Context) {
	host := c.String("host")
	port := c.String("port")

	timeout := time.Second
	conn, error := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if error != nil {
		fmt.Println("Connecting error:", error)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Println("Opened", net.JoinHostPort(host, port))
	}
}
