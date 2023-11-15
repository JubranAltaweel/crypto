/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto_go/crypto/lib"
	"fmt"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd *cobra.Command

func runBase64(cmd *cobra.Command, args []string) error {
	urlsafe, err := base64Cmd.Flags().GetBool("urlsafe")
	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}

	decode, err := base64Cmd.Flags().GetBool("decode")
	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}

	message, err := base64Cmd.Flags().GetString("message")
	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}
	if message != "" {
		if decode {
			output, err := lib.Decode(message, urlsafe)
			if err != nil {
				fmt.Println("Wrong flag")
				return err
			}
			fmt.Printf("The output is: %s\n", output)
			return nil
		}
		output := lib.Encode(message, urlsafe)
		fmt.Printf("The output is: %s\n", output)
		return nil
	} else {
		fmt.Printf("You need to have a message")
		return nil
	}

}

func init() {
	base64Cmd = &cobra.Command{
		Use:   "base64",
		Short: "Decode/Encode with base64",

		RunE: runBase64,
	}

	base64Cmd.Flags().BoolP("urlsafe", "u", false, "URL safe output")
	base64Cmd.Flags().StringP("message", "m", "", "Message to be encoded/decoded")
	base64Cmd.Flags().BoolP("decode", "d", false, "Decode")
	rootCmd.AddCommand(base64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
