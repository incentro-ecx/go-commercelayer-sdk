# POSTOrders201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBundles201ResponseDataRelationshipsMarket**](POSTBundles201ResponseDataRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**POSTCouponRecipients201ResponseDataRelationshipsCustomer**](POSTCouponRecipients201ResponseDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**POSTOrders201ResponseDataRelationshipsShippingAddress**](POSTOrders201ResponseDataRelationshipsShippingAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**POSTOrders201ResponseDataRelationshipsBillingAddress**](POSTOrders201ResponseDataRelationshipsBillingAddress.md) |  | [optional] 
**Store** | Pointer to [**POSTOrders201ResponseDataRelationshipsStore**](POSTOrders201ResponseDataRelationshipsStore.md) |  | [optional] 
**AvailablePaymentMethods** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods**](POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods.md) |  | [optional] 
**AvailableCustomerPaymentSources** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources**](POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources.md) |  | [optional] 
**AvailableFreeSkus** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailableFreeSkus**](POSTOrders201ResponseDataRelationshipsAvailableFreeSkus.md) |  | [optional] 
**AvailableFreeBundles** | Pointer to [**POSTOrders201ResponseDataRelationshipsAvailableFreeBundles**](POSTOrders201ResponseDataRelationshipsAvailableFreeBundles.md) |  | [optional] 
**PaymentMethod** | Pointer to [**POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod**](POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource.md) |  | [optional] 
**LineItems** | Pointer to [**POSTOrders201ResponseDataRelationshipsLineItems**](POSTOrders201ResponseDataRelationshipsLineItems.md) |  | [optional] 
**LineItemOptions** | Pointer to [**POSTLineItems201ResponseDataRelationshipsLineItemOptions**](POSTLineItems201ResponseDataRelationshipsLineItemOptions.md) |  | [optional] 
**StockReservations** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockReservations**](POSTLineItems201ResponseDataRelationshipsStockReservations.md) |  | [optional] 
**StockLineItems** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockLineItems**](POSTLineItems201ResponseDataRelationshipsStockLineItems.md) |  | [optional] 
**StockTransfers** | Pointer to [**POSTLineItems201ResponseDataRelationshipsStockTransfers**](POSTLineItems201ResponseDataRelationshipsStockTransfers.md) |  | [optional] 
**Shipments** | Pointer to [**POSTOrders201ResponseDataRelationshipsShipments**](POSTOrders201ResponseDataRelationshipsShipments.md) |  | [optional] 
**PaymentOptions** | Pointer to [**POSTOrders201ResponseDataRelationshipsPaymentOptions**](POSTOrders201ResponseDataRelationshipsPaymentOptions.md) |  | [optional] 
**Transactions** | Pointer to [**POSTOrders201ResponseDataRelationshipsTransactions**](POSTOrders201ResponseDataRelationshipsTransactions.md) |  | [optional] 
**Authorizations** | Pointer to [**POSTOrders201ResponseDataRelationshipsAuthorizations**](POSTOrders201ResponseDataRelationshipsAuthorizations.md) |  | [optional] 
**Captures** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids.md) |  | [optional] 
**Refunds** | Pointer to [**GETCapturesCaptureId200ResponseDataRelationshipsRefunds**](GETCapturesCaptureId200ResponseDataRelationshipsRefunds.md) |  | [optional] 
**Returns** | Pointer to [**POSTCustomers201ResponseDataRelationshipsReturns**](POSTCustomers201ResponseDataRelationshipsReturns.md) |  | [optional] 
**OrderSubscription** | Pointer to [**POSTOrderCopies201ResponseDataRelationshipsOrderSubscription**](POSTOrderCopies201ResponseDataRelationshipsOrderSubscription.md) |  | [optional] 
**OrderSubscriptions** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrderSubscriptions**](POSTCustomers201ResponseDataRelationshipsOrderSubscriptions.md) |  | [optional] 
**OrderFactories** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories**](POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories.md) |  | [optional] 
**OrderCopies** | Pointer to [**POSTOrders201ResponseDataRelationshipsOrderCopies**](POSTOrders201ResponseDataRelationshipsOrderCopies.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies**](POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Notifications** | Pointer to [**POSTLineItems201ResponseDataRelationshipsNotifications**](POSTLineItems201ResponseDataRelationshipsNotifications.md) |  | [optional] 
**Links** | Pointer to [**POSTOrders201ResponseDataRelationshipsLinks**](POSTOrders201ResponseDataRelationshipsLinks.md) |  | [optional] 
**ResourceErrors** | Pointer to [**POSTOrders201ResponseDataRelationshipsResourceErrors**](POSTOrders201ResponseDataRelationshipsResourceErrors.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**POSTAddresses201ResponseDataRelationshipsTags**](POSTAddresses201ResponseDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTOrders201ResponseDataRelationships

`func NewPOSTOrders201ResponseDataRelationships() *POSTOrders201ResponseDataRelationships`

NewPOSTOrders201ResponseDataRelationships instantiates a new POSTOrders201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrders201ResponseDataRelationshipsWithDefaults

`func NewPOSTOrders201ResponseDataRelationshipsWithDefaults() *POSTOrders201ResponseDataRelationships`

NewPOSTOrders201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrders201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrders201ResponseDataRelationships) GetMarket() POSTBundles201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrders201ResponseDataRelationships) GetMarketOk() (*POSTBundles201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrders201ResponseDataRelationships) SetMarket(v POSTBundles201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrders201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *POSTOrders201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTOrders201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTOrders201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTOrders201ResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) GetShippingAddress() POSTOrders201ResponseDataRelationshipsShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *POSTOrders201ResponseDataRelationships) GetShippingAddressOk() (*POSTOrders201ResponseDataRelationshipsShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) SetShippingAddress(v POSTOrders201ResponseDataRelationshipsShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) GetBillingAddress() POSTOrders201ResponseDataRelationshipsBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *POSTOrders201ResponseDataRelationships) GetBillingAddressOk() (*POSTOrders201ResponseDataRelationshipsBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) SetBillingAddress(v POSTOrders201ResponseDataRelationshipsBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetStore

`func (o *POSTOrders201ResponseDataRelationships) GetStore() POSTOrders201ResponseDataRelationshipsStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *POSTOrders201ResponseDataRelationships) GetStoreOk() (*POSTOrders201ResponseDataRelationshipsStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *POSTOrders201ResponseDataRelationships) SetStore(v POSTOrders201ResponseDataRelationshipsStore)`

SetStore sets Store field to given value.

### HasStore

`func (o *POSTOrders201ResponseDataRelationships) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetAvailablePaymentMethods

`func (o *POSTOrders201ResponseDataRelationships) GetAvailablePaymentMethods() POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods`

GetAvailablePaymentMethods returns the AvailablePaymentMethods field if non-nil, zero value otherwise.

### GetAvailablePaymentMethodsOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailablePaymentMethodsOk() (*POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods, bool)`

GetAvailablePaymentMethodsOk returns a tuple with the AvailablePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePaymentMethods

`func (o *POSTOrders201ResponseDataRelationships) SetAvailablePaymentMethods(v POSTOrders201ResponseDataRelationshipsAvailablePaymentMethods)`

SetAvailablePaymentMethods sets AvailablePaymentMethods field to given value.

### HasAvailablePaymentMethods

`func (o *POSTOrders201ResponseDataRelationships) HasAvailablePaymentMethods() bool`

HasAvailablePaymentMethods returns a boolean if a field has been set.

### GetAvailableCustomerPaymentSources

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableCustomerPaymentSources() POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources`

GetAvailableCustomerPaymentSources returns the AvailableCustomerPaymentSources field if non-nil, zero value otherwise.

### GetAvailableCustomerPaymentSourcesOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableCustomerPaymentSourcesOk() (*POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources, bool)`

GetAvailableCustomerPaymentSourcesOk returns a tuple with the AvailableCustomerPaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCustomerPaymentSources

`func (o *POSTOrders201ResponseDataRelationships) SetAvailableCustomerPaymentSources(v POSTOrders201ResponseDataRelationshipsAvailableCustomerPaymentSources)`

SetAvailableCustomerPaymentSources sets AvailableCustomerPaymentSources field to given value.

### HasAvailableCustomerPaymentSources

`func (o *POSTOrders201ResponseDataRelationships) HasAvailableCustomerPaymentSources() bool`

HasAvailableCustomerPaymentSources returns a boolean if a field has been set.

### GetAvailableFreeSkus

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeSkus() POSTOrders201ResponseDataRelationshipsAvailableFreeSkus`

GetAvailableFreeSkus returns the AvailableFreeSkus field if non-nil, zero value otherwise.

### GetAvailableFreeSkusOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeSkusOk() (*POSTOrders201ResponseDataRelationshipsAvailableFreeSkus, bool)`

GetAvailableFreeSkusOk returns a tuple with the AvailableFreeSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeSkus

`func (o *POSTOrders201ResponseDataRelationships) SetAvailableFreeSkus(v POSTOrders201ResponseDataRelationshipsAvailableFreeSkus)`

SetAvailableFreeSkus sets AvailableFreeSkus field to given value.

### HasAvailableFreeSkus

`func (o *POSTOrders201ResponseDataRelationships) HasAvailableFreeSkus() bool`

HasAvailableFreeSkus returns a boolean if a field has been set.

### GetAvailableFreeBundles

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeBundles() POSTOrders201ResponseDataRelationshipsAvailableFreeBundles`

GetAvailableFreeBundles returns the AvailableFreeBundles field if non-nil, zero value otherwise.

### GetAvailableFreeBundlesOk

`func (o *POSTOrders201ResponseDataRelationships) GetAvailableFreeBundlesOk() (*POSTOrders201ResponseDataRelationshipsAvailableFreeBundles, bool)`

GetAvailableFreeBundlesOk returns a tuple with the AvailableFreeBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFreeBundles

`func (o *POSTOrders201ResponseDataRelationships) SetAvailableFreeBundles(v POSTOrders201ResponseDataRelationshipsAvailableFreeBundles)`

SetAvailableFreeBundles sets AvailableFreeBundles field to given value.

### HasAvailableFreeBundles

`func (o *POSTOrders201ResponseDataRelationships) HasAvailableFreeBundles() bool`

HasAvailableFreeBundles returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentMethod() POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentMethodOk() (*POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentMethod(v POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentSource() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentSourceOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentSource(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetLineItems

`func (o *POSTOrders201ResponseDataRelationships) GetLineItems() POSTOrders201ResponseDataRelationshipsLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *POSTOrders201ResponseDataRelationships) GetLineItemsOk() (*POSTOrders201ResponseDataRelationshipsLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *POSTOrders201ResponseDataRelationships) SetLineItems(v POSTOrders201ResponseDataRelationshipsLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *POSTOrders201ResponseDataRelationships) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLineItemOptions

`func (o *POSTOrders201ResponseDataRelationships) GetLineItemOptions() POSTLineItems201ResponseDataRelationshipsLineItemOptions`

GetLineItemOptions returns the LineItemOptions field if non-nil, zero value otherwise.

### GetLineItemOptionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetLineItemOptionsOk() (*POSTLineItems201ResponseDataRelationshipsLineItemOptions, bool)`

GetLineItemOptionsOk returns a tuple with the LineItemOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemOptions

`func (o *POSTOrders201ResponseDataRelationships) SetLineItemOptions(v POSTLineItems201ResponseDataRelationshipsLineItemOptions)`

SetLineItemOptions sets LineItemOptions field to given value.

### HasLineItemOptions

`func (o *POSTOrders201ResponseDataRelationships) HasLineItemOptions() bool`

HasLineItemOptions returns a boolean if a field has been set.

### GetStockReservations

`func (o *POSTOrders201ResponseDataRelationships) GetStockReservations() POSTLineItems201ResponseDataRelationshipsStockReservations`

GetStockReservations returns the StockReservations field if non-nil, zero value otherwise.

### GetStockReservationsOk

`func (o *POSTOrders201ResponseDataRelationships) GetStockReservationsOk() (*POSTLineItems201ResponseDataRelationshipsStockReservations, bool)`

GetStockReservationsOk returns a tuple with the StockReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockReservations

`func (o *POSTOrders201ResponseDataRelationships) SetStockReservations(v POSTLineItems201ResponseDataRelationshipsStockReservations)`

SetStockReservations sets StockReservations field to given value.

### HasStockReservations

`func (o *POSTOrders201ResponseDataRelationships) HasStockReservations() bool`

HasStockReservations returns a boolean if a field has been set.

### GetStockLineItems

`func (o *POSTOrders201ResponseDataRelationships) GetStockLineItems() POSTLineItems201ResponseDataRelationshipsStockLineItems`

GetStockLineItems returns the StockLineItems field if non-nil, zero value otherwise.

### GetStockLineItemsOk

`func (o *POSTOrders201ResponseDataRelationships) GetStockLineItemsOk() (*POSTLineItems201ResponseDataRelationshipsStockLineItems, bool)`

GetStockLineItemsOk returns a tuple with the StockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLineItems

`func (o *POSTOrders201ResponseDataRelationships) SetStockLineItems(v POSTLineItems201ResponseDataRelationshipsStockLineItems)`

SetStockLineItems sets StockLineItems field to given value.

### HasStockLineItems

`func (o *POSTOrders201ResponseDataRelationships) HasStockLineItems() bool`

HasStockLineItems returns a boolean if a field has been set.

### GetStockTransfers

`func (o *POSTOrders201ResponseDataRelationships) GetStockTransfers() POSTLineItems201ResponseDataRelationshipsStockTransfers`

GetStockTransfers returns the StockTransfers field if non-nil, zero value otherwise.

### GetStockTransfersOk

`func (o *POSTOrders201ResponseDataRelationships) GetStockTransfersOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfers, bool)`

GetStockTransfersOk returns a tuple with the StockTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockTransfers

`func (o *POSTOrders201ResponseDataRelationships) SetStockTransfers(v POSTLineItems201ResponseDataRelationshipsStockTransfers)`

SetStockTransfers sets StockTransfers field to given value.

### HasStockTransfers

`func (o *POSTOrders201ResponseDataRelationships) HasStockTransfers() bool`

HasStockTransfers returns a boolean if a field has been set.

### GetShipments

`func (o *POSTOrders201ResponseDataRelationships) GetShipments() POSTOrders201ResponseDataRelationshipsShipments`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *POSTOrders201ResponseDataRelationships) GetShipmentsOk() (*POSTOrders201ResponseDataRelationshipsShipments, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *POSTOrders201ResponseDataRelationships) SetShipments(v POSTOrders201ResponseDataRelationshipsShipments)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *POSTOrders201ResponseDataRelationships) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetPaymentOptions

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentOptions() POSTOrders201ResponseDataRelationshipsPaymentOptions`

GetPaymentOptions returns the PaymentOptions field if non-nil, zero value otherwise.

### GetPaymentOptionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentOptionsOk() (*POSTOrders201ResponseDataRelationshipsPaymentOptions, bool)`

GetPaymentOptionsOk returns a tuple with the PaymentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptions

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentOptions(v POSTOrders201ResponseDataRelationshipsPaymentOptions)`

SetPaymentOptions sets PaymentOptions field to given value.

### HasPaymentOptions

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentOptions() bool`

HasPaymentOptions returns a boolean if a field has been set.

### GetTransactions

`func (o *POSTOrders201ResponseDataRelationships) GetTransactions() POSTOrders201ResponseDataRelationshipsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetTransactionsOk() (*POSTOrders201ResponseDataRelationshipsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *POSTOrders201ResponseDataRelationships) SetTransactions(v POSTOrders201ResponseDataRelationshipsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *POSTOrders201ResponseDataRelationships) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAuthorizations

`func (o *POSTOrders201ResponseDataRelationships) GetAuthorizations() POSTOrders201ResponseDataRelationshipsAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *POSTOrders201ResponseDataRelationships) GetAuthorizationsOk() (*POSTOrders201ResponseDataRelationshipsAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *POSTOrders201ResponseDataRelationships) SetAuthorizations(v POSTOrders201ResponseDataRelationshipsAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *POSTOrders201ResponseDataRelationships) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCaptures

`func (o *POSTOrders201ResponseDataRelationships) GetCaptures() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *POSTOrders201ResponseDataRelationships) GetCapturesOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *POSTOrders201ResponseDataRelationships) SetCaptures(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *POSTOrders201ResponseDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *POSTOrders201ResponseDataRelationships) GetVoids() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *POSTOrders201ResponseDataRelationships) GetVoidsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *POSTOrders201ResponseDataRelationships) SetVoids(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *POSTOrders201ResponseDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.

### GetRefunds

`func (o *POSTOrders201ResponseDataRelationships) GetRefunds() GETCapturesCaptureId200ResponseDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *POSTOrders201ResponseDataRelationships) GetRefundsOk() (*GETCapturesCaptureId200ResponseDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *POSTOrders201ResponseDataRelationships) SetRefunds(v GETCapturesCaptureId200ResponseDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *POSTOrders201ResponseDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetReturns

`func (o *POSTOrders201ResponseDataRelationships) GetReturns() POSTCustomers201ResponseDataRelationshipsReturns`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *POSTOrders201ResponseDataRelationships) GetReturnsOk() (*POSTCustomers201ResponseDataRelationshipsReturns, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *POSTOrders201ResponseDataRelationships) SetReturns(v POSTCustomers201ResponseDataRelationshipsReturns)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *POSTOrders201ResponseDataRelationships) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetOrderSubscription

`func (o *POSTOrders201ResponseDataRelationships) GetOrderSubscription() POSTOrderCopies201ResponseDataRelationshipsOrderSubscription`

GetOrderSubscription returns the OrderSubscription field if non-nil, zero value otherwise.

### GetOrderSubscriptionOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderSubscriptionOk() (*POSTOrderCopies201ResponseDataRelationshipsOrderSubscription, bool)`

GetOrderSubscriptionOk returns a tuple with the OrderSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscription

`func (o *POSTOrders201ResponseDataRelationships) SetOrderSubscription(v POSTOrderCopies201ResponseDataRelationshipsOrderSubscription)`

SetOrderSubscription sets OrderSubscription field to given value.

### HasOrderSubscription

`func (o *POSTOrders201ResponseDataRelationships) HasOrderSubscription() bool`

HasOrderSubscription returns a boolean if a field has been set.

### GetOrderSubscriptions

`func (o *POSTOrders201ResponseDataRelationships) GetOrderSubscriptions() POSTCustomers201ResponseDataRelationshipsOrderSubscriptions`

GetOrderSubscriptions returns the OrderSubscriptions field if non-nil, zero value otherwise.

### GetOrderSubscriptionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderSubscriptionsOk() (*POSTCustomers201ResponseDataRelationshipsOrderSubscriptions, bool)`

GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptions

`func (o *POSTOrders201ResponseDataRelationships) SetOrderSubscriptions(v POSTCustomers201ResponseDataRelationshipsOrderSubscriptions)`

SetOrderSubscriptions sets OrderSubscriptions field to given value.

### HasOrderSubscriptions

`func (o *POSTOrders201ResponseDataRelationships) HasOrderSubscriptions() bool`

HasOrderSubscriptions returns a boolean if a field has been set.

### GetOrderFactories

`func (o *POSTOrders201ResponseDataRelationships) GetOrderFactories() POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderFactoriesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *POSTOrders201ResponseDataRelationships) SetOrderFactories(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *POSTOrders201ResponseDataRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) GetOrderCopies() POSTOrders201ResponseDataRelationshipsOrderCopies`

GetOrderCopies returns the OrderCopies field if non-nil, zero value otherwise.

### GetOrderCopiesOk

`func (o *POSTOrders201ResponseDataRelationships) GetOrderCopiesOk() (*POSTOrders201ResponseDataRelationshipsOrderCopies, bool)`

GetOrderCopiesOk returns a tuple with the OrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) SetOrderCopies(v POSTOrders201ResponseDataRelationshipsOrderCopies)`

SetOrderCopies sets OrderCopies field to given value.

### HasOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) HasOrderCopies() bool`

HasOrderCopies returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) GetRecurringOrderCopies() POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *POSTOrders201ResponseDataRelationships) GetRecurringOrderCopiesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) SetRecurringOrderCopies(v POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *POSTOrders201ResponseDataRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTOrders201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTOrders201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTOrders201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTOrders201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetNotifications

`func (o *POSTOrders201ResponseDataRelationships) GetNotifications() POSTLineItems201ResponseDataRelationshipsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *POSTOrders201ResponseDataRelationships) GetNotificationsOk() (*POSTLineItems201ResponseDataRelationshipsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *POSTOrders201ResponseDataRelationships) SetNotifications(v POSTLineItems201ResponseDataRelationshipsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *POSTOrders201ResponseDataRelationships) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetLinks

`func (o *POSTOrders201ResponseDataRelationships) GetLinks() POSTOrders201ResponseDataRelationshipsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTOrders201ResponseDataRelationships) GetLinksOk() (*POSTOrders201ResponseDataRelationshipsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTOrders201ResponseDataRelationships) SetLinks(v POSTOrders201ResponseDataRelationshipsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTOrders201ResponseDataRelationships) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceErrors

`func (o *POSTOrders201ResponseDataRelationships) GetResourceErrors() POSTOrders201ResponseDataRelationshipsResourceErrors`

GetResourceErrors returns the ResourceErrors field if non-nil, zero value otherwise.

### GetResourceErrorsOk

`func (o *POSTOrders201ResponseDataRelationships) GetResourceErrorsOk() (*POSTOrders201ResponseDataRelationshipsResourceErrors, bool)`

GetResourceErrorsOk returns a tuple with the ResourceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceErrors

`func (o *POSTOrders201ResponseDataRelationships) SetResourceErrors(v POSTOrders201ResponseDataRelationshipsResourceErrors)`

SetResourceErrors sets ResourceErrors field to given value.

### HasResourceErrors

`func (o *POSTOrders201ResponseDataRelationships) HasResourceErrors() bool`

HasResourceErrors returns a boolean if a field has been set.

### GetEvents

`func (o *POSTOrders201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTOrders201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTOrders201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTOrders201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *POSTOrders201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *POSTOrders201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *POSTOrders201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *POSTOrders201ResponseDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *POSTOrders201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTOrders201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTOrders201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTOrders201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


