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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ModelValidateByTag(contract contracts.ModelContract, tag string, data interface{}) error {

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not a map[string]interface{}")
	}

	for _, prop := range contract.Properties {

		issetTag := contains(prop.Tags, tag)

		if !issetTag {
			continue
		}

		// Проверка наличия ключа в map:
		if _, ok := dataMap[prop.Key]; !ok {
			return fmt.Errorf("Не заполнено поле '%s'", prop.GetLabel())
		}

		err := PropValidate(prop, dataMap[prop.Key])
		if err != nil {
			return err
		}
	}

	return nil
}
