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

// ToolsAppManagementIndustryInfoListV2ApiService ToolsAppManagementIndustryInfoListV2Api service
type ToolsAppManagementIndustryInfoListV2ApiService service

type ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest struct {
	ctx         context.Context
	ApiService  *ToolsAppManagementIndustryInfoListV2ApiService
	accountId   *int64
	accountType *ToolsAppManagementIndustryInfoListV2AccountType
}

// 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
func (r *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest) AccountId(accountId int64) *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest {
	r.accountId = &accountId
	return r
}

// 账户类型
func (r *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest) AccountType(accountType ToolsAppManagementIndustryInfoListV2AccountType) *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest) Execute() (*ToolsAppManagementIndustryInfoListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementIndustryInfoListGet Method for OpenApi2ToolsAppManagementIndustryInfoListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest
*/
func (a *ToolsAppManagementIndustryInfoListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest {
	return &ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementIndustryInfoListV2Response
func (a *ToolsAppManagementIndustryInfoListV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementIndustryInfoListGetRequest) (*ToolsAppManagementIndustryInfoListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementIndustryInfoListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/industry_info/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if *r.accountId < 1 {
		return localVarReturnValue, nil, ReportError("accountId must be greater than 1")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
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
