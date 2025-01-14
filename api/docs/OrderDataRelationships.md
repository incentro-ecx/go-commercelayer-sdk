# OrderDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**BillingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**Store** | Pointer to [**MarketDataRelationshipsStores**](MarketDataRelationshipsStores.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**CustomerDataRelationshipsCustomerPaymentSources**](CustomerDataRelationshipsCustomerPaymentSources.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**LineItemDataRelationshipsBundle**](LineItemDataRelationshipsBundle.md) |  | [optional] 
**PaymentMethod** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**PaymentSource** | Pointer to [**AuthorizationDataRelationshipsPaymentSource**](AuthorizationDataRelationshipsPaymentSource.md) |  | [optional] 
**LineItems** | Pointer to [**LineItemOptionDataRelationshipsLineItem**](LineItemOptionDataRelationshipsLineItem.md) |  | [optional] 
**LineItemOptions** | Pointer to [**LineItemDataRelationshipsLineItemOptions**](LineItemDataRelationshipsLineItemOptions.md) |  | [optional] 
**StockReservations** | Pointer to [**LineItemDataRelationshipsStockReservations**](LineItemDataRelationshipsStockReservations.md) |  | [optional] 
**StockLineItems** | Pointer to [**LineItemDataRelationshipsStockLineItems**](LineItemDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**LineItemDataRelationshipsStockTransfers**](LineItemDataRelationshipsStockTransfers.md) |  | [optional] 
**Shipments** | Pointer to [**LineItemDataRelationshipsShipment**](LineItemDataRelationshipsShipment.md) |  | [optional] 
**PaymentOptions** | Pointer to [**OrderDataRelationshipsPaymentOptions**](OrderDataRelationshipsPaymentOptions.md) |  | [optional] 
**Transactions** | Pointer to [**OrderDataRelationshipsTransactions**](OrderDataRelationshipsTransactions.md) |  | [optional] 
**Authorizations** | Pointer to [**CaptureDataRelationshipsReferenceAuthorization**](CaptureDataRelationshipsReferenceAuthorization.md) |  | [optional] 
**Captures** | Pointer to [**AuthorizationDataRelationshipsCaptures**](AuthorizationDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**AuthorizationDataRelationshipsVoids**](AuthorizationDataRelationshipsVoids.md) |  | [optional] 
**Refunds** | Pointer to [**CaptureDataRelationshipsRefunds**](CaptureDataRelationshipsRefunds.md) |  | [optional] 
**Returns** | Pointer to [**CaptureDataRelationshipsReturn**](CaptureDataRelationshipsReturn.md) |  | [optional] 
**OrderSubscription** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**CustomerDataRelationshipsOrderSubscriptions**](CustomerDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderFactories** | Pointer to [**OrderSubscriptionDataRelationshipsOrderFactories**](OrderSubscriptionDataRelationshipsOrderFactories.md) |  | [optional] 
**OrderCopies** | Pointer to [**OrderDataRelationshipsOrderCopies**](OrderDataRelationshipsOrderCopies.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**OrderSubscriptionDataRelationshipsRecurringOrderCopies**](OrderSubscriptionDataRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Notifications** | Pointer to [**LineItemDataRelationshipsNotifications**](LineItemDataRelationshipsNotifications.md) |  | [optional] 
**Links** | Pointer to [**OrderDataRelationshipsLinks**](OrderDataRelationshipsLinks.md) |  | [optional] 
**ResourceErrors** | Pointer to [**OrderDataRelationshipsResourceErrors**](OrderDataRelationshipsResourceErrors.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewOrderDataRelationships

`func NewOrderDataRelationships() *OrderDataRelationships`

NewOrderDataRelationships instantiates a new OrderDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDataRelationshipsWithDefaults

`func NewOrderDataRelationshipsWithDefaults() *OrderDataRelationships`

NewOrderDataRelationshipsWithDefaults instantiates a new OrderDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *OrderDataRelationships) GetShippingAddress() BingGeocoderDataRelationshipsAddresses`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *OrderDataRelationships) GetShippingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *OrderDataRelationships) SetShippingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *OrderDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrderDataRelationships) GetBillingAddress() BingGeocoderDataRelationshipsAddresses`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrderDataRelationships) GetBillingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrderDataRelationships) SetBillingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrderDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetStore

`func (o *OrderDataRelationships) GetStore() MarketDataRelationshipsStores`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *OrderDataRelationships) GetStoreOk() (*MarketDataRelationshipsStores, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *OrderDataRelationships) SetStore(v MarketDataRelationshipsStores)`

SetStore sets Store field to given value.

### HasStore

`func (o *OrderDataRelationships) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *OrderDataRelationships) GetAvailablePaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *OrderDataRelationships) GetAvailablePaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *OrderDataRelationships) SetAvailablePaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *OrderDataRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *OrderDataRelationships) GetAvailableCustomerPaymentSources() CustomerDataRelationshipsCustomerPaymentSources`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *OrderDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *OrderDataRelationships) SetAvailableCustomerPaymentSources(v CustomerDataRelationshipsCustomerPaymentSources)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *OrderDataRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *OrderDataRelationships) GetAvailableFreeSkus() BundleDataRelationshipsSkus`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *OrderDataRelationships) GetAvailableFreeSkusOk() (*BundleDataRelationshipsSkus, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *OrderDataRelationships) SetAvailableFreeSkus(v BundleDataRelationshipsSkus)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *OrderDataRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *OrderDataRelationships) GetAvailableFreeBundles() LineItemDataRelationshipsBundle`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *OrderDataRelationships) GetAvailableFreeBundlesOk() (*LineItemDataRelationshipsBundle, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *OrderDataRelationships) SetAvailableFreeBundles(v LineItemDataRelationshipsBundle)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *OrderDataRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *OrderDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *OrderDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *OrderDataRelationships) GetPaymentSource() AuthorizationDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *OrderDataRelationships) GetPaymentSourceOk() (*AuthorizationDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *OrderDataRelationships) SetPaymentSource(v AuthorizationDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *OrderDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderDataRelationships) GetLineItems() LineItemOptionDataRelationshipsLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderDataRelationships) GetLineItemsOk() (*LineItemOptionDataRelationshipsLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderDataRelationships) SetLineItems(v LineItemOptionDataRelationshipsLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *OrderDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *OrderDataRelationships) GetLineItemOptions() LineItemDataRelationshipsLineItemOptions`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *OrderDataRelationships) GetLineItemOptionsOk() (*LineItemDataRelationshipsLineItemOptions, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *OrderDataRelationships) SetLineItemOptions(v LineItemDataRelationshipsLineItemOptions)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *OrderDataRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetStockReservations

`func (o *OrderDataRelationships) GetStockReservations() LineItemDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *OrderDataRelationships) GetStockReservationsOk() (*LineItemDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *OrderDataRelationships) SetStockReservations(v LineItemDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *OrderDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetStockLineItems

`func (o *OrderDataRelationships) GetStockLineItems() LineItemDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *OrderDataRelationships) GetStockLineItemsOk() (*LineItemDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *OrderDataRelationships) SetStockLineItems(v LineItemDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *OrderDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *OrderDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *OrderDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *OrderDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *OrderDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetShipments

`func (o *OrderDataRelationships) GetShipments() LineItemDataRelationshipsShipment`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *OrderDataRelationships) GetShipmentsOk() (*LineItemDataRelationshipsShipment, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *OrderDataRelationships) SetShipments(v LineItemDataRelationshipsShipment)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *OrderDataRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPaymentOptions

`func (o *OrderDataRelationships) GetPaymentOptions() OrderDataRelationshipsPaymentOptions`

GetPaymentOptions returns the PaymentOptions field if non-nil, zero value otherwise.

### GetPaymentOptionsOk

`func (o *OrderDataRelationships) GetPaymentOptionsOk() (*OrderDataRelationshipsPaymentOptions, bool)`

GetPaymentOptionsOk returns a tuple with the PaymentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptions

`func (o *OrderDataRelationships) SetPaymentOptions(v OrderDataRelationshipsPaymentOptions)`

SetPaymentOptions sets PaymentOptions field to given value.

### HasPaymentOptions

`func (o *OrderDataRelationships) HasPaymentOptions() bool`

HasPaymentOptions returns a boolean if a field has been set.

### GetTransactions

`func (o *OrderDataRelationships) GetTransactions() OrderDataRelationshipsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *OrderDataRelationships) GetTransactionsOk() (*OrderDataRelationshipsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *OrderDataRelationships) SetTransactions(v OrderDataRelationshipsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *OrderDataRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *OrderDataRelationships) GetAuthorizations() CaptureDataRelationshipsReferenceAuthorization`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *OrderDataRelationships) GetAuthorizationsOk() (*CaptureDataRelationshipsReferenceAuthorization, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *OrderDataRelationships) SetAuthorizations(v CaptureDataRelationshipsReferenceAuthorization)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *OrderDataRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *OrderDataRelationships) GetCaptures() AuthorizationDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *OrderDataRelationships) GetCapturesOk() (*AuthorizationDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *OrderDataRelationships) SetCaptures(v AuthorizationDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *OrderDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *OrderDataRelationships) GetVoids() AuthorizationDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *OrderDataRelationships) GetVoidsOk() (*AuthorizationDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *OrderDataRelationships) SetVoids(v AuthorizationDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *OrderDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *OrderDataRelationships) GetRefunds() CaptureDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *OrderDataRelationships) GetRefundsOk() (*CaptureDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *OrderDataRelationships) SetRefunds(v CaptureDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *OrderDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetReturns

`func (o *OrderDataRelationships) GetReturns() CaptureDataRelationshipsReturn`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *OrderDataRelationships) GetReturnsOk() (*CaptureDataRelationshipsReturn, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *OrderDataRelationships) SetReturns(v CaptureDataRelationshipsReturn)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *OrderDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetOrderSubscription

`func (o *OrderDataRelationships) GetOrderSubscription() CustomerDataRelationshipsOrderSubscriptions`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *OrderDataRelationships) GetOrderSubscriptionOk() (*CustomerDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *OrderDataRelationships) SetOrderSubscription(v CustomerDataRelationshipsOrderSubscriptions)`

SetOrderSubscription sets OrderSubscription field to given value.

### HasOrderSubscription

`func (o *OrderDataRelationships) HasOrderSubscription() bool`

HasOrderSubscription returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *OrderDataRelationships) GetOrderSubscriptions() CustomerDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *OrderDataRelationships) GetOrderSubscriptionsOk() (*CustomerDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *OrderDataRelationships) SetOrderSubscriptions(v CustomerDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *OrderDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderFactories

`func (o *OrderDataRelationships) GetOrderFactories() OrderSubscriptionDataRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *OrderDataRelationships) GetOrderFactoriesOk() (*OrderSubscriptionDataRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *OrderDataRelationships) SetOrderFactories(v OrderSubscriptionDataRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *OrderDataRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetOrderCopies

`func (o *OrderDataRelationships) GetOrderCopies() OrderDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *OrderDataRelationships) GetOrderCopiesOk() (*OrderDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *OrderDataRelationships) SetOrderCopies(v OrderDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *OrderDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *OrderDataRelationships) GetRecurringOrderCopies() OrderSubscriptionDataRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *OrderDataRelationships) GetRecurringOrderCopiesOk() (*OrderSubscriptionDataRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *OrderDataRelationships) SetRecurringOrderCopies(v OrderSubscriptionDataRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *OrderDataRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *OrderDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *OrderDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *OrderDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *OrderDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetNotifications

`func (o *OrderDataRelationships) GetNotifications() LineItemDataRelationshipsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OrderDataRelationships) GetNotificationsOk() (*LineItemDataRelationshipsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OrderDataRelationships) SetNotifications(v LineItemDataRelationshipsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OrderDataRelationships) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetLinks

`func (o *OrderDataRelationships) GetLinks() OrderDataRelationshipsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderDataRelationships) GetLinksOk() (*OrderDataRelationshipsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderDataRelationships) SetLinks(v OrderDataRelationshipsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderDataRelationships) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceErrors

`func (o *OrderDataRelationships) GetResourceErrors() OrderDataRelationshipsResourceErrors`

GetResourceErrors returns the ResourceErrors field if non-nil, zero value otherwise.

### GetResourceErrorsOk

`func (o *OrderDataRelationships) GetResourceErrorsOk() (*OrderDataRelationshipsResourceErrors, bool)`

GetResourceErrorsOk returns a tuple with the ResourceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceErrors

`func (o *OrderDataRelationships) SetResourceErrors(v OrderDataRelationshipsResourceErrors)`

SetResourceErrors sets ResourceErrors field to given value.

### HasResourceErrors

`func (o *OrderDataRelationships) HasResourceErrors() bool`

HasResourceErrors returns a boolean if a field has been set.

### GetEvents

`func (o *OrderDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrderDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrderDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrderDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *OrderDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *OrderDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *OrderDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *OrderDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *OrderDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


