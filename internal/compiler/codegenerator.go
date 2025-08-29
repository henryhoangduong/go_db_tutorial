package compiler

import (
	"strings"

	"github.com/henryhoangduong/go_db_tutorial/commons/models"
)

type Generator interface {
	Generate(tokens []Token) *models.ByteCode
}
type generator struct {
}

func (g *generator) GenerateCode(tokens []Token) *models.ByteCode {
	var bt *models.ByteCode
	if tokens[0].Type == TokenKeyword && strings.ToUpper(tokens[0].value) == "INSERT" {
		bt = g.generateCodeInsert(tokens)

	}
	if tokens[0].Type == TokenKeyword && strings.ToUpper(tokens[0].value) == "SELECT" {
		bt = g.generateCodeSelect(tokens)

	}

	return bt

}
func stringToStringPtr(s string) *string {
	return &s
}
func (g *generator) generateCodeInsert(tokens []Token) *models.ByteCode {

}
func (g *generator) generateCodeSelect(tokens []Token) *models.ByteCode {
	var bt *models.ByteCode

	ints := []models.ByteCodeValue{
		{
			Type:       models.ByteCodeOperationTypeInsert,
			Identifier: stringToStringPtr("SELECT"),
		},
	}
	varNames := []models.ByteCodeValue{}
	i := 1
	for i < len(tokens) {
		if tokens[i].Type == TokenKeyword && strings.ToUpper(tokens[i].value) == "FROM" {
			break
		}
		if tokens[i].Type == TokenIdentifier {
			v := models.ByteCodeValue{
				Type:       models.ByteCodeOperationTypeIdentifier,
				Identifier: stringToStringPtr(tokens[i].value),
			}
			varNames = append(varNames, v)
		}
		i++
	}
	i++
	tblName := models.ByteCodeValue{
		Type:       models.ByteCodeOperationTypeTableName,
		Identifier: stringToStringPtr(tokens[i].value),
	}
	countVarNams := models.ByteCodeValue{
		Type:  models.ByteCodeOperationTypeCount,
		Count: len(varNames),
	}
	ints = append(ints, tblName)
	ints = append(ints, countVarNams)
	ints = append(ints, varNames...)
	return bt
}
