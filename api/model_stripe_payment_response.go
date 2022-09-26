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

// StripePaymentResponse struct for StripePaymentResponse
type StripePaymentResponse struct {
	Data *StripePaymentResponseData `json:"data,omitempty"`
}

// NewStripePaymentResponse instantiates a new StripePaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentResponse() *StripePaymentResponse {
	this := StripePaymentResponse{}
	return &this
}

// NewStripePaymentResponseWithDefaults instantiates a new StripePaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentResponseWithDefaults() *StripePaymentResponse {
	this := StripePaymentResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StripePaymentResponse) GetData() StripePaymentResponseData {
	if o == nil || o.Data == nil {
		var ret StripePaymentResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentResponse) GetDataOk() (*StripePaymentResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StripePaymentResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given StripePaymentResponseData and assigns it to the Data field.
func (o *StripePaymentResponse) SetData(v StripePaymentResponseData) {
	o.Data = &v
}

func (o StripePaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripePaymentResponse struct {
	value *StripePaymentResponse
	isSet bool
}

func (v NullableStripePaymentResponse) Get() *StripePaymentResponse {
	return v.value
}

func (v *NullableStripePaymentResponse) Set(val *StripePaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentResponse(val *StripePaymentResponse) *NullableStripePaymentResponse {
	return &NullableStripePaymentResponse{value: val, isSet: true}
}

func (v NullableStripePaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}