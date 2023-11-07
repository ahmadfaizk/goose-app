package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "goose-app",
	Short: "Goose App is a CLI application for managing your Goose database migrations.",
	Long:  `Goose App is a CLI application for managing your Goose database migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(migrateCmd())
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.ReadInConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
