/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TaxCategoryDataRelationshipsTaxCalculator - struct for TaxCategoryDataRelationshipsTaxCalculator
type TaxCategoryDataRelationshipsTaxCalculator struct {
	AvalaraAccount        *AvalaraAccount
	ExternalTaxCalculator *ExternalTaxCalculator
	ManualTaxCalculator   *ManualTaxCalculator
	TaxjarAccount         *TaxjarAccount
}

// AvalaraAccountAsTaxCategoryDataRelationshipsTaxCalculator is a convenience function that returns AvalaraAccount wrapped in TaxCategoryDataRelationshipsTaxCalculator
func AvalaraAccountAsTaxCategoryDataRelationshipsTaxCalculator(v *AvalaraAccount) TaxCategoryDataRelationshipsTaxCalculator {
	return TaxCategoryDataRelationshipsTaxCalculator{
		AvalaraAccount: v,
	}
}

// ExternalTaxCalculatorAsTaxCategoryDataRelationshipsTaxCalculator is a convenience function that returns ExternalTaxCalculator wrapped in TaxCategoryDataRelationshipsTaxCalculator
func ExternalTaxCalculatorAsTaxCategoryDataRelationshipsTaxCalculator(v *ExternalTaxCalculator) TaxCategoryDataRelationshipsTaxCalculator {
	return TaxCategoryDataRelationshipsTaxCalculator{
		ExternalTaxCalculator: v,
	}
}

// ManualTaxCalculatorAsTaxCategoryDataRelationshipsTaxCalculator is a convenience function that returns ManualTaxCalculator wrapped in TaxCategoryDataRelationshipsTaxCalculator
func ManualTaxCalculatorAsTaxCategoryDataRelationshipsTaxCalculator(v *ManualTaxCalculator) TaxCategoryDataRelationshipsTaxCalculator {
	return TaxCategoryDataRelationshipsTaxCalculator{
		ManualTaxCalculator: v,
	}
}

// TaxjarAccountAsTaxCategoryDataRelationshipsTaxCalculator is a convenience function that returns TaxjarAccount wrapped in TaxCategoryDataRelationshipsTaxCalculator
func TaxjarAccountAsTaxCategoryDataRelationshipsTaxCalculator(v *TaxjarAccount) TaxCategoryDataRelationshipsTaxCalculator {
	return TaxCategoryDataRelationshipsTaxCalculator{
		TaxjarAccount: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TaxCategoryDataRelationshipsTaxCalculator) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AvalaraAccount
	err = newStrictDecoder(data).Decode(&dst.AvalaraAccount)
	if err == nil {
		jsonAvalaraAccount, _ := json.Marshal(dst.AvalaraAccount)
		if string(jsonAvalaraAccount) == "{}" { // empty struct
			dst.AvalaraAccount = nil
		} else {
			match++
		}
	} else {
		dst.AvalaraAccount = nil
	}

	// try to unmarshal data into ExternalTaxCalculator
	err = newStrictDecoder(data).Decode(&dst.ExternalTaxCalculator)
	if err == nil {
		jsonExternalTaxCalculator, _ := json.Marshal(dst.ExternalTaxCalculator)
		if string(jsonExternalTaxCalculator) == "{}" { // empty struct
			dst.ExternalTaxCalculator = nil
		} else {
			match++
		}
	} else {
		dst.ExternalTaxCalculator = nil
	}

	// try to unmarshal data into ManualTaxCalculator
	err = newStrictDecoder(data).Decode(&dst.ManualTaxCalculator)
	if err == nil {
		jsonManualTaxCalculator, _ := json.Marshal(dst.ManualTaxCalculator)
		if string(jsonManualTaxCalculator) == "{}" { // empty struct
			dst.ManualTaxCalculator = nil
		} else {
			match++
		}
	} else {
		dst.ManualTaxCalculator = nil
	}

	// try to unmarshal data into TaxjarAccount
	err = newStrictDecoder(data).Decode(&dst.TaxjarAccount)
	if err == nil {
		jsonTaxjarAccount, _ := json.Marshal(dst.TaxjarAccount)
		if string(jsonTaxjarAccount) == "{}" { // empty struct
			dst.TaxjarAccount = nil
		} else {
			match++
		}
	} else {
		dst.TaxjarAccount = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AvalaraAccount = nil
		dst.ExternalTaxCalculator = nil
		dst.ManualTaxCalculator = nil
		dst.TaxjarAccount = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TaxCategoryDataRelationshipsTaxCalculator)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TaxCategoryDataRelationshipsTaxCalculator)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TaxCategoryDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	if src.AvalaraAccount != nil {
		return json.Marshal(&src.AvalaraAccount)
	}

	if src.ExternalTaxCalculator != nil {
		return json.Marshal(&src.ExternalTaxCalculator)
	}

	if src.ManualTaxCalculator != nil {
		return json.Marshal(&src.ManualTaxCalculator)
	}

	if src.TaxjarAccount != nil {
		return json.Marshal(&src.TaxjarAccount)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TaxCategoryDataRelationshipsTaxCalculator) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AvalaraAccount != nil {
		return obj.AvalaraAccount
	}

	if obj.ExternalTaxCalculator != nil {
		return obj.ExternalTaxCalculator
	}

	if obj.ManualTaxCalculator != nil {
		return obj.ManualTaxCalculator
	}

	if obj.TaxjarAccount != nil {
		return obj.TaxjarAccount
	}

	// all schemas are nil
	return nil
}

type NullableTaxCategoryDataRelationshipsTaxCalculator struct {
	value *TaxCategoryDataRelationshipsTaxCalculator
	isSet bool
}

func (v NullableTaxCategoryDataRelationshipsTaxCalculator) Get() *TaxCategoryDataRelationshipsTaxCalculator {
	return v.value
}

func (v *NullableTaxCategoryDataRelationshipsTaxCalculator) Set(val *TaxCategoryDataRelationshipsTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryDataRelationshipsTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryDataRelationshipsTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryDataRelationshipsTaxCalculator(val *TaxCategoryDataRelationshipsTaxCalculator) *NullableTaxCategoryDataRelationshipsTaxCalculator {
	return &NullableTaxCategoryDataRelationshipsTaxCalculator{value: val, isSet: true}
}

func (v NullableTaxCategoryDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryDataRelationshipsTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
