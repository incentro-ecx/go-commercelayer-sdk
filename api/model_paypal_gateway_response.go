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

// PaypalGatewayResponse struct for PaypalGatewayResponse
type PaypalGatewayResponse struct {
	Data *PaypalGatewayResponseData `json:"data,omitempty"`
}

// NewPaypalGatewayResponse instantiates a new PaypalGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalGatewayResponse() *PaypalGatewayResponse {
	this := PaypalGatewayResponse{}
	return &this
}

// NewPaypalGatewayResponseWithDefaults instantiates a new PaypalGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalGatewayResponseWithDefaults() *PaypalGatewayResponse {
	this := PaypalGatewayResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaypalGatewayResponse) GetData() PaypalGatewayResponseData {
	if o == nil || o.Data == nil {
		var ret PaypalGatewayResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalGatewayResponse) GetDataOk() (*PaypalGatewayResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaypalGatewayResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PaypalGatewayResponseData and assigns it to the Data field.
func (o *PaypalGatewayResponse) SetData(v PaypalGatewayResponseData) {
	o.Data = &v
}

func (o PaypalGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePaypalGatewayResponse struct {
	value *PaypalGatewayResponse
	isSet bool
}

func (v NullablePaypalGatewayResponse) Get() *PaypalGatewayResponse {
	return v.value
}

func (v *NullablePaypalGatewayResponse) Set(val *PaypalGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalGatewayResponse(val *PaypalGatewayResponse) *NullablePaypalGatewayResponse {
	return &NullablePaypalGatewayResponse{value: val, isSet: true}
}

func (v NullablePaypalGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}