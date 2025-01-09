package cmd

import (
	"fmt"
	"log"

	"github.com/gautam24s/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(
		&priority,
		"priority",
		"p",
		2,
		"Priority:1,2,3",
	)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  "A longer description that spans multiple lines and likery...",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	var items []todo.Item
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}