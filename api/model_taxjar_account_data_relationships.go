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

// checks if the TaxjarAccountDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxjarAccountDataRelationships{}

// TaxjarAccountDataRelationships struct for TaxjarAccountDataRelationships
type TaxjarAccountDataRelationships struct {
	Markets       *AvalaraAccountDataRelationshipsMarkets       `json:"markets,omitempty"`
	Attachments   *AuthorizationDataRelationshipsAttachments    `json:"attachments,omitempty"`
	Versions      *AddressDataRelationshipsVersions             `json:"versions,omitempty"`
	TaxCategories *AvalaraAccountDataRelationshipsTaxCategories `json:"tax_categories,omitempty"`
}

// NewTaxjarAccountDataRelationships instantiates a new TaxjarAccountDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountDataRelationships() *TaxjarAccountDataRelationships {
	this := TaxjarAccountDataRelationships{}
	return &this
}

// NewTaxjarAccountDataRelationshipsWithDefaults instantiates a new TaxjarAccountDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountDataRelationshipsWithDefaults() *TaxjarAccountDataRelationships {
	this := TaxjarAccountDataRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *TaxjarAccountDataRelationships) GetMarkets() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || IsNil(o.Markets) {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountDataRelationships) GetMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *TaxjarAccountDataRelationships) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Markets field.
func (o *TaxjarAccountDataRelationships) SetMarkets(v AvalaraAccountDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TaxjarAccountDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TaxjarAccountDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *TaxjarAccountDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *TaxjarAccountDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *TaxjarAccountDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *TaxjarAccountDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *TaxjarAccountDataRelationships) GetTaxCategories() AvalaraAccountDataRelationshipsTaxCategories {
	if o == nil || IsNil(o.TaxCategories) {
		var ret AvalaraAccountDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxjarAccountDataRelationships) GetTaxCategoriesOk() (*AvalaraAccountDataRelationshipsTaxCategories, bool) {
	if o == nil || IsNil(o.TaxCategories) {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *TaxjarAccountDataRelationships) HasTaxCategories() bool {
	if o != nil && !IsNil(o.TaxCategories) {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given AvalaraAccountDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *TaxjarAccountDataRelationships) SetTaxCategories(v AvalaraAccountDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

func (o TaxjarAccountDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxjarAccountDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.TaxCategories) {
		toSerialize["tax_categories"] = o.TaxCategories
	}
	return toSerialize, nil
}

type NullableTaxjarAccountDataRelationships struct {
	value *TaxjarAccountDataRelationships
	isSet bool
}

func (v NullableTaxjarAccountDataRelationships) Get() *TaxjarAccountDataRelationships {
	return v.value
}

func (v *NullableTaxjarAccountDataRelationships) Set(val *TaxjarAccountDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountDataRelationships(val *TaxjarAccountDataRelationships) *NullableTaxjarAccountDataRelationships {
	return &NullableTaxjarAccountDataRelationships{value: val, isSet: true}
}

func (v NullableTaxjarAccountDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
