package waapi

import "github.com/Rhymen/go-whatsapp"

type WhatsappWebClient interface {
	Conn() (*whatsapp.Conn, error)
	Send(wac *whatsapp.Conn, to, message string) (string, error)
	Login(wac *whatsapp.Conn) error
}
