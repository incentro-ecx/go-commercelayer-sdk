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

// AvalaraAccountResponseDataRelationships struct for AvalaraAccountResponseDataRelationships
type AvalaraAccountResponseDataRelationships struct {
	TaxCategories *AvalaraAccountResponseDataRelationshipsTaxCategories `json:"tax_categories,omitempty"`
	Markets       *AvalaraAccountResponseDataRelationshipsMarkets       `json:"markets,omitempty"`
	Attachments   *AvalaraAccountResponseDataRelationshipsAttachments   `json:"attachments,omitempty"`
}

// NewAvalaraAccountResponseDataRelationships instantiates a new AvalaraAccountResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountResponseDataRelationships() *AvalaraAccountResponseDataRelationships {
	this := AvalaraAccountResponseDataRelationships{}
	return &this
}

// NewAvalaraAccountResponseDataRelationshipsWithDefaults instantiates a new AvalaraAccountResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountResponseDataRelationshipsWithDefaults() *AvalaraAccountResponseDataRelationships {
	this := AvalaraAccountResponseDataRelationships{}
	return &this
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *AvalaraAccountResponseDataRelationships) GetTaxCategories() AvalaraAccountResponseDataRelationshipsTaxCategories {
	if o == nil || o.TaxCategories == nil {
		var ret AvalaraAccountResponseDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountResponseDataRelationships) GetTaxCategoriesOk() (*AvalaraAccountResponseDataRelationshipsTaxCategories, bool) {
	if o == nil || o.TaxCategories == nil {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *AvalaraAccountResponseDataRelationships) HasTaxCategories() bool {
	if o != nil && o.TaxCategories != nil {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given AvalaraAccountResponseDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *AvalaraAccountResponseDataRelationships) SetTaxCategories(v AvalaraAccountResponseDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *AvalaraAccountResponseDataRelationships) GetMarkets() AvalaraAccountResponseDataRelationshipsMarkets {
	if o == nil || o.Markets == nil {
		var ret AvalaraAccountResponseDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountResponseDataRelationships) GetMarketsOk() (*AvalaraAccountResponseDataRelationshipsMarkets, bool) {
	if o == nil || o.Markets == nil {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *AvalaraAccountResponseDataRelationships) HasMarkets() bool {
	if o != nil && o.Markets != nil {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given AvalaraAccountResponseDataRelationshipsMarkets and assigns it to the Markets field.
func (o *AvalaraAccountResponseDataRelationships) SetMarkets(v AvalaraAccountResponseDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AvalaraAccountResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AvalaraAccountResponseDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *AvalaraAccountResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o AvalaraAccountResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxCategories != nil {
		toSerialize["tax_categories"] = o.TaxCategories
	}
	if o.Markets != nil {
		toSerialize["markets"] = o.Markets
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountResponseDataRelationships struct {
	value *AvalaraAccountResponseDataRelationships
	isSet bool
}

func (v NullableAvalaraAccountResponseDataRelationships) Get() *AvalaraAccountResponseDataRelationships {
	return v.value
}

func (v *NullableAvalaraAccountResponseDataRelationships) Set(val *AvalaraAccountResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountResponseDataRelationships(val *AvalaraAccountResponseDataRelationships) *NullableAvalaraAccountResponseDataRelationships {
	return &NullableAvalaraAccountResponseDataRelationships{value: val, isSet: true}
}

func (v NullableAvalaraAccountResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}