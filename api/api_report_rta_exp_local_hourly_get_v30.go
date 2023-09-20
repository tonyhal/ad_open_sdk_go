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

// ReportRtaExpLocalHourlyGetV30ApiService ReportRtaExpLocalHourlyGetV30Api service
type ReportRtaExpLocalHourlyGetV30ApiService service

type ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportRtaExpLocalHourlyGetV30ApiService
	rtaId        *int64
	advertiserId *int64
	startDate    *string
	endDate      *string
	vid          *int64
	cusVid       *int64
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) RtaId(rtaId int64) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.rtaId = &rtaId
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) StartDate(startDate string) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) EndDate(endDate string) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) Vid(vid int64) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.vid = &vid
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) CusVid(cusVid int64) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.cusVid = &cusVid
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) Execute() (*ReportRtaExpLocalHourlyGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportRtaExpLocalHourlyGetGet Method for OpenApiV30ReportRtaExpLocalHourlyGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest
*/
func (a *ReportRtaExpLocalHourlyGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest {
	return &ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportRtaExpLocalHourlyGetV30Response
func (a *ReportRtaExpLocalHourlyGetV30ApiService) getExecute(r *ApiOpenApiV30ReportRtaExpLocalHourlyGetGetRequest) (*ReportRtaExpLocalHourlyGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportRtaExpLocalHourlyGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/rta_exp_local_hourly/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rtaId == nil {
		return localVarReturnValue, nil, ReportError("rtaId is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "rta_id", r.rtaId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.vid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vid", r.vid)
	}
	if r.cusVid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cus_vid", r.cusVid)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
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