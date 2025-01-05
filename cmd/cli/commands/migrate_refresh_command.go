package commands

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var migrateRefreshCommand = &cobra.Command{
	Use:   "migrate:refresh",
	Short: "Refresh database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()
		if err != nil {
			color.Yellowln(err.Error())
			return
		}

		err = m.Down()
		if err != nil {
			color.Redf("error drop: %v \n", err)
			return
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error up: %v \n", err)
			return
		}

		color.Greenln("Migration refresh success")
	},
}

func init() {
	rootCmd.AddCommand(migrateRefreshCommand)
}
