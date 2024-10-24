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

// checks if the POSTSkus201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkus201Response{}

// POSTSkus201Response struct for POSTSkus201Response
type POSTSkus201Response struct {
	Data *POSTSkus201ResponseData `json:"data,omitempty"`
}

// NewPOSTSkus201Response instantiates a new POSTSkus201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkus201Response() *POSTSkus201Response {
	this := POSTSkus201Response{}
	return &this
}

// NewPOSTSkus201ResponseWithDefaults instantiates a new POSTSkus201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkus201ResponseWithDefaults() *POSTSkus201Response {
	this := POSTSkus201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTSkus201Response) GetData() POSTSkus201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTSkus201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkus201Response) GetDataOk() (*POSTSkus201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTSkus201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTSkus201ResponseData and assigns it to the Data field.
func (o *POSTSkus201Response) SetData(v POSTSkus201ResponseData) {
	o.Data = &v
}

func (o POSTSkus201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkus201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTSkus201Response struct {
	value *POSTSkus201Response
	isSet bool
}

func (v NullablePOSTSkus201Response) Get() *POSTSkus201Response {
	return v.value
}

func (v *NullablePOSTSkus201Response) Set(val *POSTSkus201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkus201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkus201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkus201Response(val *POSTSkus201Response) *NullablePOSTSkus201Response {
	return &NullablePOSTSkus201Response{value: val, isSet: true}
}

func (v NullablePOSTSkus201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkus201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
