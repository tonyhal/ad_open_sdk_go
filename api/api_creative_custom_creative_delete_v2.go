/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

// CreativeCustomCreativeDeleteV2ApiService CreativeCustomCreativeDeleteV2Api service
type CreativeCustomCreativeDeleteV2ApiService service

type ApiOpenApi2CreativeCustomCreativeDeletePostRequest struct {
	ctx                                   context.Context
	ApiService                            *CreativeCustomCreativeDeleteV2ApiService
	creativeCustomCreativeDeleteV2Request *CreativeCustomCreativeDeleteV2Request
}

func (r *ApiOpenApi2CreativeCustomCreativeDeletePostRequest) CreativeCustomCreativeDeleteV2Request(creativeCustomCreativeDeleteV2Request CreativeCustomCreativeDeleteV2Request) *ApiOpenApi2CreativeCustomCreativeDeletePostRequest {
	r.creativeCustomCreativeDeleteV2Request = &creativeCustomCreativeDeleteV2Request
	return r
}

func (r *ApiOpenApi2CreativeCustomCreativeDeletePostRequest) Execute() (*CreativeCustomCreativeDeleteV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2CreativeCustomCreativeDeletePostRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeCustomCreativeDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeCustomCreativeDeletePostRequest) WithLog(enable bool) *ApiOpenApi2CreativeCustomCreativeDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeCustomCreativeDeletePost Method for OpenApi2CreativeCustomCreativeDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeCustomCreativeDeletePostRequest
*/
func (a *CreativeCustomCreativeDeleteV2ApiService) Post(ctx context.Context) *ApiOpenApi2CreativeCustomCreativeDeletePostRequest {
	return &ApiOpenApi2CreativeCustomCreativeDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeCustomCreativeDeleteV2Response
func (a *CreativeCustomCreativeDeleteV2ApiService) postExecute(r *ApiOpenApi2CreativeCustomCreativeDeletePostRequest) (*CreativeCustomCreativeDeleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeCustomCreativeDeleteV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/custom_creative/delete/"

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
	localVarPostBody = r.creativeCustomCreativeDeleteV2Request
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
