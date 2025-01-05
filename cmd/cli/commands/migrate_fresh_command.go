package commands

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var migrateFreshCommand = &cobra.Command{
	Use:   "migrate:fresh",
	Short: "Reset database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()
		if err != nil {
			color.Yellowln(err.Error())
			return
		}

		if m == nil {
			color.Yellowln("Please fill database config first")
			return
		}

		err = m.Down()
		if err != nil {
			color.Redln("Migration fresh failed :: ", err.Error())
			return
		}

		color.Greenln("Migration fresh success")
	},
}

func init() {
	rootCmd.AddCommand(migrateFreshCommand)
}
