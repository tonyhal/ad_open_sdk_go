/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

// ToolsCommentReplyGetV30ApiService ToolsCommentReplyGetV30Api service
type ToolsCommentReplyGetV30ApiService service

type ApiOpenApiV30ToolsCommentReplyGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsCommentReplyGetV30ApiService
	advertiserId *int64
	commentId    *int64
	page         *int64
	pageSize     *int64
}

// 广告主id
func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 评论id
func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) CommentId(commentId int64) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	r.commentId = &commentId
	return r
}

// 页码
func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) Page(page int64) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) Execute() (*ToolsCommentReplyGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsCommentReplyGetGet Method for OpenApiV30ToolsCommentReplyGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsCommentReplyGetGetRequest
*/
func (a *ToolsCommentReplyGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsCommentReplyGetGetRequest {
	return &ApiOpenApiV30ToolsCommentReplyGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCommentReplyGetV30Response
func (a *ToolsCommentReplyGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsCommentReplyGetGetRequest) (*ToolsCommentReplyGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsCommentReplyGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/comment_reply/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.commentId == nil {
		return localVarReturnValue, nil, ReportError("commentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "comment_id", r.commentId)
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
