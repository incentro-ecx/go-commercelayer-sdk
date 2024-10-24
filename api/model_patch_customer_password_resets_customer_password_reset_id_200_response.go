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

// checks if the PATCHCustomerPasswordResetsCustomerPasswordResetId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerPasswordResetsCustomerPasswordResetId200Response{}

// PATCHCustomerPasswordResetsCustomerPasswordResetId200Response struct for PATCHCustomerPasswordResetsCustomerPasswordResetId200Response
type PATCHCustomerPasswordResetsCustomerPasswordResetId200Response struct {
	Data *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCustomerPasswordResetsCustomerPasswordResetId200Response instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200Response() *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response {
	this := PATCHCustomerPasswordResetsCustomerPasswordResetId200Response{}
	return &this
}

// NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseWithDefaults instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseWithDefaults() *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response {
	this := PATCHCustomerPasswordResetsCustomerPasswordResetId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) GetData() PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) GetDataOk() (*PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseData and assigns it to the Data field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) SetData(v PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseData) {
	o.Data = &v
}

func (o PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response struct {
	value *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response
	isSet bool
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response) Get() *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response {
	return v.value
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response) Set(val *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response(val *PATCHCustomerPasswordResetsCustomerPasswordResetId200Response) *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response {
	return &NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response{value: val, isSet: true}
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
