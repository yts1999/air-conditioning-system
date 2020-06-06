package serializer

// Center 中央空调序列化器
type Center struct {
	PowerOn     bool    `json:"power_on"`
	WorkMode    uint    `json:"work_mode"`
	DefaultTemp float32 `json:"default_temp"`
	LowestTemp  float32 `json:"lowest_temp"`
	HighestTemp float32 `json:"highest_temp"`
}

// BuildCenter 序列化中央空调
func BuildCenter(powerOn bool, workMode uint, defaultTemp float32, lowestTemp float32, highestTemp float32) Center {
	return Center{
		PowerOn:     powerOn,
		WorkMode:    workMode,
		DefaultTemp: defaultTemp,
		LowestTemp:  lowestTemp,
		HighestTemp: highestTemp,
	}
}

// BuildCenterResponse 序列化中央空调响应
func BuildCenterResponse(powerOn bool, workMode uint, defaultTemp float32, lowestTemp float32, highestTemp float32) Response {
	return Response{
		Data: BuildCenter(powerOn, workMode, defaultTemp, lowestTemp, highestTemp),
	}
}
