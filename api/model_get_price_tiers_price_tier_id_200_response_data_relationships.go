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

// checks if the GETPriceTiersPriceTierId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPriceTiersPriceTierId200ResponseDataRelationships{}

// GETPriceTiersPriceTierId200ResponseDataRelationships struct for GETPriceTiersPriceTierId200ResponseDataRelationships
type GETPriceTiersPriceTierId200ResponseDataRelationships struct {
	Price       *POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice                `json:"price,omitempty"`
	Attachments *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewGETPriceTiersPriceTierId200ResponseDataRelationships instantiates a new GETPriceTiersPriceTierId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceTiersPriceTierId200ResponseDataRelationships() *GETPriceTiersPriceTierId200ResponseDataRelationships {
	this := GETPriceTiersPriceTierId200ResponseDataRelationships{}
	return &this
}

// NewGETPriceTiersPriceTierId200ResponseDataRelationshipsWithDefaults instantiates a new GETPriceTiersPriceTierId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceTiersPriceTierId200ResponseDataRelationshipsWithDefaults() *GETPriceTiersPriceTierId200ResponseDataRelationships {
	this := GETPriceTiersPriceTierId200ResponseDataRelationships{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) GetPrice() POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice {
	if o == nil || IsNil(o.Price) {
		var ret POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) GetPriceOk() (*POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice and assigns it to the Price field.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) SetPrice(v POSTPriceFrequencyTiers201ResponseDataRelationshipsPrice) {
	o.Price = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *GETPriceTiersPriceTierId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o GETPriceTiersPriceTierId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPriceTiersPriceTierId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
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

type NullableGETPriceTiersPriceTierId200ResponseDataRelationships struct {
	value *GETPriceTiersPriceTierId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETPriceTiersPriceTierId200ResponseDataRelationships) Get() *GETPriceTiersPriceTierId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETPriceTiersPriceTierId200ResponseDataRelationships) Set(val *GETPriceTiersPriceTierId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceTiersPriceTierId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceTiersPriceTierId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceTiersPriceTierId200ResponseDataRelationships(val *GETPriceTiersPriceTierId200ResponseDataRelationships) *NullableGETPriceTiersPriceTierId200ResponseDataRelationships {
	return &NullableGETPriceTiersPriceTierId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETPriceTiersPriceTierId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceTiersPriceTierId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
