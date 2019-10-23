package cmd

import (
	"fmt"
	"log"

	"github.com/liangwj45/Service-Computing/agenda-go/entity"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register -uUsername -pPassword",
	Short: "Register a account.",
	Long:  "Register a account with the username, password, email and telephone.",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")
		if !entity.ExistUserName(username) {
			entity.AddUser(username, password, email, telephone)
			log.Println("register succesful:", username, password)
			fmt.Println("----------------------------------------")
			entity.WriteUserFile()
		} else {
			log.Println("register failed:", username)
			fmt.Println("----------------------------------------")
			fmt.Println("username has been used")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "")
	registerCmd.Flags().StringP("password", "p", "", "")
	registerCmd.Flags().StringP("email", "e", "", "")
	registerCmd.Flags().StringP("telephone", "t", "", "")
	registerCmd.MarkFlagRequired("username")
	registerCmd.MarkFlagRequired("password")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
