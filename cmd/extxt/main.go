package main

import (
	"io"
	"log"
	"os"

	"github.com/ddddddO/extxt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "xxxxx",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "xxxxx",
			},
		},
		Action: func(c *cli.Context) error {
			input := c.String("input")
			if input == "" {
				return errors.New("xxxxxx")
			}

			var output io.WriteCloser
			if c.String("output") == "" {
				output = os.Stdout
			} else {
				// file open
			}

			if err := extxt.Run(output, input); err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}