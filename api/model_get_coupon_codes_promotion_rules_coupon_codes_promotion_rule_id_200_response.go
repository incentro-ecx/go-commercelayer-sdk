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

// checks if the GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{}

// GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response struct for GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response
type GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response struct {
	Data *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseData `json:"data,omitempty"`
}

// NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response instantiates a new GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response() *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	this := GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{}
	return &this
}

// NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseWithDefaults instantiates a new GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseWithDefaults() *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	this := GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) GetData() GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) GetDataOk() (*GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseData and assigns it to the Data field.
func (o *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) SetData(v GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseData) {
	o.Data = &v
}

func (o GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response struct {
	value *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response
	isSet bool
}

func (v NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) Get() *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	return v.value
}

func (v *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) Set(val *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response(val *GETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response {
	return &NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response{value: val, isSet: true}
}

func (v NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponCodesPromotionRulesCouponCodesPromotionRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
