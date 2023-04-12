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

// checks if the POSTSubscriptionModels201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSubscriptionModels201ResponseDataRelationships{}

// POSTSubscriptionModels201ResponseDataRelationships struct for POSTSubscriptionModels201ResponseDataRelationships
type POSTSubscriptionModels201ResponseDataRelationships struct {
	Markets            *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets      `json:"markets,omitempty"`
	OrderSubscriptions *POSTCustomers201ResponseDataRelationshipsOrderSubscriptions `json:"order_subscriptions,omitempty"`
	Attachments        *POSTAvalaraAccounts201ResponseDataRelationshipsAttachments  `json:"attachments,omitempty"`
}

// NewPOSTSubscriptionModels201ResponseDataRelationships instantiates a new POSTSubscriptionModels201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSubscriptionModels201ResponseDataRelationships() *POSTSubscriptionModels201ResponseDataRelationships {
	this := POSTSubscriptionModels201ResponseDataRelationships{}
	return &this
}

// NewPOSTSubscriptionModels201ResponseDataRelationshipsWithDefaults instantiates a new POSTSubscriptionModels201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSubscriptionModels201ResponseDataRelationshipsWithDefaults() *POSTSubscriptionModels201ResponseDataRelationships {
	this := POSTSubscriptionModels201ResponseDataRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *POSTSubscriptionModels201ResponseDataRelationships) GetMarkets() POSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	if o == nil || IsNil(o.Markets) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSubscriptionModels201ResponseDataRelationships) GetMarketsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataRelationships) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsMarkets and assigns it to the Markets field.
func (o *POSTSubscriptionModels201ResponseDataRelationships) SetMarkets(v POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetOrderSubscriptions returns the OrderSubscriptions field value if set, zero value otherwise.
func (o *POSTSubscriptionModels201ResponseDataRelationships) GetOrderSubscriptions() POSTCustomers201ResponseDataRelationshipsOrderSubscriptions {
	if o == nil || IsNil(o.OrderSubscriptions) {
		var ret POSTCustomers201ResponseDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscriptions
}

// GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSubscriptionModels201ResponseDataRelationships) GetOrderSubscriptionsOk() (*POSTCustomers201ResponseDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || IsNil(o.OrderSubscriptions) {
		return nil, false
	}
	return o.OrderSubscriptions, true
}

// HasOrderSubscriptions returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataRelationships) HasOrderSubscriptions() bool {
	if o != nil && !IsNil(o.OrderSubscriptions) {
		return true
	}

	return false
}

// SetOrderSubscriptions gets a reference to the given POSTCustomers201ResponseDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscriptions field.
func (o *POSTSubscriptionModels201ResponseDataRelationships) SetOrderSubscriptions(v POSTCustomers201ResponseDataRelationshipsOrderSubscriptions) {
	o.OrderSubscriptions = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTSubscriptionModels201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSubscriptionModels201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTSubscriptionModels201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o POSTSubscriptionModels201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSubscriptionModels201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	if !IsNil(o.OrderSubscriptions) {
		toSerialize["order_subscriptions"] = o.OrderSubscriptions
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullablePOSTSubscriptionModels201ResponseDataRelationships struct {
	value *POSTSubscriptionModels201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTSubscriptionModels201ResponseDataRelationships) Get() *POSTSubscriptionModels201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTSubscriptionModels201ResponseDataRelationships) Set(val *POSTSubscriptionModels201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSubscriptionModels201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSubscriptionModels201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSubscriptionModels201ResponseDataRelationships(val *POSTSubscriptionModels201ResponseDataRelationships) *NullablePOSTSubscriptionModels201ResponseDataRelationships {
	return &NullablePOSTSubscriptionModels201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTSubscriptionModels201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSubscriptionModels201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}