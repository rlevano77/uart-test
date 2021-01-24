package sb

import (
	"time"
)

var (
	SERIAL_PORT = "/dev/ttyS0"
	SERIAL_SPEED = 9600
	SERIAL_READ_TIMEOUT = 100*time.Millisecond
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