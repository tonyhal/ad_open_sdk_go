/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
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

// ToolsSitePreviewV2ApiService ToolsSitePreviewV2Api service
type ToolsSitePreviewV2ApiService service

type ApiOpenApi2ToolsSitePreviewGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsSitePreviewV2ApiService
	advertiserId  *int64
	siteId        *string
	xOrangeCaller *string
}

// 广告主id，广告主id， 传的这个advertiser_id的数字的范围：1 &lt;&#x3D; advertiser_id &lt;&#x3D; MAX_INT64
func (r *ApiOpenApi2ToolsSitePreviewGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsSitePreviewGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 橙子建站站点id
func (r *ApiOpenApi2ToolsSitePreviewGetRequest) SiteId(siteId string) *ApiOpenApi2ToolsSitePreviewGetRequest {
	r.siteId = &siteId
	return r
}

func (r *ApiOpenApi2ToolsSitePreviewGetRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsSitePreviewGetRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

func (r *ApiOpenApi2ToolsSitePreviewGetRequest) Execute() (*ToolsSitePreviewV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSitePreviewGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSitePreviewGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSitePreviewGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSitePreviewGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSitePreviewGet Method for OpenApi2ToolsSitePreviewGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSitePreviewGetRequest
*/
func (a *ToolsSitePreviewV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSitePreviewGetRequest {
	return &ApiOpenApi2ToolsSitePreviewGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSitePreviewV2Response
func (a *ToolsSitePreviewV2ApiService) getExecute(r *ApiOpenApi2ToolsSitePreviewGetRequest) (*ToolsSitePreviewV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsSitePreviewV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site/preview/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.siteId == nil {
		return localVarReturnValue, nil, ReportError("siteId is required and must be specified")
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
	if r.xOrangeCaller != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Orange-Caller", r.xOrangeCaller)
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
