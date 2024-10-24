/*
Commerce Layer API

Testing PriceListSchedulersApiService

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

func Test_api_PriceListSchedulersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PriceListSchedulersApiService DELETEPriceListSchedulersPriceListSchedulerId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListSchedulerId interface{}

		httpRes, err := apiClient.PriceListSchedulersApi.DELETEPriceListSchedulersPriceListSchedulerId(context.Background(), priceListSchedulerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListSchedulersApiService GETMarketIdPriceListSchedulers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		httpRes, err := apiClient.PriceListSchedulersApi.GETMarketIdPriceListSchedulers(context.Background(), marketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListSchedulersApiService GETPriceListIdPriceListSchedulers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListId interface{}

		httpRes, err := apiClient.PriceListSchedulersApi.GETPriceListIdPriceListSchedulers(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListSchedulersApiService GETPriceListSchedulers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PriceListSchedulersApi.GETPriceListSchedulers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListSchedulersApiService GETPriceListSchedulersPriceListSchedulerId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListSchedulerId interface{}

		resp, httpRes, err := apiClient.PriceListSchedulersApi.GETPriceListSchedulersPriceListSchedulerId(context.Background(), priceListSchedulerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListSchedulersApiService PATCHPriceListSchedulersPriceListSchedulerId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListSchedulerId interface{}

		resp, httpRes, err := apiClient.PriceListSchedulersApi.PATCHPriceListSchedulersPriceListSchedulerId(context.Background(), priceListSchedulerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListSchedulersApiService POSTPriceListSchedulers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PriceListSchedulersApi.POSTPriceListSchedulers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
