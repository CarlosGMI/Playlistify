package cmd

import (
	"os"

	"github.com/CarlosGMI/Playlistify/cmd/account"
	"github.com/CarlosGMI/Playlistify/cmd/playlist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "playlistify",
	Short: "CLI application to look for a song or artist inside a specific Spotify playlist",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initFlags()
	initConfig()
	initCommands()
}

func initConfig() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".playlistify")
	viper.SetConfigType("json")

	_ = viper.SafeWriteConfig()
	_ = viper.ReadInConfig()
}

func initCommands() {
	// Auth commands
	rootCmd.AddCommand(account.LoginCommand())
	rootCmd.AddCommand(account.LogoutCommand())
	rootCmd.AddCommand(playlist.ListCommand())
	rootCmd.AddCommand(playlist.SearchCommand())
}

func initFlags() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
