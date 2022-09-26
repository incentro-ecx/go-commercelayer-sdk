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

// PATCHCheckoutComPaymentsCheckoutComPaymentId200Response struct for PATCHCheckoutComPaymentsCheckoutComPaymentId200Response
type PATCHCheckoutComPaymentsCheckoutComPaymentId200Response struct {
	Data *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCheckoutComPaymentsCheckoutComPaymentId200Response instantiates a new PATCHCheckoutComPaymentsCheckoutComPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCheckoutComPaymentsCheckoutComPaymentId200Response() *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response {
	this := PATCHCheckoutComPaymentsCheckoutComPaymentId200Response{}
	return &this
}

// NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseWithDefaults instantiates a new PATCHCheckoutComPaymentsCheckoutComPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseWithDefaults() *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response {
	this := PATCHCheckoutComPaymentsCheckoutComPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) GetData() PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) GetDataOk() (*PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) SetData(v PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response struct {
	value *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response
	isSet bool
}

func (v NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response) Get() *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response {
	return v.value
}

func (v *NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response) Set(val *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response(val *PATCHCheckoutComPaymentsCheckoutComPaymentId200Response) *NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response {
	return &NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCheckoutComPaymentsCheckoutComPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}