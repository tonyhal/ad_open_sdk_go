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

// LocalMultiPoiIdPoiIdsGetV30ApiService LocalMultiPoiIdPoiIdsGetV30Api service
type LocalMultiPoiIdPoiIdsGetV30ApiService service

type ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest struct {
	ctx            context.Context
	ApiService     *LocalMultiPoiIdPoiIdsGetV30ApiService
	localAccountId *int64
	multiPoiIds    *[]int64
	needEnable     *bool
}

func (r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) LocalAccountId(localAccountId int64) *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest {
	r.localAccountId = &localAccountId
	return r
}

// 门店个数
func (r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) MultiPoiIds(multiPoiIds []int64) *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest {
	r.multiPoiIds = &multiPoiIds
	return r
}

func (r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) NeedEnable(needEnable bool) *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest {
	r.needEnable = &needEnable
	return r
}

func (r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) Execute() (*LocalMultiPoiIdPoiIdsGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) WithLog(enable bool) *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalMultiPoiIdPoiIdsGetGet Method for OpenApiV30LocalMultiPoiIdPoiIdsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest
*/
func (a *LocalMultiPoiIdPoiIdsGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest {
	return &ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalMultiPoiIdPoiIdsGetV30Response
func (a *LocalMultiPoiIdPoiIdsGetV30ApiService) getExecute(r *ApiOpenApiV30LocalMultiPoiIdPoiIdsGetGetRequest) (*LocalMultiPoiIdPoiIdsGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalMultiPoiIdPoiIdsGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/multi_poi_id/poi_ids/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.localAccountId == nil {
		return localVarReturnValue, nil, ReportError("localAccountId is required and must be specified")
	}
	if r.multiPoiIds == nil {
		return localVarReturnValue, nil, ReportError("multiPoiIds is required and must be specified")
	}
	if len(*r.multiPoiIds) < 1 {
		return localVarReturnValue, nil, ReportError("multiPoiIds must have at least 1 elements")
	}
	if len(*r.multiPoiIds) > 50 {
		return localVarReturnValue, nil, ReportError("multiPoiIds must have less than 50 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "local_account_id", r.localAccountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "multi_poi_ids", r.multiPoiIds)
	if r.needEnable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "need_enable", r.needEnable)
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
