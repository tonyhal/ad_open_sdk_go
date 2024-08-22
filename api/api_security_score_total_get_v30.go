/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// SecurityScoreTotalGetV30ApiService SecurityScoreTotalGetV30Api service
type SecurityScoreTotalGetV30ApiService service

type ApiOpenApiV30SecurityScoreTotalGetGetRequest struct {
	ctx          context.Context
	ApiService   *SecurityScoreTotalGetV30ApiService
	advertiserId *int64
	businessLine *SecurityScoreTotalGetV30BusinessLine
	page         *int64
	pageSize     *int64
	filtering    *SecurityScoreTotalGetV30Filtering
}

func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 业务线
func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) BusinessLine(businessLine SecurityScoreTotalGetV30BusinessLine) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	r.businessLine = &businessLine
	return r
}

// 页码
func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) Page(page int64) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 过滤器
func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) Filtering(filtering SecurityScoreTotalGetV30Filtering) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) Execute() (*SecurityScoreTotalGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SecurityScoreTotalGetGet Method for OpenApiV30SecurityScoreTotalGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SecurityScoreTotalGetGetRequest
*/
func (a *SecurityScoreTotalGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SecurityScoreTotalGetGetRequest {
	return &ApiOpenApiV30SecurityScoreTotalGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SecurityScoreTotalGetV30Response
func (a *SecurityScoreTotalGetV30ApiService) getExecute(r *ApiOpenApiV30SecurityScoreTotalGetGetRequest) (*SecurityScoreTotalGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SecurityScoreTotalGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/security/score_total/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.businessLine == nil {
		return localVarReturnValue, nil, ReportError("businessLine is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "business_line", r.businessLine)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
