/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// QianchuanToolsShopAuthV10ApiService QianchuanToolsShopAuthV10Api service
type QianchuanToolsShopAuthV10ApiService service

type ApiOpenApiV10QianchuanToolsShopAuthPostRequest struct {
	ctx                              context.Context
	ApiService                       *QianchuanToolsShopAuthV10ApiService
	qianchuanToolsShopAuthV10Request *QianchuanToolsShopAuthV10Request
}

func (r *ApiOpenApiV10QianchuanToolsShopAuthPostRequest) QianchuanToolsShopAuthV10Request(qianchuanToolsShopAuthV10Request QianchuanToolsShopAuthV10Request) *ApiOpenApiV10QianchuanToolsShopAuthPostRequest {
	r.qianchuanToolsShopAuthV10Request = &qianchuanToolsShopAuthV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanToolsShopAuthPostRequest) Execute() (*QianchuanToolsShopAuthV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsShopAuthPostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsShopAuthPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsShopAuthPostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsShopAuthPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsShopAuthPost Method for OpenApiV10QianchuanToolsShopAuthPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsShopAuthPostRequest
*/
func (a *QianchuanToolsShopAuthV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanToolsShopAuthPostRequest {
	return &ApiOpenApiV10QianchuanToolsShopAuthPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsShopAuthV10Response
func (a *QianchuanToolsShopAuthV10ApiService) postExecute(r *ApiOpenApiV10QianchuanToolsShopAuthPostRequest) (*QianchuanToolsShopAuthV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanToolsShopAuthV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/shop_auth/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.qianchuanToolsShopAuthV10Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
