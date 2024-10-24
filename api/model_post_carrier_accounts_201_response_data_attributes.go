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

// checks if the POSTCarrierAccounts201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCarrierAccounts201ResponseDataAttributes{}

// POSTCarrierAccounts201ResponseDataAttributes struct for POSTCarrierAccounts201ResponseDataAttributes
type POSTCarrierAccounts201ResponseDataAttributes struct {
	// The carrier account internal name.
	Name interface{} `json:"name"`
	// The Easypost service carrier type.
	EasypostType interface{} `json:"easypost_type"`
	// The Easypost carrier accounts credentials fields.
	Credentials interface{} `json:"credentials"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTCarrierAccounts201ResponseDataAttributes instantiates a new POSTCarrierAccounts201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCarrierAccounts201ResponseDataAttributes(name interface{}, easypostType interface{}, credentials interface{}) *POSTCarrierAccounts201ResponseDataAttributes {
	this := POSTCarrierAccounts201ResponseDataAttributes{}
	this.Name = name
	this.EasypostType = easypostType
	this.Credentials = credentials
	return &this
}

// NewPOSTCarrierAccounts201ResponseDataAttributesWithDefaults instantiates a new POSTCarrierAccounts201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCarrierAccounts201ResponseDataAttributesWithDefaults() *POSTCarrierAccounts201ResponseDataAttributes {
	this := POSTCarrierAccounts201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTCarrierAccounts201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetEasypostType returns the EasypostType field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetEasypostType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.EasypostType
}

// GetEasypostTypeOk returns a tuple with the EasypostType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetEasypostTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EasypostType) {
		return nil, false
	}
	return &o.EasypostType, true
}

// SetEasypostType sets field value
func (o *POSTCarrierAccounts201ResponseDataAttributes) SetEasypostType(v interface{}) {
	o.EasypostType = v
}

// GetCredentials returns the Credentials field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetCredentials() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetCredentialsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *POSTCarrierAccounts201ResponseDataAttributes) SetCredentials(v interface{}) {
	o.Credentials = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCarrierAccounts201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTCarrierAccounts201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCarrierAccounts201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTCarrierAccounts201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCarrierAccounts201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCarrierAccounts201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTCarrierAccounts201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTCarrierAccounts201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCarrierAccounts201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.EasypostType != nil {
		toSerialize["easypost_type"] = o.EasypostType
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
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

type NullablePOSTCarrierAccounts201ResponseDataAttributes struct {
	value *POSTCarrierAccounts201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTCarrierAccounts201ResponseDataAttributes) Get() *POSTCarrierAccounts201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTCarrierAccounts201ResponseDataAttributes) Set(val *POSTCarrierAccounts201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCarrierAccounts201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCarrierAccounts201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCarrierAccounts201ResponseDataAttributes(val *POSTCarrierAccounts201ResponseDataAttributes) *NullablePOSTCarrierAccounts201ResponseDataAttributes {
	return &NullablePOSTCarrierAccounts201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCarrierAccounts201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCarrierAccounts201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}