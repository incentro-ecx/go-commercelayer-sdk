# NotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETNotificationsNotificationId200ResponseDataAttributes**](GETNotificationsNotificationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**NotificationDataRelationships**](NotificationDataRelationships.md) |  | [optional] 

## Methods

### NewNotificationData

`func NewNotificationData(type_ interface{}, attributes GETNotificationsNotificationId200ResponseDataAttributes, ) *NotificationData`

NewNotificationData instantiates a new NotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDataWithDefaults

`func NewNotificationDataWithDefaults() *NotificationData`

NewNotificationDataWithDefaults instantiates a new NotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *NotificationData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NotificationData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *NotificationData) GetAttributes() GETNotificationsNotificationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotificationData) GetAttributesOk() (*GETNotificationsNotificationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotificationData) SetAttributes(v GETNotificationsNotificationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *NotificationData) GetRelationships() NotificationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *NotificationData) GetRelationshipsOk() (*NotificationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *NotificationData) SetRelationships(v NotificationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *NotificationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


