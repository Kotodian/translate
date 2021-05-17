package translate

type DeviceRegistrationReq struct {
	Common `json:",inline"`
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

func (d *DeviceRegistrationReq) Com() Common {
	return Common{
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
	Common `json:",inline"`
	Status int32
}

func (d *DeviceRegistrationResp) Init(messageType int32, messageID string, action string) {
	d.MessageType = messageType
	d.MessageID = messageID
	d.Action = action

}

func (d *DeviceRegistrationResp) Comm() Common {
	return Common{
		MessageType: d.MessageType,
		MessageID:   d.MessageID,
		Action:      d.Action,
	}

}
