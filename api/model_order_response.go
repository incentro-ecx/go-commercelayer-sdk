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

// OrderResponse struct for OrderResponse
type OrderResponse struct {
	Data *OrderResponseData `json:"data,omitempty"`
}

// NewOrderResponse instantiates a new OrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponse() *OrderResponse {
	this := OrderResponse{}
	return &this
}

// NewOrderResponseWithDefaults instantiates a new OrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseWithDefaults() *OrderResponse {
	this := OrderResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponse) GetData() OrderResponseData {
	if o == nil || o.Data == nil {
		var ret OrderResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetDataOk() (*OrderResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderResponseData and assigns it to the Data field.
func (o *OrderResponse) SetData(v OrderResponseData) {
	o.Data = &v
}

func (o OrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponse struct {
	value *OrderResponse
	isSet bool
}

func (v NullableOrderResponse) Get() *OrderResponse {
	return v.value
}

func (v *NullableOrderResponse) Set(val *OrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponse(val *OrderResponse) *NullableOrderResponse {
	return &NullableOrderResponse{value: val, isSet: true}
}

func (v NullableOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}