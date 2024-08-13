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

// QianchuanUniAwemeAdCreateV10ApiService QianchuanUniAwemeAdCreateV10Api service
type QianchuanUniAwemeAdCreateV10ApiService service

type ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest struct {
	ctx                                 context.Context
	ApiService                          *QianchuanUniAwemeAdCreateV10ApiService
	qianchuanUniAwemeAdCreateV10Request *QianchuanUniAwemeAdCreateV10Request
}

func (r *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest) QianchuanUniAwemeAdCreateV10Request(qianchuanUniAwemeAdCreateV10Request QianchuanUniAwemeAdCreateV10Request) *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest {
	r.qianchuanUniAwemeAdCreateV10Request = &qianchuanUniAwemeAdCreateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest) Execute() (*QianchuanUniAwemeAdCreateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniAwemeAdCreatePost Method for OpenApiV10QianchuanUniAwemeAdCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest
*/
func (a *QianchuanUniAwemeAdCreateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest {
	return &ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniAwemeAdCreateV10Response
func (a *QianchuanUniAwemeAdCreateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanUniAwemeAdCreatePostRequest) (*QianchuanUniAwemeAdCreateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanUniAwemeAdCreateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_aweme/ad/create/"

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
	localVarPostBody = r.qianchuanUniAwemeAdCreateV10Request
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
