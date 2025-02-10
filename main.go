package main

import (
	"fmt"
	"go-property-lang/pkg/reader"
	"go-property-lang/pkg/validation"
)

func main() {

	model, err := reader.ReadModel("./proplangs/userExample.proplang.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(model)
	//fmt.Println(model.Model)
	//fmt.Println(model)
	fmt.Println(model.Properties[0].Key)

	err = validation.ModelPropKeyValidate(model, "name", "t")
	if err != nil {
		fmt.Println(err)
	}

	err = validation.ModelValidateByTag(model, "name", map[string]interface{}{
		"login":  "asfsaf",
		"name55": "asfasf",
	})
	if err != nil {
		fmt.Println(err)
	}

}
