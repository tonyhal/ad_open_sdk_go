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

// QianchuanAwemeAuthListGetV10ApiService QianchuanAwemeAuthListGetV10Api service
type QianchuanAwemeAuthListGetV10ApiService service

type ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeAuthListGetV10ApiService
	advertiserId *int64
	filtering    *QianchuanAwemeAuthListGetV10Filtering
	page         *int64
	pageSize     *QianchuanAwemeAuthListGetV10PageSize
}

func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤器
func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) Filtering(filtering QianchuanAwemeAuthListGetV10Filtering) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) Page(page int64) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) PageSize(pageSize QianchuanAwemeAuthListGetV10PageSize) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) Execute() (*QianchuanAwemeAuthListGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeAuthListGetGet Method for OpenApiV10QianchuanAwemeAuthListGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest
*/
func (a *QianchuanAwemeAuthListGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest {
	return &ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeAuthListGetV10Response
func (a *QianchuanAwemeAuthListGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeAuthListGetGetRequest) (*QianchuanAwemeAuthListGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeAuthListGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme_auth_list/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
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
