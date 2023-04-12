/*
Commerce Layer API

Testing ExternalPaymentsApiService

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

func Test_api_ExternalPaymentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExternalPaymentsApiService DELETEExternalPaymentsExternalPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPaymentId interface{}

		httpRes, err := apiClient.ExternalPaymentsApi.DELETEExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExternalPaymentsApiService GETExternalGatewayIdExternalPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalGatewayId interface{}

		httpRes, err := apiClient.ExternalPaymentsApi.GETExternalGatewayIdExternalPayments(context.Background(), externalGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExternalPaymentsApiService GETExternalPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExternalPaymentsApi.GETExternalPayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExternalPaymentsApiService GETExternalPaymentsExternalPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPaymentId interface{}

		resp, httpRes, err := apiClient.ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExternalPaymentsApiService PATCHExternalPaymentsExternalPaymentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPaymentId interface{}

		resp, httpRes, err := apiClient.ExternalPaymentsApi.PATCHExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExternalPaymentsApiService POSTExternalPayments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExternalPaymentsApi.POSTExternalPayments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}