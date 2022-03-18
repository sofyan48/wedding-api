package waapi

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
)

type waClient struct {
}

func NewWhatsappWebClient() WhatsappWebClient {
	return &waClient{}
}

func (w *waClient) readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open((os.TempDir() + "/whatsappSession.gob"))
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil

}

func (w *waClient) Conn() (*whatsapp.Conn, error) {
	return whatsapp.NewConn(10 * time.Second)
}

func (w *waClient) writeSession(session whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	// defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

func (w *waClient) Login(wac *whatsapp.Conn) error {
	session, err := w.readSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			fmt.Println("restoring failed: ", err)
			return err
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			fmt.Println("error during login: ", err)
			return err
		}
	}

	err = w.writeSession(session)
	if err != nil {
		fmt.Println("error saving session: ", err.Error())
		return err
	}
	return nil
}

func (w *waClient) Send(wac *whatsapp.Conn, to, message string) (string, error) {
	// wac, err := whatsapp.NewConn(10 * time.Second)
	// if err != nil {
	// 	return "", err
	// }

	// err = w.Login(wac)
	// if err != nil {
	// 	return "", err
	// }

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: to + "@s.whatsapp.net",
		},
		Text: message,
	}

	msgId, err := wac.Send(msg)
	if err != nil {
		return "", err
	}

	return msgId, nil

}
