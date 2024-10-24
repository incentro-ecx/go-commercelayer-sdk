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

// checks if the BillingInfoValidationRuleDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfoValidationRuleDataRelationships{}

// BillingInfoValidationRuleDataRelationships struct for BillingInfoValidationRuleDataRelationships
type BillingInfoValidationRuleDataRelationships struct {
	Market      *AvalaraAccountDataRelationshipsMarkets    `json:"market,omitempty"`
	Attachments *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
}

// NewBillingInfoValidationRuleDataRelationships instantiates a new BillingInfoValidationRuleDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleDataRelationships() *BillingInfoValidationRuleDataRelationships {
	this := BillingInfoValidationRuleDataRelationships{}
	return &this
}

// NewBillingInfoValidationRuleDataRelationshipsWithDefaults instantiates a new BillingInfoValidationRuleDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleDataRelationshipsWithDefaults() *BillingInfoValidationRuleDataRelationships {
	this := BillingInfoValidationRuleDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || IsNil(o.Market) {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *BillingInfoValidationRuleDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *BillingInfoValidationRuleDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *BillingInfoValidationRuleDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o BillingInfoValidationRuleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfoValidationRuleDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableBillingInfoValidationRuleDataRelationships struct {
	value *BillingInfoValidationRuleDataRelationships
	isSet bool
}

func (v NullableBillingInfoValidationRuleDataRelationships) Get() *BillingInfoValidationRuleDataRelationships {
	return v.value
}

func (v *NullableBillingInfoValidationRuleDataRelationships) Set(val *BillingInfoValidationRuleDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleDataRelationships(val *BillingInfoValidationRuleDataRelationships) *NullableBillingInfoValidationRuleDataRelationships {
	return &NullableBillingInfoValidationRuleDataRelationships{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
