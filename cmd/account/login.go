package account

import (
	"github.com/CarlosGMI/Playlistify/services"
	"github.com/spf13/cobra"
)

func LoginCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "login",
		Short: "A brief description of your command",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := services.Authenticate(); err != nil {
				return err
			}

			return nil
		},
	}

	return command
}
