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

// OrderResponseDataRelationships struct for OrderResponseDataRelationships
type OrderResponseDataRelationships struct {
	Market                          *BillingInfoValidationRuleResponseDataRelationshipsMarket      `json:"market,omitempty"`
	Customer                        *CouponRecipientResponseDataRelationshipsCustomer              `json:"customer,omitempty"`
	ShippingAddress                 *OrderResponseDataRelationshipsShippingAddress                 `json:"shipping_address,omitempty"`
	BillingAddress                  *OrderResponseDataRelationshipsBillingAddress                  `json:"billing_address,omitempty"`
	AvailablePaymentMethods         *OrderResponseDataRelationshipsAvailablePaymentMethods         `json:"available_payment_methods,omitempty"`
	AvailableCustomerPaymentSources *OrderResponseDataRelationshipsAvailableCustomerPaymentSources `json:"available_customer_payment_sources,omitempty"`
	AvailableFreeSkus               *OrderResponseDataRelationshipsAvailableFreeSkus               `json:"available_free_skus,omitempty"`
	AvailableFreeBundles            *OrderResponseDataRelationshipsAvailableFreeBundles            `json:"available_free_bundles,omitempty"`
	PaymentMethod                   *OrderResponseDataRelationshipsPaymentMethod                   `json:"payment_method,omitempty"`
	PaymentSource                   *CustomerPaymentSourceResponseDataRelationshipsPaymentSource   `json:"payment_source,omitempty"`
	LineItems                       *OrderResponseDataRelationshipsLineItems                       `json:"line_items,omitempty"`
	Shipments                       *OrderResponseDataRelationshipsShipments                       `json:"shipments,omitempty"`
	Transactions                    *OrderResponseDataRelationshipsTransactions                    `json:"transactions,omitempty"`
	Authorizations                  *OrderResponseDataRelationshipsAuthorizations                  `json:"authorizations,omitempty"`
	Captures                        *AuthorizationResponseDataRelationshipsCaptures                `json:"captures,omitempty"`
	Voids                           *AuthorizationResponseDataRelationshipsVoids                   `json:"voids,omitempty"`
	Refunds                         *CaptureResponseDataRelationshipsRefunds                       `json:"refunds,omitempty"`
	OrderSubscriptions              *CustomerResponseDataRelationshipsOrderSubscriptions           `json:"order_subscriptions,omitempty"`
	OrderCopies                     *OrderSubscriptionResponseDataRelationshipsOrderCopies         `json:"order_copies,omitempty"`
	Attachments                     *AvalaraAccountResponseDataRelationshipsAttachments            `json:"attachments,omitempty"`
	Events                          *CustomerAddressResponseDataRelationshipsEvents                `json:"events,omitempty"`
}

// NewOrderResponseDataRelationships instantiates a new OrderResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDataRelationships() *OrderResponseDataRelationships {
	this := OrderResponseDataRelationships{}
	return &this
}

// NewOrderResponseDataRelationshipsWithDefaults instantiates a new OrderResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDataRelationshipsWithDefaults() *OrderResponseDataRelationships {
	this := OrderResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *OrderResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientResponseDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientResponseDataRelationshipsCustomer and assigns it to the Customer field.
func (o *OrderResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetShippingAddress() OrderResponseDataRelationshipsShippingAddress {
	if o == nil || o.ShippingAddress == nil {
		var ret OrderResponseDataRelationshipsShippingAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetShippingAddressOk() (*OrderResponseDataRelationshipsShippingAddress, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given OrderResponseDataRelationshipsShippingAddress and assigns it to the ShippingAddress field.
func (o *OrderResponseDataRelationships) SetShippingAddress(v OrderResponseDataRelationshipsShippingAddress) {
	o.ShippingAddress = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetBillingAddress() OrderResponseDataRelationshipsBillingAddress {
	if o == nil || o.BillingAddress == nil {
		var ret OrderResponseDataRelationshipsBillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetBillingAddressOk() (*OrderResponseDataRelationshipsBillingAddress, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given OrderResponseDataRelationshipsBillingAddress and assigns it to the BillingAddress field.
func (o *OrderResponseDataRelationships) SetBillingAddress(v OrderResponseDataRelationshipsBillingAddress) {
	o.BillingAddress = &v
}

// GetAvailablePaymentMethods returns the AvailablePaymentMethods field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetAvailablePaymentMethods() OrderResponseDataRelationshipsAvailablePaymentMethods {
	if o == nil || o.AvailablePaymentMethods == nil {
		var ret OrderResponseDataRelationshipsAvailablePaymentMethods
		return ret
	}
	return *o.AvailablePaymentMethods
}

// GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetAvailablePaymentMethodsOk() (*OrderResponseDataRelationshipsAvailablePaymentMethods, bool) {
	if o == nil || o.AvailablePaymentMethods == nil {
		return nil, false
	}
	return o.AvailablePaymentMethods, true
}

// HasAvailablePaymentMethods returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasAvailablePaymentMethods() bool {
	if o != nil && o.AvailablePaymentMethods != nil {
		return true
	}

	return false
}

// SetAvailablePaymentMethods gets a reference to the given OrderResponseDataRelationshipsAvailablePaymentMethods and assigns it to the AvailablePaymentMethods field.
func (o *OrderResponseDataRelationships) SetAvailablePaymentMethods(v OrderResponseDataRelationshipsAvailablePaymentMethods) {
	o.AvailablePaymentMethods = &v
}

// GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetAvailableCustomerPaymentSources() OrderResponseDataRelationshipsAvailableCustomerPaymentSources {
	if o == nil || o.AvailableCustomerPaymentSources == nil {
		var ret OrderResponseDataRelationshipsAvailableCustomerPaymentSources
		return ret
	}
	return *o.AvailableCustomerPaymentSources
}

// GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*OrderResponseDataRelationshipsAvailableCustomerPaymentSources, bool) {
	if o == nil || o.AvailableCustomerPaymentSources == nil {
		return nil, false
	}
	return o.AvailableCustomerPaymentSources, true
}

// HasAvailableCustomerPaymentSources returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasAvailableCustomerPaymentSources() bool {
	if o != nil && o.AvailableCustomerPaymentSources != nil {
		return true
	}

	return false
}

// SetAvailableCustomerPaymentSources gets a reference to the given OrderResponseDataRelationshipsAvailableCustomerPaymentSources and assigns it to the AvailableCustomerPaymentSources field.
func (o *OrderResponseDataRelationships) SetAvailableCustomerPaymentSources(v OrderResponseDataRelationshipsAvailableCustomerPaymentSources) {
	o.AvailableCustomerPaymentSources = &v
}

// GetAvailableFreeSkus returns the AvailableFreeSkus field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetAvailableFreeSkus() OrderResponseDataRelationshipsAvailableFreeSkus {
	if o == nil || o.AvailableFreeSkus == nil {
		var ret OrderResponseDataRelationshipsAvailableFreeSkus
		return ret
	}
	return *o.AvailableFreeSkus
}

// GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetAvailableFreeSkusOk() (*OrderResponseDataRelationshipsAvailableFreeSkus, bool) {
	if o == nil || o.AvailableFreeSkus == nil {
		return nil, false
	}
	return o.AvailableFreeSkus, true
}

// HasAvailableFreeSkus returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasAvailableFreeSkus() bool {
	if o != nil && o.AvailableFreeSkus != nil {
		return true
	}

	return false
}

// SetAvailableFreeSkus gets a reference to the given OrderResponseDataRelationshipsAvailableFreeSkus and assigns it to the AvailableFreeSkus field.
func (o *OrderResponseDataRelationships) SetAvailableFreeSkus(v OrderResponseDataRelationshipsAvailableFreeSkus) {
	o.AvailableFreeSkus = &v
}

// GetAvailableFreeBundles returns the AvailableFreeBundles field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetAvailableFreeBundles() OrderResponseDataRelationshipsAvailableFreeBundles {
	if o == nil || o.AvailableFreeBundles == nil {
		var ret OrderResponseDataRelationshipsAvailableFreeBundles
		return ret
	}
	return *o.AvailableFreeBundles
}

// GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetAvailableFreeBundlesOk() (*OrderResponseDataRelationshipsAvailableFreeBundles, bool) {
	if o == nil || o.AvailableFreeBundles == nil {
		return nil, false
	}
	return o.AvailableFreeBundles, true
}

// HasAvailableFreeBundles returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasAvailableFreeBundles() bool {
	if o != nil && o.AvailableFreeBundles != nil {
		return true
	}

	return false
}

// SetAvailableFreeBundles gets a reference to the given OrderResponseDataRelationshipsAvailableFreeBundles and assigns it to the AvailableFreeBundles field.
func (o *OrderResponseDataRelationships) SetAvailableFreeBundles(v OrderResponseDataRelationshipsAvailableFreeBundles) {
	o.AvailableFreeBundles = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetPaymentMethod() OrderResponseDataRelationshipsPaymentMethod {
	if o == nil || o.PaymentMethod == nil {
		var ret OrderResponseDataRelationshipsPaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetPaymentMethodOk() (*OrderResponseDataRelationshipsPaymentMethod, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given OrderResponseDataRelationshipsPaymentMethod and assigns it to the PaymentMethod field.
func (o *OrderResponseDataRelationships) SetPaymentMethod(v OrderResponseDataRelationshipsPaymentMethod) {
	o.PaymentMethod = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetPaymentSource() CustomerPaymentSourceResponseDataRelationshipsPaymentSource {
	if o == nil || o.PaymentSource == nil {
		var ret CustomerPaymentSourceResponseDataRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceResponseDataRelationshipsPaymentSource, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given CustomerPaymentSourceResponseDataRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *OrderResponseDataRelationships) SetPaymentSource(v CustomerPaymentSourceResponseDataRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetLineItems() OrderResponseDataRelationshipsLineItems {
	if o == nil || o.LineItems == nil {
		var ret OrderResponseDataRelationshipsLineItems
		return ret
	}
	return *o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetLineItemsOk() (*OrderResponseDataRelationshipsLineItems, bool) {
	if o == nil || o.LineItems == nil {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasLineItems() bool {
	if o != nil && o.LineItems != nil {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given OrderResponseDataRelationshipsLineItems and assigns it to the LineItems field.
func (o *OrderResponseDataRelationships) SetLineItems(v OrderResponseDataRelationshipsLineItems) {
	o.LineItems = &v
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetShipments() OrderResponseDataRelationshipsShipments {
	if o == nil || o.Shipments == nil {
		var ret OrderResponseDataRelationshipsShipments
		return ret
	}
	return *o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetShipmentsOk() (*OrderResponseDataRelationshipsShipments, bool) {
	if o == nil || o.Shipments == nil {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasShipments() bool {
	if o != nil && o.Shipments != nil {
		return true
	}

	return false
}

// SetShipments gets a reference to the given OrderResponseDataRelationshipsShipments and assigns it to the Shipments field.
func (o *OrderResponseDataRelationships) SetShipments(v OrderResponseDataRelationshipsShipments) {
	o.Shipments = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetTransactions() OrderResponseDataRelationshipsTransactions {
	if o == nil || o.Transactions == nil {
		var ret OrderResponseDataRelationshipsTransactions
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetTransactionsOk() (*OrderResponseDataRelationshipsTransactions, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given OrderResponseDataRelationshipsTransactions and assigns it to the Transactions field.
func (o *OrderResponseDataRelationships) SetTransactions(v OrderResponseDataRelationshipsTransactions) {
	o.Transactions = &v
}

// GetAuthorizations returns the Authorizations field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetAuthorizations() OrderResponseDataRelationshipsAuthorizations {
	if o == nil || o.Authorizations == nil {
		var ret OrderResponseDataRelationshipsAuthorizations
		return ret
	}
	return *o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetAuthorizationsOk() (*OrderResponseDataRelationshipsAuthorizations, bool) {
	if o == nil || o.Authorizations == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// HasAuthorizations returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasAuthorizations() bool {
	if o != nil && o.Authorizations != nil {
		return true
	}

	return false
}

// SetAuthorizations gets a reference to the given OrderResponseDataRelationshipsAuthorizations and assigns it to the Authorizations field.
func (o *OrderResponseDataRelationships) SetAuthorizations(v OrderResponseDataRelationshipsAuthorizations) {
	o.Authorizations = &v
}

// GetCaptures returns the Captures field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetCaptures() AuthorizationResponseDataRelationshipsCaptures {
	if o == nil || o.Captures == nil {
		var ret AuthorizationResponseDataRelationshipsCaptures
		return ret
	}
	return *o.Captures
}

// GetCapturesOk returns a tuple with the Captures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetCapturesOk() (*AuthorizationResponseDataRelationshipsCaptures, bool) {
	if o == nil || o.Captures == nil {
		return nil, false
	}
	return o.Captures, true
}

// HasCaptures returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasCaptures() bool {
	if o != nil && o.Captures != nil {
		return true
	}

	return false
}

// SetCaptures gets a reference to the given AuthorizationResponseDataRelationshipsCaptures and assigns it to the Captures field.
func (o *OrderResponseDataRelationships) SetCaptures(v AuthorizationResponseDataRelationshipsCaptures) {
	o.Captures = &v
}

// GetVoids returns the Voids field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetVoids() AuthorizationResponseDataRelationshipsVoids {
	if o == nil || o.Voids == nil {
		var ret AuthorizationResponseDataRelationshipsVoids
		return ret
	}
	return *o.Voids
}

// GetVoidsOk returns a tuple with the Voids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetVoidsOk() (*AuthorizationResponseDataRelationshipsVoids, bool) {
	if o == nil || o.Voids == nil {
		return nil, false
	}
	return o.Voids, true
}

// HasVoids returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasVoids() bool {
	if o != nil && o.Voids != nil {
		return true
	}

	return false
}

// SetVoids gets a reference to the given AuthorizationResponseDataRelationshipsVoids and assigns it to the Voids field.
func (o *OrderResponseDataRelationships) SetVoids(v AuthorizationResponseDataRelationshipsVoids) {
	o.Voids = &v
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetRefunds() CaptureResponseDataRelationshipsRefunds {
	if o == nil || o.Refunds == nil {
		var ret CaptureResponseDataRelationshipsRefunds
		return ret
	}
	return *o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetRefundsOk() (*CaptureResponseDataRelationshipsRefunds, bool) {
	if o == nil || o.Refunds == nil {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasRefunds() bool {
	if o != nil && o.Refunds != nil {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given CaptureResponseDataRelationshipsRefunds and assigns it to the Refunds field.
func (o *OrderResponseDataRelationships) SetRefunds(v CaptureResponseDataRelationshipsRefunds) {
	o.Refunds = &v
}

// GetOrderSubscriptions returns the OrderSubscriptions field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetOrderSubscriptions() CustomerResponseDataRelationshipsOrderSubscriptions {
	if o == nil || o.OrderSubscriptions == nil {
		var ret CustomerResponseDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscriptions
}

// GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetOrderSubscriptionsOk() (*CustomerResponseDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || o.OrderSubscriptions == nil {
		return nil, false
	}
	return o.OrderSubscriptions, true
}

// HasOrderSubscriptions returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasOrderSubscriptions() bool {
	if o != nil && o.OrderSubscriptions != nil {
		return true
	}

	return false
}

// SetOrderSubscriptions gets a reference to the given CustomerResponseDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscriptions field.
func (o *OrderResponseDataRelationships) SetOrderSubscriptions(v CustomerResponseDataRelationshipsOrderSubscriptions) {
	o.OrderSubscriptions = &v
}

// GetOrderCopies returns the OrderCopies field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetOrderCopies() OrderSubscriptionResponseDataRelationshipsOrderCopies {
	if o == nil || o.OrderCopies == nil {
		var ret OrderSubscriptionResponseDataRelationshipsOrderCopies
		return ret
	}
	return *o.OrderCopies
}

// GetOrderCopiesOk returns a tuple with the OrderCopies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetOrderCopiesOk() (*OrderSubscriptionResponseDataRelationshipsOrderCopies, bool) {
	if o == nil || o.OrderCopies == nil {
		return nil, false
	}
	return o.OrderCopies, true
}

// HasOrderCopies returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasOrderCopies() bool {
	if o != nil && o.OrderCopies != nil {
		return true
	}

	return false
}

// SetOrderCopies gets a reference to the given OrderSubscriptionResponseDataRelationshipsOrderCopies and assigns it to the OrderCopies field.
func (o *OrderResponseDataRelationships) SetOrderCopies(v OrderSubscriptionResponseDataRelationshipsOrderCopies) {
	o.OrderCopies = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *OrderResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrderResponseDataRelationships) GetEvents() CustomerAddressResponseDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDataRelationships) GetEventsOk() (*CustomerAddressResponseDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrderResponseDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *OrderResponseDataRelationships) SetEvents(v CustomerAddressResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o OrderResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if o.AvailablePaymentMethods != nil {
		toSerialize["available_payment_methods"] = o.AvailablePaymentMethods
	}
	if o.AvailableCustomerPaymentSources != nil {
		toSerialize["available_customer_payment_sources"] = o.AvailableCustomerPaymentSources
	}
	if o.AvailableFreeSkus != nil {
		toSerialize["available_free_skus"] = o.AvailableFreeSkus
	}
	if o.AvailableFreeBundles != nil {
		toSerialize["available_free_bundles"] = o.AvailableFreeBundles
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	if o.LineItems != nil {
		toSerialize["line_items"] = o.LineItems
	}
	if o.Shipments != nil {
		toSerialize["shipments"] = o.Shipments
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	if o.Authorizations != nil {
		toSerialize["authorizations"] = o.Authorizations
	}
	if o.Captures != nil {
		toSerialize["captures"] = o.Captures
	}
	if o.Voids != nil {
		toSerialize["voids"] = o.Voids
	}
	if o.Refunds != nil {
		toSerialize["refunds"] = o.Refunds
	}
	if o.OrderSubscriptions != nil {
		toSerialize["order_subscriptions"] = o.OrderSubscriptions
	}
	if o.OrderCopies != nil {
		toSerialize["order_copies"] = o.OrderCopies
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseDataRelationships struct {
	value *OrderResponseDataRelationships
	isSet bool
}

func (v NullableOrderResponseDataRelationships) Get() *OrderResponseDataRelationships {
	return v.value
}

func (v *NullableOrderResponseDataRelationships) Set(val *OrderResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDataRelationships(val *OrderResponseDataRelationships) *NullableOrderResponseDataRelationships {
	return &NullableOrderResponseDataRelationships{value: val, isSet: true}
}

func (v NullableOrderResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}