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

// ToolsActionTextGetV2ApiService ToolsActionTextGetV2Api service
type ToolsActionTextGetV2ApiService service

type ApiOpenApi2ToolsActionTextGetGetRequest struct {
	ctx                  context.Context
	ApiService           *ToolsActionTextGetV2ApiService
	advancedCreativeType *ToolsActionTextGetV2AdvancedCreativeType
	advertiserId         *int64
	industry             *int64
	landingType          *ToolsActionTextGetV2LandingType
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) AdvancedCreativeType(advancedCreativeType ToolsActionTextGetV2AdvancedCreativeType) *ApiOpenApi2ToolsActionTextGetGetRequest {
	r.advancedCreativeType = &advancedCreativeType
	return r
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsActionTextGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) Industry(industry int64) *ApiOpenApi2ToolsActionTextGetGetRequest {
	r.industry = &industry
	return r
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) LandingType(landingType ToolsActionTextGetV2LandingType) *ApiOpenApi2ToolsActionTextGetGetRequest {
	r.landingType = &landingType
	return r
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) Execute() (*ToolsActionTextGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsActionTextGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsActionTextGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsActionTextGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsActionTextGetGet Method for OpenApi2ToolsActionTextGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsActionTextGetGetRequest
*/
func (a *ToolsActionTextGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsActionTextGetGetRequest {
	return &ApiOpenApi2ToolsActionTextGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsActionTextGetV2Response
func (a *ToolsActionTextGetV2ApiService) getExecute(r *ApiOpenApi2ToolsActionTextGetGetRequest) (*ToolsActionTextGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsActionTextGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/action_text/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advancedCreativeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advanced_creative_type", r.advancedCreativeType)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.industry != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "industry", r.industry)
	}
	if r.landingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
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
