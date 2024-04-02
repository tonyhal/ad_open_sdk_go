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

// QianchuanFileVideoDeleteV10ApiService QianchuanFileVideoDeleteV10Api service
type QianchuanFileVideoDeleteV10ApiService service

type ApiOpenApiV10QianchuanFileVideoDeletePostRequest struct {
	ctx                                context.Context
	ApiService                         *QianchuanFileVideoDeleteV10ApiService
	qianchuanFileVideoDeleteV10Request *QianchuanFileVideoDeleteV10Request
}

func (r *ApiOpenApiV10QianchuanFileVideoDeletePostRequest) QianchuanFileVideoDeleteV10Request(qianchuanFileVideoDeleteV10Request QianchuanFileVideoDeleteV10Request) *ApiOpenApiV10QianchuanFileVideoDeletePostRequest {
	r.qianchuanFileVideoDeleteV10Request = &qianchuanFileVideoDeleteV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoDeletePostRequest) Execute() (*QianchuanFileVideoDeleteV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanFileVideoDeletePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanFileVideoDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoDeletePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanFileVideoDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanFileVideoDeletePost Method for OpenApiV10QianchuanFileVideoDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanFileVideoDeletePostRequest
*/
func (a *QianchuanFileVideoDeleteV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanFileVideoDeletePostRequest {
	return &ApiOpenApiV10QianchuanFileVideoDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanFileVideoDeleteV10Response
func (a *QianchuanFileVideoDeleteV10ApiService) postExecute(r *ApiOpenApiV10QianchuanFileVideoDeletePostRequest) (*QianchuanFileVideoDeleteV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanFileVideoDeleteV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/file/video/delete/"

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
	localVarPostBody = r.qianchuanFileVideoDeleteV10Request
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
