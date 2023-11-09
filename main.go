package main

import (
	"fmt"

	"github.com/clagraff/argparse"
)

func callback(p *argparse.Parser, ns *argparse.Namespace, leftovers []string, err error) {
	if err != nil {
		switch err.(type) {
		case argparse.ShowHelpErr, argparse.ShowVersionErr:
			return
		default:
			fmt.Println(err, "\n")
			p.ShowHelp()
		}

		return
	}

}

func main() {
	p := argparse.NewParser("Welcome to a simple cypto-tool", callback).Version("0.0.1")
	p.AddHelp().AddVersion()

}
