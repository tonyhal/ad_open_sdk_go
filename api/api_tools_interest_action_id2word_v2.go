/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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

// ToolsInterestActionId2wordV2ApiService ToolsInterestActionId2wordV2Api service
type ToolsInterestActionId2wordV2ApiService service

type ApiOpenApi2ToolsInterestActionId2wordGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsInterestActionId2wordV2ApiService
	actionDays    *ToolsInterestActionId2wordV2ActionDays
	advertiserId  *int64
	ids           *[]int64
	tagType       *ToolsInterestActionId2wordV2TagType
	targetingType *ToolsInterestActionId2wordV2TargetingType
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) ActionDays(actionDays ToolsInterestActionId2wordV2ActionDays) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	r.actionDays = &actionDays
	return r
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) Ids(ids []int64) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	r.ids = &ids
	return r
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) TagType(tagType ToolsInterestActionId2wordV2TagType) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	r.tagType = &tagType
	return r
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) TargetingType(targetingType ToolsInterestActionId2wordV2TargetingType) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	r.targetingType = &targetingType
	return r
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) Execute() (*ToolsInterestActionId2wordV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsInterestActionId2wordGet Method for OpenApi2ToolsInterestActionId2wordGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsInterestActionId2wordGetRequest
*/
func (a *ToolsInterestActionId2wordV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsInterestActionId2wordGetRequest {
	return &ApiOpenApi2ToolsInterestActionId2wordGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsInterestActionId2wordV2Response
func (a *ToolsInterestActionId2wordV2ApiService) getExecute(r *ApiOpenApi2ToolsInterestActionId2wordGetRequest) (*ToolsInterestActionId2wordV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsInterestActionId2wordV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/interest_action/id2word/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.actionDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_days", r.actionDays)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.ids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ids", r.ids)
	}
	if r.tagType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag_type", r.tagType)
	}
	if r.targetingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "targeting_type", r.targetingType)
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
