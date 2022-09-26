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

// GETReturnLineItems200Response struct for GETReturnLineItems200Response
type GETReturnLineItems200Response struct {
	Data []GETReturnLineItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETReturnLineItems200Response instantiates a new GETReturnLineItems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnLineItems200Response() *GETReturnLineItems200Response {
	this := GETReturnLineItems200Response{}
	return &this
}

// NewGETReturnLineItems200ResponseWithDefaults instantiates a new GETReturnLineItems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnLineItems200ResponseWithDefaults() *GETReturnLineItems200Response {
	this := GETReturnLineItems200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturnLineItems200Response) GetData() []GETReturnLineItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETReturnLineItems200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200Response) GetDataOk() ([]GETReturnLineItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturnLineItems200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETReturnLineItems200ResponseDataInner and assigns it to the Data field.
func (o *GETReturnLineItems200Response) SetData(v []GETReturnLineItems200ResponseDataInner) {
	o.Data = v
}

func (o GETReturnLineItems200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturnLineItems200Response struct {
	value *GETReturnLineItems200Response
	isSet bool
}

func (v NullableGETReturnLineItems200Response) Get() *GETReturnLineItems200Response {
	return v.value
}

func (v *NullableGETReturnLineItems200Response) Set(val *GETReturnLineItems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnLineItems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnLineItems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnLineItems200Response(val *GETReturnLineItems200Response) *NullableGETReturnLineItems200Response {
	return &NullableGETReturnLineItems200Response{value: val, isSet: true}
}

func (v NullableGETReturnLineItems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnLineItems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}