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

// InStockSubscriptionResponseDataRelationships struct for InStockSubscriptionResponseDataRelationships
type InStockSubscriptionResponseDataRelationships struct {
	Market   *BillingInfoValidationRuleResponseDataRelationshipsMarket `json:"market,omitempty"`
	Customer *CouponRecipientResponseDataRelationshipsCustomer         `json:"customer,omitempty"`
	Sku      *InStockSubscriptionResponseDataRelationshipsSku          `json:"sku,omitempty"`
	Events   *CustomerAddressResponseDataRelationshipsEvents           `json:"events,omitempty"`
}

// NewInStockSubscriptionResponseDataRelationships instantiates a new InStockSubscriptionResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionResponseDataRelationships() *InStockSubscriptionResponseDataRelationships {
	this := InStockSubscriptionResponseDataRelationships{}
	return &this
}

// NewInStockSubscriptionResponseDataRelationshipsWithDefaults instantiates a new InStockSubscriptionResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionResponseDataRelationshipsWithDefaults() *InStockSubscriptionResponseDataRelationships {
	this := InStockSubscriptionResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *InStockSubscriptionResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *InStockSubscriptionResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseDataRelationships) GetSku() InStockSubscriptionResponseDataRelationshipsSku {
	if o == nil || o.Sku == nil {
		var ret InStockSubscriptionResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseDataRelationships) GetSkuOk() (*InStockSubscriptionResponseDataRelationshipsSku, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *InStockSubscriptionResponseDataRelationships) SetSku(v InStockSubscriptionResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *InStockSubscriptionResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *InStockSubscriptionResponseDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *InStockSubscriptionResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o InStockSubscriptionResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableInStockSubscriptionResponseDataRelationships struct {
	value *InStockSubscriptionResponseDataRelationships
	isSet bool
}

func (v NullableInStockSubscriptionResponseDataRelationships) Get() *InStockSubscriptionResponseDataRelationships {
	return v.value
}

func (v *NullableInStockSubscriptionResponseDataRelationships) Set(val *InStockSubscriptionResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionResponseDataRelationships(val *InStockSubscriptionResponseDataRelationships) *NullableInStockSubscriptionResponseDataRelationships {
	return &NullableInStockSubscriptionResponseDataRelationships{value: val, isSet: true}
}

func (v NullableInStockSubscriptionResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}