package definition

import "time"

const (
	/*RabbitMQ Queue*/
	DeviceServiceQueue              = "device.service"                // 设备服务信息
	DeviceStatusQueue               = "device.status"                 // 设备状态
	DeviceMetricDataQueue           = "device.metric.data"            // 设备（网关）各项指标数据
	DeviceMetricActionQueue         = "device.metric.action"          // 下发设备指标动作：读，写，遥控
	DeviceMetricActionResponseQueue = "device.metric.action.response" // 下发设备指标动作后设备的响应
	DeviceSettingDataQueue          = "device.setting.data"           // 设备基础配置
	DeviceAlarmSettingQueue         = "device.alarm.setting"          // 设备警报配置
	DeviceAlarmNotificationQueue    = "device.alarm.notification"     // 设备、线路警报通知
	LineStatusQueue                 = "line.status"                   // 线路状态
	LineMetricDataQueue             = "line.metric.data"              // 线路各项指标数据

	PushDeviceServiceFailed              = "推送设备服务信息失败"
	PushDeviceStatusFailed               = "推送设备状态失败"
	PushDeviceMetricDataFailed           = "推送设备指标数据失败"
	PushDeviceMetricActionFailed         = "推送设备指标动作失败"
	PushDeviceMetricActionResponseFailed = "推送设备指标动作响应失败"
	PushDeviceSettingDataFailed          = "推送设备基础配置失败"
	PushDeviceAlarmSettingFailed         = "推送设备警报配置失败"
	PushDeviceAlarmNotificationFailed    = "推送警报通知失败"
	PushLineStatusFailed                 = "推送线路状态失败"
	PushLineMetricDataFailed             = "推送线路指标数据失败"
)

type (
	// 向设备服务请求的格式
	MetricRequest struct {
		SN      string `json:"sn" bson:"sn"`
		Method  uint8  `json:"method,string" bson:"method"`
		Metrics string `json:"metrics" bson:"metrics"`
		Value   string `json:"value" bson:"value"`
		LineNo  uint8  `json:"line_no,string"`
	}

	// 设备服务对请求的响应
	MetricResponse struct {
		Code    string `json:"code" bson:"code"`
		Message string `json:"message" bson:"message"`
		Data    string `json:"data" bson:"data"`
	}

	// 设备状态
	DeviceStatus struct {
		SN         string    `json:"sn" bson:"sn"`
		DeviceType string    `json:"device_type" bson:"device_type"`
		Status     int       `json:"status" bson:"status" `
		Host       string    `json:"host" bson:"host"`
		ChangeTime time.Time `json:"change_time" bson:"change_time"` // 状态变更时间
	}

	// 线路状态
	LineStatus struct {
		SN         string    `json:"sn" bson:"sn"`
		Status     int       `json:"status" bson:"status"`
		LineNo     int       `json:"line_no" bson:"line_no"`
		ChangeTime time.Time `json:"change_time" bson:"change_time"`
	}

	// 服务信息
	// 为了防止port设置失误，导致端口冲突等问题，还是显式申明port
	Component struct {
		Name        string    `json:"name" bson:"name"`
		Port        int       `json:"port" bson:"port"`
		Host        string    `json:"host" bson:"host"`
		Version     string    `json:"version" bson:"version"`
		BuildTime   string    `json:"build_time" bson:"build_time"`     // 设备服务构建时间
		StartupTime time.Time `json:"startup_time" bson:"startup_time"` // 设备服务启动时间
	}

	// 设备警报详细信息
	DeviceAlarm struct {
		SN         string    `json:"sn" bson:"sn"`
		Metric     string    `json:"metric" bson:"metric"`
		AlarmType  int       `json:"alarm_type" bson:"alarm_type"`
		DeviceType string    `json:"device_type" bson:"device_type"`
		Current    float64   `json:"current" bson:"current"`
		SetValue   float64   `json:"set_value" bson:"set_value"`
		SendTime   time.Time `json:"send_time" bson:"send_time"` // 警报发出时间
	}

	// 线路警报详细信息
	LineAlarm struct {
		SN         string    `json:"sn" bson:"sn"`
		LineNo     int       `json:"line_no" bson:"line_no"`
		Metric     string    `json:"metric" bson:"metric"`
		AlarmType  int       `json:"alarm_type" bson:"alarm_type"`
		DeviceType string    `json:"device_type" bson:"device_type"`
		Current    float64   `json:"current" bson:"current"`
		SetValue   float64   `json:"set_value" bson:"set_value"`
		SendTime   time.Time `json:"send_time" bson:"send_time"` // 警报发出时间
	}
)
