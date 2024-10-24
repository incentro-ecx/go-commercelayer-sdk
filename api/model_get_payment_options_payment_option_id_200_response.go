/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETPaymentOptionsPaymentOptionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPaymentOptionsPaymentOptionId200Response{}

// GETPaymentOptionsPaymentOptionId200Response struct for GETPaymentOptionsPaymentOptionId200Response
type GETPaymentOptionsPaymentOptionId200Response struct {
	Data *GETPaymentOptionsPaymentOptionId200ResponseData `json:"data,omitempty"`
}

// NewGETPaymentOptionsPaymentOptionId200Response instantiates a new GETPaymentOptionsPaymentOptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaymentOptionsPaymentOptionId200Response() *GETPaymentOptionsPaymentOptionId200Response {
	this := GETPaymentOptionsPaymentOptionId200Response{}
	return &this
}

// NewGETPaymentOptionsPaymentOptionId200ResponseWithDefaults instantiates a new GETPaymentOptionsPaymentOptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaymentOptionsPaymentOptionId200ResponseWithDefaults() *GETPaymentOptionsPaymentOptionId200Response {
	this := GETPaymentOptionsPaymentOptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaymentOptionsPaymentOptionId200Response) GetData() GETPaymentOptionsPaymentOptionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETPaymentOptionsPaymentOptionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentOptionsPaymentOptionId200Response) GetDataOk() (*GETPaymentOptionsPaymentOptionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaymentOptionsPaymentOptionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPaymentOptionsPaymentOptionId200ResponseData and assigns it to the Data field.
func (o *GETPaymentOptionsPaymentOptionId200Response) SetData(v GETPaymentOptionsPaymentOptionId200ResponseData) {
	o.Data = &v
}

func (o GETPaymentOptionsPaymentOptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPaymentOptionsPaymentOptionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETPaymentOptionsPaymentOptionId200Response struct {
	value *GETPaymentOptionsPaymentOptionId200Response
	isSet bool
}

func (v NullableGETPaymentOptionsPaymentOptionId200Response) Get() *GETPaymentOptionsPaymentOptionId200Response {
	return v.value
}

func (v *NullableGETPaymentOptionsPaymentOptionId200Response) Set(val *GETPaymentOptionsPaymentOptionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaymentOptionsPaymentOptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaymentOptionsPaymentOptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaymentOptionsPaymentOptionId200Response(val *GETPaymentOptionsPaymentOptionId200Response) *NullableGETPaymentOptionsPaymentOptionId200Response {
	return &NullableGETPaymentOptionsPaymentOptionId200Response{value: val, isSet: true}
}

func (v NullableGETPaymentOptionsPaymentOptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaymentOptionsPaymentOptionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
