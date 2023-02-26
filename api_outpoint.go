/*
 * MetaSV for MVC API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 3.0.2
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// OutpointApiService OutpointApi service
type OutpointApiService service

type ApiOutpointTxidIndexGetRequest struct {
	ctx        _context.Context
	ApiService *OutpointApiService
	txid       string
	index      int32
}

func (r ApiOutpointTxidIndexGetRequest) Execute() (OutputInfo, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.OutpointTxidIndexGetExecute(r)
}

/*
 * OutpointTxidIndexGet Get tx output(outpoint for vin) spent status.
 * Get detailed info for a utxo(or txo if spent), Only outputs spent no longer than one month are available. (Premium feature will support full history)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param txid The txid of the output
 * @param index The index of the output in the tx.
 * @return ApiOutpointTxidIndexGetRequest
 */
func (a *OutpointApiService) OutpointTxidIndexGet(ctx _context.Context, txid string, index int32) ApiOutpointTxidIndexGetRequest {
	return ApiOutpointTxidIndexGetRequest{
		ApiService: a,
		ctx:        ctx,
		txid:       txid,
		index:      index,
	}
}

/*
 * Execute executes the request
 * @return OutputInfo
 */
func (a *OutpointApiService) OutpointTxidIndexGetExecute(r ApiOutpointTxidIndexGetRequest) (OutputInfo, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  OutputInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OutpointApiService.OutpointTxidIndexGet")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/outpoint/{txid}/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"txid"+"}", _neturl.PathEscape(parameterToString(r.txid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", _neturl.PathEscape(parameterToString(r.index, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}
