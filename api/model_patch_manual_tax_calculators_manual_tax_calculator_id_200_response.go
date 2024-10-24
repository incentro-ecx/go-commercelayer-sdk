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

// checks if the PATCHManualTaxCalculatorsManualTaxCalculatorId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHManualTaxCalculatorsManualTaxCalculatorId200Response{}

// PATCHManualTaxCalculatorsManualTaxCalculatorId200Response struct for PATCHManualTaxCalculatorsManualTaxCalculatorId200Response
type PATCHManualTaxCalculatorsManualTaxCalculatorId200Response struct {
	Data *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData `json:"data,omitempty"`
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200Response instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200Response() *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200Response{}
	return &this
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseWithDefaults instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseWithDefaults() *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) GetData() PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) GetDataOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData and assigns it to the Data field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) SetData(v PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseData) {
	o.Data = &v
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response struct {
	value *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response
	isSet bool
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) Get() *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	return v.value
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) Set(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response) *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response {
	return &NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response{value: val, isSet: true}
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
