/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTGiftCards201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTGiftCards201ResponseDataRelationships{}

// POSTGiftCards201ResponseDataRelationships struct for POSTGiftCards201ResponseDataRelationships
type POSTGiftCards201ResponseDataRelationships struct {
	Market            *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket   `json:"market,omitempty"`
	GiftCardRecipient *POSTGiftCards201ResponseDataRelationshipsGiftCardRecipient         `json:"gift_card_recipient,omitempty"`
	Attachments       *POSTAvalaraAccounts201ResponseDataRelationshipsAttachments         `json:"attachments,omitempty"`
	Events            *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents `json:"events,omitempty"`
}

// NewPOSTGiftCards201ResponseDataRelationships instantiates a new POSTGiftCards201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGiftCards201ResponseDataRelationships() *POSTGiftCards201ResponseDataRelationships {
	this := POSTGiftCards201ResponseDataRelationships{}
	return &this
}

// NewPOSTGiftCards201ResponseDataRelationshipsWithDefaults instantiates a new POSTGiftCards201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGiftCards201ResponseDataRelationshipsWithDefaults() *POSTGiftCards201ResponseDataRelationships {
	this := POSTGiftCards201ResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataRelationships) GetMarket() POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTGiftCards201ResponseDataRelationships) SetMarket(v POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetGiftCardRecipient returns the GiftCardRecipient field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataRelationships) GetGiftCardRecipient() POSTGiftCards201ResponseDataRelationshipsGiftCardRecipient {
	if o == nil || IsNil(o.GiftCardRecipient) {
		var ret POSTGiftCards201ResponseDataRelationshipsGiftCardRecipient
		return ret
	}
	return *o.GiftCardRecipient
}

// GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataRelationships) GetGiftCardRecipientOk() (*POSTGiftCards201ResponseDataRelationshipsGiftCardRecipient, bool) {
	if o == nil || IsNil(o.GiftCardRecipient) {
		return nil, false
	}
	return o.GiftCardRecipient, true
}

// HasGiftCardRecipient returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationships) HasGiftCardRecipient() bool {
	if o != nil && !IsNil(o.GiftCardRecipient) {
		return true
	}

	return false
}

// SetGiftCardRecipient gets a reference to the given POSTGiftCards201ResponseDataRelationshipsGiftCardRecipient and assigns it to the GiftCardRecipient field.
func (o *POSTGiftCards201ResponseDataRelationships) SetGiftCardRecipient(v POSTGiftCards201ResponseDataRelationshipsGiftCardRecipient) {
	o.GiftCardRecipient = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTGiftCards201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTGiftCards201ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o POSTGiftCards201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTGiftCards201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.GiftCardRecipient) {
		toSerialize["gift_card_recipient"] = o.GiftCardRecipient
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullablePOSTGiftCards201ResponseDataRelationships struct {
	value *POSTGiftCards201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTGiftCards201ResponseDataRelationships) Get() *POSTGiftCards201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTGiftCards201ResponseDataRelationships) Set(val *POSTGiftCards201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGiftCards201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGiftCards201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGiftCards201ResponseDataRelationships(val *POSTGiftCards201ResponseDataRelationships) *NullablePOSTGiftCards201ResponseDataRelationships {
	return &NullablePOSTGiftCards201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTGiftCards201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGiftCards201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
