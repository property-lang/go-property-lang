package repository

import (
	"go-property-lang/contracts"
	"go-property-lang/pkg/reader"
)

func LoadAllModels(path string) ([]contracts.ModelContract, error) {
	var models []contracts.ModelContract

	models, err := reader.ReadDirectory(path)
	if err != nil {
		return models, err
	}

	for _, model := range models {
		ModelsLoaded = append(ModelsLoaded, model)
	}

	return models, nil
}
