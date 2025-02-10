package validation

import (
	"fmt"
	"go-property-lang/contracts"
)

func PropValidateV1(contract contracts.PropertyContract, stringData string) error {

	if len(stringData) < contract.Min {
		return fmt.Errorf("len error: string length is less than Min")
	}

	if len(stringData) > contract.Max {
		return fmt.Errorf("len error: string length exceeds Max")
	}

	return nil
}
