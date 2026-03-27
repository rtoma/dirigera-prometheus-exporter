package devices

const DEVICE_SPEAKER_TYPE = "speaker"

type SpeakerAttributes struct {
	BaseAttributes
}

func (s *SpeakerAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(s.CustomName)
}
