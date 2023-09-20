/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

// ReportLiveRoomAudiencePortraitGetV2ApiService ReportLiveRoomAudiencePortraitGetV2Api service
type ReportLiveRoomAudiencePortraitGetV2ApiService service

type ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportLiveRoomAudiencePortraitGetV2ApiService
	advertiserId *int64
	dimension    *ReportLiveRoomAudiencePortraitGetV2Dimension
	endTime      **string
	fields       *[]string
	filtering    *ReportLiveRoomAudiencePortraitGetV2Filtering
	orderField   *ReportLiveRoomAudiencePortraitGetV2OrderField
	orderType    *ReportLiveRoomAudiencePortraitGetV2OrderType
	page         *int64
	pageSize     *int64
	startTime    **string
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) Dimension(dimension ReportLiveRoomAudiencePortraitGetV2Dimension) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.dimension = &dimension
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) EndTime(endTime *string) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) Fields(fields []string) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) Filtering(filtering ReportLiveRoomAudiencePortraitGetV2Filtering) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) OrderField(orderField ReportLiveRoomAudiencePortraitGetV2OrderField) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) OrderType(orderType ReportLiveRoomAudiencePortraitGetV2OrderType) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) Page(page int64) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) StartTime(startTime *string) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) Execute() (*ReportLiveRoomAudiencePortraitGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportLiveRoomAudiencePortraitGetGet Method for OpenApi2ReportLiveRoomAudiencePortraitGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest
*/
func (a *ReportLiveRoomAudiencePortraitGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest {
	return &ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportLiveRoomAudiencePortraitGetV2Response
func (a *ReportLiveRoomAudiencePortraitGetV2ApiService) getExecute(r *ApiOpenApi2ReportLiveRoomAudiencePortraitGetGetRequest) (*ReportLiveRoomAudiencePortraitGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportLiveRoomAudiencePortraitGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/live_room/audience/portrait/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.dimension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dimension", r.dimension)
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
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
