/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationSubmitV30RequestSubject 主体资质，相关字段见下:
type AdvertiserQualificationSubmitV30RequestSubject struct {
	// 详细地址
	Address *string `json:"address,omitempty"`
	// 资质图片附件id
	AttachmentId string                                           `json:"attachment_id"`
	CheckType    AdvertiserQualificationSubmitV30SubjectCheckType `json:"check_type"`
	// 公司名称
	CompanyName string                                             `json:"company_name"`
	CompanyType AdvertiserQualificationSubmitV30SubjectCompanyType `json:"company_type"`
	// 过期时间，格式yyyy-mm-dd
	EffectiveDate *string `json:"effective_date,omitempty"`
	// 是否有有效日期
	HasEffectiveDate bool `json:"has_effective_date"`
	// 法人
	ProprietorName string `json:"proprietor_name"`
	// 资质编号
	QualificationCode string `json:"qualification_code"`
	// 资质id
	QualificationId   *int64                                                   `json:"qualification_id,omitempty"`
	QualificationType AdvertiserQualificationSubmitV30SubjectQualificationType `json:"qualification_type"`
	// 注册城市
	RegisteredCityName *string `json:"registered_city_name,omitempty"`
	// 注册国家
	RegisteredNationName string `json:"registered_nation_name"`
	// 注册省份
	RegisteredProvinceName *string `json:"registered_province_name,omitempty"`
}
