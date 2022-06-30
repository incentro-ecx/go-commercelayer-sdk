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

// PriceTierDataRelationships struct for PriceTierDataRelationships
type PriceTierDataRelationships struct {
	Price       *PriceListDataRelationshipsPrices           `json:"price,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewPriceTierDataRelationships instantiates a new PriceTierDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTierDataRelationships() *PriceTierDataRelationships {
	this := PriceTierDataRelationships{}
	return &this
}

// NewPriceTierDataRelationshipsWithDefaults instantiates a new PriceTierDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTierDataRelationshipsWithDefaults() *PriceTierDataRelationships {
	this := PriceTierDataRelationships{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PriceTierDataRelationships) GetPrice() PriceListDataRelationshipsPrices {
	if o == nil || o.Price == nil {
		var ret PriceListDataRelationshipsPrices
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTierDataRelationships) GetPriceOk() (*PriceListDataRelationshipsPrices, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceTierDataRelationships) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PriceListDataRelationshipsPrices and assigns it to the Price field.
func (o *PriceTierDataRelationships) SetPrice(v PriceListDataRelationshipsPrices) {
	o.Price = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PriceTierDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTierDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PriceTierDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *PriceTierDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o PriceTierDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullablePriceTierDataRelationships struct {
	value *PriceTierDataRelationships
	isSet bool
}

func (v NullablePriceTierDataRelationships) Get() *PriceTierDataRelationships {
	return v.value
}

func (v *NullablePriceTierDataRelationships) Set(val *PriceTierDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTierDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTierDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTierDataRelationships(val *PriceTierDataRelationships) *NullablePriceTierDataRelationships {
	return &NullablePriceTierDataRelationships{value: val, isSet: true}
}

func (v NullablePriceTierDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTierDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
