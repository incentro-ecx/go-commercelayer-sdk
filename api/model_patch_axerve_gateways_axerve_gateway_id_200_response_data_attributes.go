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

// checks if the PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes{}

// PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes struct for PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes
type PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The merchant login code.
	Login interface{} `json:"login,omitempty"`
	// The gateway API key.
	ApiKey interface{} `json:"api_key,omitempty"`
}

// NewPATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes instantiates a new PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes() *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes {
	this := PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributesWithDefaults() *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes {
	this := PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetLogin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetLoginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return &o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) HasLogin() bool {
	if o != nil && IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given interface{} and assigns it to the Login field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) SetLogin(v interface{}) {
	o.Login = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetApiKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) GetApiKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return &o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) HasApiKey() bool {
	if o != nil && IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given interface{} and assigns it to the ApiKey field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) SetApiKey(v interface{}) {
	o.ApiKey = v
}

func (o PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes struct {
	value *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) Get() *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) Set(val *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes(val *PATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) *NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes {
	return &NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
