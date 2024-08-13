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

// StarChallengePushItemsToAdV2ApiService StarChallengePushItemsToAdV2Api service
type StarChallengePushItemsToAdV2ApiService service

type ApiOpenApi2StarChallengePushItemsToAdPostRequest struct {
	ctx                                 context.Context
	ApiService                          *StarChallengePushItemsToAdV2ApiService
	starChallengePushItemsToAdV2Request *StarChallengePushItemsToAdV2Request
}

func (r *ApiOpenApi2StarChallengePushItemsToAdPostRequest) StarChallengePushItemsToAdV2Request(starChallengePushItemsToAdV2Request StarChallengePushItemsToAdV2Request) *ApiOpenApi2StarChallengePushItemsToAdPostRequest {
	r.starChallengePushItemsToAdV2Request = &starChallengePushItemsToAdV2Request
	return r
}

func (r *ApiOpenApi2StarChallengePushItemsToAdPostRequest) Execute() (*StarChallengePushItemsToAdV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarChallengePushItemsToAdPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengePushItemsToAdPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengePushItemsToAdPostRequest) WithLog(enable bool) *ApiOpenApi2StarChallengePushItemsToAdPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengePushItemsToAdPost Method for OpenApi2StarChallengePushItemsToAdPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengePushItemsToAdPostRequest
*/
func (a *StarChallengePushItemsToAdV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarChallengePushItemsToAdPostRequest {
	return &ApiOpenApi2StarChallengePushItemsToAdPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengePushItemsToAdV2Response
func (a *StarChallengePushItemsToAdV2ApiService) postExecute(r *ApiOpenApi2StarChallengePushItemsToAdPostRequest) (*StarChallengePushItemsToAdV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengePushItemsToAdV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/push_items_to_ad/"

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
	localVarPostBody = r.starChallengePushItemsToAdV2Request
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
