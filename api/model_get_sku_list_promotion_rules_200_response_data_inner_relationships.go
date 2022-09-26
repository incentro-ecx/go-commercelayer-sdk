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

// GETSkuListPromotionRules200ResponseDataInnerRelationships struct for GETSkuListPromotionRules200ResponseDataInnerRelationships
type GETSkuListPromotionRules200ResponseDataInnerRelationships struct {
	Promotion *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion `json:"promotion,omitempty"`
	SkuList   *GETBundles200ResponseDataInnerRelationshipsSkuList                     `json:"sku_list,omitempty"`
	Skus      *GETBundles200ResponseDataInnerRelationshipsSkus                        `json:"skus,omitempty"`
}

// NewGETSkuListPromotionRules200ResponseDataInnerRelationships instantiates a new GETSkuListPromotionRules200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuListPromotionRules200ResponseDataInnerRelationships() *GETSkuListPromotionRules200ResponseDataInnerRelationships {
	this := GETSkuListPromotionRules200ResponseDataInnerRelationships{}
	return &this
}

// NewGETSkuListPromotionRules200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETSkuListPromotionRules200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuListPromotionRules200ResponseDataInnerRelationshipsWithDefaults() *GETSkuListPromotionRules200ResponseDataInnerRelationships {
	this := GETSkuListPromotionRules200ResponseDataInnerRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) GetPromotion() GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	if o == nil || o.Promotion == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) GetPromotionOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion, bool) {
	if o == nil || o.Promotion == nil {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) HasPromotion() bool {
	if o != nil && o.Promotion != nil {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion and assigns it to the Promotion field.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) SetPromotion(v GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) {
	o.Promotion = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) GetSkuList() GETBundles200ResponseDataInnerRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) GetSkuListOk() (*GETBundles200ResponseDataInnerRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkuList and assigns it to the SkuList field.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) SetSkuList(v GETBundles200ResponseDataInnerRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) GetSkus() GETBundles200ResponseDataInnerRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) GetSkusOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkus and assigns it to the Skus field.
func (o *GETSkuListPromotionRules200ResponseDataInnerRelationships) SetSkus(v GETBundles200ResponseDataInnerRelationshipsSkus) {
	o.Skus = &v
}

func (o GETSkuListPromotionRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Promotion != nil {
		toSerialize["promotion"] = o.Promotion
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuListPromotionRules200ResponseDataInnerRelationships struct {
	value *GETSkuListPromotionRules200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETSkuListPromotionRules200ResponseDataInnerRelationships) Get() *GETSkuListPromotionRules200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETSkuListPromotionRules200ResponseDataInnerRelationships) Set(val *GETSkuListPromotionRules200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuListPromotionRules200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuListPromotionRules200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuListPromotionRules200ResponseDataInnerRelationships(val *GETSkuListPromotionRules200ResponseDataInnerRelationships) *NullableGETSkuListPromotionRules200ResponseDataInnerRelationships {
	return &NullableGETSkuListPromotionRules200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETSkuListPromotionRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuListPromotionRules200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}