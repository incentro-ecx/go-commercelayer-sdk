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

// checks if the PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}

// PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct for PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes
type PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct {
	// The customer new password. This will be accepted only if a valid '_reset_password_token' is sent with the request.
	CustomerPassword interface{} `json:"customer_password,omitempty"`
	// Send the 'reset_password_token' that you got on create when updating the customer password.
	ResetPasswordToken interface{} `json:"_reset_password_token,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	this := PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	this := PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}
	return &this
}

// GetCustomerPassword returns the CustomerPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPassword() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerPassword
}

// GetCustomerPasswordOk returns a tuple with the CustomerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPasswordOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerPassword) {
		return nil, false
	}
	return &o.CustomerPassword, true
}

// HasCustomerPassword returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasCustomerPassword() bool {
	if o != nil && IsNil(o.CustomerPassword) {
		return true
	}

	return false
}

// SetCustomerPassword gets a reference to the given interface{} and assigns it to the CustomerPassword field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCustomerPassword(v interface{}) {
	o.CustomerPassword = v
}

// GetResetPasswordToken returns the ResetPasswordToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetPasswordToken
}

// GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetPasswordToken) {
		return nil, false
	}
	return &o.ResetPasswordToken, true
}

// HasResetPasswordToken returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasResetPasswordToken() bool {
	if o != nil && IsNil(o.ResetPasswordToken) {
		return true
	}

	return false
}

// SetResetPasswordToken gets a reference to the given interface{} and assigns it to the ResetPasswordToken field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordToken(v interface{}) {
	o.ResetPasswordToken = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerPassword != nil {
		toSerialize["customer_password"] = o.CustomerPassword
	}
	if o.ResetPasswordToken != nil {
		toSerialize["_reset_password_token"] = o.ResetPasswordToken
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

type NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct {
	value *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Get() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Set(val *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes(val *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	return &NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
