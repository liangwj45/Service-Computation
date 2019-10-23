package cmd

import (
	"fmt"
	"log"

	"github.com/liangwj45/Service-Computing/agenda-go/entity"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login -uUsername -pPassword",
	Short: "Log in your account.",
	Long:  "Log in your account with the username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		if entity.CheckPassword(username, password) {
			log.Println("login succesful:", username)
			fmt.Println("----------------------------------------")
			fmt.Println("welcome")
		} else {
			log.Println("login failed:", username, password)
			fmt.Println("----------------------------------------")
			fmt.Println("username or password incorrect")
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "")
	loginCmd.Flags().StringP("password", "p", "", "")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
