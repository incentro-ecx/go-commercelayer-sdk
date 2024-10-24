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

// checks if the PATCHTaxjarAccountsTaxjarAccountId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHTaxjarAccountsTaxjarAccountId200Response{}

// PATCHTaxjarAccountsTaxjarAccountId200Response struct for PATCHTaxjarAccountsTaxjarAccountId200Response
type PATCHTaxjarAccountsTaxjarAccountId200Response struct {
	Data *PATCHTaxjarAccountsTaxjarAccountId200ResponseData `json:"data,omitempty"`
}

// NewPATCHTaxjarAccountsTaxjarAccountId200Response instantiates a new PATCHTaxjarAccountsTaxjarAccountId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHTaxjarAccountsTaxjarAccountId200Response() *PATCHTaxjarAccountsTaxjarAccountId200Response {
	this := PATCHTaxjarAccountsTaxjarAccountId200Response{}
	return &this
}

// NewPATCHTaxjarAccountsTaxjarAccountId200ResponseWithDefaults instantiates a new PATCHTaxjarAccountsTaxjarAccountId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHTaxjarAccountsTaxjarAccountId200ResponseWithDefaults() *PATCHTaxjarAccountsTaxjarAccountId200Response {
	this := PATCHTaxjarAccountsTaxjarAccountId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHTaxjarAccountsTaxjarAccountId200Response) GetData() PATCHTaxjarAccountsTaxjarAccountId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHTaxjarAccountsTaxjarAccountId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHTaxjarAccountsTaxjarAccountId200Response) GetDataOk() (*PATCHTaxjarAccountsTaxjarAccountId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHTaxjarAccountsTaxjarAccountId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHTaxjarAccountsTaxjarAccountId200ResponseData and assigns it to the Data field.
func (o *PATCHTaxjarAccountsTaxjarAccountId200Response) SetData(v PATCHTaxjarAccountsTaxjarAccountId200ResponseData) {
	o.Data = &v
}

func (o PATCHTaxjarAccountsTaxjarAccountId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHTaxjarAccountsTaxjarAccountId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHTaxjarAccountsTaxjarAccountId200Response struct {
	value *PATCHTaxjarAccountsTaxjarAccountId200Response
	isSet bool
}

func (v NullablePATCHTaxjarAccountsTaxjarAccountId200Response) Get() *PATCHTaxjarAccountsTaxjarAccountId200Response {
	return v.value
}

func (v *NullablePATCHTaxjarAccountsTaxjarAccountId200Response) Set(val *PATCHTaxjarAccountsTaxjarAccountId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHTaxjarAccountsTaxjarAccountId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHTaxjarAccountsTaxjarAccountId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHTaxjarAccountsTaxjarAccountId200Response(val *PATCHTaxjarAccountsTaxjarAccountId200Response) *NullablePATCHTaxjarAccountsTaxjarAccountId200Response {
	return &NullablePATCHTaxjarAccountsTaxjarAccountId200Response{value: val, isSet: true}
}

func (v NullablePATCHTaxjarAccountsTaxjarAccountId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHTaxjarAccountsTaxjarAccountId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
