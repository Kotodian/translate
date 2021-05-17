package translate

type DeviceRegistrationReq struct {
	Header `json:",inline"`
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

func init() {
	register["BootNotificationRequest"] = NewDeviceRegistrationRequest
	register["BootNotificationResponse"] = NewDeviceRegistrationResponse
}
func NewDeviceRegistrationRequest() Translate {
	return &DeviceRegistrationReq{}
}

func (d *DeviceRegistrationReq) Head() Header {
	return Header{
		MessageType: d.MessageType,
		MessageID:   d.MessageID,
		Action:      d.Action,
	}
}

func (d *DeviceRegistrationReq) Init(messageType int32, messageID string, action string) {
	d.MessageType = messageType
	d.MessageID = messageID
	d.Action = action
}

type DeviceRegistrationResp struct {
	Header `json:",inline"`
	Status int32
}

func NewDeviceRegistrationResponse() Translate {
	return &DeviceRegistrationResp{}
}
func (d *DeviceRegistrationResp) Init(messageType int32, messageID string, action string) {
	d.MessageType = messageType
	d.MessageID = messageID
	d.Action = action

}

func (d *DeviceRegistrationResp) Head() Header {
	return Header{
		MessageType: d.MessageType,
		MessageID:   d.MessageID,
		Action:      d.Action,
	}
}
