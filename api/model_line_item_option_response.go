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

// LineItemOptionResponse struct for LineItemOptionResponse
type LineItemOptionResponse struct {
	Data *LineItemOptionResponseData `json:"data,omitempty"`
}

// NewLineItemOptionResponse instantiates a new LineItemOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionResponse() *LineItemOptionResponse {
	this := LineItemOptionResponse{}
	return &this
}

// NewLineItemOptionResponseWithDefaults instantiates a new LineItemOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionResponseWithDefaults() *LineItemOptionResponse {
	this := LineItemOptionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemOptionResponse) GetData() LineItemOptionResponseData {
	if o == nil || o.Data == nil {
		var ret LineItemOptionResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionResponse) GetDataOk() (*LineItemOptionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemOptionResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemOptionResponseData and assigns it to the Data field.
func (o *LineItemOptionResponse) SetData(v LineItemOptionResponseData) {
	o.Data = &v
}

func (o LineItemOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionResponse struct {
	value *LineItemOptionResponse
	isSet bool
}

func (v NullableLineItemOptionResponse) Get() *LineItemOptionResponse {
	return v.value
}

func (v *NullableLineItemOptionResponse) Set(val *LineItemOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionResponse(val *LineItemOptionResponse) *NullableLineItemOptionResponse {
	return &NullableLineItemOptionResponse{value: val, isSet: true}
}

func (v NullableLineItemOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}