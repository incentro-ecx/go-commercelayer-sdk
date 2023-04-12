/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TaxCategoryCreateDataRelationshipsTaxCalculator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCategoryCreateDataRelationshipsTaxCalculator{}

// TaxCategoryCreateDataRelationshipsTaxCalculator struct for TaxCategoryCreateDataRelationshipsTaxCalculator
type TaxCategoryCreateDataRelationshipsTaxCalculator struct {
	Data TaxCategoryDataRelationshipsTaxCalculatorData `json:"data"`
}

// NewTaxCategoryCreateDataRelationshipsTaxCalculator instantiates a new TaxCategoryCreateDataRelationshipsTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryCreateDataRelationshipsTaxCalculator(data TaxCategoryDataRelationshipsTaxCalculatorData) *TaxCategoryCreateDataRelationshipsTaxCalculator {
	this := TaxCategoryCreateDataRelationshipsTaxCalculator{}
	this.Data = data
	return &this
}

// NewTaxCategoryCreateDataRelationshipsTaxCalculatorWithDefaults instantiates a new TaxCategoryCreateDataRelationshipsTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryCreateDataRelationshipsTaxCalculatorWithDefaults() *TaxCategoryCreateDataRelationshipsTaxCalculator {
	this := TaxCategoryCreateDataRelationshipsTaxCalculator{}
	return &this
}

// GetData returns the Data field value
func (o *TaxCategoryCreateDataRelationshipsTaxCalculator) GetData() TaxCategoryDataRelationshipsTaxCalculatorData {
	if o == nil {
		var ret TaxCategoryDataRelationshipsTaxCalculatorData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryCreateDataRelationshipsTaxCalculator) GetDataOk() (*TaxCategoryDataRelationshipsTaxCalculatorData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxCategoryCreateDataRelationshipsTaxCalculator) SetData(v TaxCategoryDataRelationshipsTaxCalculatorData) {
	o.Data = v
}

func (o TaxCategoryCreateDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCategoryCreateDataRelationshipsTaxCalculator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTaxCategoryCreateDataRelationshipsTaxCalculator struct {
	value *TaxCategoryCreateDataRelationshipsTaxCalculator
	isSet bool
}

func (v NullableTaxCategoryCreateDataRelationshipsTaxCalculator) Get() *TaxCategoryCreateDataRelationshipsTaxCalculator {
	return v.value
}

func (v *NullableTaxCategoryCreateDataRelationshipsTaxCalculator) Set(val *TaxCategoryCreateDataRelationshipsTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryCreateDataRelationshipsTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryCreateDataRelationshipsTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryCreateDataRelationshipsTaxCalculator(val *TaxCategoryCreateDataRelationshipsTaxCalculator) *NullableTaxCategoryCreateDataRelationshipsTaxCalculator {
	return &NullableTaxCategoryCreateDataRelationshipsTaxCalculator{value: val, isSet: true}
}

func (v NullableTaxCategoryCreateDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryCreateDataRelationshipsTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
