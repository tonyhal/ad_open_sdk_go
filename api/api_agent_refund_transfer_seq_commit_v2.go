/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// AgentRefundTransferSeqCommitV2ApiService AgentRefundTransferSeqCommitV2Api service
type AgentRefundTransferSeqCommitV2ApiService service

type ApiOpenApi2AgentRefundTransferSeqCommitPostRequest struct {
	ctx                                   context.Context
	ApiService                            *AgentRefundTransferSeqCommitV2ApiService
	agentRefundTransferSeqCommitV2Request *AgentRefundTransferSeqCommitV2Request
}

func (r *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest) AgentRefundTransferSeqCommitV2Request(agentRefundTransferSeqCommitV2Request AgentRefundTransferSeqCommitV2Request) *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest {
	r.agentRefundTransferSeqCommitV2Request = &agentRefundTransferSeqCommitV2Request
	return r
}

func (r *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest) Execute() (*AgentRefundTransferSeqCommitV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest) AccessToken(accessToken string) *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest) WithLog(enable bool) *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AgentRefundTransferSeqCommitPost Method for OpenApi2AgentRefundTransferSeqCommitPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AgentRefundTransferSeqCommitPostRequest
*/
func (a *AgentRefundTransferSeqCommitV2ApiService) Post(ctx context.Context) *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest {
	return &ApiOpenApi2AgentRefundTransferSeqCommitPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentRefundTransferSeqCommitV2Response
func (a *AgentRefundTransferSeqCommitV2ApiService) postExecute(r *ApiOpenApi2AgentRefundTransferSeqCommitPostRequest) (*AgentRefundTransferSeqCommitV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AgentRefundTransferSeqCommitV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/agent/refund/transfer_seq/commit/"

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

	// body params
	localVarPostBody = r.agentRefundTransferSeqCommitV2Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
