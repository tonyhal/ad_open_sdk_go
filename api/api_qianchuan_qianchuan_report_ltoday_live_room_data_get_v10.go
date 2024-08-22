/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// QianchuanQianchuanReportLtodayLiveRoomDataGetV10ApiService QianchuanQianchuanReportLtodayLiveRoomDataGetV10Api service
type QianchuanQianchuanReportLtodayLiveRoomDataGetV10ApiService service

type ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanQianchuanReportLtodayLiveRoomDataGetV10ApiService
	advertiserId *int64
	dataTopic    *QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic
	dimensions   *[]string
	metrics      *[]string
	filters      *[]*QianchuanQianchuanReportLtodayLiveRoomDataGetV10FiltersInner
	startTime    *string
	endTime      *string
	orderBy      *[]*QianchuanQianchuanReportLtodayLiveRoomDataGetV10OrderByInner
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) DataTopic(dataTopic QianchuanQianchuanReportLtodayLiveRoomDataGetV10DataTopic) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.dataTopic = &dataTopic
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) Dimensions(dimensions []string) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.dimensions = &dimensions
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) Metrics(metrics []string) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) Filters(filters []*QianchuanQianchuanReportLtodayLiveRoomDataGetV10FiltersInner) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.filters = &filters
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) OrderBy(orderBy []*QianchuanQianchuanReportLtodayLiveRoomDataGetV10OrderByInner) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.orderBy = &orderBy
	return r
}

// 第几页
func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) Page(page int64) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.page = &page
	return r
}

// 页大小
func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) Execute() (*QianchuanQianchuanReportLtodayLiveRoomDataGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGet Method for OpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest
*/
func (a *QianchuanQianchuanReportLtodayLiveRoomDataGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest {
	return &ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanQianchuanReportLtodayLiveRoomDataGetV10Response
func (a *QianchuanQianchuanReportLtodayLiveRoomDataGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanQianchuanReportLtodayLiveRoomDataGetGetRequest) (*QianchuanQianchuanReportLtodayLiveRoomDataGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanQianchuanReportLtodayLiveRoomDataGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/qianchuan/report/ltoday_live/room/data/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.dataTopic == nil {
		return localVarReturnValue, nil, ReportError("dataTopic is required and must be specified")
	}
	if r.dimensions == nil {
		return localVarReturnValue, nil, ReportError("dimensions is required and must be specified")
	}
	if r.metrics == nil {
		return localVarReturnValue, nil, ReportError("metrics is required and must be specified")
	}
	if r.filters == nil {
		return localVarReturnValue, nil, ReportError("filters is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.orderBy == nil {
		return localVarReturnValue, nil, ReportError("orderBy is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "data_topic", r.dataTopic)
	parameterAddToHeaderOrQuery(localVarQueryParams, "dimensions", r.dimensions)
	parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filters", r.filters)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy)
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

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
