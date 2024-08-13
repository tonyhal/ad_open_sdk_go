/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
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

// BrandAdGetV30ApiService BrandAdGetV30Api service
type BrandAdGetV30ApiService service

type ApiOpenApiV30BrandAdGetGetRequest struct {
	ctx             context.Context
	ApiService      *BrandAdGetV30ApiService
	advertiserId    *int64
	page            *int64
	size            *int64
	adIds           *[]string
	adNames         *[]string
	campaignIds     *[]string
	appOrigin       *BrandAdGetV30AppOrigin
	adForm          *BrandAdGetV30AdForm
	adStatus        *BrandAdGetV30AdStatus
	createStartTime *string
	createEndTime   *string
	startTime       *string
	endTime         *string
}

// 广告主ID
func (r *ApiOpenApiV30BrandAdGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandAdGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 分页参数 默认从1开始
func (r *ApiOpenApiV30BrandAdGetGetRequest) Page(page int64) *ApiOpenApiV30BrandAdGetGetRequest {
	r.page = &page
	return r
}

// 分页参数 最大为100
func (r *ApiOpenApiV30BrandAdGetGetRequest) Size(size int64) *ApiOpenApiV30BrandAdGetGetRequest {
	r.size = &size
	return r
}

// 计划ID List
func (r *ApiOpenApiV30BrandAdGetGetRequest) AdIds(adIds []string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.adIds = &adIds
	return r
}

// 计划名称List
func (r *ApiOpenApiV30BrandAdGetGetRequest) AdNames(adNames []string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.adNames = &adNames
	return r
}

// 广告组ID List
func (r *ApiOpenApiV30BrandAdGetGetRequest) CampaignIds(campaignIds []string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.campaignIds = &campaignIds
	return r
}

// 媒体端
func (r *ApiOpenApiV30BrandAdGetGetRequest) AppOrigin(appOrigin BrandAdGetV30AppOrigin) *ApiOpenApiV30BrandAdGetGetRequest {
	r.appOrigin = &appOrigin
	return r
}

// 版位
func (r *ApiOpenApiV30BrandAdGetGetRequest) AdForm(adForm BrandAdGetV30AdForm) *ApiOpenApiV30BrandAdGetGetRequest {
	r.adForm = &adForm
	return r
}

// 广告计划状态
func (r *ApiOpenApiV30BrandAdGetGetRequest) AdStatus(adStatus BrandAdGetV30AdStatus) *ApiOpenApiV30BrandAdGetGetRequest {
	r.adStatus = &adStatus
	return r
}

// 创建起始时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandAdGetGetRequest) CreateStartTime(createStartTime string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.createStartTime = &createStartTime
	return r
}

// 创建截止时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandAdGetGetRequest) CreateEndTime(createEndTime string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.createEndTime = &createEndTime
	return r
}

// 投放起始时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandAdGetGetRequest) StartTime(startTime string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.startTime = &startTime
	return r
}

// 投放截止时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandAdGetGetRequest) EndTime(endTime string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV30BrandAdGetGetRequest) Execute() (*BrandAdGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BrandAdGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandAdGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandAdGetGetRequest) WithLog(enable bool) *ApiOpenApiV30BrandAdGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandAdGetGet Method for OpenApiV30BrandAdGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandAdGetGetRequest
*/
func (a *BrandAdGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BrandAdGetGetRequest {
	return &ApiOpenApiV30BrandAdGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandAdGetV30Response
func (a *BrandAdGetV30ApiService) getExecute(r *ApiOpenApiV30BrandAdGetGetRequest) (*BrandAdGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandAdGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/ad/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if r.size == nil {
		return localVarReturnValue, nil, ReportError("size is required and must be specified")
	}
	if *r.size > 100 {
		return localVarReturnValue, nil, ReportError("size must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.adIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	}
	if r.adNames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_names", r.adNames)
	}
	if r.campaignIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_ids", r.campaignIds)
	}
	if r.appOrigin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_origin", r.appOrigin)
	}
	if r.adForm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_form", r.adForm)
	}
	if r.adStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_status", r.adStatus)
	}
	if r.createStartTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_start_time", r.createStartTime)
	}
	if r.createEndTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_end_time", r.createEndTime)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size)
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
