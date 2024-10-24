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

// checks if the GETStockReservationsStockReservationId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETStockReservationsStockReservationId200ResponseDataAttributes{}

// GETStockReservationsStockReservationId200ResponseDataAttributes struct for GETStockReservationsStockReservationId200ResponseDataAttributes
type GETStockReservationsStockReservationId200ResponseDataAttributes struct {
	// The stock reservation status. One of 'draft' (default), or 'pending'.
	Status interface{} `json:"status,omitempty"`
	// The stock reservation quantity.
	Quantity interface{} `json:"quantity,omitempty"`
	// The expiration date/time of this stock reservation.
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETStockReservationsStockReservationId200ResponseDataAttributes instantiates a new GETStockReservationsStockReservationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockReservationsStockReservationId200ResponseDataAttributes() *GETStockReservationsStockReservationId200ResponseDataAttributes {
	this := GETStockReservationsStockReservationId200ResponseDataAttributes{}
	return &this
}

// NewGETStockReservationsStockReservationId200ResponseDataAttributesWithDefaults instantiates a new GETStockReservationsStockReservationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockReservationsStockReservationId200ResponseDataAttributesWithDefaults() *GETStockReservationsStockReservationId200ResponseDataAttributes {
	this := GETStockReservationsStockReservationId200ResponseDataAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETStockReservationsStockReservationId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETStockReservationsStockReservationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETStockReservationsStockReservationId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETStockReservationsStockReservationId200ResponseDataAttributes struct {
	value *GETStockReservationsStockReservationId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETStockReservationsStockReservationId200ResponseDataAttributes) Get() *GETStockReservationsStockReservationId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETStockReservationsStockReservationId200ResponseDataAttributes) Set(val *GETStockReservationsStockReservationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockReservationsStockReservationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockReservationsStockReservationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockReservationsStockReservationId200ResponseDataAttributes(val *GETStockReservationsStockReservationId200ResponseDataAttributes) *NullableGETStockReservationsStockReservationId200ResponseDataAttributes {
	return &NullableGETStockReservationsStockReservationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETStockReservationsStockReservationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockReservationsStockReservationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}