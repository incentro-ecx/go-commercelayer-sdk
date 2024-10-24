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

// checks if the POSTSkus201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkus201ResponseDataAttributes{}

// POSTSkus201ResponseDataAttributes struct for POSTSkus201ResponseDataAttributes
type POSTSkus201ResponseDataAttributes struct {
	// The SKU code, that uniquely identifies the SKU within the organization.
	Code interface{} `json:"code"`
	// The internal name of the SKU.
	Name interface{} `json:"name"`
	// An internal description of the SKU.
	Description interface{} `json:"description,omitempty"`
	// The URL of an image that represents the SKU.
	ImageUrl interface{} `json:"image_url,omitempty"`
	// The number of pieces that compose the SKU. This is useful to describe sets and bundles.
	PiecesPerPack interface{} `json:"pieces_per_pack,omitempty"`
	// The weight of the SKU. If present, it will be used to calculate the shipping rates.
	Weight interface{} `json:"weight,omitempty"`
	// The unit of weight. One of 'gr', 'oz', or 'lb'.
	UnitOfWeight interface{} `json:"unit_of_weight,omitempty"`
	// The Harmonized System Code used by customs to identify the products shipped across international borders.
	HsTariffNumber interface{} `json:"hs_tariff_number,omitempty"`
	// Indicates if the SKU doesn't generate shipments.
	DoNotShip interface{} `json:"do_not_ship,omitempty"`
	// Indicates if the SKU doesn't track the stock inventory.
	DoNotTrack interface{} `json:"do_not_track,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTSkus201ResponseDataAttributes instantiates a new POSTSkus201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkus201ResponseDataAttributes(code interface{}, name interface{}) *POSTSkus201ResponseDataAttributes {
	this := POSTSkus201ResponseDataAttributes{}
	this.Code = code
	this.Name = name
	return &this
}

// NewPOSTSkus201ResponseDataAttributesWithDefaults instantiates a new POSTSkus201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkus201ResponseDataAttributesWithDefaults() *POSTSkus201ResponseDataAttributes {
	this := POSTSkus201ResponseDataAttributes{}
	return &this
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTSkus201ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *POSTSkus201ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTSkus201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTSkus201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *POSTSkus201ResponseDataAttributes) SetDescription(v interface{}) {
	o.Description = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *POSTSkus201ResponseDataAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetPiecesPerPack returns the PiecesPerPack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetPiecesPerPack() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PiecesPerPack
}

// GetPiecesPerPackOk returns a tuple with the PiecesPerPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetPiecesPerPackOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PiecesPerPack) {
		return nil, false
	}
	return &o.PiecesPerPack, true
}

// HasPiecesPerPack returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasPiecesPerPack() bool {
	if o != nil && IsNil(o.PiecesPerPack) {
		return true
	}

	return false
}

// SetPiecesPerPack gets a reference to the given interface{} and assigns it to the PiecesPerPack field.
func (o *POSTSkus201ResponseDataAttributes) SetPiecesPerPack(v interface{}) {
	o.PiecesPerPack = v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return &o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasWeight() bool {
	if o != nil && IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given interface{} and assigns it to the Weight field.
func (o *POSTSkus201ResponseDataAttributes) SetWeight(v interface{}) {
	o.Weight = v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetUnitOfWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitOfWeight) {
		return nil, false
	}
	return &o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasUnitOfWeight() bool {
	if o != nil && IsNil(o.UnitOfWeight) {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given interface{} and assigns it to the UnitOfWeight field.
func (o *POSTSkus201ResponseDataAttributes) SetUnitOfWeight(v interface{}) {
	o.UnitOfWeight = v
}

// GetHsTariffNumber returns the HsTariffNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetHsTariffNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.HsTariffNumber
}

// GetHsTariffNumberOk returns a tuple with the HsTariffNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetHsTariffNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HsTariffNumber) {
		return nil, false
	}
	return &o.HsTariffNumber, true
}

// HasHsTariffNumber returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasHsTariffNumber() bool {
	if o != nil && IsNil(o.HsTariffNumber) {
		return true
	}

	return false
}

// SetHsTariffNumber gets a reference to the given interface{} and assigns it to the HsTariffNumber field.
func (o *POSTSkus201ResponseDataAttributes) SetHsTariffNumber(v interface{}) {
	o.HsTariffNumber = v
}

// GetDoNotShip returns the DoNotShip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetDoNotShip() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DoNotShip
}

// GetDoNotShipOk returns a tuple with the DoNotShip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetDoNotShipOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DoNotShip) {
		return nil, false
	}
	return &o.DoNotShip, true
}

// HasDoNotShip returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasDoNotShip() bool {
	if o != nil && IsNil(o.DoNotShip) {
		return true
	}

	return false
}

// SetDoNotShip gets a reference to the given interface{} and assigns it to the DoNotShip field.
func (o *POSTSkus201ResponseDataAttributes) SetDoNotShip(v interface{}) {
	o.DoNotShip = v
}

// GetDoNotTrack returns the DoNotTrack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetDoNotTrack() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DoNotTrack
}

// GetDoNotTrackOk returns a tuple with the DoNotTrack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetDoNotTrackOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DoNotTrack) {
		return nil, false
	}
	return &o.DoNotTrack, true
}

// HasDoNotTrack returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasDoNotTrack() bool {
	if o != nil && IsNil(o.DoNotTrack) {
		return true
	}

	return false
}

// SetDoNotTrack gets a reference to the given interface{} and assigns it to the DoNotTrack field.
func (o *POSTSkus201ResponseDataAttributes) SetDoNotTrack(v interface{}) {
	o.DoNotTrack = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTSkus201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTSkus201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkus201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkus201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTSkus201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTSkus201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTSkus201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkus201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.PiecesPerPack != nil {
		toSerialize["pieces_per_pack"] = o.PiecesPerPack
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
	}
	if o.HsTariffNumber != nil {
		toSerialize["hs_tariff_number"] = o.HsTariffNumber
	}
	if o.DoNotShip != nil {
		toSerialize["do_not_ship"] = o.DoNotShip
	}
	if o.DoNotTrack != nil {
		toSerialize["do_not_track"] = o.DoNotTrack
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePOSTSkus201ResponseDataAttributes struct {
	value *POSTSkus201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTSkus201ResponseDataAttributes) Get() *POSTSkus201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTSkus201ResponseDataAttributes) Set(val *POSTSkus201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkus201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkus201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkus201ResponseDataAttributes(val *POSTSkus201ResponseDataAttributes) *NullablePOSTSkus201ResponseDataAttributes {
	return &NullablePOSTSkus201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTSkus201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkus201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
