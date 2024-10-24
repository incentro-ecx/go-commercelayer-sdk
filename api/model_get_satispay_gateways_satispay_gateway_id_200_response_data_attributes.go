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

// checks if the GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes{}

// GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes struct for GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes
type GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
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
	// Activation code generated from the Satispay Dashboard.
	Token interface{} `json:"token,omitempty"`
	// The Satispay API key auto generated basing on activation code.
	KeyId interface{} `json:"key_id,omitempty"`
	// The gateway webhook URL, generated automatically.
	WebhookEndpointUrl interface{} `json:"webhook_endpoint_url,omitempty"`
}

// NewGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes instantiates a new GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes() *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes {
	this := GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes{}
	return &this
}

// NewGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributesWithDefaults() *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes {
	this := GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return &o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasToken() bool {
	if o != nil && IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given interface{} and assigns it to the Token field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetToken(v interface{}) {
	o.Token = v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetKeyId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetKeyIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return &o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasKeyId() bool {
	if o != nil && IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given interface{} and assigns it to the KeyId field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetKeyId(v interface{}) {
	o.KeyId = v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetWebhookEndpointUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) GetWebhookEndpointUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		return nil, false
	}
	return &o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) HasWebhookEndpointUrl() bool {
	if o != nil && IsNil(o.WebhookEndpointUrl) {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given interface{} and assigns it to the WebhookEndpointUrl field.
func (o *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) SetWebhookEndpointUrl(v interface{}) {
	o.WebhookEndpointUrl = v
}

func (o GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.KeyId != nil {
		toSerialize["key_id"] = o.KeyId
	}
	if o.WebhookEndpointUrl != nil {
		toSerialize["webhook_endpoint_url"] = o.WebhookEndpointUrl
	}
	return toSerialize, nil
}

type NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes struct {
	value *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) Get() *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) Set(val *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes(val *GETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) *NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes {
	return &NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSatispayGatewaysSatispayGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
