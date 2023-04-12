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

// checks if the PATCHAuthorizationsAuthorizationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAuthorizationsAuthorizationId200Response{}

// PATCHAuthorizationsAuthorizationId200Response struct for PATCHAuthorizationsAuthorizationId200Response
type PATCHAuthorizationsAuthorizationId200Response struct {
	Data *PATCHAuthorizationsAuthorizationId200ResponseData `json:"data,omitempty"`
}

// NewPATCHAuthorizationsAuthorizationId200Response instantiates a new PATCHAuthorizationsAuthorizationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAuthorizationsAuthorizationId200Response() *PATCHAuthorizationsAuthorizationId200Response {
	this := PATCHAuthorizationsAuthorizationId200Response{}
	return &this
}

// NewPATCHAuthorizationsAuthorizationId200ResponseWithDefaults instantiates a new PATCHAuthorizationsAuthorizationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAuthorizationsAuthorizationId200ResponseWithDefaults() *PATCHAuthorizationsAuthorizationId200Response {
	this := PATCHAuthorizationsAuthorizationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHAuthorizationsAuthorizationId200Response) GetData() PATCHAuthorizationsAuthorizationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHAuthorizationsAuthorizationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAuthorizationsAuthorizationId200Response) GetDataOk() (*PATCHAuthorizationsAuthorizationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHAuthorizationsAuthorizationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHAuthorizationsAuthorizationId200ResponseData and assigns it to the Data field.
func (o *PATCHAuthorizationsAuthorizationId200Response) SetData(v PATCHAuthorizationsAuthorizationId200ResponseData) {
	o.Data = &v
}

func (o PATCHAuthorizationsAuthorizationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAuthorizationsAuthorizationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHAuthorizationsAuthorizationId200Response struct {
	value *PATCHAuthorizationsAuthorizationId200Response
	isSet bool
}

func (v NullablePATCHAuthorizationsAuthorizationId200Response) Get() *PATCHAuthorizationsAuthorizationId200Response {
	return v.value
}

func (v *NullablePATCHAuthorizationsAuthorizationId200Response) Set(val *PATCHAuthorizationsAuthorizationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAuthorizationsAuthorizationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAuthorizationsAuthorizationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAuthorizationsAuthorizationId200Response(val *PATCHAuthorizationsAuthorizationId200Response) *NullablePATCHAuthorizationsAuthorizationId200Response {
	return &NullablePATCHAuthorizationsAuthorizationId200Response{value: val, isSet: true}
}

func (v NullablePATCHAuthorizationsAuthorizationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAuthorizationsAuthorizationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
