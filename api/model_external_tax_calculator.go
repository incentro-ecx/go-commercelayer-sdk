/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ExternalTaxCalculator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalTaxCalculator{}

// ExternalTaxCalculator struct for ExternalTaxCalculator
type ExternalTaxCalculator struct {
	Data *ExternalTaxCalculatorData `json:"data,omitempty"`
}

// NewExternalTaxCalculator instantiates a new ExternalTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculator() *ExternalTaxCalculator {
	this := ExternalTaxCalculator{}
	return &this
}

// NewExternalTaxCalculatorWithDefaults instantiates a new ExternalTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorWithDefaults() *ExternalTaxCalculator {
	this := ExternalTaxCalculator{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExternalTaxCalculator) GetData() ExternalTaxCalculatorData {
	if o == nil || IsNil(o.Data) {
		var ret ExternalTaxCalculatorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculator) GetDataOk() (*ExternalTaxCalculatorData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExternalTaxCalculator) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ExternalTaxCalculatorData and assigns it to the Data field.
func (o *ExternalTaxCalculator) SetData(v ExternalTaxCalculatorData) {
	o.Data = &v
}

func (o ExternalTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalTaxCalculator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableExternalTaxCalculator struct {
	value *ExternalTaxCalculator
	isSet bool
}

func (v NullableExternalTaxCalculator) Get() *ExternalTaxCalculator {
	return v.value
}

func (v *NullableExternalTaxCalculator) Set(val *ExternalTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculator(val *ExternalTaxCalculator) *NullableExternalTaxCalculator {
	return &NullableExternalTaxCalculator{value: val, isSet: true}
}

func (v NullableExternalTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
