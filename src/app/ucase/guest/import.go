package guest

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/orn-id/wedding-api/src/app/appctx"
	"github.com/orn-id/wedding-api/src/app/presentations"
	"github.com/orn-id/wedding-api/src/app/ucase/contract"
	"github.com/orn-id/wedding-api/src/consts"
	"github.com/orn-id/wedding-api/src/pkg/excel"
	"github.com/orn-id/wedding-api/src/pkg/logger"
	"github.com/orn-id/wedding-api/src/pkg/waapi"
)

type imports struct {
	xls  excel.Contract
	chat waapi.WhatsappWebClient
}

func NewImports() contract.UseCase {
	return &imports{
		xls:  excel.NewExcel(),
		chat: waapi.NewWhatsappWebClient(),
	}
}

func (v *imports) Serve(data *appctx.Data) appctx.Response {
	sheet := data.Request.FormValue("sheet")
	os.MkdirAll("./tmp/file", os.ModePerm)
	data.Request.ParseMultipartForm(10 << 20)
	file, _, err := data.Request.FormFile("file")
	if err != nil {
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("tmp/file", "upload-*.xlsx")
	if err != nil {
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}
	tempFile.Write(fileBytes)
	xlsx, err := v.xls.Read(tempFile.Name())
	if err != nil {
		return appctx.Response{
			Name:    consts.InternalFailure,
			Message: err.Error(),
		}
	}
	rows := xlsx.GetRows(sheet)
	reports := []presentations.ImportsReport{}
	listSave := []map[string]string{}

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

	for i := 2; i <= len(rows); i++ {
		name := xlsx.GetCellValue(sheet, fmt.Sprintf("A%d", i))
		waNumber := xlsx.GetCellValue(sheet, fmt.Sprintf("B%d", i))
		address := xlsx.GetCellValue(sheet, fmt.Sprintf("C%d", i))
		u, err := url.Parse(os.Getenv("URL_WEDDING_APP"))
		if err != nil {
			fmt.Println("Error: ", err)
		}
		q := u.Query()
		q.Set("name", name)
		q.Set("addres", address)

		u.RawQuery = q.Encode()

		message := "*Assalamualaikum wr.wb*\n\nHalo kak *" + name + "*\n\nTanpa mengurangi rasa hormat kami, kami mengundang anda di acara pernikahan kami.\nsilahkan buka link: " + u.String() + "\n\nTerima Kasih\n\nHormat Kami:\n*Sofyan Saputra & Besse Sindi Arini*."

		chatID, err := v.chat.Send(wac, waNumber, message)
		status := "send"
		if err != nil {
			fmt.Println("Sending error:> ", err)
			status = "not send"
		}

		listSave = append(listSave, map[string]string{
			"to":      waNumber,
			"name":    name,
			"status":  status,
			"chat_id": chatID,
		})

	}

	totalRows := len(reports)
	err = os.Remove(tempFile.Name())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	return appctx.Response{
		Name:    consts.Success,
		Message: fmt.Sprintf("%d", totalRows) + " Rows Imports success",
		Data:    listSave,
	}
}
