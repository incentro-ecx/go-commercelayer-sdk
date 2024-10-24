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

// checks if the InventoryModelDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryModelDataRelationships{}

// InventoryModelDataRelationships struct for InventoryModelDataRelationships
type InventoryModelDataRelationships struct {
	InventoryStockLocations  *InventoryModelDataRelationshipsInventoryStockLocations  `json:"inventory_stock_locations,omitempty"`
	InventoryReturnLocations *InventoryModelDataRelationshipsInventoryReturnLocations `json:"inventory_return_locations,omitempty"`
	Attachments              *AuthorizationDataRelationshipsAttachments               `json:"attachments,omitempty"`
	Versions                 *AddressDataRelationshipsVersions                        `json:"versions,omitempty"`
}

// NewInventoryModelDataRelationships instantiates a new InventoryModelDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelDataRelationships() *InventoryModelDataRelationships {
	this := InventoryModelDataRelationships{}
	return &this
}

// NewInventoryModelDataRelationshipsWithDefaults instantiates a new InventoryModelDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelDataRelationshipsWithDefaults() *InventoryModelDataRelationships {
	this := InventoryModelDataRelationships{}
	return &this
}

// GetInventoryStockLocations returns the InventoryStockLocations field value if set, zero value otherwise.
func (o *InventoryModelDataRelationships) GetInventoryStockLocations() InventoryModelDataRelationshipsInventoryStockLocations {
	if o == nil || IsNil(o.InventoryStockLocations) {
		var ret InventoryModelDataRelationshipsInventoryStockLocations
		return ret
	}
	return *o.InventoryStockLocations
}

// GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelDataRelationships) GetInventoryStockLocationsOk() (*InventoryModelDataRelationshipsInventoryStockLocations, bool) {
	if o == nil || IsNil(o.InventoryStockLocations) {
		return nil, false
	}
	return o.InventoryStockLocations, true
}

// HasInventoryStockLocations returns a boolean if a field has been set.
func (o *InventoryModelDataRelationships) HasInventoryStockLocations() bool {
	if o != nil && !IsNil(o.InventoryStockLocations) {
		return true
	}

	return false
}

// SetInventoryStockLocations gets a reference to the given InventoryModelDataRelationshipsInventoryStockLocations and assigns it to the InventoryStockLocations field.
func (o *InventoryModelDataRelationships) SetInventoryStockLocations(v InventoryModelDataRelationshipsInventoryStockLocations) {
	o.InventoryStockLocations = &v
}

// GetInventoryReturnLocations returns the InventoryReturnLocations field value if set, zero value otherwise.
func (o *InventoryModelDataRelationships) GetInventoryReturnLocations() InventoryModelDataRelationshipsInventoryReturnLocations {
	if o == nil || IsNil(o.InventoryReturnLocations) {
		var ret InventoryModelDataRelationshipsInventoryReturnLocations
		return ret
	}
	return *o.InventoryReturnLocations
}

// GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelDataRelationships) GetInventoryReturnLocationsOk() (*InventoryModelDataRelationshipsInventoryReturnLocations, bool) {
	if o == nil || IsNil(o.InventoryReturnLocations) {
		return nil, false
	}
	return o.InventoryReturnLocations, true
}

// HasInventoryReturnLocations returns a boolean if a field has been set.
func (o *InventoryModelDataRelationships) HasInventoryReturnLocations() bool {
	if o != nil && !IsNil(o.InventoryReturnLocations) {
		return true
	}

	return false
}

// SetInventoryReturnLocations gets a reference to the given InventoryModelDataRelationshipsInventoryReturnLocations and assigns it to the InventoryReturnLocations field.
func (o *InventoryModelDataRelationships) SetInventoryReturnLocations(v InventoryModelDataRelationshipsInventoryReturnLocations) {
	o.InventoryReturnLocations = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *InventoryModelDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *InventoryModelDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *InventoryModelDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *InventoryModelDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryModelDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *InventoryModelDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *InventoryModelDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o InventoryModelDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryModelDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InventoryStockLocations) {
		toSerialize["inventory_stock_locations"] = o.InventoryStockLocations
	}
	if !IsNil(o.InventoryReturnLocations) {
		toSerialize["inventory_return_locations"] = o.InventoryReturnLocations
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableInventoryModelDataRelationships struct {
	value *InventoryModelDataRelationships
	isSet bool
}

func (v NullableInventoryModelDataRelationships) Get() *InventoryModelDataRelationships {
	return v.value
}

func (v *NullableInventoryModelDataRelationships) Set(val *InventoryModelDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelDataRelationships(val *InventoryModelDataRelationships) *NullableInventoryModelDataRelationships {
	return &NullableInventoryModelDataRelationships{value: val, isSet: true}
}

func (v NullableInventoryModelDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
