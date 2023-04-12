/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETEventCallbacksEventCallbackId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETEventCallbacksEventCallbackId200ResponseDataAttributes{}

// GETEventCallbacksEventCallbackId200ResponseDataAttributes struct for GETEventCallbacksEventCallbackId200ResponseDataAttributes
type GETEventCallbacksEventCallbackId200ResponseDataAttributes struct {
	// The URI of the callback, inherited by the associated webhook.
	CallbackUrl interface{} `json:"callback_url,omitempty"`
	// The payload sent to the callback endpoint, including the event affected resource and the specified includes.
	Payload interface{} `json:"payload,omitempty"`
	// The HTTP response code of the callback endpoint.
	ResponseCode interface{} `json:"response_code,omitempty"`
	// The HTTP response message of the callback endpoint.
	ResponseMessage interface{} `json:"response_message,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETEventCallbacksEventCallbackId200ResponseDataAttributes instantiates a new GETEventCallbacksEventCallbackId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacksEventCallbackId200ResponseDataAttributes() *GETEventCallbacksEventCallbackId200ResponseDataAttributes {
	this := GETEventCallbacksEventCallbackId200ResponseDataAttributes{}
	return &this
}

// NewGETEventCallbacksEventCallbackId200ResponseDataAttributesWithDefaults instantiates a new GETEventCallbacksEventCallbackId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacksEventCallbackId200ResponseDataAttributesWithDefaults() *GETEventCallbacksEventCallbackId200ResponseDataAttributes {
	this := GETEventCallbacksEventCallbackId200ResponseDataAttributes{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetCallbackUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetCallbackUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasCallbackUrl() bool {
	if o != nil && IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given interface{} and assigns it to the CallbackUrl field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetCallbackUrl(v interface{}) {
	o.CallbackUrl = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetPayload() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetPayloadOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasPayload() bool {
	if o != nil && IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given interface{} and assigns it to the Payload field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetPayload(v interface{}) {
	o.Payload = v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetResponseCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetResponseCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResponseCode) {
		return nil, false
	}
	return &o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasResponseCode() bool {
	if o != nil && IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given interface{} and assigns it to the ResponseCode field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetResponseCode(v interface{}) {
	o.ResponseCode = v
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetResponseMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetResponseMessageOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResponseMessage) {
		return nil, false
	}
	return &o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasResponseMessage() bool {
	if o != nil && IsNil(o.ResponseMessage) {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given interface{} and assigns it to the ResponseMessage field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetResponseMessage(v interface{}) {
	o.ResponseMessage = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETEventCallbacksEventCallbackId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETEventCallbacksEventCallbackId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETEventCallbacksEventCallbackId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.ResponseCode != nil {
		toSerialize["response_code"] = o.ResponseCode
	}
	if o.ResponseMessage != nil {
		toSerialize["response_message"] = o.ResponseMessage
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

type NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes struct {
	value *GETEventCallbacksEventCallbackId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes) Get() *GETEventCallbacksEventCallbackId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes) Set(val *GETEventCallbacksEventCallbackId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacksEventCallbackId200ResponseDataAttributes(val *GETEventCallbacksEventCallbackId200ResponseDataAttributes) *NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes {
	return &NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacksEventCallbackId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
