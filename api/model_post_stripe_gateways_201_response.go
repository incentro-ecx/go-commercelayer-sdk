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

// checks if the POSTStripeGateways201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStripeGateways201Response{}

// POSTStripeGateways201Response struct for POSTStripeGateways201Response
type POSTStripeGateways201Response struct {
	Data *POSTStripeGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTStripeGateways201Response instantiates a new POSTStripeGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStripeGateways201Response() *POSTStripeGateways201Response {
	this := POSTStripeGateways201Response{}
	return &this
}

// NewPOSTStripeGateways201ResponseWithDefaults instantiates a new POSTStripeGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStripeGateways201ResponseWithDefaults() *POSTStripeGateways201Response {
	this := POSTStripeGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTStripeGateways201Response) GetData() POSTStripeGateways201ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret POSTStripeGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201Response) GetDataOk() (*POSTStripeGateways201ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTStripeGateways201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTStripeGateways201ResponseData and assigns it to the Data field.
func (o *POSTStripeGateways201Response) SetData(v POSTStripeGateways201ResponseData) {
	o.Data = &v
}

func (o POSTStripeGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStripeGateways201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTStripeGateways201Response struct {
	value *POSTStripeGateways201Response
	isSet bool
}

func (v NullablePOSTStripeGateways201Response) Get() *POSTStripeGateways201Response {
	return v.value
}

func (v *NullablePOSTStripeGateways201Response) Set(val *POSTStripeGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStripeGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStripeGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStripeGateways201Response(val *POSTStripeGateways201Response) *NullablePOSTStripeGateways201Response {
	return &NullablePOSTStripeGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTStripeGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStripeGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
