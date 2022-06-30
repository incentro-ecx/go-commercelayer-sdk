/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalGatewayDataAttributes struct for ExternalGatewayDataAttributes
type ExternalGatewayDataAttributes struct {
	// The payment gateway's internal name.
	Name *string `json:"name,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The shared secret used to sign gateway payloads.
	SharedSecret *string `json:"shared_secret,omitempty"`
	// The endpoint used by the external gateway to authorize payments.
	AuthorizeUrl *string `json:"authorize_url,omitempty"`
	// The endpoint used by the external gateway to capture payments.
	CaptureUrl *string `json:"capture_url,omitempty"`
	// The endpoint used by the external gateway to void payments.
	VoidUrl *string `json:"void_url,omitempty"`
	// The endpoint used by the external gateway to refund payments.
	RefundUrl *string `json:"refund_url,omitempty"`
	// The endpoint used by the external gateway to create a customer payment token.
	TokenUrl *string `json:"token_url,omitempty"`
}

// NewExternalGatewayDataAttributes instantiates a new ExternalGatewayDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGatewayDataAttributes() *ExternalGatewayDataAttributes {
	this := ExternalGatewayDataAttributes{}
	return &this
}

// NewExternalGatewayDataAttributesWithDefaults instantiates a new ExternalGatewayDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGatewayDataAttributesWithDefaults() *ExternalGatewayDataAttributes {
	this := ExternalGatewayDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalGatewayDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalGatewayDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ExternalGatewayDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ExternalGatewayDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ExternalGatewayDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ExternalGatewayDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ExternalGatewayDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *ExternalGatewayDataAttributes) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetAuthorizeUrl returns the AuthorizeUrl field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetAuthorizeUrl() string {
	if o == nil || o.AuthorizeUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthorizeUrl
}

// GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetAuthorizeUrlOk() (*string, bool) {
	if o == nil || o.AuthorizeUrl == nil {
		return nil, false
	}
	return o.AuthorizeUrl, true
}

// HasAuthorizeUrl returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasAuthorizeUrl() bool {
	if o != nil && o.AuthorizeUrl != nil {
		return true
	}

	return false
}

// SetAuthorizeUrl gets a reference to the given string and assigns it to the AuthorizeUrl field.
func (o *ExternalGatewayDataAttributes) SetAuthorizeUrl(v string) {
	o.AuthorizeUrl = &v
}

// GetCaptureUrl returns the CaptureUrl field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetCaptureUrl() string {
	if o == nil || o.CaptureUrl == nil {
		var ret string
		return ret
	}
	return *o.CaptureUrl
}

// GetCaptureUrlOk returns a tuple with the CaptureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetCaptureUrlOk() (*string, bool) {
	if o == nil || o.CaptureUrl == nil {
		return nil, false
	}
	return o.CaptureUrl, true
}

// HasCaptureUrl returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasCaptureUrl() bool {
	if o != nil && o.CaptureUrl != nil {
		return true
	}

	return false
}

// SetCaptureUrl gets a reference to the given string and assigns it to the CaptureUrl field.
func (o *ExternalGatewayDataAttributes) SetCaptureUrl(v string) {
	o.CaptureUrl = &v
}

// GetVoidUrl returns the VoidUrl field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetVoidUrl() string {
	if o == nil || o.VoidUrl == nil {
		var ret string
		return ret
	}
	return *o.VoidUrl
}

// GetVoidUrlOk returns a tuple with the VoidUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetVoidUrlOk() (*string, bool) {
	if o == nil || o.VoidUrl == nil {
		return nil, false
	}
	return o.VoidUrl, true
}

// HasVoidUrl returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasVoidUrl() bool {
	if o != nil && o.VoidUrl != nil {
		return true
	}

	return false
}

// SetVoidUrl gets a reference to the given string and assigns it to the VoidUrl field.
func (o *ExternalGatewayDataAttributes) SetVoidUrl(v string) {
	o.VoidUrl = &v
}

// GetRefundUrl returns the RefundUrl field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetRefundUrl() string {
	if o == nil || o.RefundUrl == nil {
		var ret string
		return ret
	}
	return *o.RefundUrl
}

// GetRefundUrlOk returns a tuple with the RefundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetRefundUrlOk() (*string, bool) {
	if o == nil || o.RefundUrl == nil {
		return nil, false
	}
	return o.RefundUrl, true
}

// HasRefundUrl returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasRefundUrl() bool {
	if o != nil && o.RefundUrl != nil {
		return true
	}

	return false
}

// SetRefundUrl gets a reference to the given string and assigns it to the RefundUrl field.
func (o *ExternalGatewayDataAttributes) SetRefundUrl(v string) {
	o.RefundUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *ExternalGatewayDataAttributes) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataAttributes) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *ExternalGatewayDataAttributes) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *ExternalGatewayDataAttributes) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

func (o ExternalGatewayDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	if o.AuthorizeUrl != nil {
		toSerialize["authorize_url"] = o.AuthorizeUrl
	}
	if o.CaptureUrl != nil {
		toSerialize["capture_url"] = o.CaptureUrl
	}
	if o.VoidUrl != nil {
		toSerialize["void_url"] = o.VoidUrl
	}
	if o.RefundUrl != nil {
		toSerialize["refund_url"] = o.RefundUrl
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGatewayDataAttributes struct {
	value *ExternalGatewayDataAttributes
	isSet bool
}

func (v NullableExternalGatewayDataAttributes) Get() *ExternalGatewayDataAttributes {
	return v.value
}

func (v *NullableExternalGatewayDataAttributes) Set(val *ExternalGatewayDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGatewayDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGatewayDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGatewayDataAttributes(val *ExternalGatewayDataAttributes) *NullableExternalGatewayDataAttributes {
	return &NullableExternalGatewayDataAttributes{value: val, isSet: true}
}

func (v NullableExternalGatewayDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGatewayDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
