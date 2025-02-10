package validation

import (
	"fmt"
	"go-property-lang/contracts"
)

func ModelPropKeyValidate(contract contracts.ModelContract, keyProperty string, stringData string) error {
	var foundProperty *contracts.PropertyContract
	for _, prop := range contract.Properties {
		if prop.Key == keyProperty {
			foundProperty = &prop
			break
		}
	}

	if foundProperty == nil {
		return fmt.Errorf("свойство с ключом '%s' не найдено в модели", keyProperty)
	}

	// Валидация свойства через PropValidate
	return PropValidate(*foundProperty, stringData)
}
