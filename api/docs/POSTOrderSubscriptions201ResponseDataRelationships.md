# POSTOrderSubscriptions201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBundles201ResponseDataRelationshipsMarket**](POSTBundles201ResponseDataRelationshipsMarket.md) |  | [optional] 
**SubscriptionModel** | Pointer to [**POSTMarkets201ResponseDataRelationshipsSubscriptionModel**](POSTMarkets201ResponseDataRelationshipsSubscriptionModel.md) |  | [optional] 
**SourceOrder** | Pointer to [**POSTOrderCopies201ResponseDataRelationshipsSourceOrder**](POSTOrderCopies201ResponseDataRelationshipsSourceOrder.md) |  | [optional] 
**Customer** | Pointer to [**POSTCouponRecipients201ResponseDataRelationshipsCustomer**](POSTCouponRecipients201ResponseDataRelationshipsCustomer.md) |  | [optional] 
**CustomerPaymentSource** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource**](POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource.md) |  | [optional] 
**OrderSubscriptionItems** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems**](POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems.md) |  | [optional] 
**OrderFactories** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories**](POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories.md) |  | [optional] 
**RecurringOrderCopies** | Pointer to [**POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies**](POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies.md) |  | [optional] 
**Orders** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrders**](POSTCustomers201ResponseDataRelationshipsOrders.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**POSTAddresses201ResponseDataRelationshipsTags**](POSTAddresses201ResponseDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTOrderSubscriptions201ResponseDataRelationships

`func NewPOSTOrderSubscriptions201ResponseDataRelationships() *POSTOrderSubscriptions201ResponseDataRelationships`

NewPOSTOrderSubscriptions201ResponseDataRelationships instantiates a new POSTOrderSubscriptions201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults

`func NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationships`

NewPOSTOrderSubscriptions201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetMarket() POSTBundles201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetMarketOk() (*POSTBundles201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetMarket(v POSTBundles201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSubscriptionModel

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSubscriptionModel() POSTMarkets201ResponseDataRelationshipsSubscriptionModel`

GetSubscriptionModel returns the SubscriptionModel field if non-nil, zero value otherwise.

### GetSubscriptionModelOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSubscriptionModelOk() (*POSTMarkets201ResponseDataRelationshipsSubscriptionModel, bool)`

GetSubscriptionModelOk returns a tuple with the SubscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionModel

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetSubscriptionModel(v POSTMarkets201ResponseDataRelationshipsSubscriptionModel)`

SetSubscriptionModel sets SubscriptionModel field to given value.

### HasSubscriptionModel

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasSubscriptionModel() bool`

HasSubscriptionModel returns a boolean if a field has been set.

### GetSourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSourceOrder() POSTOrderCopies201ResponseDataRelationshipsSourceOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetSourceOrderOk() (*POSTOrderCopies201ResponseDataRelationshipsSourceOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetSourceOrder(v POSTOrderCopies201ResponseDataRelationshipsSourceOrder)`

SetSourceOrder sets SourceOrder field to given value.

### HasSourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasSourceOrder() bool`

HasSourceOrder returns a boolean if a field has been set.

### GetCustomer

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomer() POSTCouponRecipients201ResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomerOk() (*POSTCouponRecipients201ResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetCustomer(v POSTCouponRecipients201ResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerPaymentSource

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomerPaymentSource() POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource`

GetCustomerPaymentSource returns the CustomerPaymentSource field if non-nil, zero value otherwise.

### GetCustomerPaymentSourceOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetCustomerPaymentSourceOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource, bool)`

GetCustomerPaymentSourceOk returns a tuple with the CustomerPaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPaymentSource

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetCustomerPaymentSource(v POSTOrderSubscriptions201ResponseDataRelationshipsCustomerPaymentSource)`

SetCustomerPaymentSource sets CustomerPaymentSource field to given value.

### HasCustomerPaymentSource

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasCustomerPaymentSource() bool`

HasCustomerPaymentSource returns a boolean if a field has been set.

### GetOrderSubscriptionItems

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderSubscriptionItems() POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems`

GetOrderSubscriptionItems returns the OrderSubscriptionItems field if non-nil, zero value otherwise.

### GetOrderSubscriptionItemsOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderSubscriptionItemsOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems, bool)`

GetOrderSubscriptionItemsOk returns a tuple with the OrderSubscriptionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSubscriptionItems

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetOrderSubscriptionItems(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderSubscriptionItems)`

SetOrderSubscriptionItems sets OrderSubscriptionItems field to given value.

### HasOrderSubscriptionItems

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasOrderSubscriptionItems() bool`

HasOrderSubscriptionItems returns a boolean if a field has been set.

### GetOrderFactories

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderFactories() POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories`

GetOrderFactories returns the OrderFactories field if non-nil, zero value otherwise.

### GetOrderFactoriesOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrderFactoriesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories, bool)`

GetOrderFactoriesOk returns a tuple with the OrderFactories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFactories

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetOrderFactories(v POSTOrderSubscriptions201ResponseDataRelationshipsOrderFactories)`

SetOrderFactories sets OrderFactories field to given value.

### HasOrderFactories

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasOrderFactories() bool`

HasOrderFactories returns a boolean if a field has been set.

### GetRecurringOrderCopies

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetRecurringOrderCopies() POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies`

GetRecurringOrderCopies returns the RecurringOrderCopies field if non-nil, zero value otherwise.

### GetRecurringOrderCopiesOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetRecurringOrderCopiesOk() (*POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies, bool)`

GetRecurringOrderCopiesOk returns a tuple with the RecurringOrderCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringOrderCopies

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetRecurringOrderCopies(v POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopies)`

SetRecurringOrderCopies sets RecurringOrderCopies field to given value.

### HasRecurringOrderCopies

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasRecurringOrderCopies() bool`

HasRecurringOrderCopies returns a boolean if a field has been set.

### GetOrders

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrders() POSTCustomers201ResponseDataRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetOrdersOk() (*POSTCustomers201ResponseDataRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetOrders(v POSTCustomers201ResponseDataRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetEvents

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetTags() POSTAddresses201ResponseDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetTagsOk() (*POSTAddresses201ResponseDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetTags(v POSTAddresses201ResponseDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTOrderSubscriptions201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


