package definition

import "time"

const (
	/*RabbitMQ Queue*/
	ComponentInfoQueue               = "component.info"                 // 设备网关组件信息
	DeviceStatusQueue                = "device.status"                  // 设备状态
	DeviceMetricsDataQueue           = "device.metrics.data"            // 设备（网关）各项指标数据
	DeviceMetricsActionQueue         = "device.metrics.action"          // 下发设备指标动作：读，写，遥控
	DeviceMetricsActionResponseQueue = "device.metrics.action.response" // 下发设备指标动作后设备的响应
	DeviceSettingDataQueue           = "device.setting.data"            // 设备基础配置
	DeviceAlarmSettingQueue          = "device.alarm.setting"           // 设备警报配置
	DeviceAlarmNotificationQueue     = "device.alarm.notifications"     // 设备、线路警报通知
	LineStatusQueue                  = "line.status"                    // 线路状态
	LineMetricsDataQueue             = "line.metrics.data"              // 线路各项指标数据
	LineResetQueue                   = "line.reset"                     // 线路重置通知
	LineOfflineQueue                 = "line.offline"                   // 网关离线通知

	PushComponentInfoFailed               = "推送设备组件信息失败"
	PushDeviceStatusFailed                = "推送设备状态失败"
	PushDeviceMetricsDataFailed           = "推送设备指标数据失败"
	PushDeviceMetricsActionFailed         = "推送设备指标动作失败"
	PushDeviceMetricsActionResponseFailed = "推送设备指标动作响应失败"
	PushDeviceSettingDataFailed           = "推送设备基础配置失败"
	PushDeviceAlarmSettingFailed          = "推送设备警报配置失败"
	PushDeviceAlarmNotificationFailed     = "推送警报通知失败"
	PushLineStatusFailed                  = "推送线路状态失败"
	PushLineMetricsDataFailed             = "推送线路指标数据失败"
	PushLineResetFailed                   = "推送线路重置通知失败"
	PushLineOfflineFailed                 = "推送网关离线通知失败"
)

type (
	// 向设备组件请求的格式
	MetricsRequest struct {
		SN      string
		Method  uint8 `json:"Method,string"`
		Metrics string
		Value   string
		LineNo  uint8 `json:"LineNo,string"`
	}

	// 设备组件对请求的响应
	MetricsResponse struct {
		Code    string
		Message string
		Data    string
	}

	// 设备状态
	DeviceStatus struct {
		Timestamp time.Time `bson:"Timestamp"`
		SN        string    `bson:"SN"`
		Status    int       `bson:"Status"`
		Host      string    `bson:"Host"`
	}

	// 线路状态
	LineStatus struct {
		Timestamp time.Time `bson:"Timestamp"`
		SN        string    `bson:"SN"`
		Status    int       `bson:"Status"`
		LineNo    int       `json:"LineNo" bson:"LineNo"`
	}

	// 组件信息
	// 为了防止port设置失误，导致端口冲突等问题，还是显式申明port
	Component struct {
		Timestamp time.Time `bson:"Timestamp"`
		Name      string    `bson:"Name"`
		Port      int       `bson:"Port"`
		Host      string    `bson:"Host"`
		Version   string    `bson:"Version"`
		BuildTime string    `bson:"BuildTime"`
	}

	// 设备警报详细信息
	DeviceAlarm struct {
		Timestamp  time.Time `bson:"Timestamp"`
		SN         string    `bson:"SN"`
		Metric     string    `bson:"Metric"`
		AlarmType  int       `bson:"AlarmType"`
		DeviceType string    `bson:"DeviceType"`
		Current    float64   `bson:"Current"`
		SetValue   float64   `bson:"SetValue"`
	}

	// 线路警报详细信息
	LineAlarm struct {
		Timestamp  time.Time `bson:"Timestamp"`
		SN         string    `bson:"SN"`
		LineNo     int       `json:"LineNo,string" bson:"LineNo"`
		Metric     string    `bson:"Metric"`
		AlarmType  int       `bson:"AlarmType"`
		DeviceType string    `bson:"DeviceType"`
		Current    float64   `bson:"Current"`
		SetValue   float64   `bson:"SetValue"`
	}
)
