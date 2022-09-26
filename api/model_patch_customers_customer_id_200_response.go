/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHCustomersCustomerId200Response struct for PATCHCustomersCustomerId200Response
type PATCHCustomersCustomerId200Response struct {
	Data *PATCHCustomersCustomerId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCustomersCustomerId200Response instantiates a new PATCHCustomersCustomerId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomersCustomerId200Response() *PATCHCustomersCustomerId200Response {
	this := PATCHCustomersCustomerId200Response{}
	return &this
}

// NewPATCHCustomersCustomerId200ResponseWithDefaults instantiates a new PATCHCustomersCustomerId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomersCustomerId200ResponseWithDefaults() *PATCHCustomersCustomerId200Response {
	this := PATCHCustomersCustomerId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200Response) GetData() PATCHCustomersCustomerId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHCustomersCustomerId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200Response) GetDataOk() (*PATCHCustomersCustomerId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCustomersCustomerId200ResponseData and assigns it to the Data field.
func (o *PATCHCustomersCustomerId200Response) SetData(v PATCHCustomersCustomerId200ResponseData) {
	o.Data = &v
}

func (o PATCHCustomersCustomerId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHCustomersCustomerId200Response struct {
	value *PATCHCustomersCustomerId200Response
	isSet bool
}

func (v NullablePATCHCustomersCustomerId200Response) Get() *PATCHCustomersCustomerId200Response {
	return v.value
}

func (v *NullablePATCHCustomersCustomerId200Response) Set(val *PATCHCustomersCustomerId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomersCustomerId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomersCustomerId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomersCustomerId200Response(val *PATCHCustomersCustomerId200Response) *NullablePATCHCustomersCustomerId200Response {
	return &NullablePATCHCustomersCustomerId200Response{value: val, isSet: true}
}

func (v NullablePATCHCustomersCustomerId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomersCustomerId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}