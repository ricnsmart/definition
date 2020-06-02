package definition

const (

	// 设备类型
	RCNVJ   = "RCNVJ"
	RCN350F = "RCN350F"
	RCN210  = "RCN210"
	PMC350  = "PMC350"
	PMC350F = "PMC350F"
	ASCP200 = "ASCP200"
	GS524N  = "GS524N"
	GS894C  = "GS894C"
	TK82    = "TK82"
	TK83    = "TK83"

	/*设备/线路状态*/
	OFFLINE = 0 // 离线
	ONLINE  = 1 // 在线
	WARN    = 2 // 预警
	ALERT   = 3 // 报警
	Fault   = 4 // 故障
)

// 用于设备迁移
type Migration struct {
	IP      string
	Port    int
	Pattern string // 正则模式
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
