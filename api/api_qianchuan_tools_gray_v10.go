/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
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

// QianchuanToolsGrayV10ApiService QianchuanToolsGrayV10Api service
type QianchuanToolsGrayV10ApiService service

type ApiOpenApiV10QianchuanToolsGrayGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanToolsGrayV10ApiService
	advertiserId *int64
	grayKeys     *[]string
	awemeIds     *[]int64
}

func (r *ApiOpenApiV10QianchuanToolsGrayGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanToolsGrayGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanToolsGrayGetRequest) GrayKeys(grayKeys []string) *ApiOpenApiV10QianchuanToolsGrayGetRequest {
	r.grayKeys = &grayKeys
	return r
}

func (r *ApiOpenApiV10QianchuanToolsGrayGetRequest) AwemeIds(awemeIds []int64) *ApiOpenApiV10QianchuanToolsGrayGetRequest {
	r.awemeIds = &awemeIds
	return r
}

func (r *ApiOpenApiV10QianchuanToolsGrayGetRequest) Execute() (*QianchuanToolsGrayV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsGrayGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsGrayGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsGrayGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsGrayGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsGrayGet Method for OpenApiV10QianchuanToolsGrayGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsGrayGetRequest
*/
func (a *QianchuanToolsGrayV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanToolsGrayGetRequest {
	return &ApiOpenApiV10QianchuanToolsGrayGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsGrayV10Response
func (a *QianchuanToolsGrayV10ApiService) getExecute(r *ApiOpenApiV10QianchuanToolsGrayGetRequest) (*QianchuanToolsGrayV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanToolsGrayV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/gray/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.grayKeys == nil {
		return localVarReturnValue, nil, ReportError("grayKeys is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "gray_keys", r.grayKeys)
	if r.awemeIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_ids", r.awemeIds)
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
