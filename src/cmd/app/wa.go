package app

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/orn-id/wedding-api/src/pkg/waapi"
)

func LoginClient() {
	wac, err := whatsapp.NewConn(10 * time.Second)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	client := waapi.NewWhatsappWebClient()
	client.Login(wac)
}
