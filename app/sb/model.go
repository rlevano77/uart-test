package sb

import (
	"time"
)

var (
	SERIAL_PORT = "/dev/ttyS0"
	SERIAL_SPEED = 9600
	SERIAL_READ_TIMEOUT = 250*time.Millisecond
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