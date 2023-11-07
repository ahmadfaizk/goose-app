package cmd

import (
	"fmt"

	"github.com/ahmadfaizk/goose-app/config"
	"github.com/ahmadfaizk/goose-app/db/migrate"
	"github.com/spf13/cobra"
)

func migrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate your database",
		Long:  `Migrate your database`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(migrateUpCmd)
	cmd.AddCommand(migrateDownCmd)
	cmd.AddCommand(migrateResetCmd)
	cmd.AddCommand(migrateRedoCmd)
	cmd.AddCommand(migrateVersionCmd)

	return cmd
}

func getMigrateInstance() *migrate.Migrate {
	cfg := config.Load()
	m, err := migrate.New(cfg.Database)
	if err != nil {
		panic(err)
	}
	return m
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Runs all available migrations",
	Long:  `Runs all available migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigrateInstance()
		if err := m.Up(); err != nil {
			panic(err)
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rolls back a single migration",
	Long:  `Rolls back a single migration`,
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigrateInstance()
		if err := m.Down(); err != nil {
			panic(err)
		}
	},
}

var migrateResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Rolls back all migrations",
	Long:  `Rolls back all migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigrateInstance()
		if err := m.Reset(); err != nil {
			panic(err)
		}
	},
}

var migrateRedoCmd = &cobra.Command{
	Use:   "redo",
	Short: "Rolls back a single migration and then runs it again",
	Long:  `Rolls back a single migration and then runs it again`,
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigrateInstance()
		if err := m.Redo(); err != nil {
			panic(err)
		}
	},
}

var migrateVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of the database",
	Long:  `Print the current version of the database`,
	Run: func(cmd *cobra.Command, args []string) {
		m := getMigrateInstance()
		version, err := m.Version()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Current version: %d\n", version)
	},
}
