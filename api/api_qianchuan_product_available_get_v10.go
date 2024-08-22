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

// QianchuanProductAvailableGetV10ApiService QianchuanProductAvailableGetV10Api service
type QianchuanProductAvailableGetV10ApiService service

type ApiOpenApiV10QianchuanProductAvailableGetGetRequest struct {
	ctx           context.Context
	ApiService    *QianchuanProductAvailableGetV10ApiService
	advertiserId  *int64
	awemeId       *int64
	filter        *QianchuanProductAvailableGetV10Filter
	campaignScene *QianchuanProductAvailableGetV10CampaignScene
	page          *int32
	pageSize      *int32
	cursor        *int64
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) Filter(filter QianchuanProductAvailableGetV10Filter) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) CampaignScene(campaignScene QianchuanProductAvailableGetV10CampaignScene) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.campaignScene = &campaignScene
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 页码游标值
func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) Cursor(cursor int64) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.cursor = &cursor
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) Execute() (*QianchuanProductAvailableGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanProductAvailableGetGet Method for OpenApiV10QianchuanProductAvailableGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanProductAvailableGetGetRequest
*/
func (a *QianchuanProductAvailableGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanProductAvailableGetGetRequest {
	return &ApiOpenApiV10QianchuanProductAvailableGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanProductAvailableGetV10Response
func (a *QianchuanProductAvailableGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanProductAvailableGetGetRequest) (*QianchuanProductAvailableGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanProductAvailableGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/product/available/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 0 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 0")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
	if r.campaignScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_scene", r.campaignScene)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
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
