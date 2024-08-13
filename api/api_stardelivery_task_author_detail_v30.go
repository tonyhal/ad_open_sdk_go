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

// StardeliveryTaskAuthorDetailV30ApiService StardeliveryTaskAuthorDetailV30Api service
type StardeliveryTaskAuthorDetailV30ApiService service

type ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest struct {
	ctx          context.Context
	ApiService   *StardeliveryTaskAuthorDetailV30ApiService
	advertiserId *int64
	starTaskId   *int64
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) StarTaskId(starTaskId int64) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	r.starTaskId = &starTaskId
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) Page(page int64) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) PageSize(pageSize int64) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) Execute() (*StardeliveryTaskAuthorDetailV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) AccessToken(accessToken string) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) WithLog(enable bool) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30StardeliveryTaskAuthorDetailGet Method for OpenApiV30StardeliveryTaskAuthorDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest
*/
func (a *StardeliveryTaskAuthorDetailV30ApiService) Get(ctx context.Context) *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest {
	return &ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StardeliveryTaskAuthorDetailV30Response
func (a *StardeliveryTaskAuthorDetailV30ApiService) getExecute(r *ApiOpenApiV30StardeliveryTaskAuthorDetailGetRequest) (*StardeliveryTaskAuthorDetailV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StardeliveryTaskAuthorDetailV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/stardelivery/task_author/detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.starTaskId == nil {
		return localVarReturnValue, nil, ReportError("starTaskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_id", r.starTaskId)
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
