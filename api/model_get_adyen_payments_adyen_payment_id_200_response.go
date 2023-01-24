/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETAdyenPaymentsAdyenPaymentId200Response struct for GETAdyenPaymentsAdyenPaymentId200Response
type GETAdyenPaymentsAdyenPaymentId200Response struct {
	Data *GETAdyenPayments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETAdyenPaymentsAdyenPaymentId200Response instantiates a new GETAdyenPaymentsAdyenPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenPaymentsAdyenPaymentId200Response() *GETAdyenPaymentsAdyenPaymentId200Response {
	this := GETAdyenPaymentsAdyenPaymentId200Response{}
	return &this
}

// NewGETAdyenPaymentsAdyenPaymentId200ResponseWithDefaults instantiates a new GETAdyenPaymentsAdyenPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenPaymentsAdyenPaymentId200ResponseWithDefaults() *GETAdyenPaymentsAdyenPaymentId200Response {
	this := GETAdyenPaymentsAdyenPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) GetData() GETAdyenPayments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETAdyenPayments200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) GetDataOk() (*GETAdyenPayments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAdyenPayments200ResponseDataInner and assigns it to the Data field.
func (o *GETAdyenPaymentsAdyenPaymentId200Response) SetData(v GETAdyenPayments200ResponseDataInner) {
	o.Data = &v
}

func (o GETAdyenPaymentsAdyenPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenPaymentsAdyenPaymentId200Response struct {
	value *GETAdyenPaymentsAdyenPaymentId200Response
	isSet bool
}

func (v NullableGETAdyenPaymentsAdyenPaymentId200Response) Get() *GETAdyenPaymentsAdyenPaymentId200Response {
	return v.value
}

func (v *NullableGETAdyenPaymentsAdyenPaymentId200Response) Set(val *GETAdyenPaymentsAdyenPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenPaymentsAdyenPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenPaymentsAdyenPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenPaymentsAdyenPaymentId200Response(val *GETAdyenPaymentsAdyenPaymentId200Response) *NullableGETAdyenPaymentsAdyenPaymentId200Response {
	return &NullableGETAdyenPaymentsAdyenPaymentId200Response{value: val, isSet: true}
}

func (v NullableGETAdyenPaymentsAdyenPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenPaymentsAdyenPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
