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

// ReportLiveRoomFlowCategoryGetV2ApiService ReportLiveRoomFlowCategoryGetV2Api service
type ReportLiveRoomFlowCategoryGetV2ApiService service

type ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportLiveRoomFlowCategoryGetV2ApiService
	advertiserId *int64
	endTime      **string
	fields       *[]string
	filtering    *ReportLiveRoomFlowCategoryGetV2Filtering
	orderField   *ReportLiveRoomFlowCategoryGetV2OrderField
	orderType    *ReportLiveRoomFlowCategoryGetV2OrderType
	startTime    **string
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) EndTime(endTime *string) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) Fields(fields []string) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) Filtering(filtering ReportLiveRoomFlowCategoryGetV2Filtering) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) OrderField(orderField ReportLiveRoomFlowCategoryGetV2OrderField) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) OrderType(orderType ReportLiveRoomFlowCategoryGetV2OrderType) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) StartTime(startTime *string) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) Execute() (*ReportLiveRoomFlowCategoryGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportLiveRoomFlowCategoryGetGet Method for OpenApi2ReportLiveRoomFlowCategoryGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest
*/
func (a *ReportLiveRoomFlowCategoryGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest {
	return &ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportLiveRoomFlowCategoryGetV2Response
func (a *ReportLiveRoomFlowCategoryGetV2ApiService) getExecute(r *ApiOpenApi2ReportLiveRoomFlowCategoryGetGetRequest) (*ReportLiveRoomFlowCategoryGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportLiveRoomFlowCategoryGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/live_room/flow_category/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
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
