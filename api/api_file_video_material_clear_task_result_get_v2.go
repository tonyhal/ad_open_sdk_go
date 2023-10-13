/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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

// FileVideoMaterialClearTaskResultGetV2ApiService FileVideoMaterialClearTaskResultGetV2Api service
type FileVideoMaterialClearTaskResultGetV2ApiService service

type ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest struct {
	ctx          context.Context
	ApiService   *FileVideoMaterialClearTaskResultGetV2ApiService
	advertiserId *int64
	clearId      *int64
	page         *int64
	pageSize     *int64
}

// 广告主id
func (r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 清理任务id
func (r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) ClearId(clearId int64) *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest {
	r.clearId = &clearId
	return r
}

// 页码，默认1
func (r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) Page(page int64) *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认10，最大100
func (r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) Execute() (*FileVideoMaterialClearTaskResultGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) WithLog(enable bool) *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoMaterialClearTaskResultGetGet Method for OpenApi2FileVideoMaterialClearTaskResultGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest
*/
func (a *FileVideoMaterialClearTaskResultGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest {
	return &ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoMaterialClearTaskResultGetV2Response
func (a *FileVideoMaterialClearTaskResultGetV2ApiService) getExecute(r *ApiOpenApi2FileVideoMaterialClearTaskResultGetGetRequest) (*FileVideoMaterialClearTaskResultGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *FileVideoMaterialClearTaskResultGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/material/clear_task_result/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.clearId == nil {
		return localVarReturnValue, nil, ReportError("clearId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "clear_id", r.clearId)
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
