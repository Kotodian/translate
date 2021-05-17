package translate

type DeviceRegistrationReq struct {
	// SerialNumber 设备sn
	SerialNumber string `json:"serial_number"`
	// Model
	// urn:x-oca:ocpp:uid:1:569325
	Model string `json:"model"`
	// FirmwareVersion 固件版本
	FirmwareVersion string `json:"firmware_version"`
	// VendorName 厂商名称
	VendorName string `json:"vendor_name"`
}

type DeviceRegistrationResp struct {
	Status int32
}
