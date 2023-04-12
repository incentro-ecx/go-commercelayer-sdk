/*
Commerce Layer API

Testing OrderSubscriptionItemsApiService

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

func Test_api_OrderSubscriptionItemsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrderSubscriptionItemsApiService DELETEOrderSubscriptionItemsOrderSubscriptionItemId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionItemId interface{}

		httpRes, err := apiClient.OrderSubscriptionItemsApi.DELETEOrderSubscriptionItemsOrderSubscriptionItemId(context.Background(), orderSubscriptionItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderSubscriptionItemsApiService GETOrderSubscriptionIdOrderSubscriptionItems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionId interface{}

		httpRes, err := apiClient.OrderSubscriptionItemsApi.GETOrderSubscriptionIdOrderSubscriptionItems(context.Background(), orderSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderSubscriptionItemsApiService GETOrderSubscriptionItems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OrderSubscriptionItemsApi.GETOrderSubscriptionItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderSubscriptionItemsApiService GETOrderSubscriptionItemsOrderSubscriptionItemId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionItemId interface{}

		resp, httpRes, err := apiClient.OrderSubscriptionItemsApi.GETOrderSubscriptionItemsOrderSubscriptionItemId(context.Background(), orderSubscriptionItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderSubscriptionItemsApiService PATCHOrderSubscriptionItemsOrderSubscriptionItemId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionItemId interface{}

		resp, httpRes, err := apiClient.OrderSubscriptionItemsApi.PATCHOrderSubscriptionItemsOrderSubscriptionItemId(context.Background(), orderSubscriptionItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderSubscriptionItemsApiService POSTOrderSubscriptionItems", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OrderSubscriptionItemsApi.POSTOrderSubscriptionItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}