/*
Commerce Layer API

Testing CapturesApiService

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

func Test_api_CapturesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CapturesApiService GETAuthorizationIdCaptures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authorizationId interface{}

		httpRes, err := apiClient.CapturesApi.GETAuthorizationIdCaptures(context.Background(), authorizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CapturesApiService GETCaptures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CapturesApi.GETCaptures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CapturesApiService GETCapturesCaptureId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var captureId interface{}

		resp, httpRes, err := apiClient.CapturesApi.GETCapturesCaptureId(context.Background(), captureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CapturesApiService GETOrderIdCaptures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.CapturesApi.GETOrderIdCaptures(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CapturesApiService GETRefundIdReferenceCapture", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var refundId interface{}

		httpRes, err := apiClient.CapturesApi.GETRefundIdReferenceCapture(context.Background(), refundId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CapturesApiService PATCHCapturesCaptureId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var captureId interface{}

		resp, httpRes, err := apiClient.CapturesApi.PATCHCapturesCaptureId(context.Background(), captureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}