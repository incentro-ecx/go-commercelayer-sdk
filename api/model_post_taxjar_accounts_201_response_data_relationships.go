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

// checks if the POSTTaxjarAccounts201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTTaxjarAccounts201ResponseDataRelationships{}

// POSTTaxjarAccounts201ResponseDataRelationships struct for POSTTaxjarAccounts201ResponseDataRelationships
type POSTTaxjarAccounts201ResponseDataRelationships struct {
	Markets       *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets                  `json:"markets,omitempty"`
	Attachments   *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions      *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
	TaxCategories *POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories            `json:"tax_categories,omitempty"`
}

// NewPOSTTaxjarAccounts201ResponseDataRelationships instantiates a new POSTTaxjarAccounts201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxjarAccounts201ResponseDataRelationships() *POSTTaxjarAccounts201ResponseDataRelationships {
	this := POSTTaxjarAccounts201ResponseDataRelationships{}
	return &this
}

// NewPOSTTaxjarAccounts201ResponseDataRelationshipsWithDefaults instantiates a new POSTTaxjarAccounts201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxjarAccounts201ResponseDataRelationshipsWithDefaults() *POSTTaxjarAccounts201ResponseDataRelationships {
	this := POSTTaxjarAccounts201ResponseDataRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetMarkets() POSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	if o == nil || IsNil(o.Markets) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetMarketsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsMarkets, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsMarkets and assigns it to the Markets field.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) SetMarkets(v POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetTaxCategories() POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories {
	if o == nil || IsNil(o.TaxCategories) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) GetTaxCategoriesOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories, bool) {
	if o == nil || IsNil(o.TaxCategories) {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) HasTaxCategories() bool {
	if o != nil && !IsNil(o.TaxCategories) {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *POSTTaxjarAccounts201ResponseDataRelationships) SetTaxCategories(v POSTAvalaraAccounts201ResponseDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

func (o POSTTaxjarAccounts201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTTaxjarAccounts201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTTaxjarAccounts201ResponseDataRelationships struct {
	value *POSTTaxjarAccounts201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTTaxjarAccounts201ResponseDataRelationships) Get() *POSTTaxjarAccounts201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTTaxjarAccounts201ResponseDataRelationships) Set(val *POSTTaxjarAccounts201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxjarAccounts201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxjarAccounts201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxjarAccounts201ResponseDataRelationships(val *POSTTaxjarAccounts201ResponseDataRelationships) *NullablePOSTTaxjarAccounts201ResponseDataRelationships {
	return &NullablePOSTTaxjarAccounts201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTTaxjarAccounts201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxjarAccounts201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}