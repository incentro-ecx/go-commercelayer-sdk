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

// checks if the SkuListPromotionRuleUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListPromotionRuleUpdateDataRelationships{}

// SkuListPromotionRuleUpdateDataRelationships struct for SkuListPromotionRuleUpdateDataRelationships
type SkuListPromotionRuleUpdateDataRelationships struct {
	Promotion *CouponCodesPromotionRuleCreateDataRelationshipsPromotion `json:"promotion,omitempty"`
	SkuList   *BundleCreateDataRelationshipsSkuList                     `json:"sku_list,omitempty"`
}

// NewSkuListPromotionRuleUpdateDataRelationships instantiates a new SkuListPromotionRuleUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleUpdateDataRelationships() *SkuListPromotionRuleUpdateDataRelationships {
	this := SkuListPromotionRuleUpdateDataRelationships{}
	return &this
}

// NewSkuListPromotionRuleUpdateDataRelationshipsWithDefaults instantiates a new SkuListPromotionRuleUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleUpdateDataRelationshipsWithDefaults() *SkuListPromotionRuleUpdateDataRelationships {
	this := SkuListPromotionRuleUpdateDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *SkuListPromotionRuleUpdateDataRelationships) GetPromotion() CouponCodesPromotionRuleCreateDataRelationshipsPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret CouponCodesPromotionRuleCreateDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdateDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleCreateDataRelationshipsPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *SkuListPromotionRuleUpdateDataRelationships) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given CouponCodesPromotionRuleCreateDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *SkuListPromotionRuleUpdateDataRelationships) SetPromotion(v CouponCodesPromotionRuleCreateDataRelationshipsPromotion) {
	o.Promotion = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *SkuListPromotionRuleUpdateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret BundleCreateDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *SkuListPromotionRuleUpdateDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleCreateDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *SkuListPromotionRuleUpdateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList) {
	o.SkuList = &v
}

func (o SkuListPromotionRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListPromotionRuleUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	return toSerialize, nil
}

type NullableSkuListPromotionRuleUpdateDataRelationships struct {
	value *SkuListPromotionRuleUpdateDataRelationships
	isSet bool
}

func (v NullableSkuListPromotionRuleUpdateDataRelationships) Get() *SkuListPromotionRuleUpdateDataRelationships {
	return v.value
}

func (v *NullableSkuListPromotionRuleUpdateDataRelationships) Set(val *SkuListPromotionRuleUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleUpdateDataRelationships(val *SkuListPromotionRuleUpdateDataRelationships) *NullableSkuListPromotionRuleUpdateDataRelationships {
	return &NullableSkuListPromotionRuleUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
