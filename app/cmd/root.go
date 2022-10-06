package cmd

import (
	"fmt"
	"os"

	"github.com/koalaylj/zkpass-alice/app/typing"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var settings typing.Settings

var rootCmd = &cobra.Command{
	Use:   "alice",
	Short: "zkpass test client",
	Long: `zkpass test client. For example:

alice version
alice ping
.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	},
}

func Execute() error {
	err := rootCmd.Execute()
	return err
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./settings.json)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName("settings")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		viper.Unmarshal(&settings)
	} else {
		fmt.Println("initConfig failed", err)
		os.Exit(1)
	}
}
