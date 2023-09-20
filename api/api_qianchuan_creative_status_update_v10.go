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

// QianchuanCreativeStatusUpdateV10ApiService QianchuanCreativeStatusUpdateV10Api service
type QianchuanCreativeStatusUpdateV10ApiService service

type ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest struct {
	ctx                                     context.Context
	ApiService                              *QianchuanCreativeStatusUpdateV10ApiService
	qianchuanCreativeStatusUpdateV10Request *QianchuanCreativeStatusUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest) QianchuanCreativeStatusUpdateV10Request(qianchuanCreativeStatusUpdateV10Request QianchuanCreativeStatusUpdateV10Request) *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest {
	r.qianchuanCreativeStatusUpdateV10Request = &qianchuanCreativeStatusUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest) Execute() (*QianchuanCreativeStatusUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanCreativeStatusUpdatePost Method for OpenApiV10QianchuanCreativeStatusUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest
*/
func (a *QianchuanCreativeStatusUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest {
	return &ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanCreativeStatusUpdateV10Response
func (a *QianchuanCreativeStatusUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanCreativeStatusUpdatePostRequest) (*QianchuanCreativeStatusUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanCreativeStatusUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/creative/status/update/"

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
	localVarPostBody = r.qianchuanCreativeStatusUpdateV10Request
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
