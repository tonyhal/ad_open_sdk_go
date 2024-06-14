/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

// DecorationCouponGetV30ApiService DecorationCouponGetV30Api service
type DecorationCouponGetV30ApiService service

type ApiOpenApiV30DecorationCouponGetGetRequest struct {
	ctx          context.Context
	ApiService   *DecorationCouponGetV30ApiService
	advertiserId *int64
	filtering    *DecorationCouponGetV30Filtering
	page         *int32
	pageSize     *int32
}

// 广告账户id
func (r *ApiOpenApiV30DecorationCouponGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30DecorationCouponGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件
func (r *ApiOpenApiV30DecorationCouponGetGetRequest) Filtering(filtering DecorationCouponGetV30Filtering) *ApiOpenApiV30DecorationCouponGetGetRequest {
	r.filtering = &filtering
	return r
}

// 页码，默认值为1
func (r *ApiOpenApiV30DecorationCouponGetGetRequest) Page(page int32) *ApiOpenApiV30DecorationCouponGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值为10，最大不超过50
func (r *ApiOpenApiV30DecorationCouponGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30DecorationCouponGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30DecorationCouponGetGetRequest) Execute() (*DecorationCouponGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30DecorationCouponGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30DecorationCouponGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30DecorationCouponGetGetRequest) WithLog(enable bool) *ApiOpenApiV30DecorationCouponGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30DecorationCouponGetGet Method for OpenApiV30DecorationCouponGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30DecorationCouponGetGetRequest
*/
func (a *DecorationCouponGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30DecorationCouponGetGetRequest {
	return &ApiOpenApiV30DecorationCouponGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DecorationCouponGetV30Response
func (a *DecorationCouponGetV30ApiService) getExecute(r *ApiOpenApiV30DecorationCouponGetGetRequest) (*DecorationCouponGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DecorationCouponGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/decoration/coupon/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
