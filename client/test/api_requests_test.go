/*
Wasp API

Testing RequestsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_RequestsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RequestsApiService CallView", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RequestsApi.CallView(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RequestsApiService GetReceipt", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var chainID string
        var requestID string

        resp, httpRes, err := apiClient.RequestsApi.GetReceipt(context.Background(), chainID, requestID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RequestsApiService OffLedger", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RequestsApi.OffLedger(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RequestsApiService WaitForTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var chainID string
        var requestID string

        resp, httpRes, err := apiClient.RequestsApi.WaitForTransaction(context.Background(), chainID, requestID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
