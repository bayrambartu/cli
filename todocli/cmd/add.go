package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("you must a write todo content.")
			return
		}
		todos := readTodos()
		todo := args[0]
		todos = append(todos, todo)
		saveTodos(todos)
		fmt.Println("Todo added:", todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
