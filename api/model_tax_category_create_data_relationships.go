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

// checks if the TaxCategoryCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCategoryCreateDataRelationships{}

// TaxCategoryCreateDataRelationships struct for TaxCategoryCreateDataRelationships
type TaxCategoryCreateDataRelationships struct {
	Sku           InStockSubscriptionCreateDataRelationshipsSku   `json:"sku"`
	TaxCalculator TaxCategoryCreateDataRelationshipsTaxCalculator `json:"tax_calculator"`
}

// NewTaxCategoryCreateDataRelationships instantiates a new TaxCategoryCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryCreateDataRelationships(sku InStockSubscriptionCreateDataRelationshipsSku, taxCalculator TaxCategoryCreateDataRelationshipsTaxCalculator) *TaxCategoryCreateDataRelationships {
	this := TaxCategoryCreateDataRelationships{}
	this.Sku = sku
	this.TaxCalculator = taxCalculator
	return &this
}

// NewTaxCategoryCreateDataRelationshipsWithDefaults instantiates a new TaxCategoryCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryCreateDataRelationshipsWithDefaults() *TaxCategoryCreateDataRelationships {
	this := TaxCategoryCreateDataRelationships{}
	return &this
}

// GetSku returns the Sku field value
func (o *TaxCategoryCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *TaxCategoryCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = v
}

// GetTaxCalculator returns the TaxCalculator field value
func (o *TaxCategoryCreateDataRelationships) GetTaxCalculator() TaxCategoryCreateDataRelationshipsTaxCalculator {
	if o == nil {
		var ret TaxCategoryCreateDataRelationshipsTaxCalculator
		return ret
	}

	return o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryCreateDataRelationships) GetTaxCalculatorOk() (*TaxCategoryCreateDataRelationshipsTaxCalculator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxCalculator, true
}

// SetTaxCalculator sets field value
func (o *TaxCategoryCreateDataRelationships) SetTaxCalculator(v TaxCategoryCreateDataRelationshipsTaxCalculator) {
	o.TaxCalculator = v
}

func (o TaxCategoryCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCategoryCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["tax_calculator"] = o.TaxCalculator
	return toSerialize, nil
}

type NullableTaxCategoryCreateDataRelationships struct {
	value *TaxCategoryCreateDataRelationships
	isSet bool
}

func (v NullableTaxCategoryCreateDataRelationships) Get() *TaxCategoryCreateDataRelationships {
	return v.value
}

func (v *NullableTaxCategoryCreateDataRelationships) Set(val *TaxCategoryCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryCreateDataRelationships(val *TaxCategoryCreateDataRelationships) *NullableTaxCategoryCreateDataRelationships {
	return &NullableTaxCategoryCreateDataRelationships{value: val, isSet: true}
}

func (v NullableTaxCategoryCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
