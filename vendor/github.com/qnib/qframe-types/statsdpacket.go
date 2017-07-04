package qtypes


type StatsdPacket struct {
	Bucket   	string
	ValFlt   	float64
	ValStr   	string
	Modifier 	string
	Sampling 	float32
	Group		Dimensions
	Dimensions	Dimensions
}

func NewStatsdPacket(bucket string, val float64, modifier string) *StatsdPacket {
	return &StatsdPacket{
		Bucket: bucket,
		ValFlt: val,
		Modifier: modifier,
		Sampling: float32(1),
		Group: NewDimensions(),
		Dimensions: NewDimensions(),
	}
}

func (sd *StatsdPacket) AddDimension(key, val string) {
	sd.Dimensions.Add(key, val)
}

func (sd *StatsdPacket) DimensionString() string {
	return sd.Dimensions.String()
}
