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

// SecurityScoreViolationEventGetV30ApiService SecurityScoreViolationEventGetV30Api service
type SecurityScoreViolationEventGetV30ApiService service

type ApiOpenApiV30SecurityScoreViolationEventGetGetRequest struct {
	ctx          context.Context
	ApiService   *SecurityScoreViolationEventGetV30ApiService
	advertiserId *int64
	businessLine *SecurityScoreViolationEventGetV30BusinessLine
	page         *int64
	pageSize     *int64
	filtering    *SecurityScoreViolationEventGetV30Filtering
}

// 广告主id
func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 业务线
func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) BusinessLine(businessLine SecurityScoreViolationEventGetV30BusinessLine) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	r.businessLine = &businessLine
	return r
}

// 第几页
func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) Page(page int64) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	r.page = &page
	return r
}

// 每页的数目
func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 过滤器
func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) Filtering(filtering SecurityScoreViolationEventGetV30Filtering) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) Execute() (*SecurityScoreViolationEventGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) WithLog(enable bool) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SecurityScoreViolationEventGetGet Method for OpenApiV30SecurityScoreViolationEventGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SecurityScoreViolationEventGetGetRequest
*/
func (a *SecurityScoreViolationEventGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest {
	return &ApiOpenApiV30SecurityScoreViolationEventGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SecurityScoreViolationEventGetV30Response
func (a *SecurityScoreViolationEventGetV30ApiService) getExecute(r *ApiOpenApiV30SecurityScoreViolationEventGetGetRequest) (*SecurityScoreViolationEventGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SecurityScoreViolationEventGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/security/score_violation_event/get/"

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
	parameterAddToHeaderOrQuery(localVarQueryParams, "business_line", r.businessLine)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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
