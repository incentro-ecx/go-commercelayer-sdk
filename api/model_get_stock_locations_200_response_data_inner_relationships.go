/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETStockLocations200ResponseDataInnerRelationships struct for GETStockLocations200ResponseDataInnerRelationships
type GETStockLocations200ResponseDataInnerRelationships struct {
	Address                  *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress                `json:"address,omitempty"`
	InventoryStockLocations  *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations  `json:"inventory_stock_locations,omitempty"`
	InventoryReturnLocations *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations `json:"inventory_return_locations,omitempty"`
	StockItems               *GETSkus200ResponseDataInnerRelationshipsStockItems                          `json:"stock_items,omitempty"`
	StockTransfers           *GETLineItems200ResponseDataInnerRelationshipsStockTransfers                 `json:"stock_transfers,omitempty"`
	Attachments              *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments              `json:"attachments,omitempty"`
}

// NewGETStockLocations200ResponseDataInnerRelationships instantiates a new GETStockLocations200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockLocations200ResponseDataInnerRelationships() *GETStockLocations200ResponseDataInnerRelationships {
	this := GETStockLocations200ResponseDataInnerRelationships{}
	return &this
}

// NewGETStockLocations200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStockLocations200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockLocations200ResponseDataInnerRelationshipsWithDefaults() *GETStockLocations200ResponseDataInnerRelationships {
	this := GETStockLocations200ResponseDataInnerRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetAddress() GETCustomerAddresses200ResponseDataInnerRelationshipsAddress {
	if o == nil || o.Address == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetAddressOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsAddress and assigns it to the Address field.
func (o *GETStockLocations200ResponseDataInnerRelationships) SetAddress(v GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) {
	o.Address = &v
}

// GetInventoryStockLocations returns the InventoryStockLocations field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryStockLocations() GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations {
	if o == nil || o.InventoryStockLocations == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations
		return ret
	}
	return *o.InventoryStockLocations
}

// GetInventoryStockLocationsOk returns a tuple with the InventoryStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryStockLocationsOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations, bool) {
	if o == nil || o.InventoryStockLocations == nil {
		return nil, false
	}
	return o.InventoryStockLocations, true
}

// HasInventoryStockLocations returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) HasInventoryStockLocations() bool {
	if o != nil && o.InventoryStockLocations != nil {
		return true
	}

	return false
}

// SetInventoryStockLocations gets a reference to the given GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations and assigns it to the InventoryStockLocations field.
func (o *GETStockLocations200ResponseDataInnerRelationships) SetInventoryStockLocations(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) {
	o.InventoryStockLocations = &v
}

// GetInventoryReturnLocations returns the InventoryReturnLocations field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryReturnLocations() GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations {
	if o == nil || o.InventoryReturnLocations == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations
		return ret
	}
	return *o.InventoryReturnLocations
}

// GetInventoryReturnLocationsOk returns a tuple with the InventoryReturnLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetInventoryReturnLocationsOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations, bool) {
	if o == nil || o.InventoryReturnLocations == nil {
		return nil, false
	}
	return o.InventoryReturnLocations, true
}

// HasInventoryReturnLocations returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) HasInventoryReturnLocations() bool {
	if o != nil && o.InventoryReturnLocations != nil {
		return true
	}

	return false
}

// SetInventoryReturnLocations gets a reference to the given GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations and assigns it to the InventoryReturnLocations field.
func (o *GETStockLocations200ResponseDataInnerRelationships) SetInventoryReturnLocations(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) {
	o.InventoryReturnLocations = &v
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockItems() GETSkus200ResponseDataInnerRelationshipsStockItems {
	if o == nil || o.StockItems == nil {
		var ret GETSkus200ResponseDataInnerRelationshipsStockItems
		return ret
	}
	return *o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockItemsOk() (*GETSkus200ResponseDataInnerRelationshipsStockItems, bool) {
	if o == nil || o.StockItems == nil {
		return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) HasStockItems() bool {
	if o != nil && o.StockItems != nil {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given GETSkus200ResponseDataInnerRelationshipsStockItems and assigns it to the StockItems field.
func (o *GETStockLocations200ResponseDataInnerRelationships) SetStockItems(v GETSkus200ResponseDataInnerRelationshipsStockItems) {
	o.StockItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockTransfers() GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	if o == nil || o.StockTransfers == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetStockTransfersOk() (*GETLineItems200ResponseDataInnerRelationshipsStockTransfers, bool) {
	if o == nil || o.StockTransfers == nil {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) HasStockTransfers() bool {
	if o != nil && o.StockTransfers != nil {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *GETStockLocations200ResponseDataInnerRelationships) SetStockTransfers(v GETLineItems200ResponseDataInnerRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETStockLocations200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETStockLocations200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.InventoryStockLocations != nil {
		toSerialize["inventory_stock_locations"] = o.InventoryStockLocations
	}
	if o.InventoryReturnLocations != nil {
		toSerialize["inventory_return_locations"] = o.InventoryReturnLocations
	}
	if o.StockItems != nil {
		toSerialize["stock_items"] = o.StockItems
	}
	if o.StockTransfers != nil {
		toSerialize["stock_transfers"] = o.StockTransfers
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockLocations200ResponseDataInnerRelationships struct {
	value *GETStockLocations200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETStockLocations200ResponseDataInnerRelationships) Get() *GETStockLocations200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETStockLocations200ResponseDataInnerRelationships) Set(val *GETStockLocations200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockLocations200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockLocations200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockLocations200ResponseDataInnerRelationships(val *GETStockLocations200ResponseDataInnerRelationships) *NullableGETStockLocations200ResponseDataInnerRelationships {
	return &NullableGETStockLocations200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETStockLocations200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockLocations200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}