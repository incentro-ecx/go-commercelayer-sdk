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

// checks if the GETBundlesBundleId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETBundlesBundleId200Response{}

// GETBundlesBundleId200Response struct for GETBundlesBundleId200Response
type GETBundlesBundleId200Response struct {
	Data *GETBundlesBundleId200ResponseData `json:"data,omitempty"`
}

// NewGETBundlesBundleId200Response instantiates a new GETBundlesBundleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBundlesBundleId200Response() *GETBundlesBundleId200Response {
	this := GETBundlesBundleId200Response{}
	return &this
}

// NewGETBundlesBundleId200ResponseWithDefaults instantiates a new GETBundlesBundleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBundlesBundleId200ResponseWithDefaults() *GETBundlesBundleId200Response {
	this := GETBundlesBundleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBundlesBundleId200Response) GetData() GETBundlesBundleId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETBundlesBundleId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundlesBundleId200Response) GetDataOk() (*GETBundlesBundleId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBundlesBundleId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETBundlesBundleId200ResponseData and assigns it to the Data field.
func (o *GETBundlesBundleId200Response) SetData(v GETBundlesBundleId200ResponseData) {
	o.Data = &v
}

func (o GETBundlesBundleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETBundlesBundleId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETBundlesBundleId200Response struct {
	value *GETBundlesBundleId200Response
	isSet bool
}

func (v NullableGETBundlesBundleId200Response) Get() *GETBundlesBundleId200Response {
	return v.value
}

func (v *NullableGETBundlesBundleId200Response) Set(val *GETBundlesBundleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBundlesBundleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBundlesBundleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBundlesBundleId200Response(val *GETBundlesBundleId200Response) *NullableGETBundlesBundleId200Response {
	return &NullableGETBundlesBundleId200Response{value: val, isSet: true}
}

func (v NullableGETBundlesBundleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBundlesBundleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
