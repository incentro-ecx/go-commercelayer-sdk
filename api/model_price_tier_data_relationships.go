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

// checks if the PriceTierDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceTierDataRelationships{}

// PriceTierDataRelationships struct for PriceTierDataRelationships
type PriceTierDataRelationships struct {
	Price       *PriceFrequencyTierDataRelationshipsPrice  `json:"price,omitempty"`
	Attachments *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
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
func (o *PriceTierDataRelationships) GetPrice() PriceFrequencyTierDataRelationshipsPrice {
	if o == nil || IsNil(o.Price) {
		var ret PriceFrequencyTierDataRelationshipsPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTierDataRelationships) GetPriceOk() (*PriceFrequencyTierDataRelationshipsPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceTierDataRelationships) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PriceFrequencyTierDataRelationshipsPrice and assigns it to the Price field.
func (o *PriceTierDataRelationships) SetPrice(v PriceFrequencyTierDataRelationshipsPrice) {
	o.Price = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PriceTierDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTierDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PriceTierDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *PriceTierDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *PriceTierDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTierDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *PriceTierDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *PriceTierDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o PriceTierDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceTierDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
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
