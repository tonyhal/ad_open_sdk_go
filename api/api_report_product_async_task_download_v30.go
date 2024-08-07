/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// ReportProductAsyncTaskDownloadV30ApiService ReportProductAsyncTaskDownloadV30Api service
type ReportProductAsyncTaskDownloadV30ApiService service

type ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest struct {
	ctx          context.Context
	ApiService   *ReportProductAsyncTaskDownloadV30ApiService
	advertiserId *int64
	taskId       *int64
	rangeFrom    *int64
	rangeTo      *int64
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) TaskId(taskId int64) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	r.taskId = &taskId
	return r
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) RangeFrom(rangeFrom int64) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	r.rangeFrom = &rangeFrom
	return r
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) RangeTo(rangeTo int64) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	r.rangeTo = &rangeTo
	return r
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) Execute() (*ReportProductAsyncTaskDownloadV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportProductAsyncTaskDownloadGet Method for OpenApiV30ReportProductAsyncTaskDownloadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest
*/
func (a *ReportProductAsyncTaskDownloadV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest {
	return &ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportProductAsyncTaskDownloadV30Response
func (a *ReportProductAsyncTaskDownloadV30ApiService) getExecute(r *ApiOpenApiV30ReportProductAsyncTaskDownloadGetRequest) (*ReportProductAsyncTaskDownloadV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportProductAsyncTaskDownloadV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/product/async_task/download/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.taskId == nil {
		return localVarReturnValue, nil, ReportError("taskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "task_id", r.taskId)
	if r.rangeFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range_from", r.rangeFrom)
	}
	if r.rangeTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range_to", r.rangeTo)
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
