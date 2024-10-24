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

// checks if the GETInventoryModels200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETInventoryModels200Response{}

// GETInventoryModels200Response struct for GETInventoryModels200Response
type GETInventoryModels200Response struct {
	Data interface{} `json:"data,omitempty"`
}

// NewGETInventoryModels200Response instantiates a new GETInventoryModels200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModels200Response() *GETInventoryModels200Response {
	this := GETInventoryModels200Response{}
	return &this
}

// NewGETInventoryModels200ResponseWithDefaults instantiates a new GETInventoryModels200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModels200ResponseWithDefaults() *GETInventoryModels200Response {
	this := GETInventoryModels200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModels200Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModels200Response) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryModels200Response) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *GETInventoryModels200Response) SetData(v interface{}) {
	o.Data = v
}

func (o GETInventoryModels200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETInventoryModels200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETInventoryModels200Response struct {
	value *GETInventoryModels200Response
	isSet bool
}

func (v NullableGETInventoryModels200Response) Get() *GETInventoryModels200Response {
	return v.value
}

func (v *NullableGETInventoryModels200Response) Set(val *GETInventoryModels200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModels200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModels200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModels200Response(val *GETInventoryModels200Response) *NullableGETInventoryModels200Response {
	return &NullableGETInventoryModels200Response{value: val, isSet: true}
}

func (v NullableGETInventoryModels200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModels200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
