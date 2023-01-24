/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETFreeGiftPromotions200Response struct for GETFreeGiftPromotions200Response
type GETFreeGiftPromotions200Response struct {
	Data []GETFreeGiftPromotions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETFreeGiftPromotions200Response instantiates a new GETFreeGiftPromotions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFreeGiftPromotions200Response() *GETFreeGiftPromotions200Response {
	this := GETFreeGiftPromotions200Response{}
	return &this
}

// NewGETFreeGiftPromotions200ResponseWithDefaults instantiates a new GETFreeGiftPromotions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFreeGiftPromotions200ResponseWithDefaults() *GETFreeGiftPromotions200Response {
	this := GETFreeGiftPromotions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETFreeGiftPromotions200Response) GetData() []GETFreeGiftPromotions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETFreeGiftPromotions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotions200Response) GetDataOk() ([]GETFreeGiftPromotions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETFreeGiftPromotions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETFreeGiftPromotions200ResponseDataInner and assigns it to the Data field.
func (o *GETFreeGiftPromotions200Response) SetData(v []GETFreeGiftPromotions200ResponseDataInner) {
	o.Data = v
}

func (o GETFreeGiftPromotions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETFreeGiftPromotions200Response struct {
	value *GETFreeGiftPromotions200Response
	isSet bool
}

func (v NullableGETFreeGiftPromotions200Response) Get() *GETFreeGiftPromotions200Response {
	return v.value
}

func (v *NullableGETFreeGiftPromotions200Response) Set(val *GETFreeGiftPromotions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFreeGiftPromotions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFreeGiftPromotions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFreeGiftPromotions200Response(val *GETFreeGiftPromotions200Response) *NullableGETFreeGiftPromotions200Response {
	return &NullableGETFreeGiftPromotions200Response{value: val, isSet: true}
}

func (v NullableGETFreeGiftPromotions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFreeGiftPromotions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
