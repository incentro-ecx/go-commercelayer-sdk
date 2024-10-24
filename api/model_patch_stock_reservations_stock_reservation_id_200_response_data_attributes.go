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

// checks if the PATCHStockReservationsStockReservationId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHStockReservationsStockReservationId200ResponseDataAttributes{}

// PATCHStockReservationsStockReservationId200ResponseDataAttributes struct for PATCHStockReservationsStockReservationId200ResponseDataAttributes
type PATCHStockReservationsStockReservationId200ResponseDataAttributes struct {
	// The stock reservation quantity.
	Quantity interface{} `json:"quantity,omitempty"`
	// Send this attribute if you want to mark this stock reservation as pending.
	Pending interface{} `json:"_pending,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHStockReservationsStockReservationId200ResponseDataAttributes instantiates a new PATCHStockReservationsStockReservationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStockReservationsStockReservationId200ResponseDataAttributes() *PATCHStockReservationsStockReservationId200ResponseDataAttributes {
	this := PATCHStockReservationsStockReservationId200ResponseDataAttributes{}
	return &this
}

// NewPATCHStockReservationsStockReservationId200ResponseDataAttributesWithDefaults instantiates a new PATCHStockReservationsStockReservationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStockReservationsStockReservationId200ResponseDataAttributesWithDefaults() *PATCHStockReservationsStockReservationId200ResponseDataAttributes {
	this := PATCHStockReservationsStockReservationId200ResponseDataAttributes{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetPending returns the Pending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetPending() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetPendingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return &o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) HasPending() bool {
	if o != nil && IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given interface{} and assigns it to the Pending field.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) SetPending(v interface{}) {
	o.Pending = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHStockReservationsStockReservationId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHStockReservationsStockReservationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHStockReservationsStockReservationId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Pending != nil {
		toSerialize["_pending"] = o.Pending
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

type NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes struct {
	value *PATCHStockReservationsStockReservationId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes) Get() *PATCHStockReservationsStockReservationId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes) Set(val *PATCHStockReservationsStockReservationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStockReservationsStockReservationId200ResponseDataAttributes(val *PATCHStockReservationsStockReservationId200ResponseDataAttributes) *NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes {
	return &NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStockReservationsStockReservationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
