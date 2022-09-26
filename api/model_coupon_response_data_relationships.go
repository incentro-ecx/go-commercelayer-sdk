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

// CouponResponseDataRelationships struct for CouponResponseDataRelationships
type CouponResponseDataRelationships struct {
	PromotionRule *CouponResponseDataRelationshipsPromotionRule `json:"promotion_rule,omitempty"`
}

// NewCouponResponseDataRelationships instantiates a new CouponResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponResponseDataRelationships() *CouponResponseDataRelationships {
	this := CouponResponseDataRelationships{}
	return &this
}

// NewCouponResponseDataRelationshipsWithDefaults instantiates a new CouponResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponResponseDataRelationshipsWithDefaults() *CouponResponseDataRelationships {
	this := CouponResponseDataRelationships{}
	return &this
}

// GetPromotionRule returns the PromotionRule field value if set, zero value otherwise.
func (o *CouponResponseDataRelationships) GetPromotionRule() CouponResponseDataRelationshipsPromotionRule {
	if o == nil || o.PromotionRule == nil {
		var ret CouponResponseDataRelationshipsPromotionRule
		return ret
	}
	return *o.PromotionRule
}

// GetPromotionRuleOk returns a tuple with the PromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponResponseDataRelationships) GetPromotionRuleOk() (*CouponResponseDataRelationshipsPromotionRule, bool) {
	if o == nil || o.PromotionRule == nil {
		return nil, false
	}
	return o.PromotionRule, true
}

// HasPromotionRule returns a boolean if a field has been set.
func (o *CouponResponseDataRelationships) HasPromotionRule() bool {
	if o != nil && o.PromotionRule != nil {
		return true
	}

	return false
}

// SetPromotionRule gets a reference to the given CouponResponseDataRelationshipsPromotionRule and assigns it to the PromotionRule field.
func (o *CouponResponseDataRelationships) SetPromotionRule(v CouponResponseDataRelationshipsPromotionRule) {
	o.PromotionRule = &v
}

func (o CouponResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PromotionRule != nil {
		toSerialize["promotion_rule"] = o.PromotionRule
	}
	return json.Marshal(toSerialize)
}

type NullableCouponResponseDataRelationships struct {
	value *CouponResponseDataRelationships
	isSet bool
}

func (v NullableCouponResponseDataRelationships) Get() *CouponResponseDataRelationships {
	return v.value
}

func (v *NullableCouponResponseDataRelationships) Set(val *CouponResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponResponseDataRelationships(val *CouponResponseDataRelationships) *NullableCouponResponseDataRelationships {
	return &NullableCouponResponseDataRelationships{value: val, isSet: true}
}

func (v NullableCouponResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}