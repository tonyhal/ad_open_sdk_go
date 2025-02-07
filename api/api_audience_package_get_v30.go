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

// AudiencePackageGetV30ApiService AudiencePackageGetV30Api service
type AudiencePackageGetV30ApiService service

type ApiOpenApiV30AudiencePackageGetGetRequest struct {
	ctx          context.Context
	ApiService   *AudiencePackageGetV30ApiService
	advertiserId *int64
	filtering    *AudiencePackageGetV30Filtering
	page         *int64
	pageSize     *int64
}

// 广告主ID， advertiserId 鉴权，看其有无访问/操作数据的权限
func (r *ApiOpenApiV30AudiencePackageGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30AudiencePackageGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤器
func (r *ApiOpenApiV30AudiencePackageGetGetRequest) Filtering(filtering AudiencePackageGetV30Filtering) *ApiOpenApiV30AudiencePackageGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30AudiencePackageGetGetRequest) Page(page int64) *ApiOpenApiV30AudiencePackageGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30AudiencePackageGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30AudiencePackageGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30AudiencePackageGetGetRequest) Execute() (*AudiencePackageGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30AudiencePackageGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30AudiencePackageGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AudiencePackageGetGetRequest) WithLog(enable bool) *ApiOpenApiV30AudiencePackageGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AudiencePackageGetGet Method for OpenApiV30AudiencePackageGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AudiencePackageGetGetRequest
*/
func (a *AudiencePackageGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30AudiencePackageGetGetRequest {
	return &ApiOpenApiV30AudiencePackageGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AudiencePackageGetV30Response
func (a *AudiencePackageGetV30ApiService) getExecute(r *ApiOpenApiV30AudiencePackageGetGetRequest) (*AudiencePackageGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AudiencePackageGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/audience_package/get/"

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
