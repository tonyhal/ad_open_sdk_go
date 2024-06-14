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

// StarMcnGetUnparticipatedTaskV2ApiService StarMcnGetUnparticipatedTaskV2Api service
type StarMcnGetUnparticipatedTaskV2ApiService service

type ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest struct {
	ctx                context.Context
	ApiService         *StarMcnGetUnparticipatedTaskV2ApiService
	starId             *int64
	page               *int32
	pageSize           *int32
	payType            *int32
	minCreateTimeStamp *int64
	maxCreateTimeStamp *int64
	developerId        *int64
	settlementType     *int32
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) StarId(starId int64) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) Page(page int32) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) PageSize(pageSize int32) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) PayType(payType int32) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.payType = &payType
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) MinCreateTimeStamp(minCreateTimeStamp int64) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.minCreateTimeStamp = &minCreateTimeStamp
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) MaxCreateTimeStamp(maxCreateTimeStamp int64) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.maxCreateTimeStamp = &maxCreateTimeStamp
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) DeveloperId(developerId int64) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.developerId = &developerId
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) SettlementType(settlementType int32) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.settlementType = &settlementType
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) Execute() (*StarMcnGetUnparticipatedTaskV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) WithLog(enable bool) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarMcnGetUnparticipatedTaskGet Method for OpenApi2StarMcnGetUnparticipatedTaskGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest
*/
func (a *StarMcnGetUnparticipatedTaskV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest {
	return &ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarMcnGetUnparticipatedTaskV2Response
func (a *StarMcnGetUnparticipatedTaskV2ApiService) getExecute(r *ApiOpenApi2StarMcnGetUnparticipatedTaskGetRequest) (*StarMcnGetUnparticipatedTaskV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarMcnGetUnparticipatedTaskV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/mcn/get_unparticipated_task/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.payType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pay_type", r.payType)
	}
	if r.minCreateTimeStamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "min_create_time_stamp", r.minCreateTimeStamp)
	}
	if r.maxCreateTimeStamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_create_time_stamp", r.maxCreateTimeStamp)
	}
	if r.developerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer_id", r.developerId)
	}
	if r.settlementType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "settlement_type", r.settlementType)
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
