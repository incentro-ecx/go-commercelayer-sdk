/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule struct for GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule
type GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData `json:"data,omitempty"`
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule{}
	return &this
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleWithDefaults instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleWithDefaults() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetData() GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData {
	if o == nil || o.Data == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetDataOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData and assigns it to the Data field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) SetData(v GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleData) {
	o.Data = &v
}

func (o GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule struct {
	value *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule
	isSet bool
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) Get() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	return v.value
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) Set(val *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule(val *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	return &NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
