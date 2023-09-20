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

// ToolsRubeexPlayableListV2ApiService ToolsRubeexPlayableListV2Api service
type ToolsRubeexPlayableListV2ApiService service

type ApiOpenApi2ToolsRubeexPlayableListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsRubeexPlayableListV2ApiService
	projectId    *int32
	advertiserId *int64
	page         *int32
	pageSize     *int32
	filtering    *ToolsRubeexPlayableListV2Filtering
}

// 互动作品id
func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) ProjectId(projectId int32) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	r.projectId = &projectId
	return r
}

// 广告主id
func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 页码。默认为1。
func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) Page(page int32) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	r.page = &page
	return r
}

// 页面大小。默认为1，最大为100
func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) Filtering(filtering ToolsRubeexPlayableListV2Filtering) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) Execute() (*ToolsRubeexPlayableListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRubeexPlayableListGet Method for OpenApi2ToolsRubeexPlayableListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRubeexPlayableListGetRequest
*/
func (a *ToolsRubeexPlayableListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRubeexPlayableListGetRequest {
	return &ApiOpenApi2ToolsRubeexPlayableListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRubeexPlayableListV2Response
func (a *ToolsRubeexPlayableListV2ApiService) getExecute(r *ApiOpenApi2ToolsRubeexPlayableListGetRequest) (*ToolsRubeexPlayableListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsRubeexPlayableListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/rubeex_playable/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.projectId == nil {
		return localVarReturnValue, nil, ReportError("projectId is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
