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

// checks if the POSTTaxCategories201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTTaxCategories201Response{}

// POSTTaxCategories201Response struct for POSTTaxCategories201Response
type POSTTaxCategories201Response struct {
	Data *POSTTaxCategories201ResponseData `json:"data,omitempty"`
}

// NewPOSTTaxCategories201Response instantiates a new POSTTaxCategories201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxCategories201Response() *POSTTaxCategories201Response {
	this := POSTTaxCategories201Response{}
	return &this
}

// NewPOSTTaxCategories201ResponseWithDefaults instantiates a new POSTTaxCategories201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxCategories201ResponseWithDefaults() *POSTTaxCategories201Response {
	this := POSTTaxCategories201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTTaxCategories201Response) GetData() POSTTaxCategories201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTTaxCategories201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxCategories201Response) GetDataOk() (*POSTTaxCategories201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTTaxCategories201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTTaxCategories201ResponseData and assigns it to the Data field.
func (o *POSTTaxCategories201Response) SetData(v POSTTaxCategories201ResponseData) {
	o.Data = &v
}

func (o POSTTaxCategories201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTTaxCategories201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTTaxCategories201Response struct {
	value *POSTTaxCategories201Response
	isSet bool
}

func (v NullablePOSTTaxCategories201Response) Get() *POSTTaxCategories201Response {
	return v.value
}

func (v *NullablePOSTTaxCategories201Response) Set(val *POSTTaxCategories201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxCategories201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxCategories201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxCategories201Response(val *POSTTaxCategories201Response) *NullablePOSTTaxCategories201Response {
	return &NullablePOSTTaxCategories201Response{value: val, isSet: true}
}

func (v NullablePOSTTaxCategories201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxCategories201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
