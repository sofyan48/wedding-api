package excel

import "github.com/360EntSecGroup-Skylar/excelize"

type Contract interface {
	Read(path string) (*excelize.File, error)
}
