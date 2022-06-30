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

// CheckoutComGatewayUpdateData struct for CheckoutComGatewayUpdateData
type CheckoutComGatewayUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                     `json:"id"`
	Attributes    CheckoutComGatewayUpdateDataAttributes     `json:"attributes"`
	Relationships *CheckoutComGatewayCreateDataRelationships `json:"relationships,omitempty"`
}

// NewCheckoutComGatewayUpdateData instantiates a new CheckoutComGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayUpdateData(type_ string, id string, attributes CheckoutComGatewayUpdateDataAttributes) *CheckoutComGatewayUpdateData {
	this := CheckoutComGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCheckoutComGatewayUpdateDataWithDefaults instantiates a new CheckoutComGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayUpdateDataWithDefaults() *CheckoutComGatewayUpdateData {
	this := CheckoutComGatewayUpdateData{}
	var type_ string = "checkout_com_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CheckoutComGatewayUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComGatewayUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CheckoutComGatewayUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutComGatewayUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComGatewayUpdateData) GetAttributes() CheckoutComGatewayUpdateDataAttributes {
	if o == nil {
		var ret CheckoutComGatewayUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayUpdateData) GetAttributesOk() (*CheckoutComGatewayUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComGatewayUpdateData) SetAttributes(v CheckoutComGatewayUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComGatewayUpdateData) GetRelationships() CheckoutComGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CheckoutComGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayUpdateData) GetRelationshipsOk() (*CheckoutComGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComGatewayUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CheckoutComGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComGatewayUpdateData) SetRelationships(v CheckoutComGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComGatewayUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComGatewayUpdateData struct {
	value *CheckoutComGatewayUpdateData
	isSet bool
}

func (v NullableCheckoutComGatewayUpdateData) Get() *CheckoutComGatewayUpdateData {
	return v.value
}

func (v *NullableCheckoutComGatewayUpdateData) Set(val *CheckoutComGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayUpdateData(val *CheckoutComGatewayUpdateData) *NullableCheckoutComGatewayUpdateData {
	return &NullableCheckoutComGatewayUpdateData{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
