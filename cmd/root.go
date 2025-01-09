package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile string
var cfgFile string

func initConfig() {
	viper.SetConfigName(".tri")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("tri")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set dta file using --datafile.")
	}
	RootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store todos")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.tri.yaml)")

}

var RootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long:  "Tri will help you get more done in less time. It's designed to be as simple as possible to help you accomplish your goals.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("Error starting CLI: %v", err)
		os.Exit(-1)
	}
}
