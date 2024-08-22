/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkg 推广产品资质
type AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkg struct {
	// 来自【推广产品资质规则配置查询接口】，行业的资质规则中的config_id
	ConfigId         int64                                                             `json:"config_id"`
	NecessaryCombine *AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgNecessaryCombine `json:"necessary_combine,omitempty"`
	// 推广产品组id ① 针对新增的场景传0即可  ② 提交推广产品资质时系统会返回生成的pkg_id，针对编辑的场景，再此传入对应需要编辑的pkg_id即可
	PkgId int64 `json:"pkg_id"`
	// 推广产品资质的产品名称，字符长度1~128
	ProductName string `json:"product_name"`
	// 选填资质模块，数组长度0~30
	UnnecessaryCombines []*AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInner `json:"unnecessary_combines,omitempty"`
}
