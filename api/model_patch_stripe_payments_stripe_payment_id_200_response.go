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

// PATCHStripePaymentsStripePaymentId200Response struct for PATCHStripePaymentsStripePaymentId200Response
type PATCHStripePaymentsStripePaymentId200Response struct {
	Data *PATCHStripePaymentsStripePaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHStripePaymentsStripePaymentId200Response instantiates a new PATCHStripePaymentsStripePaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStripePaymentsStripePaymentId200Response() *PATCHStripePaymentsStripePaymentId200Response {
	this := PATCHStripePaymentsStripePaymentId200Response{}
	return &this
}

// NewPATCHStripePaymentsStripePaymentId200ResponseWithDefaults instantiates a new PATCHStripePaymentsStripePaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStripePaymentsStripePaymentId200ResponseWithDefaults() *PATCHStripePaymentsStripePaymentId200Response {
	this := PATCHStripePaymentsStripePaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHStripePaymentsStripePaymentId200Response) GetData() PATCHStripePaymentsStripePaymentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHStripePaymentsStripePaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStripePaymentsStripePaymentId200Response) GetDataOk() (*PATCHStripePaymentsStripePaymentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHStripePaymentsStripePaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHStripePaymentsStripePaymentId200Response) SetData(v PATCHStripePaymentsStripePaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHStripePaymentsStripePaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHStripePaymentsStripePaymentId200Response struct {
	value *PATCHStripePaymentsStripePaymentId200Response
	isSet bool
}

func (v NullablePATCHStripePaymentsStripePaymentId200Response) Get() *PATCHStripePaymentsStripePaymentId200Response {
	return v.value
}

func (v *NullablePATCHStripePaymentsStripePaymentId200Response) Set(val *PATCHStripePaymentsStripePaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStripePaymentsStripePaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStripePaymentsStripePaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStripePaymentsStripePaymentId200Response(val *PATCHStripePaymentsStripePaymentId200Response) *NullablePATCHStripePaymentsStripePaymentId200Response {
	return &NullablePATCHStripePaymentsStripePaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHStripePaymentsStripePaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStripePaymentsStripePaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}