/*
Commerce Layer API

Testing SkuListPromotionRulesApiService

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

func Test_api_SkuListPromotionRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SkuListPromotionRulesApiService DELETESkuListPromotionRulesSkuListPromotionRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListPromotionRuleId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.DELETESkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETBuyXPayYPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var buyXPayYPromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETBuyXPayYPromotionIdSkuListPromotionRule(context.Background(), buyXPayYPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETExternalPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETExternalPromotionIdSkuListPromotionRule(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETFixedAmountPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETFixedAmountPromotionIdSkuListPromotionRule(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETFixedPricePromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETFixedPricePromotionIdSkuListPromotionRule(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETFreeGiftPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETFreeGiftPromotionIdSkuListPromotionRule(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETFreeShippingPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETFreeShippingPromotionIdSkuListPromotionRule(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETPercentageDiscountPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETPercentageDiscountPromotionIdSkuListPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETPromotionIdSkuListPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.SkuListPromotionRulesApi.GETPromotionIdSkuListPromotionRule(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETSkuListPromotionRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SkuListPromotionRulesApi.GETSkuListPromotionRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService GETSkuListPromotionRulesSkuListPromotionRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListPromotionRuleId interface{}

		resp, httpRes, err := apiClient.SkuListPromotionRulesApi.GETSkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService PATCHSkuListPromotionRulesSkuListPromotionRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuListPromotionRuleId interface{}

		resp, httpRes, err := apiClient.SkuListPromotionRulesApi.PATCHSkuListPromotionRulesSkuListPromotionRuleId(context.Background(), skuListPromotionRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkuListPromotionRulesApiService POSTSkuListPromotionRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SkuListPromotionRulesApi.POSTSkuListPromotionRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
