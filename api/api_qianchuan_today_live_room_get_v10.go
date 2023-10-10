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

// QianchuanTodayLiveRoomGetV10ApiService QianchuanTodayLiveRoomGetV10Api service
type QianchuanTodayLiveRoomGetV10ApiService service

type ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanTodayLiveRoomGetV10ApiService
	advertiserId *int64
	awemeId      *int64
	dateTime     *string
	fields       *[]string
	roomStatus   *QianchuanTodayLiveRoomGetV10RoomStatus
	adStatus     *QianchuanTodayLiveRoomGetV10AdStatus
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.awemeId = &awemeId
	return r
}

// 开播日期，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) DateTime(dateTime string) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.dateTime = &dateTime
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) RoomStatus(roomStatus QianchuanTodayLiveRoomGetV10RoomStatus) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.roomStatus = &roomStatus
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) AdStatus(adStatus QianchuanTodayLiveRoomGetV10AdStatus) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.adStatus = &adStatus
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) Execute() (*QianchuanTodayLiveRoomGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanTodayLiveRoomGetGet Method for OpenApiV10QianchuanTodayLiveRoomGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest
*/
func (a *QianchuanTodayLiveRoomGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest {
	return &ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanTodayLiveRoomGetV10Response
func (a *QianchuanTodayLiveRoomGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanTodayLiveRoomGetGetRequest) (*QianchuanTodayLiveRoomGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanTodayLiveRoomGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/today_live/room/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if r.dateTime == nil {
		return localVarReturnValue, nil, ReportError("dateTime is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "date_time", r.dateTime)
	if r.roomStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "room_status", r.roomStatus)
	}
	if r.adStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_status", r.adStatus)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
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
