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

// GETAdjustments200Response struct for GETAdjustments200Response
type GETAdjustments200Response struct {
	Data []GETAdjustments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETAdjustments200Response instantiates a new GETAdjustments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdjustments200Response() *GETAdjustments200Response {
	this := GETAdjustments200Response{}
	return &this
}

// NewGETAdjustments200ResponseWithDefaults instantiates a new GETAdjustments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdjustments200ResponseWithDefaults() *GETAdjustments200Response {
	this := GETAdjustments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdjustments200Response) GetData() []GETAdjustments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETAdjustments200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdjustments200Response) GetDataOk() ([]GETAdjustments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdjustments200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETAdjustments200ResponseDataInner and assigns it to the Data field.
func (o *GETAdjustments200Response) SetData(v []GETAdjustments200ResponseDataInner) {
	o.Data = v
}

func (o GETAdjustments200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdjustments200Response struct {
	value *GETAdjustments200Response
	isSet bool
}

func (v NullableGETAdjustments200Response) Get() *GETAdjustments200Response {
	return v.value
}

func (v *NullableGETAdjustments200Response) Set(val *GETAdjustments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdjustments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdjustments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdjustments200Response(val *GETAdjustments200Response) *NullableGETAdjustments200Response {
	return &NullableGETAdjustments200Response{value: val, isSet: true}
}

func (v NullableGETAdjustments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdjustments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}