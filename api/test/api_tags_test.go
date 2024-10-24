/*
Commerce Layer API

Testing TagsApiService

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

func Test_api_TagsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsApiService DELETETagsTagId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tagId interface{}

		httpRes, err := apiClient.TagsApi.DELETETagsTagId(context.Background(), tagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETAddressIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var addressId interface{}

		httpRes, err := apiClient.TagsApi.GETAddressIdTags(context.Background(), addressId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETBundleIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId interface{}

		httpRes, err := apiClient.TagsApi.GETBundleIdTags(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETBuyXPayYPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var buyXPayYPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETBuyXPayYPromotionIdTags(context.Background(), buyXPayYPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETCouponIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var couponId interface{}

		httpRes, err := apiClient.TagsApi.GETCouponIdTags(context.Background(), couponId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETCustomerIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId interface{}

		httpRes, err := apiClient.TagsApi.GETCustomerIdTags(context.Background(), customerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETExternalPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETExternalPromotionIdTags(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETFixedAmountPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETFixedAmountPromotionIdTags(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETFixedPricePromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETFixedPricePromotionIdTags(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETFlexPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var flexPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETFlexPromotionIdTags(context.Background(), flexPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETFreeGiftPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETFreeGiftPromotionIdTags(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETFreeShippingPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETFreeShippingPromotionIdTags(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETGiftCardIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardId interface{}

		httpRes, err := apiClient.TagsApi.GETGiftCardIdTags(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETLineItemIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineItemId interface{}

		httpRes, err := apiClient.TagsApi.GETLineItemIdTags(context.Background(), lineItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETLineItemOptionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineItemOptionId interface{}

		httpRes, err := apiClient.TagsApi.GETLineItemOptionIdTags(context.Background(), lineItemOptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETOrderIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.TagsApi.GETOrderIdTags(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETOrderSubscriptionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionId interface{}

		httpRes, err := apiClient.TagsApi.GETOrderSubscriptionIdTags(context.Background(), orderSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETPercentageDiscountPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.TagsApi.GETPercentageDiscountPromotionIdTags(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETPromotionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.TagsApi.GETPromotionIdTags(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETReturnIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var returnId interface{}

		httpRes, err := apiClient.TagsApi.GETReturnIdTags(context.Background(), returnId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETShipmentIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.TagsApi.GETShipmentIdTags(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETSkuIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuId interface{}

		httpRes, err := apiClient.TagsApi.GETSkuIdTags(context.Background(), skuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETSkuOptionIdTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuOptionId interface{}

		httpRes, err := apiClient.TagsApi.GETSkuOptionIdTags(context.Background(), skuOptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.GETTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GETTagsTagId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tagId interface{}

		resp, httpRes, err := apiClient.TagsApi.GETTagsTagId(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService PATCHTagsTagId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tagId interface{}

		resp, httpRes, err := apiClient.TagsApi.PATCHTagsTagId(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService POSTTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.POSTTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
