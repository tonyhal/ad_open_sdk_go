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

// ToolsAdRaiseResultGetV2V2ApiService ToolsAdRaiseResultGetV2V2Api service
type ToolsAdRaiseResultGetV2V2ApiService service

type ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest struct {
	ctx            context.Context
	ApiService     *ToolsAdRaiseResultGetV2V2ApiService
	adId           *int64
	adRaiseVersion *int64
	advertiserId   *int64
	endTime        **string
	orderField     *string
	orderType      *ToolsAdRaiseResultGetV2V2OrderType
	page           *int64
	pageSize       *int64
	startTime      **string
	timeDimension  *ToolsAdRaiseResultGetV2V2TimeDimension
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) AdId(adId int64) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) AdRaiseVersion(adRaiseVersion int64) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.adRaiseVersion = &adRaiseVersion
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) EndTime(endTime *string) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) OrderField(orderField string) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) OrderType(orderType ToolsAdRaiseResultGetV2V2OrderType) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) Page(page int64) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) StartTime(startTime *string) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) TimeDimension(timeDimension ToolsAdRaiseResultGetV2V2TimeDimension) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.timeDimension = &timeDimension
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) Execute() (*ToolsAdRaiseResultGetV2V2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdRaiseResultGetV2Get Method for OpenApi2ToolsAdRaiseResultGetV2Get

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest
*/
func (a *ToolsAdRaiseResultGetV2V2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest {
	return &ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdRaiseResultGetV2V2Response
func (a *ToolsAdRaiseResultGetV2V2ApiService) getExecute(r *ApiOpenApi2ToolsAdRaiseResultGetV2GetRequest) (*ToolsAdRaiseResultGetV2V2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdRaiseResultGetV2V2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_raise_result/get_v2/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.adId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	}
	if r.adRaiseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_raise_version", r.adRaiseVersion)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
	if r.timeDimension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_dimension", r.timeDimension)
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
