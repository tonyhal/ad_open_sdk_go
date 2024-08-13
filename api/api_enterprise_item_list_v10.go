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

// EnterpriseItemListV10ApiService EnterpriseItemListV10Api service
type EnterpriseItemListV10ApiService service

type ApiOpenApiV10EnterpriseItemListGetRequest struct {
	ctx         context.Context
	ApiService  *EnterpriseItemListV10ApiService
	ccAccountId *int64
	eDouyinId   *string
	endTime     *string
	filter      *EnterpriseItemListV10Filter
	page        *int64
	pageSize    *int64
	startTime   *string
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) CcAccountId(ccAccountId int64) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.ccAccountId = &ccAccountId
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) EDouyinId(eDouyinId string) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.eDouyinId = &eDouyinId
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) EndTime(endTime string) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) Filter(filter EnterpriseItemListV10Filter) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) Page(page int64) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) PageSize(pageSize int64) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) StartTime(startTime string) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) Execute() (*EnterpriseItemListV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseItemListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseItemListGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseItemListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseItemListGet Method for OpenApiV10EnterpriseItemListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseItemListGetRequest
*/
func (a *EnterpriseItemListV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseItemListGetRequest {
	return &ApiOpenApiV10EnterpriseItemListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseItemListV10Response
func (a *EnterpriseItemListV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseItemListGetRequest) (*EnterpriseItemListV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EnterpriseItemListV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/item/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ccAccountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cc_account_id", r.ccAccountId)
	}
	if r.eDouyinId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "e_douyin_id", r.eDouyinId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
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
