package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"git.sr.ht/~ocelotsloth/goqrz"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goqrz"
	app.Usage = "Search QRZ.com database via CLI"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "login to generate session token (saves to environment)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "QRZ.com Username",
				},
				cli.StringFlag{
					Name:  "pass, p",
					Usage: "QRZ.com Password",
				},
			},
			Action: func(c *cli.Context) error {
				sessionKey, err := goqrz.GetSessionKey(c.String("user"), c.String("pass"), "goqrz")
				if err != nil {
					return err
				}
				fmt.Println(sessionKey)
				return nil
			},
		},
		{
			Name:    "callsign",
			Aliases: []string{"c"},
			Usage:   "callsign [--key=session_key] Call1 [Call2...]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Usage: "Session key from login command (or use GOQRZ_KEY environment variable)",
				},
			},
			Action: func(c *cli.Context) error {
				key := c.String("key")
				if key == "" {
					key = os.Getenv("GOQRZ_KEY")
					if key == "" {
						return errors.New("no session key specified")
					}
				}
				for _, arg := range c.Args() {
					callsign, err := goqrz.GetCallsign(key, arg, "goqrz-cli")
					if err != nil {
						return err
					}
					callJSON, err := json.Marshal(callsign)
					if err != nil {
						return err
					}
					fmt.Println(string(callJSON))
				}

				return nil
			},
		},
		{
			Name:    "dxcc",
			Aliases: []string{"d"},
			Usage:   "dxcc [--key=session_key] dxccID1 [dxccID2...]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Usage: "Session key from login command (or use GOQRZ_KEY environment variable)",
				},
			},
			Action: func(c *cli.Context) error {
				key := c.String("key")
				if key == "" {
					key = os.Getenv("GOQRZ_KEY")
					if key == "" {
						return errors.New("no session key specified")
					}
				}
				for _, arg := range c.Args() {
					dxcc, err := goqrz.GetDXCC(key, arg, "goqrz-cli")
					if err != nil {
						return err
					}
					dxccJSON, err := json.Marshal(dxcc)
					if err != nil {
						return err
					}
					fmt.Println(string(dxccJSON))
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
