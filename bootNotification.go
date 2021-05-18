package translate

type BootNotificationRequest struct {
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
	register["BootNotificationRequest"] = NewBootNotificationRequest
	register["BootNotificationResponse"] = NewBootNotificationResponse
}
func NewBootNotificationRequest() Translate {
	return &BootNotificationRequest{}
}

func (d *BootNotificationRequest) Head() Header {
	return Header{
		MessageType: d.MessageType,
		MessageID:   d.MessageID,
		Action:      d.Action,
	}
}

func (d *BootNotificationRequest) Init(messageType int32, messageID string, action string) {
	d.MessageType = messageType
	d.MessageID = messageID
	d.Action = action
}

const (
	RegistrationStatusRejected = 0
	RegistrationStatusAccepted = 1
	RegistrationStatusPending  = 2
)

type BootNotificationResponse struct {
	Header `json:",inline"`
	Status int32
}

func NewBootNotificationResponse() Translate {
	return &BootNotificationResponse{}
}
func (d *BootNotificationResponse) Init(messageType int32, messageID string, action string) {
	d.MessageType = messageType
	d.MessageID = messageID
	d.Action = action

}

func (d *BootNotificationResponse) Head() Header {
	return Header{
		MessageType: d.MessageType,
		MessageID:   d.MessageID,
		Action:      d.Action,
	}
}
