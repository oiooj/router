package query

import (
	"math"
	"strings"
)

const (
	// DAY is all seconds in a day
	DAY int64 = 86400000
)

// SetPrecision set a precision for a float64
func SetPrecision(from float64, precision int) float64 {
	base := math.Pow10(precision)
	return float64(int64(from*base)) / base
}

func transAgg(agg string) string {
	switch agg {
	case "avg":
		return "AVERAGE"
	case "min":
		return "MIN"
	case "max":
		return "MAX"
	case "sum":
		return "SUM"
	default:
		return ""
	}
}

func transKey(key string) string {
	switch strings.ToLower(key) {
	case "cpu":
		return "CPU"
	case "mem":
		return "内存"
	case "net":
		return "网络"
	case "disk":
		return "硬盘"
	case "fs":
		return "文件系统"
	case "io":
		return "IO"
	case "port":
		return "端口监控"
	case "plugin":
		return "插件监控"
	case "proc":
		return "进程监控"
	case "run":
		return "监控上报"
	default:
		return "其他系统监控"
	}
}

// Detail infludes Measurement detail info
type Detail struct {
	Unit      string `json:"unit"`
	Mode      string `json:"mode"`
	Aggregate string `json:"aggregate"`
	Fill      string `json:"fill"`
}

// MeasurementDetail include measurements detail
func MeasurementDetail(key string) Detail {
	var d Detail

	// switch unit
	if strings.HasPrefix(key, "RUN.net.traffic.") {
		d.Unit = "bit"
		return d
	}

	switch strings.ToLower(key) {
	case "cpu.idle":
		d.Unit = "%"
	case "mem.buffers":
		d.Unit = "MB"
	case "mem.cached":
		d.Unit = "MB"
	case "mem.free":
		d.Unit = "MB"
	case "mem.total":
		d.Unit = "MB"
	case "mem.used":
		d.Unit = "MB"
	case "mem.used.percent":
		d.Unit = "%"
	case "fs.inodes.used.percent":
		d.Unit = "%"
	case "fs.space.used.percent":
		d.Unit = "%"
	case "fs.space.used":
		d.Unit = "MB"
	case "fs.space.free":
		d.Unit = "MB"
	case "fs.space.total":
		d.Unit = "MB"
	case "fs.files.rw":
		d.Mode = "bar"
	case "disk.io.util":
		d.Unit = "%"
	case "disk.io.read_requests":
		d.Unit = "次/秒"
	case "disk.io.write_requests":
		d.Unit = "次/秒"
	case "time.offset":
		d.Unit = "s"
	case "net.out":
		d.Unit = "bit"
	case "net.in":
		d.Unit = "bit"
	case "kernel.files.allocated.percent":
		d.Unit = "%"
	case "run.ping.loss":
		d.Unit = "%"
	}
	return d
}
