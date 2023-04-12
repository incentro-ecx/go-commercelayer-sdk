/*
Commerce Layer API

Testing CustomerAddressesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_CustomerAddressesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomerAddressesApiService DELETECustomerAddressesCustomerAddressId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerAddressId interface{}

		httpRes, err := apiClient.CustomerAddressesApi.DELETECustomerAddressesCustomerAddressId(context.Background(), customerAddressId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerAddressesApiService GETCustomerAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomerAddressesApi.GETCustomerAddresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerAddressesApiService GETCustomerAddressesCustomerAddressId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerAddressId interface{}

		resp, httpRes, err := apiClient.CustomerAddressesApi.GETCustomerAddressesCustomerAddressId(context.Background(), customerAddressId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerAddressesApiService GETCustomerIdCustomerAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId interface{}

		httpRes, err := apiClient.CustomerAddressesApi.GETCustomerIdCustomerAddresses(context.Background(), customerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerAddressesApiService PATCHCustomerAddressesCustomerAddressId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerAddressId interface{}

		resp, httpRes, err := apiClient.CustomerAddressesApi.PATCHCustomerAddressesCustomerAddressId(context.Background(), customerAddressId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerAddressesApiService POSTCustomerAddresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CustomerAddressesApi.POSTCustomerAddresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}