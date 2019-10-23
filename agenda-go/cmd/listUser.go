package cmd

import (
	"fmt"
	"log"

	"github.com/liangwj45/Service-Computing/agenda-go/entity"
	"github.com/spf13/cobra"
)

var listUserCmd = &cobra.Command{
	Use:   "listUser",
	Short: "List all registered users.",
	Long:  "List the information of all registered users.",
	Run: func(cmd *cobra.Command, args []string) {
		users := entity.GetAllUser()
		log.Println("query all users")
		fmt.Printf("Query all users\n")
		fmt.Printf("Username\t Email\tTelephone\n")
		for i := range users {
			fmt.Printf("%s\t %s\t%s\n", users[i].Username, users[i].Email, users[i].Telephone)
		}
	},
}

func init() {
	rootCmd.AddCommand(listUserCmd)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
