/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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

// ReportAudienceProvinceV2ApiService ReportAudienceProvinceV2Api service
type ReportAudienceProvinceV2ApiService service

type ApiOpenApi2ReportAudienceProvinceGetRequest struct {
	ctx          context.Context
	ApiService   *ReportAudienceProvinceV2ApiService
	advertiserId *int64
	endDate      **string
	idType       *ReportAudienceProvinceV2IdType
	ids          *[]int64
	metrics      *[]string
	startDate    **string
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) IdType(idType ReportAudienceProvinceV2IdType) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	r.idType = &idType
	return r
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) Ids(ids []int64) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	r.ids = &ids
	return r
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) Metrics(metrics []string) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) Execute() (*ReportAudienceProvinceV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportAudienceProvinceGetRequest) WithLog(enable bool) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportAudienceProvinceGet Method for OpenApi2ReportAudienceProvinceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportAudienceProvinceGetRequest
*/
func (a *ReportAudienceProvinceV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportAudienceProvinceGetRequest {
	return &ApiOpenApi2ReportAudienceProvinceGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportAudienceProvinceV2Response
func (a *ReportAudienceProvinceV2ApiService) getExecute(r *ApiOpenApi2ReportAudienceProvinceGetRequest) (*ReportAudienceProvinceV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportAudienceProvinceV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/audience/province/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	}
	if r.idType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id_type", r.idType)
	}
	if r.ids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ids", r.ids)
	}
	if r.metrics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
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
