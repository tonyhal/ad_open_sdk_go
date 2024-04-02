/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

// StarReportOrderOverviewV2ApiService StarReportOrderOverviewV2Api service
type StarReportOrderOverviewV2ApiService service

type ApiOpenApi2StarReportOrderOverviewGetRequest struct {
	ctx        context.Context
	ApiService *StarReportOrderOverviewV2ApiService
	starId     *int64
	orderIds   *[]int64
}

// 客户id
func (r *ApiOpenApi2StarReportOrderOverviewGetRequest) StarId(starId int64) *ApiOpenApi2StarReportOrderOverviewGetRequest {
	r.starId = &starId
	return r
}

// 订单id列表
func (r *ApiOpenApi2StarReportOrderOverviewGetRequest) OrderIds(orderIds []int64) *ApiOpenApi2StarReportOrderOverviewGetRequest {
	r.orderIds = &orderIds
	return r
}

func (r *ApiOpenApi2StarReportOrderOverviewGetRequest) Execute() (*StarReportOrderOverviewV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarReportOrderOverviewGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarReportOrderOverviewGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarReportOrderOverviewGetRequest) WithLog(enable bool) *ApiOpenApi2StarReportOrderOverviewGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarReportOrderOverviewGet Method for OpenApi2StarReportOrderOverviewGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarReportOrderOverviewGetRequest
*/
func (a *StarReportOrderOverviewV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarReportOrderOverviewGetRequest {
	return &ApiOpenApi2StarReportOrderOverviewGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarReportOrderOverviewV2Response
func (a *StarReportOrderOverviewV2ApiService) getExecute(r *ApiOpenApi2StarReportOrderOverviewGetRequest) (*StarReportOrderOverviewV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarReportOrderOverviewV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/report/order_overview/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.orderIds == nil {
		return localVarReturnValue, nil, ReportError("orderIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "order_ids", r.orderIds)
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
