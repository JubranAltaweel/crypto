/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto_go/crypto/lib"
	"fmt"

	"github.com/spf13/cobra"
)

// hashCmd represents the hash command
var hashCmd *cobra.Command

func runHash(cmd *cobra.Command, args []string) error {
	raw, err := hashCmd.Flags().GetBool("raw")
	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}

	message, err := hashCmd.Flags().GetString("message")
	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}

	if message != "" {
		output := lib.Hash(message, raw)
		fmt.Printf("The output is: %s\n", output)
	}

	checksum, err := hashCmd.Flags().GetString("checksum")

	if err != nil {
		fmt.Println("Wrong flag")
		return err
	}

	if checksum != "" {
		output, err := lib.Checksum(checksum, raw)
		if err != nil {
			fmt.Println("Wrong flag")
			return err
		}
		fmt.Printf("The output is: %s\n", output)
	}
	return nil
}

func init() {
	hashCmd = &cobra.Command{
		Use:   "hash",
		Short: "Hash and checksum with SHA256",
		RunE:  runHash,
	}
	hashCmd.Flags().BoolP("raw", "r", false, "Raw output")
	hashCmd.Flags().StringP("message", "m", "", "Message to be hashed")
	hashCmd.Flags().StringP("checksum", "c", "", "Filepath to check")

	rootCmd.AddCommand(hashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
