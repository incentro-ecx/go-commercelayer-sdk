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

// POSTSkuListPromotionRules201Response struct for POSTSkuListPromotionRules201Response
type POSTSkuListPromotionRules201Response struct {
	Data *POSTSkuListPromotionRules201ResponseData `json:"data,omitempty"`
}

// NewPOSTSkuListPromotionRules201Response instantiates a new POSTSkuListPromotionRules201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuListPromotionRules201Response() *POSTSkuListPromotionRules201Response {
	this := POSTSkuListPromotionRules201Response{}
	return &this
}

// NewPOSTSkuListPromotionRules201ResponseWithDefaults instantiates a new POSTSkuListPromotionRules201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuListPromotionRules201ResponseWithDefaults() *POSTSkuListPromotionRules201Response {
	this := POSTSkuListPromotionRules201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTSkuListPromotionRules201Response) GetData() POSTSkuListPromotionRules201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTSkuListPromotionRules201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201Response) GetDataOk() (*POSTSkuListPromotionRules201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTSkuListPromotionRules201ResponseData and assigns it to the Data field.
func (o *POSTSkuListPromotionRules201Response) SetData(v POSTSkuListPromotionRules201ResponseData) {
	o.Data = &v
}

func (o POSTSkuListPromotionRules201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTSkuListPromotionRules201Response struct {
	value *POSTSkuListPromotionRules201Response
	isSet bool
}

func (v NullablePOSTSkuListPromotionRules201Response) Get() *POSTSkuListPromotionRules201Response {
	return v.value
}

func (v *NullablePOSTSkuListPromotionRules201Response) Set(val *POSTSkuListPromotionRules201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuListPromotionRules201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuListPromotionRules201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuListPromotionRules201Response(val *POSTSkuListPromotionRules201Response) *NullablePOSTSkuListPromotionRules201Response {
	return &NullablePOSTSkuListPromotionRules201Response{value: val, isSet: true}
}

func (v NullablePOSTSkuListPromotionRules201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuListPromotionRules201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}