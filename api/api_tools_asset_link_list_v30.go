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

// ToolsAssetLinkListV30ApiService ToolsAssetLinkListV30Api service
type ToolsAssetLinkListV30ApiService service

type ApiOpenApiV30ToolsAssetLinkListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAssetLinkListV30ApiService
	advertiserId *int64
	filtering    *ToolsAssetLinkListV30Filtering
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) Filtering(filtering ToolsAssetLinkListV30Filtering) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) Page(page int32) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) Execute() (*ToolsAssetLinkListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsAssetLinkListGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsAssetLinkListGet Method for OpenApiV30ToolsAssetLinkListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsAssetLinkListGetRequest
*/
func (a *ToolsAssetLinkListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsAssetLinkListGetRequest {
	return &ApiOpenApiV30ToolsAssetLinkListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAssetLinkListV30Response
func (a *ToolsAssetLinkListV30ApiService) getExecute(r *ApiOpenApiV30ToolsAssetLinkListGetRequest) (*ToolsAssetLinkListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAssetLinkListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/asset_link/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
