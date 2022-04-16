package app

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/orn-id/wedding-api/src/pkg/waapi"
	"github.com/spf13/cobra"
)

func Scan() *cobra.Command {
	return &cobra.Command{
		Use:   "wa",
		Short: "Login whatsapp web client",
		Run: func(cmd *cobra.Command, args []string) {
			wac, err := whatsapp.NewConn(10 * time.Second)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				os.Exit(1)
			}

			client := waapi.NewWhatsappWebClient()
			client.Login(wac)
		},
	}
}
