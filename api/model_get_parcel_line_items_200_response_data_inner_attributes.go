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

// GETParcelLineItems200ResponseDataInnerAttributes struct for GETParcelLineItems200ResponseDataInnerAttributes
type GETParcelLineItems200ResponseDataInnerAttributes struct {
	// The code of the SKU of the associated shipment_line_item.
	SkuCode *string `json:"sku_code,omitempty"`
	// The parcel line item quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// The internal name of the associated line item.
	Name *string `json:"name,omitempty"`
	// The image_url of the associated line item.
	ImageUrl *string `json:"image_url,omitempty"`
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

// NewGETParcelLineItems200ResponseDataInnerAttributes instantiates a new GETParcelLineItems200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelLineItems200ResponseDataInnerAttributes() *GETParcelLineItems200ResponseDataInnerAttributes {
	this := GETParcelLineItems200ResponseDataInnerAttributes{}
	return &this
}

// NewGETParcelLineItems200ResponseDataInnerAttributesWithDefaults instantiates a new GETParcelLineItems200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelLineItems200ResponseDataInnerAttributesWithDefaults() *GETParcelLineItems200ResponseDataInnerAttributes {
	this := GETParcelLineItems200ResponseDataInnerAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETParcelLineItems200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETParcelLineItems200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
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

type NullableGETParcelLineItems200ResponseDataInnerAttributes struct {
	value *GETParcelLineItems200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETParcelLineItems200ResponseDataInnerAttributes) Get() *GETParcelLineItems200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETParcelLineItems200ResponseDataInnerAttributes) Set(val *GETParcelLineItems200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelLineItems200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelLineItems200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelLineItems200ResponseDataInnerAttributes(val *GETParcelLineItems200ResponseDataInnerAttributes) *NullableGETParcelLineItems200ResponseDataInnerAttributes {
	return &NullableGETParcelLineItems200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETParcelLineItems200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelLineItems200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
