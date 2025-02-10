package tests

import (
	"go-property-lang/contracts"
	"go-property-lang/pkg/validation"
	"testing"
)

func TestModelByTag(t *testing.T) {

	var props []contracts.PropertyContract

	props = append(props, contracts.PropertyContract{Key: "name", Label: "Имя", Type: "string", Min: 3, Max: 10, Required: true, Tags: []string{"a", "b"}})
	props = append(props, contracts.PropertyContract{Key: "login", Type: "string", Min: 3, Max: 10, Required: true, Tags: []string{"a", "byLogin"}})
	props = append(props, contracts.PropertyContract{Key: "name3", Type: "string", Min: 3, Max: 10, Required: true, Tags: []string{"other"}})

	model := contracts.ModelContract{Properties: props}

	err := validation.ModelValidateByTag(model, "a", map[string]interface{}{
		"name55": "asfasf",
		"login":  "asfsaf",
	})

	if err == nil || err.Error() != "Не заполнено поле 'Имя'" {
		t.Errorf("ожидалось, что вернётся ошибка: %s", err)
	}

	err = validation.ModelValidateByTag(model, "byLogin", map[string]interface{}{
		"login":  "asfsaf",
		"name55": "asfasf",
	})

	if err != nil {
		t.Errorf("ожидалось, что будет ок: %s", err)
	}

}
