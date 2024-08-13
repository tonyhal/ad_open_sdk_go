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

// StarChallengeGetCustomTaskDataV2ApiService StarChallengeGetCustomTaskDataV2Api service
type StarChallengeGetCustomTaskDataV2ApiService service

type ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest struct {
	ctx             context.Context
	ApiService      *StarChallengeGetCustomTaskDataV2ApiService
	starId          *int64
	challengeTaskId *int64
	page            *int32
	pageSize        *int32
}

// 客户星图ID
func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) StarId(starId int64) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	r.starId = &starId
	return r
}

// 投稿任务ID
func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

// page
func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) Page(page int32) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	r.page = &page
	return r
}

// 小于100
func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) PageSize(pageSize int32) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) Execute() (*StarChallengeGetCustomTaskDataV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeGetCustomTaskDataGet Method for OpenApi2StarChallengeGetCustomTaskDataGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest
*/
func (a *StarChallengeGetCustomTaskDataV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest {
	return &ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeGetCustomTaskDataV2Response
func (a *StarChallengeGetCustomTaskDataV2ApiService) getExecute(r *ApiOpenApi2StarChallengeGetCustomTaskDataGetRequest) (*StarChallengeGetCustomTaskDataV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeGetCustomTaskDataV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/get_custom_task_data/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.challengeTaskId == nil {
		return localVarReturnValue, nil, ReportError("challengeTaskId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize > 100 {
		return localVarReturnValue, nil, ReportError("pageSize must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
