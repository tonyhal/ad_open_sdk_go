/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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

// DmpCustomAudienceDeleteV2ApiService DmpCustomAudienceDeleteV2Api service
type DmpCustomAudienceDeleteV2ApiService service

type ApiOpenApi2DmpCustomAudienceDeletePostRequest struct {
	ctx                              context.Context
	ApiService                       *DmpCustomAudienceDeleteV2ApiService
	dmpCustomAudienceDeleteV2Request *DmpCustomAudienceDeleteV2Request
}

func (r *ApiOpenApi2DmpCustomAudienceDeletePostRequest) DmpCustomAudienceDeleteV2Request(dmpCustomAudienceDeleteV2Request DmpCustomAudienceDeleteV2Request) *ApiOpenApi2DmpCustomAudienceDeletePostRequest {
	r.dmpCustomAudienceDeleteV2Request = &dmpCustomAudienceDeleteV2Request
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceDeletePostRequest) Execute() (*DmpCustomAudienceDeleteV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2DmpCustomAudienceDeletePostRequest) WithLog(enable bool) *ApiOpenApi2DmpCustomAudienceDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DmpCustomAudienceDeletePost Method for OpenApi2DmpCustomAudienceDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DmpCustomAudienceDeletePostRequest
*/
func (a *DmpCustomAudienceDeleteV2ApiService) Post(ctx context.Context) *ApiOpenApi2DmpCustomAudienceDeletePostRequest {
	return &ApiOpenApi2DmpCustomAudienceDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DmpCustomAudienceDeleteV2Response
func (a *DmpCustomAudienceDeleteV2ApiService) postExecute(r *ApiOpenApi2DmpCustomAudienceDeletePostRequest) (*DmpCustomAudienceDeleteV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DmpCustomAudienceDeleteV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dmp/custom_audience/delete/"

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
	localVarPostBody = r.dmpCustomAudienceDeleteV2Request
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
