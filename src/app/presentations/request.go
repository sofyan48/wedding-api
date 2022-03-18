package presentations

//go:generate easytags $GOFILE form,json

type SendInvitation struct {
	Name     string `json:"name"`
	WaNumber string `json:"to"`
	Address  string `json:"address"`
}

type ImportsReport struct {
	Row       int    `json:"row"`
	Name      string `json:"name"`
	Validates bool   `json:"is_valid"`
}
