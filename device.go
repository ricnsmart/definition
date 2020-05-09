package definition

const (

	/*设备/线路状态*/
	OFFLINE = 0 // 离线
	ONLINE  = 1 // 在线
	WARN    = 2 // 预警
	ALERT   = 3 // 报警
	Fault   = 4 // 故障
)

// 用于设备迁移
type Migration struct {
	IP   string
	Port int
}

func DeviceStatusToString(status int) (result string) {
	switch status {
	case OFFLINE:
		result = "离线"
	case ONLINE:
		result = "正常"
	case WARN:
		result = "预警"
	case ALERT:
		result = "报警"
	case Fault:
		result = "故障"
	}
	return
}
