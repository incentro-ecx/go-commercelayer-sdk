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

// OrderCopyResponse struct for OrderCopyResponse
type OrderCopyResponse struct {
	Data *OrderCopyResponseData `json:"data,omitempty"`
}

// NewOrderCopyResponse instantiates a new OrderCopyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyResponse() *OrderCopyResponse {
	this := OrderCopyResponse{}
	return &this
}

// NewOrderCopyResponseWithDefaults instantiates a new OrderCopyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyResponseWithDefaults() *OrderCopyResponse {
	this := OrderCopyResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderCopyResponse) GetData() OrderCopyResponseData {
	if o == nil || o.Data == nil {
		var ret OrderCopyResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyResponse) GetDataOk() (*OrderCopyResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderCopyResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderCopyResponseData and assigns it to the Data field.
func (o *OrderCopyResponse) SetData(v OrderCopyResponseData) {
	o.Data = &v
}

func (o OrderCopyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCopyResponse struct {
	value *OrderCopyResponse
	isSet bool
}

func (v NullableOrderCopyResponse) Get() *OrderCopyResponse {
	return v.value
}

func (v *NullableOrderCopyResponse) Set(val *OrderCopyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyResponse(val *OrderCopyResponse) *NullableOrderCopyResponse {
	return &NullableOrderCopyResponse{value: val, isSet: true}
}

func (v NullableOrderCopyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}