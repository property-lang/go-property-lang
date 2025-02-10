package tests

import (
	"go-property-lang/contracts"
	"go-property-lang/pkg/validation"
	"testing"
)

func TestModelPropKeyValidate(t *testing.T) {

	prop := contracts.PropertyContract{Key: "name", Type: "string", Min: 3, Max: 10, Required: true}

	var props []contracts.PropertyContract
	props = append(props, prop)

	model := contracts.ModelContract{Properties: props}

	err := validation.ModelPropKeyValidate(model, "name", "hello")
	if err != nil {
		t.Errorf("ожидалось, что ошибка не вернётся, но получили: %s", err)
	}

	err = validation.ModelPropKeyValidate(model, "name2", "hello")
	if err == nil {
		t.Errorf("ожидалось, что ошибка не вернётся, но получили: %s", err)
	}

}
