/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETLineItemOptionsLineItemOptionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETLineItemOptionsLineItemOptionId200Response{}

// GETLineItemOptionsLineItemOptionId200Response struct for GETLineItemOptionsLineItemOptionId200Response
type GETLineItemOptionsLineItemOptionId200Response struct {
	Data *GETLineItemOptionsLineItemOptionId200ResponseData `json:"data,omitempty"`
}

// NewGETLineItemOptionsLineItemOptionId200Response instantiates a new GETLineItemOptionsLineItemOptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptionsLineItemOptionId200Response() *GETLineItemOptionsLineItemOptionId200Response {
	this := GETLineItemOptionsLineItemOptionId200Response{}
	return &this
}

// NewGETLineItemOptionsLineItemOptionId200ResponseWithDefaults instantiates a new GETLineItemOptionsLineItemOptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptionsLineItemOptionId200ResponseWithDefaults() *GETLineItemOptionsLineItemOptionId200Response {
	this := GETLineItemOptionsLineItemOptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItemOptionsLineItemOptionId200Response) GetData() GETLineItemOptionsLineItemOptionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETLineItemOptionsLineItemOptionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptionsLineItemOptionId200Response) GetDataOk() (*GETLineItemOptionsLineItemOptionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItemOptionsLineItemOptionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItemOptionsLineItemOptionId200ResponseData and assigns it to the Data field.
func (o *GETLineItemOptionsLineItemOptionId200Response) SetData(v GETLineItemOptionsLineItemOptionId200ResponseData) {
	o.Data = &v
}

func (o GETLineItemOptionsLineItemOptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETLineItemOptionsLineItemOptionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETLineItemOptionsLineItemOptionId200Response struct {
	value *GETLineItemOptionsLineItemOptionId200Response
	isSet bool
}

func (v NullableGETLineItemOptionsLineItemOptionId200Response) Get() *GETLineItemOptionsLineItemOptionId200Response {
	return v.value
}

func (v *NullableGETLineItemOptionsLineItemOptionId200Response) Set(val *GETLineItemOptionsLineItemOptionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptionsLineItemOptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptionsLineItemOptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptionsLineItemOptionId200Response(val *GETLineItemOptionsLineItemOptionId200Response) *NullableGETLineItemOptionsLineItemOptionId200Response {
	return &NullableGETLineItemOptionsLineItemOptionId200Response{value: val, isSet: true}
}

func (v NullableGETLineItemOptionsLineItemOptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptionsLineItemOptionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
