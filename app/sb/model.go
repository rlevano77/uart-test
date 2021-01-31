package sb

import (
	"time"
)

var (
	SERIAL_PORT = "/dev/ttyS0"
	SERIAL_SPEED = 9600
	SERIAL_READ_TIMEOUT = 250*time.Millisecond
	PIC18_SERIAL_READ_TIMEOUT = 500*time.Millisecond
)

// Command Struct
type SCmd struct {
	Uri     string  `json:"uri"`
	Cmd     string  `json:"cmd"`
	Payload string `json:"payload"`
}

// Beep Payload Struct
type SBeepPayload struct {
	Duration int32 `json:"duration"`
}

// Offset Payload Struct
type SOffsetPayload struct {
	Offset int32 `json:"offset"`
}

// Set Config Settings Struct
type SSetConfigSettings struct {
	HeatAuxType    string `json:"heat_aux_type"`
	HeatStage      string `json:"heat_stage"`
	CoolStage      string `json:"cool_stage"`
	HeatPumpConfig string `json:"heat_pump_config"`
}

// Set Test Ambient Conditions
type SSetAmbientCond struct {
	Temperature int32 `json:"temperature"`
	Humidity    int32 `json:"humidity"`
}

// Set Target
type SSetTarget struct {
	THeat int32 `json:"t_heat"`
	TCool int32 `json:"t_cool"`
}

// Set Time
type SSetTime struct {
	Hour    int32    `json:"hour"`
	Minutes int32    `json:"minutes"`
	Weekday string `json:"weekday"`
}

// Set Stage-to-Stage Delay
type SStage2StageDelay struct {
	Delay int32 `json:"delay"`
}

// Set Swing Temperature
type SSwingTemperature struct {
	Swing     int32    `json:"swing"`
	SwingMode string `json:"swing_mode"`
}

// Set Temperature Differential
type SSetTempDiff struct {
	HeatDiff int32 `json:"heat_diff"`
	CoolDiff int32 `json:"cool_diff"`
}

// Set Humidity Setpoint
type SSetHumidity struct {
	Humidity int32 `json:"humidity"`
}

// Set De-Humidifier Setpoint
type SSetDehumidifier struct {
	Dehumidifier int32 `json:"dehumidifier"`
}

// Set External De-Humidifier Setpoint
type SSetExtDehumidifier struct {
	Extdehumi int32 `json:"extdehumi"`
}

// Set Thermostat Control relay independently
type SSRelays struct {
	W  int32 `json:"W"`
	W2 int32 `json:"W2"`
	W3 int32 `json:"W3"`
	Y  int32 `json:"Y"`
	Y2 int32 `json:"Y2"`
	G  int32 `json:"G"`
	OB int32 `json:"OB"`
	H  int32 `json:"H"`
	DH int32 `json:"DH"`
	EX int32 `json:"EX"`
}

// Erase specified page in flash / Set page address point
type SEraseSet struct {
	Mode    string `json:"mode"`
	Address int32    `json:"address"`
}

// CRC struct
type S_CRC struct {
	CRCH byte `json:"CRC_H"`
	CRCL byte `json:"CRC_L"`
}

type S_QMod struct {
	Q   int32 `json:"q"`
	Mod int32 `json:"mod"`
}

type S_Page struct {
	Idx  int32    `json:"idx"`
	Data string `json:"data"`
}

// OTA image status
type S_OTAStatus struct {
	OtaImageStatus int `json:"ota_image_status"`
}












