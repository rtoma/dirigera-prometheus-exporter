package devices

const DEVICE_GATEWAY_TYPE = "gateway"

type GatewayAttributes struct {
	BaseAttributes
}

func (g *GatewayAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(g.CustomName)
}
