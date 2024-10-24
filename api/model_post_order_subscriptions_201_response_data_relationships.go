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

// checks if the POSTOrderSubscriptions201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataRelationships{}

// POSTOrderSubscriptions201ResponseDataRelationships struct for POSTOrderSubscriptions201ResponseDataRelationships
type POSTOrderSubscriptions201ResponseDataRelationships struct {
	Market                 *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket         `json:"market,omitempty"`
	SubscriptionModel      *POSTMarkets201ResponseDataRelationshipsSubscriptionModel                 `json:"subscription_model,omitempty"`
	SourceOrder            *POSTOrderCopies201ResponseDataRelationshipsSourceOrder                   `json:"source_order,omitempty"`
	Customer               *POSTCouponRecipients201ResponseDataRelationshipsCustomer                 `json:"customer,omitempty"`
	CustomerPaymentSource  *POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource  `json:"customer_payment_source,omitempty"`
	OrderSubscriptionItems *POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems `json:"order_subscription_items,omitempty"`
	OrderFactories         *POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories         `json:"order_factories,omitempty"`
	RecurringOrderCopies   *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies   `json:"recurring_order_copies,omitempty"`
	Orders                 *POSTCustomers201ResponseDataRelationshipsOrders                          `json:"orders,omitempty"`
	Events                 *POSTAddresses201ResponseDataRelationshipsEvents                          `json:"events,omitempty"`
	Tags                   *POSTAddresses201ResponseDataRelationshipsTags                            `json:"tags,omitempty"`
	Versions               *POSTAddresses201ResponseDataRelationshipsVersions                        `json:"versions,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataRelationships instantiates a new POSTOrderSubscriptions201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataRelationships() *POSTOrderSubscriptions201ResponseDataRelationships {
	this := POSTOrderSubscriptions201ResponseDataRelationships{}
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationships {
	this := POSTOrderSubscriptions201ResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetMarket() POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetMarket(v POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetSubscriptionModel returns the SubscriptionModel field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSubscriptionModel() POSTMarkets201ResponseDataRelationshipsSubscriptionModel {
	if o == nil || IsNil(o.SubscriptionModel) {
		var ret POSTMarkets201ResponseDataRelationshipsSubscriptionModel
		return ret
	}
	return *o.SubscriptionModel
}

// GetSubscriptionModelOk returns a tuple with the SubscriptionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSubscriptionModelOk() (*POSTMarkets201ResponseDataRelationshipsSubscriptionModel, bool) {
	if o == nil || IsNil(o.SubscriptionModel) {
		return nil, false
	}
	return o.SubscriptionModel, true
}

// HasSubscriptionModel returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasSubscriptionModel() bool {
	if o != nil && !IsNil(o.SubscriptionModel) {
		return true
	}

	return false
}

// SetSubscriptionModel gets a reference to the given POSTMarkets201ResponseDataRelationshipsSubscriptionModel and assigns it to the SubscriptionModel field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetSubscriptionModel(v POSTMarkets201ResponseDataRelationshipsSubscriptionModel) {
	o.SubscriptionModel = &v
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSourceOrder() POSTOrderCopies201ResponseDataRelationshipsSourceOrder {
	if o == nil || IsNil(o.SourceOrder) {
		var ret POSTOrderCopies201ResponseDataRelationshipsSourceOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSourceOrderOk() (*POSTOrderCopies201ResponseDataRelationshipsSourceOrder, bool) {
	if o == nil || IsNil(o.SourceOrder) {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasSourceOrder() bool {
	if o != nil && !IsNil(o.SourceOrder) {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given POSTOrderCopies201ResponseDataRelationshipsSourceOrder and assigns it to the SourceOrder field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetSourceOrder(v POSTOrderCopies201ResponseDataRelationshipsSourceOrder) {
	o.SourceOrder = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret POSTCouponRecipients201ResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given POSTCouponRecipients201ResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetCustomerPaymentSource returns the CustomerPaymentSource field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomerPaymentSource() POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource {
	if o == nil || IsNil(o.CustomerPaymentSource) {
		var ret POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource
		return ret
	}
	return *o.CustomerPaymentSource
}

// GetCustomerPaymentSourceOk returns a tuple with the CustomerPaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomerPaymentSourceOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource, bool) {
	if o == nil || IsNil(o.CustomerPaymentSource) {
		return nil, false
	}
	return o.CustomerPaymentSource, true
}

// HasCustomerPaymentSource returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasCustomerPaymentSource() bool {
	if o != nil && !IsNil(o.CustomerPaymentSource) {
		return true
	}

	return false
}

// SetCustomerPaymentSource gets a reference to the given POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource and assigns it to the CustomerPaymentSource field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetCustomerPaymentSource(v POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource) {
	o.CustomerPaymentSource = &v
}

// GetOrderSubscriptionItems returns the OrderSubscriptionItems field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderSubscriptionItems() POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems {
	if o == nil || IsNil(o.OrderSubscriptionItems) {
		var ret POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems
		return ret
	}
	return *o.OrderSubscriptionItems
}

// GetOrderSubscriptionItemsOk returns a tuple with the OrderSubscriptionItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderSubscriptionItemsOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems, bool) {
	if o == nil || IsNil(o.OrderSubscriptionItems) {
		return nil, false
	}
	return o.OrderSubscriptionItems, true
}

// HasOrderSubscriptionItems returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasOrderSubscriptionItems() bool {
	if o != nil && !IsNil(o.OrderSubscriptionItems) {
		return true
	}

	return false
}

// SetOrderSubscriptionItems gets a reference to the given POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems and assigns it to the OrderSubscriptionItems field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetOrderSubscriptionItems(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems) {
	o.OrderSubscriptionItems = &v
}

// GetOrderFactories returns the OrderFactories field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderFactories() POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories {
	if o == nil || IsNil(o.OrderFactories) {
		var ret POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories
		return ret
	}
	return *o.OrderFactories
}

// GetOrderFactoriesOk returns a tuple with the OrderFactories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderFactoriesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories, bool) {
	if o == nil || IsNil(o.OrderFactories) {
		return nil, false
	}
	return o.OrderFactories, true
}

// HasOrderFactories returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasOrderFactories() bool {
	if o != nil && !IsNil(o.OrderFactories) {
		return true
	}

	return false
}

// SetOrderFactories gets a reference to the given POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories and assigns it to the OrderFactories field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetOrderFactories(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories) {
	o.OrderFactories = &v
}

// GetRecurringOrderCopies returns the RecurringOrderCopies field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetRecurringOrderCopies() POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies {
	if o == nil || IsNil(o.RecurringOrderCopies) {
		var ret POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies
		return ret
	}
	return *o.RecurringOrderCopies
}

// GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetRecurringOrderCopiesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies, bool) {
	if o == nil || IsNil(o.RecurringOrderCopies) {
		return nil, false
	}
	return o.RecurringOrderCopies, true
}

// HasRecurringOrderCopies returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasRecurringOrderCopies() bool {
	if o != nil && !IsNil(o.RecurringOrderCopies) {
		return true
	}

	return false
}

// SetRecurringOrderCopies gets a reference to the given POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies and assigns it to the RecurringOrderCopies field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetRecurringOrderCopies(v POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies) {
	o.RecurringOrderCopies = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrders() POSTCustomers201ResponseDataRelationshipsOrders {
	if o == nil || IsNil(o.Orders) {
		var ret POSTCustomers201ResponseDataRelationshipsOrders
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrdersOk() (*POSTCustomers201ResponseDataRelationshipsOrders, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given POSTCustomers201ResponseDataRelationshipsOrders and assigns it to the Orders field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetOrders(v POSTCustomers201ResponseDataRelationshipsOrders) {
	o.Orders = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret POSTAddresses201ResponseDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given POSTAddresses201ResponseDataRelationshipsTags and assigns it to the Tags field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTOrderSubscriptions201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.SubscriptionModel) {
		toSerialize["subscription_model"] = o.SubscriptionModel
	}
	if !IsNil(o.SourceOrder) {
		toSerialize["source_order"] = o.SourceOrder
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.CustomerPaymentSource) {
		toSerialize["customer_payment_source"] = o.CustomerPaymentSource
	}
	if !IsNil(o.OrderSubscriptionItems) {
		toSerialize["order_subscription_items"] = o.OrderSubscriptionItems
	}
	if !IsNil(o.OrderFactories) {
		toSerialize["order_factories"] = o.OrderFactories
	}
	if !IsNil(o.RecurringOrderCopies) {
		toSerialize["recurring_order_copies"] = o.RecurringOrderCopies
	}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptions201ResponseDataRelationships struct {
	value *POSTOrderSubscriptions201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationships) Get() *POSTOrderSubscriptions201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationships) Set(val *POSTOrderSubscriptions201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataRelationships(val *POSTOrderSubscriptions201ResponseDataRelationships) *NullablePOSTOrderSubscriptions201ResponseDataRelationships {
	return &NullablePOSTOrderSubscriptions201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
