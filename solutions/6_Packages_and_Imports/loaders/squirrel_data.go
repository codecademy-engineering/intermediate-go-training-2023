package loaders

import (
	"encoding/json"
	"os"

	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
)

func LoadSquirrelData() ([]models.Squirrel, error) {

	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		return nil, err
	}

	squirrels := []models.Squirrel{}
	err = json.Unmarshal(fileContents, &squirrels)
	if err != nil {
		return nil, err
	}

	return squirrels, nil
}
