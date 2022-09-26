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

// GETAuthorizations200Response struct for GETAuthorizations200Response
type GETAuthorizations200Response struct {
	Data []GETAuthorizations200ResponseDataInner `json:"data,omitempty"`
}

// NewGETAuthorizations200Response instantiates a new GETAuthorizations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizations200Response() *GETAuthorizations200Response {
	this := GETAuthorizations200Response{}
	return &this
}

// NewGETAuthorizations200ResponseWithDefaults instantiates a new GETAuthorizations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizations200ResponseWithDefaults() *GETAuthorizations200Response {
	this := GETAuthorizations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAuthorizations200Response) GetData() []GETAuthorizations200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETAuthorizations200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizations200Response) GetDataOk() ([]GETAuthorizations200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAuthorizations200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETAuthorizations200ResponseDataInner and assigns it to the Data field.
func (o *GETAuthorizations200Response) SetData(v []GETAuthorizations200ResponseDataInner) {
	o.Data = v
}

func (o GETAuthorizations200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAuthorizations200Response struct {
	value *GETAuthorizations200Response
	isSet bool
}

func (v NullableGETAuthorizations200Response) Get() *GETAuthorizations200Response {
	return v.value
}

func (v *NullableGETAuthorizations200Response) Set(val *GETAuthorizations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizations200Response(val *GETAuthorizations200Response) *NullableGETAuthorizations200Response {
	return &NullableGETAuthorizations200Response{value: val, isSet: true}
}

func (v NullableGETAuthorizations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}