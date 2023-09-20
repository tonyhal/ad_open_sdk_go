/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

// AssetsCreativeComponentUpdateV2ApiService AssetsCreativeComponentUpdateV2Api service
type AssetsCreativeComponentUpdateV2ApiService service

type ApiOpenApi2AssetsCreativeComponentUpdatePostRequest struct {
	ctx                                    context.Context
	ApiService                             *AssetsCreativeComponentUpdateV2ApiService
	assetsCreativeComponentUpdateV2Request *AssetsCreativeComponentUpdateV2Request
}

func (r *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest) AssetsCreativeComponentUpdateV2Request(assetsCreativeComponentUpdateV2Request AssetsCreativeComponentUpdateV2Request) *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest {
	r.assetsCreativeComponentUpdateV2Request = &assetsCreativeComponentUpdateV2Request
	return r
}

func (r *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest) Execute() (*AssetsCreativeComponentUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AssetsCreativeComponentUpdatePost Method for OpenApi2AssetsCreativeComponentUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AssetsCreativeComponentUpdatePostRequest
*/
func (a *AssetsCreativeComponentUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest {
	return &ApiOpenApi2AssetsCreativeComponentUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AssetsCreativeComponentUpdateV2Response
func (a *AssetsCreativeComponentUpdateV2ApiService) postExecute(r *ApiOpenApi2AssetsCreativeComponentUpdatePostRequest) (*AssetsCreativeComponentUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *AssetsCreativeComponentUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/assets/creative_component/update/"

	localVarHeaderParams := make(map[string]string)
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
	localVarPostBody = r.assetsCreativeComponentUpdateV2Request
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
