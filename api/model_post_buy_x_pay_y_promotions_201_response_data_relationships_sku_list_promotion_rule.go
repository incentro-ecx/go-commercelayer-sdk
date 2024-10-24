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

// checks if the POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule{}

// POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule struct for POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule
type POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleData `json:"data,omitempty"`
}

// NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule instantiates a new POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule() *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule {
	this := POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule{}
	return &this
}

// NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleWithDefaults instantiates a new POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleWithDefaults() *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule {
	this := POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) GetData() POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) GetDataOk() (*POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleData and assigns it to the Data field.
func (o *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) SetData(v POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRuleData) {
	o.Data = &v
}

func (o POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule struct {
	value *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule
	isSet bool
}

func (v NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) Get() *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule {
	return v.value
}

func (v *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) Set(val *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule(val *POSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule {
	return &NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule{value: val, isSet: true}
}

func (v NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBuyXPayYPromotions201ResponseDataRelationshipsSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
