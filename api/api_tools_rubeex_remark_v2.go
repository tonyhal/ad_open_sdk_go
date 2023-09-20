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

// ToolsRubeexRemarkV2ApiService ToolsRubeexRemarkV2Api service
type ToolsRubeexRemarkV2ApiService service

type ApiOpenApi2ToolsRubeexRemarkGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsRubeexRemarkV2ApiService
	advertiserId *float64
	projectId    *float64
	scene        *ToolsRubeexRemarkV2Scene
	filtering    *ToolsRubeexRemarkV2Filtering
}

// 广告主ID
func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) AdvertiserId(advertiserId float64) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 作品id即试玩素材id，属于试玩素材唯一标识
func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) ProjectId(projectId float64) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	r.projectId = &projectId
	return r
}

// &#x60;1&#x60;：分场景&lt;br&gt;&#x60;2&#x60;：分场景并且分区域
func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) Scene(scene ToolsRubeexRemarkV2Scene) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	r.scene = &scene
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) Filtering(filtering ToolsRubeexRemarkV2Filtering) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) Execute() (*ToolsRubeexRemarkV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRubeexRemarkGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRubeexRemarkGet Method for OpenApi2ToolsRubeexRemarkGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRubeexRemarkGetRequest
*/
func (a *ToolsRubeexRemarkV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRubeexRemarkGetRequest {
	return &ApiOpenApi2ToolsRubeexRemarkGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRubeexRemarkV2Response
func (a *ToolsRubeexRemarkV2ApiService) getExecute(r *ApiOpenApi2ToolsRubeexRemarkGetRequest) (*ToolsRubeexRemarkV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsRubeexRemarkV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/rubeex/remark/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1.0 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1.0")
	}
	if *r.advertiserId > 9.223372036854776e+18 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9.223372036854776E+18")
	}
	if r.projectId == nil {
		return localVarReturnValue, nil, ReportError("projectId is required and must be specified")
	}
	if *r.projectId < 1.0 {
		return localVarReturnValue, nil, ReportError("projectId must be greater than 1.0")
	}
	if *r.projectId > 2147483647 {
		return localVarReturnValue, nil, ReportError("projectId must be less than 2147483647")
	}
	if r.scene == nil {
		return localVarReturnValue, nil, ReportError("scene is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "scene", r.scene)
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
