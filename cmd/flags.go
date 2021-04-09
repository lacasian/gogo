package cmd

import "github.com/spf13/cobra"

func addDBFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("db.connection-string", "", "Postgres connection string.")
	cmd.PersistentFlags().String("db.host", "localhost", "Database host")
	cmd.PersistentFlags().String("db.port", "5432", "Database port")
	cmd.PersistentFlags().String("db.sslmode", "disable", "Database sslmode")
	cmd.PersistentFlags().String("db.dbname", "simulator", "Database name")
	cmd.PersistentFlags().String("db.user", "core", "Database user")
	cmd.PersistentFlags().String("db.password", "password", "Database password")
	cmd.PersistentFlags().Bool("db.automigrate", true, "Auto run database migrations")
}

func addAPIFlags(cmd *cobra.Command) {
	cmd.Flags().String("api.port", "3001", "HTTP API port")
	cmd.Flags().Bool("api.dev-cors", false, "Enable development cors for HTTP API")
	cmd.Flags().String("api.dev-cors-host", "", "Allowed host for HTTP API dev cors")
}