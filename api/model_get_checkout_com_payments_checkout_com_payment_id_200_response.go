/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETCheckoutComPaymentsCheckoutComPaymentId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCheckoutComPaymentsCheckoutComPaymentId200Response{}

// GETCheckoutComPaymentsCheckoutComPaymentId200Response struct for GETCheckoutComPaymentsCheckoutComPaymentId200Response
type GETCheckoutComPaymentsCheckoutComPaymentId200Response struct {
	Data *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseData `json:"data,omitempty"`
}

// NewGETCheckoutComPaymentsCheckoutComPaymentId200Response instantiates a new GETCheckoutComPaymentsCheckoutComPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComPaymentsCheckoutComPaymentId200Response() *GETCheckoutComPaymentsCheckoutComPaymentId200Response {
	this := GETCheckoutComPaymentsCheckoutComPaymentId200Response{}
	return &this
}

// NewGETCheckoutComPaymentsCheckoutComPaymentId200ResponseWithDefaults instantiates a new GETCheckoutComPaymentsCheckoutComPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComPaymentsCheckoutComPaymentId200ResponseWithDefaults() *GETCheckoutComPaymentsCheckoutComPaymentId200Response {
	this := GETCheckoutComPaymentsCheckoutComPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200Response) GetData() GETCheckoutComPaymentsCheckoutComPaymentId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCheckoutComPaymentsCheckoutComPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200Response) GetDataOk() (*GETCheckoutComPaymentsCheckoutComPaymentId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCheckoutComPaymentsCheckoutComPaymentId200ResponseData and assigns it to the Data field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200Response) SetData(v GETCheckoutComPaymentsCheckoutComPaymentId200ResponseData) {
	o.Data = &v
}

func (o GETCheckoutComPaymentsCheckoutComPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCheckoutComPaymentsCheckoutComPaymentId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response struct {
	value *GETCheckoutComPaymentsCheckoutComPaymentId200Response
	isSet bool
}

func (v NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response) Get() *GETCheckoutComPaymentsCheckoutComPaymentId200Response {
	return v.value
}

func (v *NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response) Set(val *GETCheckoutComPaymentsCheckoutComPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComPaymentsCheckoutComPaymentId200Response(val *GETCheckoutComPaymentsCheckoutComPaymentId200Response) *NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response {
	return &NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response{value: val, isSet: true}
}

func (v NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComPaymentsCheckoutComPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
