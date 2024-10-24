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

// checks if the POSTExports201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExports201ResponseDataAttributes{}

// POSTExports201ResponseDataAttributes struct for POSTExports201ResponseDataAttributes
type POSTExports201ResponseDataAttributes struct {
	// The type of resource being exported.
	ResourceType interface{} `json:"resource_type"`
	// The format of the export one of 'json' (default) or 'csv'.
	Format interface{} `json:"format,omitempty"`
	// List of related resources that should be included in the export.
	Includes interface{} `json:"includes,omitempty"`
	// The filters used to select the records to be exported.
	Filters interface{} `json:"filters,omitempty"`
	// Send this attribute if you want to skip exporting redundant attributes (IDs, timestamps, blanks, etc.), useful when combining export and import to duplicate your dataset.
	DryData interface{} `json:"dry_data,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTExports201ResponseDataAttributes instantiates a new POSTExports201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExports201ResponseDataAttributes(resourceType interface{}) *POSTExports201ResponseDataAttributes {
	this := POSTExports201ResponseDataAttributes{}
	this.ResourceType = resourceType
	return &this
}

// NewPOSTExports201ResponseDataAttributesWithDefaults instantiates a new POSTExports201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExports201ResponseDataAttributesWithDefaults() *POSTExports201ResponseDataAttributes {
	this := POSTExports201ResponseDataAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTExports201ResponseDataAttributes) GetResourceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetResourceTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *POSTExports201ResponseDataAttributes) SetResourceType(v interface{}) {
	o.ResourceType = v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetFormat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetFormatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasFormat() bool {
	if o != nil && IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given interface{} and assigns it to the Format field.
func (o *POSTExports201ResponseDataAttributes) SetFormat(v interface{}) {
	o.Format = v
}

// GetIncludes returns the Includes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetIncludes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetIncludesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return &o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasIncludes() bool {
	if o != nil && IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given interface{} and assigns it to the Includes field.
func (o *POSTExports201ResponseDataAttributes) SetIncludes(v interface{}) {
	o.Includes = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetFilters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetFiltersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasFilters() bool {
	if o != nil && IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *POSTExports201ResponseDataAttributes) SetFilters(v interface{}) {
	o.Filters = v
}

// GetDryData returns the DryData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetDryData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DryData
}

// GetDryDataOk returns a tuple with the DryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetDryDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DryData) {
		return nil, false
	}
	return &o.DryData, true
}

// HasDryData returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasDryData() bool {
	if o != nil && IsNil(o.DryData) {
		return true
	}

	return false
}

// SetDryData gets a reference to the given interface{} and assigns it to the DryData field.
func (o *POSTExports201ResponseDataAttributes) SetDryData(v interface{}) {
	o.DryData = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTExports201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTExports201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExports201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExports201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTExports201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTExports201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTExports201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExports201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Includes != nil {
		toSerialize["includes"] = o.Includes
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.DryData != nil {
		toSerialize["dry_data"] = o.DryData
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

type NullablePOSTExports201ResponseDataAttributes struct {
	value *POSTExports201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTExports201ResponseDataAttributes) Get() *POSTExports201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTExports201ResponseDataAttributes) Set(val *POSTExports201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExports201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExports201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExports201ResponseDataAttributes(val *POSTExports201ResponseDataAttributes) *NullablePOSTExports201ResponseDataAttributes {
	return &NullablePOSTExports201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTExports201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExports201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
