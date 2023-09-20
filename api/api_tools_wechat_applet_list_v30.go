/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

// ToolsWechatAppletListV30ApiService ToolsWechatAppletListV30Api service
type ToolsWechatAppletListV30ApiService service

type ApiOpenApiV30ToolsWechatAppletListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsWechatAppletListV30ApiService
	advertiserId *int64
	filtering    *ToolsWechatAppletListV30Filtering
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) Filtering(filtering ToolsWechatAppletListV30Filtering) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) Page(page int32) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) Execute() (*ToolsWechatAppletListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsWechatAppletListGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsWechatAppletListGet Method for OpenApiV30ToolsWechatAppletListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsWechatAppletListGetRequest
*/
func (a *ToolsWechatAppletListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsWechatAppletListGetRequest {
	return &ApiOpenApiV30ToolsWechatAppletListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsWechatAppletListV30Response
func (a *ToolsWechatAppletListV30ApiService) getExecute(r *ApiOpenApiV30ToolsWechatAppletListGetRequest) (*ToolsWechatAppletListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsWechatAppletListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/wechat_applet/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
