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

// checks if the PriceListSchedulerCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListSchedulerCreateDataRelationships{}

// PriceListSchedulerCreateDataRelationships struct for PriceListSchedulerCreateDataRelationships
type PriceListSchedulerCreateDataRelationships struct {
	Market    BillingInfoValidationRuleCreateDataRelationshipsMarket `json:"market"`
	PriceList MarketCreateDataRelationshipsPriceList                 `json:"price_list"`
}

// NewPriceListSchedulerCreateDataRelationships instantiates a new PriceListSchedulerCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListSchedulerCreateDataRelationships(market BillingInfoValidationRuleCreateDataRelationshipsMarket, priceList MarketCreateDataRelationshipsPriceList) *PriceListSchedulerCreateDataRelationships {
	this := PriceListSchedulerCreateDataRelationships{}
	this.Market = market
	this.PriceList = priceList
	return &this
}

// NewPriceListSchedulerCreateDataRelationshipsWithDefaults instantiates a new PriceListSchedulerCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListSchedulerCreateDataRelationshipsWithDefaults() *PriceListSchedulerCreateDataRelationships {
	this := PriceListSchedulerCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value
func (o *PriceListSchedulerCreateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerCreateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *PriceListSchedulerCreateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = v
}

// GetPriceList returns the PriceList field value
func (o *PriceListSchedulerCreateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList {
	if o == nil {
		var ret MarketCreateDataRelationshipsPriceList
		return ret
	}

	return o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value
// and a boolean to check if the value has been set.
func (o *PriceListSchedulerCreateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceList, true
}

// SetPriceList sets field value
func (o *PriceListSchedulerCreateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList) {
	o.PriceList = v
}

func (o PriceListSchedulerCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListSchedulerCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["market"] = o.Market
	toSerialize["price_list"] = o.PriceList
	return toSerialize, nil
}

type NullablePriceListSchedulerCreateDataRelationships struct {
	value *PriceListSchedulerCreateDataRelationships
	isSet bool
}

func (v NullablePriceListSchedulerCreateDataRelationships) Get() *PriceListSchedulerCreateDataRelationships {
	return v.value
}

func (v *NullablePriceListSchedulerCreateDataRelationships) Set(val *PriceListSchedulerCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListSchedulerCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListSchedulerCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListSchedulerCreateDataRelationships(val *PriceListSchedulerCreateDataRelationships) *NullablePriceListSchedulerCreateDataRelationships {
	return &NullablePriceListSchedulerCreateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceListSchedulerCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListSchedulerCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
