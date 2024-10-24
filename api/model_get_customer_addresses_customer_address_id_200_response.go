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

// checks if the GETCustomerAddressesCustomerAddressId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCustomerAddressesCustomerAddressId200Response{}

// GETCustomerAddressesCustomerAddressId200Response struct for GETCustomerAddressesCustomerAddressId200Response
type GETCustomerAddressesCustomerAddressId200Response struct {
	Data *GETCustomerAddressesCustomerAddressId200ResponseData `json:"data,omitempty"`
}

// NewGETCustomerAddressesCustomerAddressId200Response instantiates a new GETCustomerAddressesCustomerAddressId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerAddressesCustomerAddressId200Response() *GETCustomerAddressesCustomerAddressId200Response {
	this := GETCustomerAddressesCustomerAddressId200Response{}
	return &this
}

// NewGETCustomerAddressesCustomerAddressId200ResponseWithDefaults instantiates a new GETCustomerAddressesCustomerAddressId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerAddressesCustomerAddressId200ResponseWithDefaults() *GETCustomerAddressesCustomerAddressId200Response {
	this := GETCustomerAddressesCustomerAddressId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerAddressesCustomerAddressId200Response) GetData() GETCustomerAddressesCustomerAddressId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCustomerAddressesCustomerAddressId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddressesCustomerAddressId200Response) GetDataOk() (*GETCustomerAddressesCustomerAddressId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerAddressesCustomerAddressId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomerAddressesCustomerAddressId200ResponseData and assigns it to the Data field.
func (o *GETCustomerAddressesCustomerAddressId200Response) SetData(v GETCustomerAddressesCustomerAddressId200ResponseData) {
	o.Data = &v
}

func (o GETCustomerAddressesCustomerAddressId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCustomerAddressesCustomerAddressId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCustomerAddressesCustomerAddressId200Response struct {
	value *GETCustomerAddressesCustomerAddressId200Response
	isSet bool
}

func (v NullableGETCustomerAddressesCustomerAddressId200Response) Get() *GETCustomerAddressesCustomerAddressId200Response {
	return v.value
}

func (v *NullableGETCustomerAddressesCustomerAddressId200Response) Set(val *GETCustomerAddressesCustomerAddressId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerAddressesCustomerAddressId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerAddressesCustomerAddressId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerAddressesCustomerAddressId200Response(val *GETCustomerAddressesCustomerAddressId200Response) *NullableGETCustomerAddressesCustomerAddressId200Response {
	return &NullableGETCustomerAddressesCustomerAddressId200Response{value: val, isSet: true}
}

func (v NullableGETCustomerAddressesCustomerAddressId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerAddressesCustomerAddressId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
