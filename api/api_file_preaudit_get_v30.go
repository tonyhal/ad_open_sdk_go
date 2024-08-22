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

// FilePreauditGetV30ApiService FilePreauditGetV30Api service
type FilePreauditGetV30ApiService service

type ApiOpenApiV30FilePreauditGetGetRequest struct {
	ctx          context.Context
	ApiService   *FilePreauditGetV30ApiService
	advertiserId *int64
	filtering    *FilePreauditGetV30Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30FilePreauditGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) Filtering(filtering FilePreauditGetV30Filtering) *ApiOpenApiV30FilePreauditGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) Page(page int64) *ApiOpenApiV30FilePreauditGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30FilePreauditGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) Execute() (*FilePreauditGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30FilePreauditGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30FilePreauditGetGetRequest) WithLog(enable bool) *ApiOpenApiV30FilePreauditGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30FilePreauditGetGet Method for OpenApiV30FilePreauditGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30FilePreauditGetGetRequest
*/
func (a *FilePreauditGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30FilePreauditGetGetRequest {
	return &ApiOpenApiV30FilePreauditGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FilePreauditGetV30Response
func (a *FilePreauditGetV30ApiService) getExecute(r *ApiOpenApiV30FilePreauditGetGetRequest) (*FilePreauditGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FilePreauditGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/file/preaudit/get/"

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
