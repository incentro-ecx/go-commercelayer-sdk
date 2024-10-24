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

// checks if the POSTInventoryStockLocations201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInventoryStockLocations201ResponseDataAttributes{}

// POSTInventoryStockLocations201ResponseDataAttributes struct for POSTInventoryStockLocations201ResponseDataAttributes
type POSTInventoryStockLocations201ResponseDataAttributes struct {
	// The stock location priority within the associated invetory model.
	Priority interface{} `json:"priority"`
	// Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled.
	OnHold interface{} `json:"on_hold,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTInventoryStockLocations201ResponseDataAttributes instantiates a new POSTInventoryStockLocations201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryStockLocations201ResponseDataAttributes(priority interface{}) *POSTInventoryStockLocations201ResponseDataAttributes {
	this := POSTInventoryStockLocations201ResponseDataAttributes{}
	this.Priority = priority
	return &this
}

// NewPOSTInventoryStockLocations201ResponseDataAttributesWithDefaults instantiates a new POSTInventoryStockLocations201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryStockLocations201ResponseDataAttributesWithDefaults() *POSTInventoryStockLocations201ResponseDataAttributes {
	this := POSTInventoryStockLocations201ResponseDataAttributes{}
	return &this
}

// GetPriority returns the Priority field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *POSTInventoryStockLocations201ResponseDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetOnHold() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetOnHoldOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return &o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) HasOnHold() bool {
	if o != nil && IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given interface{} and assigns it to the OnHold field.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) SetOnHold(v interface{}) {
	o.OnHold = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryStockLocations201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTInventoryStockLocations201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTInventoryStockLocations201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInventoryStockLocations201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.OnHold != nil {
		toSerialize["on_hold"] = o.OnHold
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

type NullablePOSTInventoryStockLocations201ResponseDataAttributes struct {
	value *POSTInventoryStockLocations201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTInventoryStockLocations201ResponseDataAttributes) Get() *POSTInventoryStockLocations201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTInventoryStockLocations201ResponseDataAttributes) Set(val *POSTInventoryStockLocations201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryStockLocations201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryStockLocations201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryStockLocations201ResponseDataAttributes(val *POSTInventoryStockLocations201ResponseDataAttributes) *NullablePOSTInventoryStockLocations201ResponseDataAttributes {
	return &NullablePOSTInventoryStockLocations201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTInventoryStockLocations201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryStockLocations201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
