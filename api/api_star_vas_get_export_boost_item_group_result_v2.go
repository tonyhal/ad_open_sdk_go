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

// StarVasGetExportBoostItemGroupResultV2ApiService StarVasGetExportBoostItemGroupResultV2Api service
type StarVasGetExportBoostItemGroupResultV2ApiService service

type ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest struct {
	ctx        context.Context
	ApiService *StarVasGetExportBoostItemGroupResultV2ApiService
	starId     *int64
	ticketId   *int64
}

// 客户ID
func (r *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest) StarId(starId int64) *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest {
	r.starId = &starId
	return r
}

// 导出任务唯一标识
func (r *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest) TicketId(ticketId int64) *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest {
	r.ticketId = &ticketId
	return r
}

func (r *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest) Execute() (*StarVasGetExportBoostItemGroupResultV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest) WithLog(enable bool) *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarVasGetExportBoostItemGroupResultGet Method for OpenApi2StarVasGetExportBoostItemGroupResultGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest
*/
func (a *StarVasGetExportBoostItemGroupResultV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest {
	return &ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarVasGetExportBoostItemGroupResultV2Response
func (a *StarVasGetExportBoostItemGroupResultV2ApiService) getExecute(r *ApiOpenApi2StarVasGetExportBoostItemGroupResultGetRequest) (*StarVasGetExportBoostItemGroupResultV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarVasGetExportBoostItemGroupResultV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/vas/get_export_boost_item_group_result/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.ticketId == nil {
		return localVarReturnValue, nil, ReportError("ticketId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ticket_id", r.ticketId)
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
