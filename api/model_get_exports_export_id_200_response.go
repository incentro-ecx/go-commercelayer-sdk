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

// checks if the GETExportsExportId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExportsExportId200Response{}

// GETExportsExportId200Response struct for GETExportsExportId200Response
type GETExportsExportId200Response struct {
	Data *GETExportsExportId200ResponseData `json:"data,omitempty"`
}

// NewGETExportsExportId200Response instantiates a new GETExportsExportId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExportsExportId200Response() *GETExportsExportId200Response {
	this := GETExportsExportId200Response{}
	return &this
}

// NewGETExportsExportId200ResponseWithDefaults instantiates a new GETExportsExportId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExportsExportId200ResponseWithDefaults() *GETExportsExportId200Response {
	this := GETExportsExportId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExportsExportId200Response) GetData() GETExportsExportId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETExportsExportId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExportsExportId200Response) GetDataOk() (*GETExportsExportId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExportsExportId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExportsExportId200ResponseData and assigns it to the Data field.
func (o *GETExportsExportId200Response) SetData(v GETExportsExportId200ResponseData) {
	o.Data = &v
}

func (o GETExportsExportId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExportsExportId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETExportsExportId200Response struct {
	value *GETExportsExportId200Response
	isSet bool
}

func (v NullableGETExportsExportId200Response) Get() *GETExportsExportId200Response {
	return v.value
}

func (v *NullableGETExportsExportId200Response) Set(val *GETExportsExportId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExportsExportId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExportsExportId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExportsExportId200Response(val *GETExportsExportId200Response) *NullableGETExportsExportId200Response {
	return &NullableGETExportsExportId200Response{value: val, isSet: true}
}

func (v NullableGETExportsExportId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExportsExportId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
