package excel

import "github.com/360EntSecGroup-Skylar/excelize"

type excel struct {
}

func NewExcel() Contract {
	return &excel{}
}

func (e *excel) Read(path string) (*excelize.File, error) {
	return excelize.OpenFile(path)
}
