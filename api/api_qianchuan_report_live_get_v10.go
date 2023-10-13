/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
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

// QianchuanReportLiveGetV10ApiService QianchuanReportLiveGetV10Api service
type QianchuanReportLiveGetV10ApiService service

type ApiOpenApiV10QianchuanReportLiveGetGetRequest struct {
	ctx            context.Context
	ApiService     *QianchuanReportLiveGetV10ApiService
	advertiserId   *int64
	awemeId        *int64
	startTime      *string
	endTime        *string
	fields         *[]string
	statsAuthority *QianchuanReportLiveGetV10StatsAuthority
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) StatsAuthority(statsAuthority QianchuanReportLiveGetV10StatsAuthority) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.statsAuthority = &statsAuthority
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) Execute() (*QianchuanReportLiveGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportLiveGetGet Method for OpenApiV10QianchuanReportLiveGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportLiveGetGetRequest
*/
func (a *QianchuanReportLiveGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportLiveGetGetRequest {
	return &ApiOpenApiV10QianchuanReportLiveGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportLiveGetV10Response
func (a *QianchuanReportLiveGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportLiveGetGetRequest) (*QianchuanReportLiveGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanReportLiveGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/live/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	if r.statsAuthority != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stats_authority", r.statsAuthority)
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
