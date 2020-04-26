package definition

import "time"

const (
	/*RabbitMQ Queue*/
	ComponentInfoQueue               = "gateway/component/info"         // 设备网关组件信息
	DeviceStatusQueue                = "device/status"                  // 设备状态
	DeviceMetricsDataQueue           = "device/metrics/data"            // 设备（网关）各项指标数据
	DeviceMetricsActionQueue         = "device/metrics/action"          // 设备指标动作：读，写，遥控
	DeviceMetricsActionResponseQueue = "device/metrics/action/response" // 设备指标动作结果响应
	DeviceSettingDataQueue           = "device/setting/data"            // 设备基础配置
	DeviceAlarmSettingQueue          = "device/alarm/setting"           // 设备警报配置
	DeviceAlarmNotificationQueue     = "device/alarm/notifications"     // 设备警报通知

	LineStatusQueue      = "line/status"       // 线路状态
	LineMetricsDataQueue = "line/metrics/data" // 线路各项指标数据
	LineResetQueue       = "line/reset"        // 线路重置通知
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

	// 设备/线路状态
	DeviceStatus struct {
		Timestamp time.Time `bson:"Timestamp"`
		SN        string    `bson:"SN"`
		Status    int       `bson:"Status"`
		Host      string    `bson:"Host"`
		LineNo    uint8     `json:"LineNo,string" bson:"LineNo"`
	}

	// 组件信息
	Component struct {
		Timestamp time.Time `bson:"Timestamp"`
		Name      string    `bson:"Name"`
		Host      string    `bson:"Host"`
		Port      string    `bson:"Port"`
		Version   string    `bson:"Version"`
		BuildTime string    `bson:"BuildTime"`
	}

	// 警报详细信息
	Alarm struct {
		Timestamp  time.Time `bson:"Timestamp"`
		SN         string
		LineNo     uint16
		Metric     string
		AlarmType  int
		DeviceType string
		Current    float32
		SetValue   float32
	}

	Alarms []*Alarm
)
