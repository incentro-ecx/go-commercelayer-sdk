# AddressDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geocoder** | Pointer to [**AddressDataRelationshipsGeocoder**](AddressDataRelationshipsGeocoder.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewAddressDataRelationships

`func NewAddressDataRelationships() *AddressDataRelationships`

NewAddressDataRelationships instantiates a new AddressDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDataRelationshipsWithDefaults

`func NewAddressDataRelationshipsWithDefaults() *AddressDataRelationships`

NewAddressDataRelationshipsWithDefaults instantiates a new AddressDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeocoder

`func (o *AddressDataRelationships) GetGeocoder() AddressDataRelationshipsGeocoder`

GetGeocoder returns the Geocoder field if non-nil, zero value otherwise.

### GetGeocoderOk

`func (o *AddressDataRelationships) GetGeocoderOk() (*AddressDataRelationshipsGeocoder, bool)`

GetGeocoderOk returns a tuple with the Geocoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoder

`func (o *AddressDataRelationships) SetGeocoder(v AddressDataRelationshipsGeocoder)`

SetGeocoder sets Geocoder field to given value.

### HasGeocoder

`func (o *AddressDataRelationships) HasGeocoder() bool`

HasGeocoder returns a boolean if a field has been set.

### GetEvents

`func (o *AddressDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AddressDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AddressDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AddressDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *AddressDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddressDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddressDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddressDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *AddressDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AddressDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AddressDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AddressDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


