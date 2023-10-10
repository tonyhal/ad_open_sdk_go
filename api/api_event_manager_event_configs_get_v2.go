/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
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

// EventManagerEventConfigsGetV2ApiService EventManagerEventConfigsGetV2Api service
type EventManagerEventConfigsGetV2ApiService service

type ApiOpenApi2EventManagerEventConfigsGetGetRequest struct {
	ctx          context.Context
	ApiService   *EventManagerEventConfigsGetV2ApiService
	advertiserId *int64
	assetId      *int64
	sortType     *EventManagerEventConfigsGetV2SortType
}

// 广告主ID
func (r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2EventManagerEventConfigsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 资产ID
func (r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) AssetId(assetId int64) *ApiOpenApi2EventManagerEventConfigsGetGetRequest {
	r.assetId = &assetId
	return r
}

// 创建时间排序方式，允许值：&#x60;DESC&#x60; 降序、&#x60;ASC&#x60; 升序。默认：&#x60;ASC&#x60;
func (r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) SortType(sortType EventManagerEventConfigsGetV2SortType) *ApiOpenApi2EventManagerEventConfigsGetGetRequest {
	r.sortType = &sortType
	return r
}

func (r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) Execute() (*EventManagerEventConfigsGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2EventManagerEventConfigsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) WithLog(enable bool) *ApiOpenApi2EventManagerEventConfigsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2EventManagerEventConfigsGetGet Method for OpenApi2EventManagerEventConfigsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2EventManagerEventConfigsGetGetRequest
*/
func (a *EventManagerEventConfigsGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2EventManagerEventConfigsGetGetRequest {
	return &ApiOpenApi2EventManagerEventConfigsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerEventConfigsGetV2Response
func (a *EventManagerEventConfigsGetV2ApiService) getExecute(r *ApiOpenApi2EventManagerEventConfigsGetGetRequest) (*EventManagerEventConfigsGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *EventManagerEventConfigsGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/event_manager/event_configs/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if *r.advertiserId > -9223372036854775616 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than -9223372036854775616")
	}
	if r.assetId == nil {
		return localVarReturnValue, nil, ReportError("assetId is required and must be specified")
	}
	if *r.assetId < 1 {
		return localVarReturnValue, nil, ReportError("assetId must be greater than 1")
	}
	if *r.assetId > -9223372036854775616 {
		return localVarReturnValue, nil, ReportError("assetId must be less than -9223372036854775616")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId)
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_type", r.sortType)
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
