package app

import (
	"fmt"
	"log"
	"net"
	"net/http"

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
			Value: "nilsonvieira.com.br",
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

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status Code:", res.StatusCode)
}
