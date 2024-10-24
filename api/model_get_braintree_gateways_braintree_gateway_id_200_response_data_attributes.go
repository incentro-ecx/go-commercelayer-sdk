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

// checks if the GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes{}

// GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes struct for GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes
type GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes struct {
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
	// The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars.
	DescriptorName interface{} `json:"descriptor_name,omitempty"`
	// The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods.
	DescriptorPhone interface{} `json:"descriptor_phone,omitempty"`
	// The dynamic descriptor URL.
	DescriptorUrl interface{} `json:"descriptor_url,omitempty"`
	// The gateway webhook URL, generated automatically.
	WebhookEndpointUrl interface{} `json:"webhook_endpoint_url,omitempty"`
}

// NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes instantiates a new GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes() *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	this := GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes{}
	return &this
}

// NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults instantiates a new GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributesWithDefaults() *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	this := GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetDescriptorName returns the DescriptorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DescriptorName
}

// GetDescriptorNameOk returns a tuple with the DescriptorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DescriptorName) {
		return nil, false
	}
	return &o.DescriptorName, true
}

// HasDescriptorName returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorName() bool {
	if o != nil && IsNil(o.DescriptorName) {
		return true
	}

	return false
}

// SetDescriptorName gets a reference to the given interface{} and assigns it to the DescriptorName field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorName(v interface{}) {
	o.DescriptorName = v
}

// GetDescriptorPhone returns the DescriptorPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhone() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DescriptorPhone
}

// GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorPhoneOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DescriptorPhone) {
		return nil, false
	}
	return &o.DescriptorPhone, true
}

// HasDescriptorPhone returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorPhone() bool {
	if o != nil && IsNil(o.DescriptorPhone) {
		return true
	}

	return false
}

// SetDescriptorPhone gets a reference to the given interface{} and assigns it to the DescriptorPhone field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorPhone(v interface{}) {
	o.DescriptorPhone = v
}

// GetDescriptorUrl returns the DescriptorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DescriptorUrl
}

// GetDescriptorUrlOk returns a tuple with the DescriptorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetDescriptorUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DescriptorUrl) {
		return nil, false
	}
	return &o.DescriptorUrl, true
}

// HasDescriptorUrl returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasDescriptorUrl() bool {
	if o != nil && IsNil(o.DescriptorUrl) {
		return true
	}

	return false
}

// SetDescriptorUrl gets a reference to the given interface{} and assigns it to the DescriptorUrl field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetDescriptorUrl(v interface{}) {
	o.DescriptorUrl = v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetWebhookEndpointUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) GetWebhookEndpointUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		return nil, false
	}
	return &o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) HasWebhookEndpointUrl() bool {
	if o != nil && IsNil(o.WebhookEndpointUrl) {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given interface{} and assigns it to the WebhookEndpointUrl field.
func (o *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) SetWebhookEndpointUrl(v interface{}) {
	o.WebhookEndpointUrl = v
}

func (o GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.DescriptorName != nil {
		toSerialize["descriptor_name"] = o.DescriptorName
	}
	if o.DescriptorPhone != nil {
		toSerialize["descriptor_phone"] = o.DescriptorPhone
	}
	if o.DescriptorUrl != nil {
		toSerialize["descriptor_url"] = o.DescriptorUrl
	}
	if o.WebhookEndpointUrl != nil {
		toSerialize["webhook_endpoint_url"] = o.WebhookEndpointUrl
	}
	return toSerialize, nil
}

type NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes struct {
	value *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) Get() *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) Set(val *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes(val *GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) *NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	return &NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
