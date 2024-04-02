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

// CreativeTemplateDetailGetV2ApiService CreativeTemplateDetailGetV2Api service
type CreativeTemplateDetailGetV2ApiService service

type ApiOpenApi2CreativeTemplateDetailGetGetRequest struct {
	ctx          context.Context
	ApiService   *CreativeTemplateDetailGetV2ApiService
	advertiserId *int64
	templateId   *int64
	templateType *CreativeTemplateDetailGetV2TemplateType
}

// 广告主ID
func (r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeTemplateDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 模板ID
func (r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) TemplateId(templateId int64) *ApiOpenApi2CreativeTemplateDetailGetGetRequest {
	r.templateId = &templateId
	return r
}

// 模板样式,eg:
func (r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) TemplateType(templateType CreativeTemplateDetailGetV2TemplateType) *ApiOpenApi2CreativeTemplateDetailGetGetRequest {
	r.templateType = &templateType
	return r
}

func (r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) Execute() (*CreativeTemplateDetailGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeTemplateDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeTemplateDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeTemplateDetailGetGet Method for OpenApi2CreativeTemplateDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeTemplateDetailGetGetRequest
*/
func (a *CreativeTemplateDetailGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeTemplateDetailGetGetRequest {
	return &ApiOpenApi2CreativeTemplateDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeTemplateDetailGetV2Response
func (a *CreativeTemplateDetailGetV2ApiService) getExecute(r *ApiOpenApi2CreativeTemplateDetailGetGetRequest) (*CreativeTemplateDetailGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeTemplateDetailGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/template/detail/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.templateId == nil {
		return localVarReturnValue, nil, ReportError("templateId is required and must be specified")
	}
	if r.templateType == nil {
		return localVarReturnValue, nil, ReportError("templateType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "template_id", r.templateId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "template_type", r.templateType)
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
