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

// ManualGatewayResponse struct for ManualGatewayResponse
type ManualGatewayResponse struct {
	Data *ManualGatewayResponseData `json:"data,omitempty"`
}

// NewManualGatewayResponse instantiates a new ManualGatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayResponse() *ManualGatewayResponse {
	this := ManualGatewayResponse{}
	return &this
}

// NewManualGatewayResponseWithDefaults instantiates a new ManualGatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayResponseWithDefaults() *ManualGatewayResponse {
	this := ManualGatewayResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ManualGatewayResponse) GetData() ManualGatewayResponseData {
	if o == nil || o.Data == nil {
		var ret ManualGatewayResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGatewayResponse) GetDataOk() (*ManualGatewayResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ManualGatewayResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ManualGatewayResponseData and assigns it to the Data field.
func (o *ManualGatewayResponse) SetData(v ManualGatewayResponseData) {
	o.Data = &v
}

func (o ManualGatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualGatewayResponse struct {
	value *ManualGatewayResponse
	isSet bool
}

func (v NullableManualGatewayResponse) Get() *ManualGatewayResponse {
	return v.value
}

func (v *NullableManualGatewayResponse) Set(val *ManualGatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayResponse(val *ManualGatewayResponse) *NullableManualGatewayResponse {
	return &NullableManualGatewayResponse{value: val, isSet: true}
}

func (v NullableManualGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}