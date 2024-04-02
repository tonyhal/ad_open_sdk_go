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

// ReportAdGetV2ApiService ReportAdGetV2Api service
type ReportAdGetV2ApiService service

type ApiOpenApi2ReportAdGetGetRequest struct {
	ctx             context.Context
	ApiService      *ReportAdGetV2ApiService
	advertiserId    *int64
	endDate         **string
	fields          *[]string
	filtering       *ReportAdGetV2Filtering
	groupBy         *[]*ReportAdGetV2GroupBy
	orderField      *ReportAdGetV2OrderField
	orderType       *ReportAdGetV2OrderType
	page            *int64
	pageSize        *int64
	startDate       **string
	timeGranularity *ReportAdGetV2TimeGranularity
}

func (r *ApiOpenApi2ReportAdGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportAdGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportAdGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) Fields(fields []string) *ApiOpenApi2ReportAdGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) Filtering(filtering ReportAdGetV2Filtering) *ApiOpenApi2ReportAdGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) GroupBy(groupBy []*ReportAdGetV2GroupBy) *ApiOpenApi2ReportAdGetGetRequest {
	r.groupBy = &groupBy
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) OrderField(orderField ReportAdGetV2OrderField) *ApiOpenApi2ReportAdGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) OrderType(orderType ReportAdGetV2OrderType) *ApiOpenApi2ReportAdGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) Page(page int64) *ApiOpenApi2ReportAdGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportAdGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportAdGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) TimeGranularity(timeGranularity ReportAdGetV2TimeGranularity) *ApiOpenApi2ReportAdGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) Execute() (*ReportAdGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportAdGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportAdGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportAdGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportAdGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportAdGetGet Method for OpenApi2ReportAdGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportAdGetGetRequest
*/
func (a *ReportAdGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportAdGetGetRequest {
	return &ApiOpenApi2ReportAdGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportAdGetV2Response
func (a *ReportAdGetV2ApiService) getExecute(r *ApiOpenApi2ReportAdGetGetRequest) (*ReportAdGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportAdGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/ad/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_by", r.groupBy)
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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	}
	if r.timeGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_granularity", r.timeGranularity)
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
