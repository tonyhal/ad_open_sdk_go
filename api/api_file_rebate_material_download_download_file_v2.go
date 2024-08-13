/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
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

// FileRebateMaterialDownloadDownloadFileV2ApiService FileRebateMaterialDownloadDownloadFileV2Api service
type FileRebateMaterialDownloadDownloadFileV2ApiService service

type ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest struct {
	ctx        context.Context
	ApiService *FileRebateMaterialDownloadDownloadFileV2ApiService
	agentId    *int64
	taskId     *string
}

// 代理商帐户ID
func (r *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest) AgentId(agentId int64) *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest {
	r.agentId = &agentId
	return r
}

// 任务ID
func (r *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest) TaskId(taskId string) *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest {
	r.taskId = &taskId
	return r
}

func (r *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest) Execute() (*FileRebateMaterialDownloadDownloadFileV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest) WithLog(enable bool) *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileRebateMaterialDownloadDownloadFileGet Method for OpenApi2FileRebateMaterialDownloadDownloadFileGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest
*/
func (a *FileRebateMaterialDownloadDownloadFileV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest {
	return &ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileRebateMaterialDownloadDownloadFileV2Response
func (a *FileRebateMaterialDownloadDownloadFileV2ApiService) getExecute(r *ApiOpenApi2FileRebateMaterialDownloadDownloadFileGetRequest) (*FileRebateMaterialDownloadDownloadFileV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileRebateMaterialDownloadDownloadFileV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/rebate/material_download/download_file/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.taskId == nil {
		return localVarReturnValue, nil, ReportError("taskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "task_id", r.taskId)
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
