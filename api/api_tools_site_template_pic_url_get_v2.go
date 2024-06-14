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

// ToolsSiteTemplatePicUrlGetV2ApiService ToolsSiteTemplatePicUrlGetV2Api service
type ToolsSiteTemplatePicUrlGetV2ApiService service

type ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsSiteTemplatePicUrlGetV2ApiService
	advertiserId *int64
	siteId       *int64
	templateId   *int64
}

// 广告主
func (r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 站点id
func (r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) SiteId(siteId int64) *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest {
	r.siteId = &siteId
	return r
}

// 模板id
func (r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) TemplateId(templateId int64) *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest {
	r.templateId = &templateId
	return r
}

func (r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) Execute() (*ToolsSiteTemplatePicUrlGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteTemplatePicUrlGetGet Method for OpenApi2ToolsSiteTemplatePicUrlGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest
*/
func (a *ToolsSiteTemplatePicUrlGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest {
	return &ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteTemplatePicUrlGetV2Response
func (a *ToolsSiteTemplatePicUrlGetV2ApiService) getExecute(r *ApiOpenApi2ToolsSiteTemplatePicUrlGetGetRequest) (*ToolsSiteTemplatePicUrlGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteTemplatePicUrlGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site_template/pic_url/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId)
	}
	if r.templateId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "template_id", r.templateId)
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
