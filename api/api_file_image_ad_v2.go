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

// FileImageAdV2ApiService FileImageAdV2Api service
type FileImageAdV2ApiService service

type ApiOpenApi2FileImageAdPostRequest struct {
	ctx            context.Context
	ApiService     *FileImageAdV2ApiService
	advertiserId   *int64
	filename       *string
	imageFile      *FormFileInfo
	imageSignature *string
	imageUrl       *string
	isAigc         *bool
	uploadType     *FileImageAdV2UploadType
}

// 广告主ID
func (r *ApiOpenApi2FileImageAdPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileImageAdPostRequest {
	r.advertiserId = &advertiserId
	return r
}

// 素材的文件名，可自定义素材名，不传则默认取文件名，最长255个字符 注：若同一素材已进行上传，重新上传不会改名
func (r *ApiOpenApi2FileImageAdPostRequest) Filename(filename string) *ApiOpenApi2FileImageAdPostRequest {
	r.filename = &filename
	return r
}

// 图片文件 upload_type为UPLOAD_BY_FILE必填 格式: jpg、jpeg、png、bmp、gif, 大小1.5M内
func (r *ApiOpenApi2FileImageAdPostRequest) ImageFile(imageFile *FormFileInfo) *ApiOpenApi2FileImageAdPostRequest {
	r.imageFile = imageFile
	return r
}

// 图片的md5值(用于服务端校验) upload_type为UPLOAD_BY_FILE
func (r *ApiOpenApi2FileImageAdPostRequest) ImageSignature(imageSignature string) *ApiOpenApi2FileImageAdPostRequest {
	r.imageSignature = &imageSignature
	return r
}

// 图片url地址，如http://xxx.xxx upload_type为UPLOAD_BY_URL必填
func (r *ApiOpenApi2FileImageAdPostRequest) ImageUrl(imageUrl string) *ApiOpenApi2FileImageAdPostRequest {
	r.imageUrl = &imageUrl
	return r
}

// 图片素材是否为AIGC生成
func (r *ApiOpenApi2FileImageAdPostRequest) IsAigc(isAigc bool) *ApiOpenApi2FileImageAdPostRequest {
	r.isAigc = &isAigc
	return r
}

func (r *ApiOpenApi2FileImageAdPostRequest) UploadType(uploadType FileImageAdV2UploadType) *ApiOpenApi2FileImageAdPostRequest {
	r.uploadType = &uploadType
	return r
}

func (r *ApiOpenApi2FileImageAdPostRequest) Execute() (*FileImageAdV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2FileImageAdPostRequest) AccessToken(accessToken string) *ApiOpenApi2FileImageAdPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileImageAdPostRequest) WithLog(enable bool) *ApiOpenApi2FileImageAdPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileImageAdPost Method for OpenApi2FileImageAdPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileImageAdPostRequest
*/
func (a *FileImageAdV2ApiService) Post(ctx context.Context) *ApiOpenApi2FileImageAdPostRequest {
	return &ApiOpenApi2FileImageAdPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileImageAdV2Response
func (a *FileImageAdV2ApiService) postExecute(r *ApiOpenApi2FileImageAdPostRequest) (*FileImageAdV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileImageAdV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/image/ad/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "filename", r.filename)
	}
	if r.imageFile != nil {
		formFiles["image_file"] = r.imageFile
	}
	if r.imageSignature != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "image_signature", r.imageSignature)
	}
	if r.imageUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "image_url", r.imageUrl)
	}
	if r.isAigc != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_aigc", r.isAigc)
	}
	if r.uploadType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload_type", r.uploadType)
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
