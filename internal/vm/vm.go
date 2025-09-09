package vm

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/henryhoangduong/go_db_tutorial/commons/models"
)

const FILE_NAME = "db.txt"

type VM interface {
	ExecuteByteCode() *models.VMResult
}
type MiniVM struct {
	TableName string
	byteCode  *models.ByteCode
}

func NewVM(TableName string, byteCode *models.ByteCode) VM {
	return &MiniVM{TableName, byteCode}
}

func (vm_ *MiniVM) ExecuteByteCode() *models.VMResult {
	tbl, err := vm_.dbOpen()
	res := &models.VMResult{}
	if err != nil {
		res.Err = err
		return res
	}
	typeB := vm_.byteCode.Instructions[0].Type
	if typeB == models.ByteCodeOperationTypeInsert {
		str, err := vm_.executeInsert()
		res.MSG = str
		res.Err = err
		return res
	}
	res.Err = fmt.Errorf("%s", "Operation NOT_FOUND")
	return res
}
func checkFileExists(fPath string) bool {
	_, err := os.Open(fPath)
	return err == nil
}

func (vm_ *MiniVM) write(p *models.Pager) error {
	f, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	defer f.Close()
	bts, err := json.Marshal()

	if err != nil {
		return err
	}
	f.Write(bts)
	f.Sync()
	return nil
}
func (vm_ *MiniVM) dbOpen() (*models.Table, error) {
	if !checkFileExists(FILE_NAME) {
		f, _ := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_APPEND, 0644)
		f.Close()
	}
	dat, err := os.ReadFile(FILE_NAME)

	if err != nil {
		return nil, err
	}
	var p models.Pager
	if len(dat) > 0 {
		if err := json.Unmarshal(dat, &p); err != nil {
			return nil, err
		}
	}
	t := models.Table{
		Name:    vm_.TableName,
		NumRows: len(p.Pages),
		Pager:   &p,
	}
	return &t, nil
}
