/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestAdInfo 广告计划维度信息
type AdlabGroupCreateV30RequestAdInfo struct {
	AdTarget AdlabGroupCreateV30AdInfoAdTarget `json:"ad_target"`
	AppType  *AdlabGroupCreateV30AdInfoAppType `json:"app_type,omitempty"`
	// 资产 id，当使用事件管理时必填 字节小程序链路下，需要传入两个资产id（字节小程序资产+备用落地页资产），其他场景仅需传入一个资产id
	AssetsIds []int64                                   `json:"assets_ids,omitempty"`
	Audience  *AdlabGroupCreateV30RequestAdInfoAudience `json:"audience,omitempty"`
	// 目标转化出价/预期成本
	CpaBid      *float64                              `json:"cpa_bid,omitempty"`
	DeepBidType *AdlabGroupCreateV30AdInfoDeepBidType `json:"deep_bid_type,omitempty"`
	// 深度优化出价 设定深度转化目标时才会有效。 取值范围：0.1~10000元
	DeepCpaBid         *float64                                       `json:"deep_cpa_bid,omitempty"`
	DeepExternalAction *AdlabGroupCreateV30AdInfoDeepExternalAction   `json:"deep_external_action,omitempty"`
	DeliveryRange      *AdlabGroupCreateV30RequestAdInfoDeliveryRange `json:"delivery_range,omitempty"`
	DownloadMode       *AdlabGroupCreateV30AdInfoDownloadMode         `json:"download_mode,omitempty"`
	DownloadType       *AdlabGroupCreateV30AdInfoDownloadType         `json:"download_type,omitempty"`
	// 下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
	// 投放结束时间 当schedule_type为SCHEDULE_START_END时必填，形式如：2017-01-01
	EndTime        *string                                  `json:"end_time,omitempty"`
	ExternalAction *AdlabGroupCreateV30AdInfoExternalAction `json:"external_action,omitempty"`
	// 直达链接。（点击唤起APP）直达链接仅支持部分App唤起，点击创意将优先跳转App，再根据投放内容跳转相关链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 应用包名，需要与应用下载链接中包名一致（格式例如：com.xx.xx）
	Package *string `json:"package,omitempty"`
	// 商品信息，最多允许传入1个
	ProductInfo []*AdlabGroupCreateV30RequestAdInfoProductInfoInner `json:"product_info,omitempty"`
	// 深度转化ROI系数 范围(0,5]，精度：保留小数点后四位
	RoiGoal      *float64                               `json:"roi_goal,omitempty"`
	ScheduleType *AdlabGroupCreateV30AdInfoScheduleType `json:"schedule_type,omitempty"`
	// 投放起始时间 当schedule_type为SCHEDULE_START_END时必填，形式如：2017-01-01
	StartTime       *string                                          `json:"start_time,omitempty"`
	TrackUrlSetting *AdlabGroupCreateV30RequestAdInfoTrackUrlSetting `json:"track_url_setting,omitempty"`
	// 投放时段 使用要求： 1、 数组为空，表示表示全时段投放 2、 第一层数组不为空时，必须为7个元素，依次表示每周从周一到周日7天的投放时间段信息 3、 第二层数组最少0个元素，最多48个元素。每个元素的范围为 [0,47]，表示全天每半个小时是否投放。 - 例如：包含0表示 00:00-00:29 需要进行投放，包含47表示 23:00-23:59需要进行投放 - 例如：周一0点到2点投放，周三0点到1点半投放，其他时段都不投放，表示为：[[0,1,2,3], [], [0,1,2], [], [], [], []]
	WeekSchedule []*[]int64 `json:"week_schedule,omitempty"`
}
