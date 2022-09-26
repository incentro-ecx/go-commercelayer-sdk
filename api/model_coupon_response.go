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

// CouponResponse struct for CouponResponse
type CouponResponse struct {
	Data *CouponResponseData `json:"data,omitempty"`
}

// NewCouponResponse instantiates a new CouponResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponResponse() *CouponResponse {
	this := CouponResponse{}
	return &this
}

// NewCouponResponseWithDefaults instantiates a new CouponResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponResponseWithDefaults() *CouponResponse {
	this := CouponResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CouponResponse) GetData() CouponResponseData {
	if o == nil || o.Data == nil {
		var ret CouponResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponResponse) GetDataOk() (*CouponResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CouponResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CouponResponseData and assigns it to the Data field.
func (o *CouponResponse) SetData(v CouponResponseData) {
	o.Data = &v
}

func (o CouponResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCouponResponse struct {
	value *CouponResponse
	isSet bool
}

func (v NullableCouponResponse) Get() *CouponResponse {
	return v.value
}

func (v *NullableCouponResponse) Set(val *CouponResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponResponse(val *CouponResponse) *NullableCouponResponse {
	return &NullableCouponResponse{value: val, isSet: true}
}

func (v NullableCouponResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}