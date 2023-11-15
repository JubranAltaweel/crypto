/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto_go/crypto/lib"
	"fmt"

	"github.com/spf13/cobra"
)

// aesCmd represents the aes command
var aesCmd *cobra.Command

func runAES(cmd *cobra.Command, args []string) error {
	decrypt, err := aesCmd.Flags().GetBool("decrypt")

	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}

	text, err := aesCmd.Flags().GetString("text")
	if err != nil {
		fmt.Println("Wrong value for text")
		return err
	}

	key, err := aesCmd.Flags().GetString("key")
	if err != nil {
		fmt.Println("Wrong value for key")
		return err
	}

	outfile, err := aesCmd.Flags().GetString("outfile")
	if err != nil {
		fmt.Println("Wrong value for outfile")
		return err
	}

	if decrypt {
		message, new_err := lib.Decrypt(text, key, outfile)
		if new_err != nil {
			fmt.Println(new_err)
			return new_err
		}

		fmt.Printf("The message is: %s\n", message)
		return nil

	} else {
		message, new_err := lib.Encrypt(text, key, outfile)
		if new_err != nil {
			fmt.Println(new_err)
			return new_err
		}

		fmt.Printf("The message is: %s\n", message)
		return nil
	}
}

func init() {
	aesCmd = &cobra.Command{
		Use:   "aes",
		Short: "Encrypt/Decrypt with AES",

		RunE: runAES,
	}
	aesCmd.Flags().BoolP("decrypt", "d", false, "Decrypt the cipher")
	aesCmd.Flags().StringP("text", "t", "", "Encrypted text")
	aesCmd.Flags().StringP("key", "k", "", "Key to the cipher")
	aesCmd.Flags().StringP("outfile", "o", "", "Outfile to save into")
	rootCmd.AddCommand(aesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
