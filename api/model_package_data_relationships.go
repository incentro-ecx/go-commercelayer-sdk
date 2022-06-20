/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PackageDataRelationships struct for PackageDataRelationships
type PackageDataRelationships struct {
	StockLocation *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Parcels       *PackageDataRelationshipsParcels                `json:"parcels,omitempty"`
	Attachments   *AvalaraAccountDataRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewPackageDataRelationships instantiates a new PackageDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDataRelationships() *PackageDataRelationships {
	this := PackageDataRelationships{}
	return &this
}

// NewPackageDataRelationshipsWithDefaults instantiates a new PackageDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDataRelationshipsWithDefaults() *PackageDataRelationships {
	this := PackageDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *PackageDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *PackageDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *PackageDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetParcels returns the Parcels field value if set, zero value otherwise.
func (o *PackageDataRelationships) GetParcels() PackageDataRelationshipsParcels {
	if o == nil || o.Parcels == nil {
		var ret PackageDataRelationshipsParcels
		return ret
	}
	return *o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDataRelationships) GetParcelsOk() (*PackageDataRelationshipsParcels, bool) {
	if o == nil || o.Parcels == nil {
		return nil, false
	}
	return o.Parcels, true
}

// HasParcels returns a boolean if a field has been set.
func (o *PackageDataRelationships) HasParcels() bool {
	if o != nil && o.Parcels != nil {
		return true
	}

	return false
}

// SetParcels gets a reference to the given PackageDataRelationshipsParcels and assigns it to the Parcels field.
func (o *PackageDataRelationships) SetParcels(v PackageDataRelationshipsParcels) {
	o.Parcels = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PackageDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PackageDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *PackageDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o PackageDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.Parcels != nil {
		toSerialize["parcels"] = o.Parcels
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullablePackageDataRelationships struct {
	value *PackageDataRelationships
	isSet bool
}

func (v NullablePackageDataRelationships) Get() *PackageDataRelationships {
	return v.value
}

func (v *NullablePackageDataRelationships) Set(val *PackageDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDataRelationships(val *PackageDataRelationships) *NullablePackageDataRelationships {
	return &NullablePackageDataRelationships{value: val, isSet: true}
}

func (v NullablePackageDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}