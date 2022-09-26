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

// GETReturns200Response struct for GETReturns200Response
type GETReturns200Response struct {
	Data []GETReturns200ResponseDataInner `json:"data,omitempty"`
}

// NewGETReturns200Response instantiates a new GETReturns200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturns200Response() *GETReturns200Response {
	this := GETReturns200Response{}
	return &this
}

// NewGETReturns200ResponseWithDefaults instantiates a new GETReturns200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturns200ResponseWithDefaults() *GETReturns200Response {
	this := GETReturns200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturns200Response) GetData() []GETReturns200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETReturns200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200Response) GetDataOk() ([]GETReturns200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturns200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETReturns200ResponseDataInner and assigns it to the Data field.
func (o *GETReturns200Response) SetData(v []GETReturns200ResponseDataInner) {
	o.Data = v
}

func (o GETReturns200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturns200Response struct {
	value *GETReturns200Response
	isSet bool
}

func (v NullableGETReturns200Response) Get() *GETReturns200Response {
	return v.value
}

func (v *NullableGETReturns200Response) Set(val *GETReturns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturns200Response(val *GETReturns200Response) *NullableGETReturns200Response {
	return &NullableGETReturns200Response{value: val, isSet: true}
}

func (v NullableGETReturns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}