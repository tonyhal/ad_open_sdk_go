/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

// ToolsThirdSitePreviewV2ApiService ToolsThirdSitePreviewV2Api service
type ToolsThirdSitePreviewV2ApiService service

type ApiOpenApi2ToolsThirdSitePreviewGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsThirdSitePreviewV2ApiService
	advertiserId *int64
	siteId       *int64
}

func (r *ApiOpenApi2ToolsThirdSitePreviewGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsThirdSitePreviewGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsThirdSitePreviewGetRequest) SiteId(siteId int64) *ApiOpenApi2ToolsThirdSitePreviewGetRequest {
	r.siteId = &siteId
	return r
}

func (r *ApiOpenApi2ToolsThirdSitePreviewGetRequest) Execute() (*ToolsThirdSitePreviewV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsThirdSitePreviewGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsThirdSitePreviewGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsThirdSitePreviewGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsThirdSitePreviewGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsThirdSitePreviewGet Method for OpenApi2ToolsThirdSitePreviewGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsThirdSitePreviewGetRequest
*/
func (a *ToolsThirdSitePreviewV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsThirdSitePreviewGetRequest {
	return &ApiOpenApi2ToolsThirdSitePreviewGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsThirdSitePreviewV2Response
func (a *ToolsThirdSitePreviewV2ApiService) getExecute(r *ApiOpenApi2ToolsThirdSitePreviewGetRequest) (*ToolsThirdSitePreviewV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsThirdSitePreviewV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/third_site/preview/"

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
	if *r.advertiserId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9223372036854775807")
	}
	if r.siteId == nil {
		return localVarReturnValue, nil, ReportError("siteId is required and must be specified")
	}
	if *r.siteId < 1 {
		return localVarReturnValue, nil, ReportError("siteId must be greater than 1")
	}
	if *r.siteId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("siteId must be less than 9223372036854775807")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId)
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
