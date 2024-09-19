package cmd

import (
	"fmt"
	"math/rand/v2"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate passwords",
	Long: `generate paswords for the user with customization options

	For Example:
	passwordGen generate -l 12 -d -s
	`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "to give the length of the password to be generated")
	generateCmd.Flags().BoolP("digits", "d", false, "to add digits to the password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "to add special charachters to the password")
}

func generatePassword(cmd *cobra.Command, args []string) {

	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()-=[]\\;',./_+{}|:<>?"
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.IntN(len(charset))]
	}
	fmt.Println("Generated Password")
	fmt.Println(string(password))

}
