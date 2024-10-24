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

// checks if the PATCHVoidsVoidId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHVoidsVoidId200Response{}

// PATCHVoidsVoidId200Response struct for PATCHVoidsVoidId200Response
type PATCHVoidsVoidId200Response struct {
	Data *PATCHVoidsVoidId200ResponseData `json:"data,omitempty"`
}

// NewPATCHVoidsVoidId200Response instantiates a new PATCHVoidsVoidId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHVoidsVoidId200Response() *PATCHVoidsVoidId200Response {
	this := PATCHVoidsVoidId200Response{}
	return &this
}

// NewPATCHVoidsVoidId200ResponseWithDefaults instantiates a new PATCHVoidsVoidId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHVoidsVoidId200ResponseWithDefaults() *PATCHVoidsVoidId200Response {
	this := PATCHVoidsVoidId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHVoidsVoidId200Response) GetData() PATCHVoidsVoidId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHVoidsVoidId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHVoidsVoidId200Response) GetDataOk() (*PATCHVoidsVoidId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHVoidsVoidId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHVoidsVoidId200ResponseData and assigns it to the Data field.
func (o *PATCHVoidsVoidId200Response) SetData(v PATCHVoidsVoidId200ResponseData) {
	o.Data = &v
}

func (o PATCHVoidsVoidId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHVoidsVoidId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHVoidsVoidId200Response struct {
	value *PATCHVoidsVoidId200Response
	isSet bool
}

func (v NullablePATCHVoidsVoidId200Response) Get() *PATCHVoidsVoidId200Response {
	return v.value
}

func (v *NullablePATCHVoidsVoidId200Response) Set(val *PATCHVoidsVoidId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHVoidsVoidId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHVoidsVoidId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHVoidsVoidId200Response(val *PATCHVoidsVoidId200Response) *NullablePATCHVoidsVoidId200Response {
	return &NullablePATCHVoidsVoidId200Response{value: val, isSet: true}
}

func (v NullablePATCHVoidsVoidId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHVoidsVoidId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
