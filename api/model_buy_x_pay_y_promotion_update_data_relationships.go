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

// checks if the BuyXPayYPromotionUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyXPayYPromotionUpdateDataRelationships{}

// BuyXPayYPromotionUpdateDataRelationships struct for BuyXPayYPromotionUpdateDataRelationships
type BuyXPayYPromotionUpdateDataRelationships struct {
	Market                   *BillingInfoValidationRuleCreateDataRelationshipsMarket           `json:"market,omitempty"`
	OrderAmountPromotionRule *BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule     *BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule     `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	CustomPromotionRule      *BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule      `json:"custom_promotion_rule,omitempty"`
	SkuList                  *BundleCreateDataRelationshipsSkuList                             `json:"sku_list,omitempty"`
	Tags                     *AddressCreateDataRelationshipsTags                               `json:"tags,omitempty"`
}

// NewBuyXPayYPromotionUpdateDataRelationships instantiates a new BuyXPayYPromotionUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyXPayYPromotionUpdateDataRelationships() *BuyXPayYPromotionUpdateDataRelationships {
	this := BuyXPayYPromotionUpdateDataRelationships{}
	return &this
}

// NewBuyXPayYPromotionUpdateDataRelationshipsWithDefaults instantiates a new BuyXPayYPromotionUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyXPayYPromotionUpdateDataRelationshipsWithDefaults() *BuyXPayYPromotionUpdateDataRelationships {
	this := BuyXPayYPromotionUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetOrderAmountPromotionRule() BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule {
	if o == nil || IsNil(o.OrderAmountPromotionRule) {
		var ret BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetOrderAmountPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || IsNil(o.OrderAmountPromotionRule) {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && !IsNil(o.OrderAmountPromotionRule) {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetOrderAmountPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetSkuListPromotionRule() BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule {
	if o == nil || IsNil(o.SkuListPromotionRule) {
		var ret BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetSkuListPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule, bool) {
	if o == nil || IsNil(o.SkuListPromotionRule) {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasSkuListPromotionRule() bool {
	if o != nil && !IsNil(o.SkuListPromotionRule) {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetSkuListPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetCouponCodesPromotionRule() BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		var ret BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetCouponCodesPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || IsNil(o.CouponCodesPromotionRule) {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && !IsNil(o.CouponCodesPromotionRule) {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetCouponCodesPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetCustomPromotionRule returns the CustomPromotionRule field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetCustomPromotionRule() BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule {
	if o == nil || IsNil(o.CustomPromotionRule) {
		var ret BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule
		return ret
	}
	return *o.CustomPromotionRule
}

// GetCustomPromotionRuleOk returns a tuple with the CustomPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetCustomPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule, bool) {
	if o == nil || IsNil(o.CustomPromotionRule) {
		return nil, false
	}
	return o.CustomPromotionRule, true
}

// HasCustomPromotionRule returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasCustomPromotionRule() bool {
	if o != nil && !IsNil(o.CustomPromotionRule) {
		return true
	}

	return false
}

// SetCustomPromotionRule gets a reference to the given BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule and assigns it to the CustomPromotionRule field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetCustomPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule) {
	o.CustomPromotionRule = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret BundleCreateDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleCreateDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BuyXPayYPromotionUpdateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *BuyXPayYPromotionUpdateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o BuyXPayYPromotionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyXPayYPromotionUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.OrderAmountPromotionRule) {
		toSerialize["order_amount_promotion_rule"] = o.OrderAmountPromotionRule
	}
	if !IsNil(o.SkuListPromotionRule) {
		toSerialize["sku_list_promotion_rule"] = o.SkuListPromotionRule
	}
	if !IsNil(o.CouponCodesPromotionRule) {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if !IsNil(o.CustomPromotionRule) {
		toSerialize["custom_promotion_rule"] = o.CustomPromotionRule
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableBuyXPayYPromotionUpdateDataRelationships struct {
	value *BuyXPayYPromotionUpdateDataRelationships
	isSet bool
}

func (v NullableBuyXPayYPromotionUpdateDataRelationships) Get() *BuyXPayYPromotionUpdateDataRelationships {
	return v.value
}

func (v *NullableBuyXPayYPromotionUpdateDataRelationships) Set(val *BuyXPayYPromotionUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyXPayYPromotionUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyXPayYPromotionUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyXPayYPromotionUpdateDataRelationships(val *BuyXPayYPromotionUpdateDataRelationships) *NullableBuyXPayYPromotionUpdateDataRelationships {
	return &NullableBuyXPayYPromotionUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableBuyXPayYPromotionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyXPayYPromotionUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}