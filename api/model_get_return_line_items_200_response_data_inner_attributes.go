/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETReturnLineItems200ResponseDataInnerAttributes struct for GETReturnLineItems200ResponseDataInnerAttributes
type GETReturnLineItems200ResponseDataInnerAttributes struct {
	// The code of the associated SKU.
	SkuCode *string `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode *string `json:"bundle_code,omitempty"`
	// The name of the line item.
	Name *string `json:"name,omitempty"`
	// The line item quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// Set of key-value pairs that you can use to add details about return reason.
	ReturnReason map[string]interface{} `json:"return_reason,omitempty"`
	// Time at which the return line item was restocked.
	RestockedAt *string `json:"restocked_at,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETReturnLineItems200ResponseDataInnerAttributes instantiates a new GETReturnLineItems200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnLineItems200ResponseDataInnerAttributes() *GETReturnLineItems200ResponseDataInnerAttributes {
	this := GETReturnLineItems200ResponseDataInnerAttributes{}
	return &this
}

// NewGETReturnLineItems200ResponseDataInnerAttributesWithDefaults instantiates a new GETReturnLineItems200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnLineItems200ResponseDataInnerAttributesWithDefaults() *GETReturnLineItems200ResponseDataInnerAttributes {
	this := GETReturnLineItems200ResponseDataInnerAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetBundleCode() string {
	if o == nil || o.BundleCode == nil {
		var ret string
		return ret
	}
	return *o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetBundleCodeOk() (*string, bool) {
	if o == nil || o.BundleCode == nil {
		return nil, false
	}
	return o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasBundleCode() bool {
	if o != nil && o.BundleCode != nil {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given string and assigns it to the BundleCode field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetBundleCode(v string) {
	o.BundleCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetReturnReason() map[string]interface{} {
	if o == nil || o.ReturnReason == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetReturnReasonOk() (map[string]interface{}, bool) {
	if o == nil || o.ReturnReason == nil {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasReturnReason() bool {
	if o != nil && o.ReturnReason != nil {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given map[string]interface{} and assigns it to the ReturnReason field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetReturnReason(v map[string]interface{}) {
	o.ReturnReason = v
}

// GetRestockedAt returns the RestockedAt field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetRestockedAt() string {
	if o == nil || o.RestockedAt == nil {
		var ret string
		return ret
	}
	return *o.RestockedAt
}

// GetRestockedAtOk returns a tuple with the RestockedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetRestockedAtOk() (*string, bool) {
	if o == nil || o.RestockedAt == nil {
		return nil, false
	}
	return o.RestockedAt, true
}

// HasRestockedAt returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasRestockedAt() bool {
	if o != nil && o.RestockedAt != nil {
		return true
	}

	return false
}

// SetRestockedAt gets a reference to the given string and assigns it to the RestockedAt field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetRestockedAt(v string) {
	o.RestockedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETReturnLineItems200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETReturnLineItems200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.BundleCode != nil {
		toSerialize["bundle_code"] = o.BundleCode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ReturnReason != nil {
		toSerialize["return_reason"] = o.ReturnReason
	}
	if o.RestockedAt != nil {
		toSerialize["restocked_at"] = o.RestockedAt
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
	return json.Marshal(toSerialize)
}

type NullableGETReturnLineItems200ResponseDataInnerAttributes struct {
	value *GETReturnLineItems200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETReturnLineItems200ResponseDataInnerAttributes) Get() *GETReturnLineItems200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETReturnLineItems200ResponseDataInnerAttributes) Set(val *GETReturnLineItems200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnLineItems200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnLineItems200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnLineItems200ResponseDataInnerAttributes(val *GETReturnLineItems200ResponseDataInnerAttributes) *NullableGETReturnLineItems200ResponseDataInnerAttributes {
	return &NullableGETReturnLineItems200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETReturnLineItems200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnLineItems200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
