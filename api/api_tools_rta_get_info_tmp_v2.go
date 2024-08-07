/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// ToolsRtaGetInfoTmpV2ApiService ToolsRtaGetInfoTmpV2Api service
type ToolsRtaGetInfoTmpV2ApiService service

type ApiOpenApi2ToolsRtaGetInfoTmpGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsRtaGetInfoTmpV2ApiService
	advertiserId *int64
	campaignId   *int64
}

// 广告账户id
func (r *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告组id，若传入，则拉取的是组维度的RTA策略
func (r *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest) CampaignId(campaignId int64) *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest {
	r.campaignId = &campaignId
	return r
}

func (r *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest) Execute() (*ToolsRtaGetInfoTmpV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRtaGetInfoTmpGet Method for OpenApi2ToolsRtaGetInfoTmpGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRtaGetInfoTmpGetRequest
*/
func (a *ToolsRtaGetInfoTmpV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest {
	return &ApiOpenApi2ToolsRtaGetInfoTmpGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRtaGetInfoTmpV2Response
func (a *ToolsRtaGetInfoTmpV2ApiService) getExecute(r *ApiOpenApi2ToolsRtaGetInfoTmpGetRequest) (*ToolsRtaGetInfoTmpV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsRtaGetInfoTmpV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/rta/get_info_tmp/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.campaignId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_id", r.campaignId)
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
