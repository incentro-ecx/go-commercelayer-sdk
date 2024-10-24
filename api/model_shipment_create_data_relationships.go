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

// checks if the ShipmentCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentCreateDataRelationships{}

// ShipmentCreateDataRelationships struct for ShipmentCreateDataRelationships
type ShipmentCreateDataRelationships struct {
	Order                  AdyenPaymentCreateDataRelationshipsOrder               `json:"order"`
	ShippingCategory       *ShipmentCreateDataRelationshipsShippingCategory       `json:"shipping_category,omitempty"`
	InventoryStockLocation ShipmentCreateDataRelationshipsInventoryStockLocation  `json:"inventory_stock_location"`
	ShippingAddress        *CustomerAddressCreateDataRelationshipsAddress         `json:"shipping_address,omitempty"`
	ShippingMethod         *DeliveryLeadTimeCreateDataRelationshipsShippingMethod `json:"shipping_method,omitempty"`
	Tags                   *AddressCreateDataRelationshipsTags                    `json:"tags,omitempty"`
}

// NewShipmentCreateDataRelationships instantiates a new ShipmentCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentCreateDataRelationships(order AdyenPaymentCreateDataRelationshipsOrder, inventoryStockLocation ShipmentCreateDataRelationshipsInventoryStockLocation) *ShipmentCreateDataRelationships {
	this := ShipmentCreateDataRelationships{}
	this.Order = order
	this.InventoryStockLocation = inventoryStockLocation
	return &this
}

// NewShipmentCreateDataRelationshipsWithDefaults instantiates a new ShipmentCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentCreateDataRelationshipsWithDefaults() *ShipmentCreateDataRelationships {
	this := ShipmentCreateDataRelationships{}
	return &this
}

// GetOrder returns the Order field value
func (o *ShipmentCreateDataRelationships) GetOrder() AdyenPaymentCreateDataRelationshipsOrder {
	if o == nil {
		var ret AdyenPaymentCreateDataRelationshipsOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationships) GetOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ShipmentCreateDataRelationships) SetOrder(v AdyenPaymentCreateDataRelationshipsOrder) {
	o.Order = v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *ShipmentCreateDataRelationships) GetShippingCategory() ShipmentCreateDataRelationshipsShippingCategory {
	if o == nil || IsNil(o.ShippingCategory) {
		var ret ShipmentCreateDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationships) GetShippingCategoryOk() (*ShipmentCreateDataRelationshipsShippingCategory, bool) {
	if o == nil || IsNil(o.ShippingCategory) {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *ShipmentCreateDataRelationships) HasShippingCategory() bool {
	if o != nil && !IsNil(o.ShippingCategory) {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentCreateDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *ShipmentCreateDataRelationships) SetShippingCategory(v ShipmentCreateDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetInventoryStockLocation returns the InventoryStockLocation field value
func (o *ShipmentCreateDataRelationships) GetInventoryStockLocation() ShipmentCreateDataRelationshipsInventoryStockLocation {
	if o == nil {
		var ret ShipmentCreateDataRelationshipsInventoryStockLocation
		return ret
	}

	return o.InventoryStockLocation
}

// GetInventoryStockLocationOk returns a tuple with the InventoryStockLocation field value
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationships) GetInventoryStockLocationOk() (*ShipmentCreateDataRelationshipsInventoryStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryStockLocation, true
}

// SetInventoryStockLocation sets field value
func (o *ShipmentCreateDataRelationships) SetInventoryStockLocation(v ShipmentCreateDataRelationshipsInventoryStockLocation) {
	o.InventoryStockLocation = v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *ShipmentCreateDataRelationships) GetShippingAddress() CustomerAddressCreateDataRelationshipsAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret CustomerAddressCreateDataRelationshipsAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationships) GetShippingAddressOk() (*CustomerAddressCreateDataRelationshipsAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *ShipmentCreateDataRelationships) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given CustomerAddressCreateDataRelationshipsAddress and assigns it to the ShippingAddress field.
func (o *ShipmentCreateDataRelationships) SetShippingAddress(v CustomerAddressCreateDataRelationshipsAddress) {
	o.ShippingAddress = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *ShipmentCreateDataRelationships) GetShippingMethod() DeliveryLeadTimeCreateDataRelationshipsShippingMethod {
	if o == nil || IsNil(o.ShippingMethod) {
		var ret DeliveryLeadTimeCreateDataRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeCreateDataRelationshipsShippingMethod, bool) {
	if o == nil || IsNil(o.ShippingMethod) {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *ShipmentCreateDataRelationships) HasShippingMethod() bool {
	if o != nil && !IsNil(o.ShippingMethod) {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *ShipmentCreateDataRelationships) SetShippingMethod(v DeliveryLeadTimeCreateDataRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ShipmentCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressCreateDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ShipmentCreateDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressCreateDataRelationshipsTags and assigns it to the Tags field.
func (o *ShipmentCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags) {
	o.Tags = &v
}

func (o ShipmentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	if !IsNil(o.ShippingCategory) {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	toSerialize["inventory_stock_location"] = o.InventoryStockLocation
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if !IsNil(o.ShippingMethod) {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableShipmentCreateDataRelationships struct {
	value *ShipmentCreateDataRelationships
	isSet bool
}

func (v NullableShipmentCreateDataRelationships) Get() *ShipmentCreateDataRelationships {
	return v.value
}

func (v *NullableShipmentCreateDataRelationships) Set(val *ShipmentCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentCreateDataRelationships(val *ShipmentCreateDataRelationships) *NullableShipmentCreateDataRelationships {
	return &NullableShipmentCreateDataRelationships{value: val, isSet: true}
}

func (v NullableShipmentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
