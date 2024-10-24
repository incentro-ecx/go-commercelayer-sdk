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

// checks if the POSTStockTransfers201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockTransfers201Response{}

// POSTStockTransfers201Response struct for POSTStockTransfers201Response
type POSTStockTransfers201Response struct {
	Data *POSTStockTransfers201ResponseData `json:"data,omitempty"`
}

// NewPOSTStockTransfers201Response instantiates a new POSTStockTransfers201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockTransfers201Response() *POSTStockTransfers201Response {
	this := POSTStockTransfers201Response{}
	return &this
}

// NewPOSTStockTransfers201ResponseWithDefaults instantiates a new POSTStockTransfers201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockTransfers201ResponseWithDefaults() *POSTStockTransfers201Response {
	this := POSTStockTransfers201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTStockTransfers201Response) GetData() POSTStockTransfers201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTStockTransfers201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201Response) GetDataOk() (*POSTStockTransfers201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTStockTransfers201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTStockTransfers201ResponseData and assigns it to the Data field.
func (o *POSTStockTransfers201Response) SetData(v POSTStockTransfers201ResponseData) {
	o.Data = &v
}

func (o POSTStockTransfers201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockTransfers201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTStockTransfers201Response struct {
	value *POSTStockTransfers201Response
	isSet bool
}

func (v NullablePOSTStockTransfers201Response) Get() *POSTStockTransfers201Response {
	return v.value
}

func (v *NullablePOSTStockTransfers201Response) Set(val *POSTStockTransfers201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockTransfers201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockTransfers201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockTransfers201Response(val *POSTStockTransfers201Response) *NullablePOSTStockTransfers201Response {
	return &NullablePOSTStockTransfers201Response{value: val, isSet: true}
}

func (v NullablePOSTStockTransfers201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockTransfers201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
