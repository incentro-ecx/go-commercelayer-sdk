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

// checks if the BundleCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleCreateDataRelationships{}

// BundleCreateDataRelationships struct for BundleCreateDataRelationships
type BundleCreateDataRelationships struct {
	Market  *BillingInfoValidationRuleCreateDataRelationshipsMarket `json:"market,omitempty"`
	SkuList BundleCreateDataRelationshipsSkuList                    `json:"sku_list"`
	Tags    *AddressCreateDataRelationshipsTags                     `json:"tags,omitempty"`
}

// NewBundleCreateDataRelationships instantiates a new BundleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleCreateDataRelationships(skuList BundleCreateDataRelationshipsSkuList) *BundleCreateDataRelationships {
	this := BundleCreateDataRelationships{}
	this.SkuList = skuList
	return &this
}

// NewBundleCreateDataRelationshipsWithDefaults instantiates a new BundleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleCreateDataRelationshipsWithDefaults() *BundleCreateDataRelationships {
	this := BundleCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *BundleCreateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *BundleCreateDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *BundleCreateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetSkuList returns the SkuList field value
func (o *BundleCreateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList {
	if o == nil {
		var ret BundleCreateDataRelationshipsSkuList
		return ret
	}

	return o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuList, true
}

// SetSkuList sets field value
func (o *BundleCreateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList) {
	o.SkuList = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BundleCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BundleCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *BundleCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o BundleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	toSerialize["sku_list"] = o.SkuList
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableBundleCreateDataRelationships struct {
	value *BundleCreateDataRelationships
	isSet bool
}

func (v NullableBundleCreateDataRelationships) Get() *BundleCreateDataRelationships {
	return v.value
}

func (v *NullableBundleCreateDataRelationships) Set(val *BundleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleCreateDataRelationships(val *BundleCreateDataRelationships) *NullableBundleCreateDataRelationships {
	return &NullableBundleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableBundleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
