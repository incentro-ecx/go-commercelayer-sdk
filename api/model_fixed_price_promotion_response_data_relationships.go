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

// FixedPricePromotionResponseDataRelationships struct for FixedPricePromotionResponseDataRelationships
type FixedPricePromotionResponseDataRelationships struct {
	Market                   *BillingInfoValidationRuleResponseDataRelationshipsMarket           `json:"market,omitempty"`
	PromotionRules           *ExternalPromotionResponseDataRelationshipsPromotionRules           `json:"promotion_rules,omitempty"`
	OrderAmountPromotionRule *ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule `json:"order_amount_promotion_rule,omitempty"`
	SkuListPromotionRule     *ExternalPromotionResponseDataRelationshipsSkuListPromotionRule     `json:"sku_list_promotion_rule,omitempty"`
	CouponCodesPromotionRule *ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule `json:"coupon_codes_promotion_rule,omitempty"`
	Attachments              *AvalaraAccountResponseDataRelationshipsAttachments                 `json:"attachments,omitempty"`
	SkuList                  *BundleResponseDataRelationshipsSkuList                             `json:"sku_list,omitempty"`
	Skus                     *BundleResponseDataRelationshipsSkus                                `json:"skus,omitempty"`
}

// NewFixedPricePromotionResponseDataRelationships instantiates a new FixedPricePromotionResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotionResponseDataRelationships() *FixedPricePromotionResponseDataRelationships {
	this := FixedPricePromotionResponseDataRelationships{}
	return &this
}

// NewFixedPricePromotionResponseDataRelationshipsWithDefaults instantiates a new FixedPricePromotionResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionResponseDataRelationshipsWithDefaults() *FixedPricePromotionResponseDataRelationships {
	this := FixedPricePromotionResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *FixedPricePromotionResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetPromotionRules returns the PromotionRules field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetPromotionRules() ExternalPromotionResponseDataRelationshipsPromotionRules {
	if o == nil || o.PromotionRules == nil {
		var ret ExternalPromotionResponseDataRelationshipsPromotionRules
		return ret
	}
	return *o.PromotionRules
}

// GetPromotionRulesOk returns a tuple with the PromotionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetPromotionRulesOk() (*ExternalPromotionResponseDataRelationshipsPromotionRules, bool) {
	if o == nil || o.PromotionRules == nil {
		return nil, false
	}
	return o.PromotionRules, true
}

// HasPromotionRules returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasPromotionRules() bool {
	if o != nil && o.PromotionRules != nil {
		return true
	}

	return false
}

// SetPromotionRules gets a reference to the given ExternalPromotionResponseDataRelationshipsPromotionRules and assigns it to the PromotionRules field.
func (o *FixedPricePromotionResponseDataRelationships) SetPromotionRules(v ExternalPromotionResponseDataRelationshipsPromotionRules) {
	o.PromotionRules = &v
}

// GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetOrderAmountPromotionRule() ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule {
	if o == nil || o.OrderAmountPromotionRule == nil {
		var ret ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule
		return ret
	}
	return *o.OrderAmountPromotionRule
}

// GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetOrderAmountPromotionRuleOk() (*ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule, bool) {
	if o == nil || o.OrderAmountPromotionRule == nil {
		return nil, false
	}
	return o.OrderAmountPromotionRule, true
}

// HasOrderAmountPromotionRule returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasOrderAmountPromotionRule() bool {
	if o != nil && o.OrderAmountPromotionRule != nil {
		return true
	}

	return false
}

// SetOrderAmountPromotionRule gets a reference to the given ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule and assigns it to the OrderAmountPromotionRule field.
func (o *FixedPricePromotionResponseDataRelationships) SetOrderAmountPromotionRule(v ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule) {
	o.OrderAmountPromotionRule = &v
}

// GetSkuListPromotionRule returns the SkuListPromotionRule field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetSkuListPromotionRule() ExternalPromotionResponseDataRelationshipsSkuListPromotionRule {
	if o == nil || o.SkuListPromotionRule == nil {
		var ret ExternalPromotionResponseDataRelationshipsSkuListPromotionRule
		return ret
	}
	return *o.SkuListPromotionRule
}

// GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetSkuListPromotionRuleOk() (*ExternalPromotionResponseDataRelationshipsSkuListPromotionRule, bool) {
	if o == nil || o.SkuListPromotionRule == nil {
		return nil, false
	}
	return o.SkuListPromotionRule, true
}

// HasSkuListPromotionRule returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasSkuListPromotionRule() bool {
	if o != nil && o.SkuListPromotionRule != nil {
		return true
	}

	return false
}

// SetSkuListPromotionRule gets a reference to the given ExternalPromotionResponseDataRelationshipsSkuListPromotionRule and assigns it to the SkuListPromotionRule field.
func (o *FixedPricePromotionResponseDataRelationships) SetSkuListPromotionRule(v ExternalPromotionResponseDataRelationshipsSkuListPromotionRule) {
	o.SkuListPromotionRule = &v
}

// GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetCouponCodesPromotionRule() ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule {
	if o == nil || o.CouponCodesPromotionRule == nil {
		var ret ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule
		return ret
	}
	return *o.CouponCodesPromotionRule
}

// GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetCouponCodesPromotionRuleOk() (*ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule, bool) {
	if o == nil || o.CouponCodesPromotionRule == nil {
		return nil, false
	}
	return o.CouponCodesPromotionRule, true
}

// HasCouponCodesPromotionRule returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasCouponCodesPromotionRule() bool {
	if o != nil && o.CouponCodesPromotionRule != nil {
		return true
	}

	return false
}

// SetCouponCodesPromotionRule gets a reference to the given ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule and assigns it to the CouponCodesPromotionRule field.
func (o *FixedPricePromotionResponseDataRelationships) SetCouponCodesPromotionRule(v ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule) {
	o.CouponCodesPromotionRule = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *FixedPricePromotionResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetSkuList() BundleResponseDataRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret BundleResponseDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetSkuListOk() (*BundleResponseDataRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleResponseDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *FixedPricePromotionResponseDataRelationships) SetSkuList(v BundleResponseDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *FixedPricePromotionResponseDataRelationships) GetSkus() BundleResponseDataRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret BundleResponseDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionResponseDataRelationships) GetSkusOk() (*BundleResponseDataRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *FixedPricePromotionResponseDataRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given BundleResponseDataRelationshipsSkus and assigns it to the Skus field.
func (o *FixedPricePromotionResponseDataRelationships) SetSkus(v BundleResponseDataRelationshipsSkus) {
	o.Skus = &v
}

func (o FixedPricePromotionResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.PromotionRules != nil {
		toSerialize["promotion_rules"] = o.PromotionRules
	}
	if o.OrderAmountPromotionRule != nil {
		toSerialize["order_amount_promotion_rule"] = o.OrderAmountPromotionRule
	}
	if o.SkuListPromotionRule != nil {
		toSerialize["sku_list_promotion_rule"] = o.SkuListPromotionRule
	}
	if o.CouponCodesPromotionRule != nil {
		toSerialize["coupon_codes_promotion_rule"] = o.CouponCodesPromotionRule
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	return json.Marshal(toSerialize)
}

type NullableFixedPricePromotionResponseDataRelationships struct {
	value *FixedPricePromotionResponseDataRelationships
	isSet bool
}

func (v NullableFixedPricePromotionResponseDataRelationships) Get() *FixedPricePromotionResponseDataRelationships {
	return v.value
}

func (v *NullableFixedPricePromotionResponseDataRelationships) Set(val *FixedPricePromotionResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotionResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotionResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotionResponseDataRelationships(val *FixedPricePromotionResponseDataRelationships) *NullableFixedPricePromotionResponseDataRelationships {
	return &NullableFixedPricePromotionResponseDataRelationships{value: val, isSet: true}
}

func (v NullableFixedPricePromotionResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotionResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}