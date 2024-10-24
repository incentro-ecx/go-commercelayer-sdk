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

// checks if the GETApplicationApplicationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETApplicationApplicationId200Response{}

// GETApplicationApplicationId200Response struct for GETApplicationApplicationId200Response
type GETApplicationApplicationId200Response struct {
	Data *GETApplicationApplicationId200ResponseData `json:"data,omitempty"`
}

// NewGETApplicationApplicationId200Response instantiates a new GETApplicationApplicationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETApplicationApplicationId200Response() *GETApplicationApplicationId200Response {
	this := GETApplicationApplicationId200Response{}
	return &this
}

// NewGETApplicationApplicationId200ResponseWithDefaults instantiates a new GETApplicationApplicationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETApplicationApplicationId200ResponseWithDefaults() *GETApplicationApplicationId200Response {
	this := GETApplicationApplicationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200Response) GetData() GETApplicationApplicationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETApplicationApplicationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200Response) GetDataOk() (*GETApplicationApplicationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETApplicationApplicationId200ResponseData and assigns it to the Data field.
func (o *GETApplicationApplicationId200Response) SetData(v GETApplicationApplicationId200ResponseData) {
	o.Data = &v
}

func (o GETApplicationApplicationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETApplicationApplicationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETApplicationApplicationId200Response struct {
	value *GETApplicationApplicationId200Response
	isSet bool
}

func (v NullableGETApplicationApplicationId200Response) Get() *GETApplicationApplicationId200Response {
	return v.value
}

func (v *NullableGETApplicationApplicationId200Response) Set(val *GETApplicationApplicationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETApplicationApplicationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETApplicationApplicationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETApplicationApplicationId200Response(val *GETApplicationApplicationId200Response) *NullableGETApplicationApplicationId200Response {
	return &NullableGETApplicationApplicationId200Response{value: val, isSet: true}
}

func (v NullableGETApplicationApplicationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETApplicationApplicationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
