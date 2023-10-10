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

// FileVideoGetV2ApiService FileVideoGetV2Api service
type FileVideoGetV2ApiService service

type ApiOpenApi2FileVideoGetGetRequest struct {
	ctx          context.Context
	ApiService   *FileVideoGetV2ApiService
	advertiserId *int64
	filtering    *FileVideoGetV2Filtering
	page         *int64
	pageSize     *int64
}

// 广告主ID
func (r *ApiOpenApi2FileVideoGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileVideoGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 视频过滤条件
func (r *ApiOpenApi2FileVideoGetGetRequest) Filtering(filtering FileVideoGetV2Filtering) *ApiOpenApi2FileVideoGetGetRequest {
	r.filtering = &filtering
	return r
}

// 页码，默认值1
func (r *ApiOpenApi2FileVideoGetGetRequest) Page(page int64) *ApiOpenApi2FileVideoGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值20
func (r *ApiOpenApi2FileVideoGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2FileVideoGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2FileVideoGetGetRequest) Execute() (*FileVideoGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileVideoGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoGetGetRequest) WithLog(enable bool) *ApiOpenApi2FileVideoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoGetGet Method for OpenApi2FileVideoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoGetGetRequest
*/
func (a *FileVideoGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileVideoGetGetRequest {
	return &ApiOpenApi2FileVideoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoGetV2Response
func (a *FileVideoGetV2ApiService) getExecute(r *ApiOpenApi2FileVideoGetGetRequest) (*FileVideoGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *FileVideoGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/get/"

	localVarHeaderParams := make(map[string]string)
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
