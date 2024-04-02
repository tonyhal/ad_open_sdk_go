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

// EnterpriseCommentListGetV10ApiService EnterpriseCommentListGetV10Api service
type EnterpriseCommentListGetV10ApiService service

type ApiOpenApiV10EnterpriseCommentListGetGetRequest struct {
	ctx         context.Context
	ApiService  *EnterpriseCommentListGetV10ApiService
	ccAccountId *int64
	eDouyinId   *string
	endTime     *string
	filter      *EnterpriseCommentListGetV10Filter
	orderField  *string
	orderType   *EnterpriseCommentListGetV10OrderType
	page        *int64
	pageSize    *int64
	startTime   *string
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) CcAccountId(ccAccountId int64) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.ccAccountId = &ccAccountId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) EDouyinId(eDouyinId string) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.eDouyinId = &eDouyinId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) EndTime(endTime string) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) Filter(filter EnterpriseCommentListGetV10Filter) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) OrderField(orderField string) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) OrderType(orderType EnterpriseCommentListGetV10OrderType) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) Page(page int64) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) StartTime(startTime string) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) Execute() (*EnterpriseCommentListGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseCommentListGetGet Method for OpenApiV10EnterpriseCommentListGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseCommentListGetGetRequest
*/
func (a *EnterpriseCommentListGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseCommentListGetGetRequest {
	return &ApiOpenApiV10EnterpriseCommentListGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseCommentListGetV10Response
func (a *EnterpriseCommentListGetV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseCommentListGetGetRequest) (*EnterpriseCommentListGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EnterpriseCommentListGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/comment/list/get/"

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
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
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
