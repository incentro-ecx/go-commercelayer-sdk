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

// checks if the PriceUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceUpdateDataRelationships{}

// PriceUpdateDataRelationships struct for PriceUpdateDataRelationships
type PriceUpdateDataRelationships struct {
	PriceList  *MarketCreateDataRelationshipsPriceList        `json:"price_list,omitempty"`
	Sku        *InStockSubscriptionCreateDataRelationshipsSku `json:"sku,omitempty"`
	PriceTiers *PriceCreateDataRelationshipsPriceTiers        `json:"price_tiers,omitempty"`
}

// NewPriceUpdateDataRelationships instantiates a new PriceUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceUpdateDataRelationships() *PriceUpdateDataRelationships {
	this := PriceUpdateDataRelationships{}
	return &this
}

// NewPriceUpdateDataRelationshipsWithDefaults instantiates a new PriceUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceUpdateDataRelationshipsWithDefaults() *PriceUpdateDataRelationships {
	this := PriceUpdateDataRelationships{}
	return &this
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *PriceUpdateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList {
	if o == nil || IsNil(o.PriceList) {
		var ret MarketCreateDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceUpdateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool) {
	if o == nil || IsNil(o.PriceList) {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *PriceUpdateDataRelationships) HasPriceList() bool {
	if o != nil && !IsNil(o.PriceList) {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given MarketCreateDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *PriceUpdateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList) {
	o.PriceList = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *PriceUpdateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceUpdateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *PriceUpdateDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *PriceUpdateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

// GetPriceTiers returns the PriceTiers field value if set, zero value otherwise.
func (o *PriceUpdateDataRelationships) GetPriceTiers() PriceCreateDataRelationshipsPriceTiers {
	if o == nil || IsNil(o.PriceTiers) {
		var ret PriceCreateDataRelationshipsPriceTiers
		return ret
	}
	return *o.PriceTiers
}

// GetPriceTiersOk returns a tuple with the PriceTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceUpdateDataRelationships) GetPriceTiersOk() (*PriceCreateDataRelationshipsPriceTiers, bool) {
	if o == nil || IsNil(o.PriceTiers) {
		return nil, false
	}
	return o.PriceTiers, true
}

// HasPriceTiers returns a boolean if a field has been set.
func (o *PriceUpdateDataRelationships) HasPriceTiers() bool {
	if o != nil && !IsNil(o.PriceTiers) {
		return true
	}

	return false
}

// SetPriceTiers gets a reference to the given PriceCreateDataRelationshipsPriceTiers and assigns it to the PriceTiers field.
func (o *PriceUpdateDataRelationships) SetPriceTiers(v PriceCreateDataRelationshipsPriceTiers) {
	o.PriceTiers = &v
}

func (o PriceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceList) {
		toSerialize["price_list"] = o.PriceList
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.PriceTiers) {
		toSerialize["price_tiers"] = o.PriceTiers
	}
	return toSerialize, nil
}

type NullablePriceUpdateDataRelationships struct {
	value *PriceUpdateDataRelationships
	isSet bool
}

func (v NullablePriceUpdateDataRelationships) Get() *PriceUpdateDataRelationships {
	return v.value
}

func (v *NullablePriceUpdateDataRelationships) Set(val *PriceUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceUpdateDataRelationships(val *PriceUpdateDataRelationships) *NullablePriceUpdateDataRelationships {
	return &NullablePriceUpdateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
