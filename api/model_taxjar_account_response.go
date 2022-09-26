/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TaxjarAccountResponse struct for TaxjarAccountResponse
type TaxjarAccountResponse struct {
	Data *TaxjarAccountResponseData `json:"data,omitempty"`
}

// NewTaxjarAccountResponse instantiates a new TaxjarAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountResponse() *TaxjarAccountResponse {
	this := TaxjarAccountResponse{}
	return &this
}

// NewTaxjarAccountResponseWithDefaults instantiates a new TaxjarAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountResponseWithDefaults() *TaxjarAccountResponse {
	this := TaxjarAccountResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TaxjarAccountResponse) GetData() TaxjarAccountResponseData {
	if o == nil || o.Data == nil {
		var ret TaxjarAccountResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountResponse) GetDataOk() (*TaxjarAccountResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TaxjarAccountResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TaxjarAccountResponseData and assigns it to the Data field.
func (o *TaxjarAccountResponse) SetData(v TaxjarAccountResponseData) {
	o.Data = &v
}

func (o TaxjarAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxjarAccountResponse struct {
	value *TaxjarAccountResponse
	isSet bool
}

func (v NullableTaxjarAccountResponse) Get() *TaxjarAccountResponse {
	return v.value
}

func (v *NullableTaxjarAccountResponse) Set(val *TaxjarAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountResponse(val *TaxjarAccountResponse) *NullableTaxjarAccountResponse {
	return &NullableTaxjarAccountResponse{value: val, isSet: true}
}

func (v NullableTaxjarAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}