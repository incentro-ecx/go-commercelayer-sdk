/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GiftCardCreateDataRelationships struct for GiftCardCreateDataRelationships
type GiftCardCreateDataRelationships struct {
	Market            *AvalaraAccountDataRelationshipsMarkets     `json:"market,omitempty"`
	GiftCardRecipient *GiftCardDataRelationshipsGiftCardRecipient `json:"gift_card_recipient,omitempty"`
}

// NewGiftCardCreateDataRelationships instantiates a new GiftCardCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardCreateDataRelationships() *GiftCardCreateDataRelationships {
	this := GiftCardCreateDataRelationships{}
	return &this
}

// NewGiftCardCreateDataRelationshipsWithDefaults instantiates a new GiftCardCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardCreateDataRelationshipsWithDefaults() *GiftCardCreateDataRelationships {
	this := GiftCardCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GiftCardCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GiftCardCreateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *GiftCardCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetGiftCardRecipient returns the GiftCardRecipient field value if set, zero value otherwise.
func (o *GiftCardCreateDataRelationships) GetGiftCardRecipient() GiftCardDataRelationshipsGiftCardRecipient {
	if o == nil || o.GiftCardRecipient == nil {
		var ret GiftCardDataRelationshipsGiftCardRecipient
		return ret
	}
	return *o.GiftCardRecipient
}

// GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardCreateDataRelationships) GetGiftCardRecipientOk() (*GiftCardDataRelationshipsGiftCardRecipient, bool) {
	if o == nil || o.GiftCardRecipient == nil {
		return nil, false
	}
	return o.GiftCardRecipient, true
}

// HasGiftCardRecipient returns a boolean if a field has been set.
func (o *GiftCardCreateDataRelationships) HasGiftCardRecipient() bool {
	if o != nil && o.GiftCardRecipient != nil {
		return true
	}

	return false
}

// SetGiftCardRecipient gets a reference to the given GiftCardDataRelationshipsGiftCardRecipient and assigns it to the GiftCardRecipient field.
func (o *GiftCardCreateDataRelationships) SetGiftCardRecipient(v GiftCardDataRelationshipsGiftCardRecipient) {
	o.GiftCardRecipient = &v
}

func (o GiftCardCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.GiftCardRecipient != nil {
		toSerialize["gift_card_recipient"] = o.GiftCardRecipient
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardCreateDataRelationships struct {
	value *GiftCardCreateDataRelationships
	isSet bool
}

func (v NullableGiftCardCreateDataRelationships) Get() *GiftCardCreateDataRelationships {
	return v.value
}

func (v *NullableGiftCardCreateDataRelationships) Set(val *GiftCardCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardCreateDataRelationships(val *GiftCardCreateDataRelationships) *NullableGiftCardCreateDataRelationships {
	return &NullableGiftCardCreateDataRelationships{value: val, isSet: true}
}

func (v NullableGiftCardCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
