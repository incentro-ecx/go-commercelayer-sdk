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

// BundleResponse struct for BundleResponse
type BundleResponse struct {
	Data *BundleResponseData `json:"data,omitempty"`
}

// NewBundleResponse instantiates a new BundleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleResponse() *BundleResponse {
	this := BundleResponse{}
	return &this
}

// NewBundleResponseWithDefaults instantiates a new BundleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleResponseWithDefaults() *BundleResponse {
	this := BundleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BundleResponse) GetData() BundleResponseData {
	if o == nil || o.Data == nil {
		var ret BundleResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleResponse) GetDataOk() (*BundleResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BundleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BundleResponseData and assigns it to the Data field.
func (o *BundleResponse) SetData(v BundleResponseData) {
	o.Data = &v
}

func (o BundleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBundleResponse struct {
	value *BundleResponse
	isSet bool
}

func (v NullableBundleResponse) Get() *BundleResponse {
	return v.value
}

func (v *NullableBundleResponse) Set(val *BundleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleResponse(val *BundleResponse) *NullableBundleResponse {
	return &NullableBundleResponse{value: val, isSet: true}
}

func (v NullableBundleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}