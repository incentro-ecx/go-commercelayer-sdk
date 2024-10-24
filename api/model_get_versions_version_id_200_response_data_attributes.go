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

// checks if the GETVersionsVersionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETVersionsVersionId200ResponseDataAttributes{}

// GETVersionsVersionId200ResponseDataAttributes struct for GETVersionsVersionId200ResponseDataAttributes
type GETVersionsVersionId200ResponseDataAttributes struct {
	// The type of the versioned resource.
	ResourceType interface{} `json:"resource_type,omitempty"`
	// The versioned resource id.
	ResourceId interface{} `json:"resource_id,omitempty"`
	// The event which generates the version.
	Event interface{} `json:"event,omitempty"`
	// The object changes payload.
	Changes interface{} `json:"changes,omitempty"`
	// Information about who triggered the change.
	Who interface{} `json:"who,omitempty"`
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

// NewGETVersionsVersionId200ResponseDataAttributes instantiates a new GETVersionsVersionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETVersionsVersionId200ResponseDataAttributes() *GETVersionsVersionId200ResponseDataAttributes {
	this := GETVersionsVersionId200ResponseDataAttributes{}
	return &this
}

// NewGETVersionsVersionId200ResponseDataAttributesWithDefaults instantiates a new GETVersionsVersionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETVersionsVersionId200ResponseDataAttributesWithDefaults() *GETVersionsVersionId200ResponseDataAttributes {
	this := GETVersionsVersionId200ResponseDataAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetResourceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return &o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasResourceType() bool {
	if o != nil && IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given interface{} and assigns it to the ResourceType field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetResourceType(v interface{}) {
	o.ResourceType = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetResourceId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetResourceIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return &o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasResourceId() bool {
	if o != nil && IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given interface{} and assigns it to the ResourceId field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetResourceId(v interface{}) {
	o.ResourceId = v
}

// GetEvent returns the Event field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetEvent() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetEventOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return &o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasEvent() bool {
	if o != nil && IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given interface{} and assigns it to the Event field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetEvent(v interface{}) {
	o.Event = v
}

// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetChanges() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetChangesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return &o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasChanges() bool {
	if o != nil && IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given interface{} and assigns it to the Changes field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetChanges(v interface{}) {
	o.Changes = v
}

// GetWho returns the Who field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetWho() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Who
}

// GetWhoOk returns a tuple with the Who field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetWhoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Who) {
		return nil, false
	}
	return &o.Who, true
}

// HasWho returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasWho() bool {
	if o != nil && IsNil(o.Who) {
		return true
	}

	return false
}

// SetWho gets a reference to the given interface{} and assigns it to the Who field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetWho(v interface{}) {
	o.Who = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETVersionsVersionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETVersionsVersionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETVersionsVersionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETVersionsVersionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETVersionsVersionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETVersionsVersionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.Who != nil {
		toSerialize["who"] = o.Who
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

type NullableGETVersionsVersionId200ResponseDataAttributes struct {
	value *GETVersionsVersionId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETVersionsVersionId200ResponseDataAttributes) Get() *GETVersionsVersionId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETVersionsVersionId200ResponseDataAttributes) Set(val *GETVersionsVersionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETVersionsVersionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETVersionsVersionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETVersionsVersionId200ResponseDataAttributes(val *GETVersionsVersionId200ResponseDataAttributes) *NullableGETVersionsVersionId200ResponseDataAttributes {
	return &NullableGETVersionsVersionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETVersionsVersionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETVersionsVersionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}