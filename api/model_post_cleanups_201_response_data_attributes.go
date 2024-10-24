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

// checks if the POSTCleanups201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCleanups201ResponseDataAttributes{}

// POSTCleanups201ResponseDataAttributes struct for POSTCleanups201ResponseDataAttributes
type POSTCleanups201ResponseDataAttributes struct {
	// The type of resource being cleaned.
	ResourceType interface{} `json:"resource_type"`
	// The filters used to select the records to be cleaned.
	Filters interface{} `json:"filters,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTCleanups201ResponseDataAttributes instantiates a new POSTCleanups201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCleanups201ResponseDataAttributes(resourceType interface{}) *POSTCleanups201ResponseDataAttributes {
	this := POSTCleanups201ResponseDataAttributes{}
	this.ResourceType = resourceType
	return &this
}

// NewPOSTCleanups201ResponseDataAttributesWithDefaults instantiates a new POSTCleanups201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCleanups201ResponseDataAttributesWithDefaults() *POSTCleanups201ResponseDataAttributes {
	this := POSTCleanups201ResponseDataAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCleanups201ResponseDataAttributes) GetResourceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCleanups201ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *POSTCleanups201ResponseDataAttributes) SetResourceType(v interface{}) {
	o.ResourceType = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCleanups201ResponseDataAttributes) GetFilters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCleanups201ResponseDataAttributes) GetFiltersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *POSTCleanups201ResponseDataAttributes) HasFilters() bool {
	if o != nil && IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *POSTCleanups201ResponseDataAttributes) SetFilters(v interface{}) {
	o.Filters = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCleanups201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCleanups201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCleanups201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTCleanups201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCleanups201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCleanups201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCleanups201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTCleanups201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCleanups201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCleanups201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCleanups201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTCleanups201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTCleanups201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCleanups201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
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

type NullablePOSTCleanups201ResponseDataAttributes struct {
	value *POSTCleanups201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTCleanups201ResponseDataAttributes) Get() *POSTCleanups201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTCleanups201ResponseDataAttributes) Set(val *POSTCleanups201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCleanups201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCleanups201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCleanups201ResponseDataAttributes(val *POSTCleanups201ResponseDataAttributes) *NullablePOSTCleanups201ResponseDataAttributes {
	return &NullablePOSTCleanups201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCleanups201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCleanups201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
