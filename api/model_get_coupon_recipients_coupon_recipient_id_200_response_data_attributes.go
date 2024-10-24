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

// checks if the GETCouponRecipientsCouponRecipientId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCouponRecipientsCouponRecipientId200ResponseDataAttributes{}

// GETCouponRecipientsCouponRecipientId200ResponseDataAttributes struct for GETCouponRecipientsCouponRecipientId200ResponseDataAttributes
type GETCouponRecipientsCouponRecipientId200ResponseDataAttributes struct {
	// The recipient email address.
	Email interface{} `json:"email,omitempty"`
	// The recipient first name.
	FirstName interface{} `json:"first_name,omitempty"`
	// The recipient last name.
	LastName interface{} `json:"last_name,omitempty"`
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

// NewGETCouponRecipientsCouponRecipientId200ResponseDataAttributes instantiates a new GETCouponRecipientsCouponRecipientId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponRecipientsCouponRecipientId200ResponseDataAttributes() *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	this := GETCouponRecipientsCouponRecipientId200ResponseDataAttributes{}
	return &this
}

// NewGETCouponRecipientsCouponRecipientId200ResponseDataAttributesWithDefaults instantiates a new GETCouponRecipientsCouponRecipientId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponRecipientsCouponRecipientId200ResponseDataAttributesWithDefaults() *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	this := GETCouponRecipientsCouponRecipientId200ResponseDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetEmail(v interface{}) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetFirstName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetFirstNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return &o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasFirstName() bool {
	if o != nil && IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given interface{} and assigns it to the FirstName field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetFirstName(v interface{}) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetLastName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetLastNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return &o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasLastName() bool {
	if o != nil && IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given interface{} and assigns it to the LastName field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetLastName(v interface{}) {
	o.LastName = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
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

type NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes struct {
	value *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes) Get() *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes) Set(val *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes(val *GETCouponRecipientsCouponRecipientId200ResponseDataAttributes) *NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	return &NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponRecipientsCouponRecipientId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
