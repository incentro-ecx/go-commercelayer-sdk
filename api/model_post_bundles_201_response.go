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

// checks if the POSTBundles201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBundles201Response{}

// POSTBundles201Response struct for POSTBundles201Response
type POSTBundles201Response struct {
	Data *POSTBundles201ResponseData `json:"data,omitempty"`
}

// NewPOSTBundles201Response instantiates a new POSTBundles201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBundles201Response() *POSTBundles201Response {
	this := POSTBundles201Response{}
	return &this
}

// NewPOSTBundles201ResponseWithDefaults instantiates a new POSTBundles201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBundles201ResponseWithDefaults() *POSTBundles201Response {
	this := POSTBundles201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBundles201Response) GetData() POSTBundles201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBundles201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201Response) GetDataOk() (*POSTBundles201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBundles201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBundles201ResponseData and assigns it to the Data field.
func (o *POSTBundles201Response) SetData(v POSTBundles201ResponseData) {
	o.Data = &v
}

func (o POSTBundles201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBundles201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBundles201Response struct {
	value *POSTBundles201Response
	isSet bool
}

func (v NullablePOSTBundles201Response) Get() *POSTBundles201Response {
	return v.value
}

func (v *NullablePOSTBundles201Response) Set(val *POSTBundles201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBundles201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBundles201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBundles201Response(val *POSTBundles201Response) *NullablePOSTBundles201Response {
	return &NullablePOSTBundles201Response{value: val, isSet: true}
}

func (v NullablePOSTBundles201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBundles201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
