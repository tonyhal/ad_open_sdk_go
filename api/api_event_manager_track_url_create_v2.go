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

// EventManagerTrackUrlCreateV2ApiService EventManagerTrackUrlCreateV2Api service
type EventManagerTrackUrlCreateV2ApiService service

type ApiOpenApi2EventManagerTrackUrlCreatePostRequest struct {
	ctx                                 context.Context
	ApiService                          *EventManagerTrackUrlCreateV2ApiService
	eventManagerTrackUrlCreateV2Request *EventManagerTrackUrlCreateV2Request
}

func (r *ApiOpenApi2EventManagerTrackUrlCreatePostRequest) EventManagerTrackUrlCreateV2Request(eventManagerTrackUrlCreateV2Request EventManagerTrackUrlCreateV2Request) *ApiOpenApi2EventManagerTrackUrlCreatePostRequest {
	r.eventManagerTrackUrlCreateV2Request = &eventManagerTrackUrlCreateV2Request
	return r
}

func (r *ApiOpenApi2EventManagerTrackUrlCreatePostRequest) Execute() (*EventManagerTrackUrlCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2EventManagerTrackUrlCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2EventManagerTrackUrlCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2EventManagerTrackUrlCreatePostRequest) WithLog(enable bool) *ApiOpenApi2EventManagerTrackUrlCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2EventManagerTrackUrlCreatePost Method for OpenApi2EventManagerTrackUrlCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2EventManagerTrackUrlCreatePostRequest
*/
func (a *EventManagerTrackUrlCreateV2ApiService) Post(ctx context.Context) *ApiOpenApi2EventManagerTrackUrlCreatePostRequest {
	return &ApiOpenApi2EventManagerTrackUrlCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerTrackUrlCreateV2Response
func (a *EventManagerTrackUrlCreateV2ApiService) postExecute(r *ApiOpenApi2EventManagerTrackUrlCreatePostRequest) (*EventManagerTrackUrlCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EventManagerTrackUrlCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/event_manager/track_url/create/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.eventManagerTrackUrlCreateV2Request
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
