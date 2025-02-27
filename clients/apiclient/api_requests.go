/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// RequestsApiService RequestsApi service
type RequestsApiService service

type ApiCallViewRequest struct {
	ctx context.Context
	ApiService *RequestsApiService
	contractCallViewRequest *ContractCallViewRequest
}

// Parameters
func (r ApiCallViewRequest) ContractCallViewRequest(contractCallViewRequest ContractCallViewRequest) ApiCallViewRequest {
	r.contractCallViewRequest = &contractCallViewRequest
	return r
}

func (r ApiCallViewRequest) Execute() (*JSONDict, *http.Response, error) {
	return r.ApiService.CallViewExecute(r)
}

/*
CallView Call a view function on a contract by Hname

Execute a view call. Either use HName or Name properties. If both are supplied, HName are used.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCallViewRequest
*/
func (a *RequestsApiService) CallView(ctx context.Context) ApiCallViewRequest {
	return ApiCallViewRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JSONDict
func (a *RequestsApiService) CallViewExecute(r ApiCallViewRequest) (*JSONDict, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JSONDict
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.CallView")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests/callview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contractCallViewRequest == nil {
		return localVarReturnValue, nil, reportError("contractCallViewRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.contractCallViewRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetReceiptRequest struct {
	ctx context.Context
	ApiService *RequestsApiService
	chainID string
	requestID string
}

func (r ApiGetReceiptRequest) Execute() (*ReceiptResponse, *http.Response, error) {
	return r.ApiService.GetReceiptExecute(r)
}

/*
GetReceipt Get a receipt from a request ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chainID ChainID (Bech32)
 @param requestID RequestID (Hex)
 @return ApiGetReceiptRequest
*/
func (a *RequestsApiService) GetReceipt(ctx context.Context, chainID string, requestID string) ApiGetReceiptRequest {
	return ApiGetReceiptRequest{
		ApiService: a,
		ctx: ctx,
		chainID: chainID,
		requestID: requestID,
	}
}

// Execute executes the request
//  @return ReceiptResponse
func (a *RequestsApiService) GetReceiptExecute(r ApiGetReceiptRequest) (*ReceiptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReceiptResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.GetReceipt")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chainID}/receipts/{requestID}"
	localVarPath = strings.Replace(localVarPath, "{"+"chainID"+"}", url.PathEscape(parameterValueToString(r.chainID, "chainID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestID"+"}", url.PathEscape(parameterValueToString(r.requestID, "requestID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOffLedgerRequest struct {
	ctx context.Context
	ApiService *RequestsApiService
	offLedgerRequest *OffLedgerRequest
}

// Offledger request as JSON. Request encoded in Hex
func (r ApiOffLedgerRequest) OffLedgerRequest(offLedgerRequest OffLedgerRequest) ApiOffLedgerRequest {
	r.offLedgerRequest = &offLedgerRequest
	return r
}

func (r ApiOffLedgerRequest) Execute() (*http.Response, error) {
	return r.ApiService.OffLedgerExecute(r)
}

/*
OffLedger Post an off-ledger request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOffLedgerRequest
*/
func (a *RequestsApiService) OffLedger(ctx context.Context) ApiOffLedgerRequest {
	return ApiOffLedgerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RequestsApiService) OffLedgerExecute(r ApiOffLedgerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.OffLedger")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/requests/offledger"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offLedgerRequest == nil {
		return nil, reportError("offLedgerRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.offLedgerRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWaitForRequestRequest struct {
	ctx context.Context
	ApiService *RequestsApiService
	chainID string
	requestID string
	timeoutSeconds *int32
}

// The timeout in seconds
func (r ApiWaitForRequestRequest) TimeoutSeconds(timeoutSeconds int32) ApiWaitForRequestRequest {
	r.timeoutSeconds = &timeoutSeconds
	return r
}

func (r ApiWaitForRequestRequest) Execute() (*ReceiptResponse, *http.Response, error) {
	return r.ApiService.WaitForRequestExecute(r)
}

/*
WaitForRequest Wait until the given request has been processed by the node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chainID ChainID (Bech32)
 @param requestID RequestID (Hex)
 @return ApiWaitForRequestRequest
*/
func (a *RequestsApiService) WaitForRequest(ctx context.Context, chainID string, requestID string) ApiWaitForRequestRequest {
	return ApiWaitForRequestRequest{
		ApiService: a,
		ctx: ctx,
		chainID: chainID,
		requestID: requestID,
	}
}

// Execute executes the request
//  @return ReceiptResponse
func (a *RequestsApiService) WaitForRequestExecute(r ApiWaitForRequestRequest) (*ReceiptResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReceiptResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.WaitForRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chainID}/requests/{requestID}/wait"
	localVarPath = strings.Replace(localVarPath, "{"+"chainID"+"}", url.PathEscape(parameterValueToString(r.chainID, "chainID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestID"+"}", url.PathEscape(parameterValueToString(r.requestID, "requestID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.timeoutSeconds != nil {
		parameterAddToQuery(localVarQueryParams, "timeoutSeconds", r.timeoutSeconds, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
