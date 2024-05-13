package gameTemplateService

import (
	"encoding/json"
	"errors"
	"luckyChess/entities"
	"os"
)

type GameTemplateService struct {
}

func New() *GameTemplateService {
	return &GameTemplateService{}
}

func (s GameTemplateService) GetTemplate(name string) entities.BoardTemplate {
	bytes, err := os.ReadFile("gameTemplates/" + name + ".json")

	defer func() {
		if err == nil {
			return
		}
		panic(err)
	}()

	jsonOut := entities.BoardTemplate{}

	json.Unmarshal(bytes, &jsonOut.Template)

	return jsonOut
}

func (s GameTemplateService) NewTemplate(name string, template entities.BoardTemplate) error {
	return errors.New("not implemented")
}
func (s GameTemplateService) AtlerTemplate(name string, template entities.BoardTemplate) error {
	return errors.New("not implemented")
}

func (s GameTemplateService) RemoveTemplate(name string, template entities.BoardTemplate) error {
	return errors.New("not implemented")
}
