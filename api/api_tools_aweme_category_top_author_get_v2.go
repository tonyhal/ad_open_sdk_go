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

// ToolsAwemeCategoryTopAuthorGetV2ApiService ToolsAwemeCategoryTopAuthorGetV2Api service
type ToolsAwemeCategoryTopAuthorGetV2ApiService service

type ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAwemeCategoryTopAuthorGetV2ApiService
	advertiserId *int64
	behaviors    *[]*ToolsAwemeCategoryTopAuthorGetV2Behaviors
	categoryId   *int64
}

func (r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) Behaviors(behaviors []*ToolsAwemeCategoryTopAuthorGetV2Behaviors) *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest {
	r.behaviors = &behaviors
	return r
}

func (r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) CategoryId(categoryId int64) *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest {
	r.categoryId = &categoryId
	return r
}

func (r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) Execute() (*ToolsAwemeCategoryTopAuthorGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAwemeCategoryTopAuthorGetGet Method for OpenApi2ToolsAwemeCategoryTopAuthorGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest
*/
func (a *ToolsAwemeCategoryTopAuthorGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest {
	return &ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAwemeCategoryTopAuthorGetV2Response
func (a *ToolsAwemeCategoryTopAuthorGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAwemeCategoryTopAuthorGetGetRequest) (*ToolsAwemeCategoryTopAuthorGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAwemeCategoryTopAuthorGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aweme_category_top_author/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.behaviors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "behaviors", r.behaviors)
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category_id", r.categoryId)
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
