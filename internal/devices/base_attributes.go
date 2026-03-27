package devices

type BaseAttributes struct {
	CustomName      string `json:"customName"`
	Model           string `json:"model"`
	Manufacturer    string `json:"manufacturer"`
	FirmwareVersion string `json:"firmwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	SerialNumber    string `json:"serialNumber"`
	ProductCode     string `json:"productCode"`
}
