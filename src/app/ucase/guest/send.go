package guest

import (
	"fmt"
	"net/url"
	"os"

	"github.com/orn-id/wedding-api/src/app/appctx"
	"github.com/orn-id/wedding-api/src/app/presentations"
	"github.com/orn-id/wedding-api/src/app/ucase/contract"
	"github.com/orn-id/wedding-api/src/consts"
	"github.com/orn-id/wedding-api/src/pkg/excel"
	"github.com/orn-id/wedding-api/src/pkg/logger"
	"github.com/orn-id/wedding-api/src/pkg/waapi"
	"github.com/orn-id/wedding-api/src/validator"
)

type sendInvitation struct {
	xls       excel.Contract
	chat      waapi.WhatsappWebClient
	validates validator.Validator
}

func NewSendInvitation(validates validator.Validator) contract.UseCase {
	return &sendInvitation{
		xls:       excel.NewExcel(),
		chat:      waapi.NewWhatsappWebClient(),
		validates: validates,
	}
}

func (v *sendInvitation) Serve(data *appctx.Data) appctx.Response {
	req := presentations.SendInvitation{}
	e := data.Cast(&req)
	if e != nil {
		logger.Error(logger.SetMessageFormat("[send-invitation] parsing body request error: %s", e.Error()))
		return appctx.Response{
			Name: consts.ValidationFailure,
		}
	}
	e = v.validates.Request(req)
	if e != nil {
		logger.Error(logger.SetMessageFormat("[send-invitation] validation request error: %s", e.Error()))
		return appctx.Response{
			Name: consts.ValidationFailure,
		}
	}

	u, err := url.Parse(os.Getenv("URL_WEDDING_APP"))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	q := u.Query()
	q.Set("name", req.Name)
	q.Set("addres", req.Address)

	u.RawQuery = q.Encode()

	message := "*Assalamualaikum wr.wb*\n\nHalo kak *" + req.Name + "*\n\nTanpa mengurangi rasa hormat kami, kami mengundang anda di acara pernikahan kami.\nsilahkan buka link: " + u.String() + "\n\nTerima Kasih\n\nHormat Kami:\n*Sofyan Saputra & Besse Sindi Arini*."

	wac, err := v.chat.Conn()
	if err != nil {
		logger.Error(logger.SetMessageFormat("[send-invitation] error: %s", err.Error()))
		return appctx.Response{
			Name:    consts.UnprocessAbleEntity,
			Message: err.Error(),
		}
	}

	err = v.chat.Login(wac)
	if err != nil {
		logger.Error(logger.SetMessageFormat("[send-invitation] error: %s", err.Error()))
		return appctx.Response{
			Name:    consts.UnprocessAbleEntity,
			Message: err.Error(),
		}
	}

	chatID, err := v.chat.Send(wac, req.WaNumber, message)
	status := "send"
	if err != nil {
		status = "not send"
	}

	return appctx.Response{
		Name:    consts.Success,
		Message: "Invitation send",
		Data: map[string]string{
			"to":      req.WaNumber,
			"name":    req.Name,
			"status":  status,
			"chat_id": chatID,
		},
	}
}
