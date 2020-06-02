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

	PushServiceFailed                    = "推送设备服务信息失败"
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
		SN     string
		Method uint8 `json:"method,string"`
		Metric string
		Value  string
		LineNo uint8 `json:"line_no,string"`
	}

	// 设备服务对请求的响应
	MetricResponse struct {
		Code    string
		Message string
		Data    string
	}

	// 设备状态
	DeviceStatus struct {
		ChangeTime time.Time `bson:"change_time"` // 状态变更时间
		SN         string    `bson:"sn"`
		DeviceType string    `bson:"device_type"`
		Status     int       `bson:"status" `
		Host       string    `bson:"host"`
	}

	// 线路状态
	LineStatus struct {
		ChangeTime time.Time `bson:"change_time"`
		SN         string    `bson:"sn"`
		Status     int       `bson:"status"`
		LineNo     int       `json:"line_no" bson:"line_no"`
	}

	// 服务信息
	// 为了防止port设置失误，导致端口冲突等问题，还是显式申明port
	Component struct {
		StartupTime time.Time `bson:"startup_time"` // 设备服务启动时间
		Name        string    `bson:"name"`
		Port        int       `bson:"port"`
		Host        string    `bson:"host"`
		Version     string    `bson:"version"`
		BuildTime   string    `bson:"build_time"` // 设备服务构建时间
	}

	// 设备警报详细信息
	DeviceAlarm struct {
		SendTime   time.Time `bson:"send_time"` // 警报发出时间
		SN         string    `bson:"sn"`
		Metric     string    `bson:"metric"`
		AlarmType  int       `bson:"alarm_type"`
		DeviceType string    `bson:"device_type"`
		Current    float64   `bson:"current"`
		SetValue   float64   `bson:"set_value"`
	}

	// 线路警报详细信息
	LineAlarm struct {
		SendTime   time.Time `bson:"send_time"` // 警报发出时间
		SN         string    `bson:"sn"`
		LineNo     int       `json:"line_no" bson:"line_no"`
		Metric     string    `bson:"metric"`
		AlarmType  int       `bson:"alarm_type"`
		DeviceType string    `bson:"device_type"`
		Current    float64   `bson:"current"`
		SetValue   float64   `bson:"set_value"`
	}
)
