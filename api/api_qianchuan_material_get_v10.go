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

// QianchuanMaterialGetV10ApiService QianchuanMaterialGetV10Api service
type QianchuanMaterialGetV10ApiService service

type ApiOpenApiV10QianchuanMaterialGetGetRequest struct {
	ctx            context.Context
	ApiService     *QianchuanMaterialGetV10ApiService
	advertiserId   *int64
	marketingGoal  *QianchuanMaterialGetV10MarketingGoal
	marketingScene *QianchuanMaterialGetV10MarketingScene
	filtering      *QianchuanMaterialGetV10Filtering
	fields         *[]string
	campaignScene  *[]*QianchuanMaterialGetV10CampaignScene
	page           *int32
	pageSize       *QianchuanMaterialGetV10PageSize
	orderType      *QianchuanMaterialGetV10OrderType
	orderField     *string
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) MarketingGoal(marketingGoal QianchuanMaterialGetV10MarketingGoal) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) MarketingScene(marketingScene QianchuanMaterialGetV10MarketingScene) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.marketingScene = &marketingScene
	return r
}

// 过滤器
func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) Filtering(filtering QianchuanMaterialGetV10Filtering) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.filtering = &filtering
	return r
}

// 查询指标
func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) CampaignScene(campaignScene []*QianchuanMaterialGetV10CampaignScene) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.campaignScene = &campaignScene
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) PageSize(pageSize QianchuanMaterialGetV10PageSize) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 排序方式
func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) OrderType(orderType QianchuanMaterialGetV10OrderType) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.orderType = &orderType
	return r
}

// 排序字段
func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) Execute() (*QianchuanMaterialGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanMaterialGetGet Method for OpenApiV10QianchuanMaterialGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanMaterialGetGetRequest
*/
func (a *QianchuanMaterialGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanMaterialGetGetRequest {
	return &ApiOpenApiV10QianchuanMaterialGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanMaterialGetV10Response
func (a *QianchuanMaterialGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanMaterialGetGetRequest) (*QianchuanMaterialGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanMaterialGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/material/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.marketingScene == nil {
		return localVarReturnValue, nil, ReportError("marketingScene is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	if r.campaignScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_scene", r.campaignScene)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_scene", r.marketingScene)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
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
