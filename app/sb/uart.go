package sb

import (
	"github.com/tarm/serial"
	"encoding/json"
	"bufio"
	"sync"
	"bytes"
	"log"
	"errors"
)

var (
	serial_port *serial.Port
	sb_mutex = &sync.Mutex{}
	Err_cnt = 0	/* errors count */
)

func Open() error{
	c := &serial.Config{Name: SERIAL_PORT, Baud: SERIAL_SPEED, ReadTimeout: SERIAL_READ_TIMEOUT}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	serial_port = s
	return nil
}

func request(cmd []byte)([]byte, error) {
	/* Write command */
	_, err := serial_port.Write([]byte("<" + string(cmd) + ">"))
	if err != nil {
		serial_port.Flush()
		log.Println("Error sending request.")
		return []byte{},err
	}

	/* Reading response */
	reader := bufio.NewReader(serial_port)
	reply, err := reader.ReadBytes('\x0a')
	if err != nil {
		serial_port.Flush()
		return []byte{},err
	}

	response := bytes.TrimRight(reply,"\n")

	/* If error message from MCU => increase errors count */
	if string(response) == "{\"result\":\"ERROR\"}" {
		Err_cnt = Err_cnt + 1
		if Err_cnt > 100 {
			Err_cnt = 0
		}	
	}

	return response, err
}

/* Middleware */

func menable_api_test_mode() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_ENABLE_API_TEST_MODE,
		Cmd:  POST,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mdisable_api_test_mode() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_DISABLE_API_TEST_MODE,
		Cmd:  POST,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mread_sensors() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	SB_SENSORS_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mread_proximity() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	PROXIMITY_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mread_iaq() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	IAQ_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_offset(offset int32) ([]byte,error){
	p, err := json.Marshal(SOffsetPayload{
		Offset: offset*10,
	})
	b, err := json.Marshal(SCmd{
		Uri:	OFFSET_URI,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_offset() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	OFFSET_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mpsoc5_reset() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	RESET_URI,
		Cmd:  POST,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_psoc5_version() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	PSOC5_FW_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_psoc5_die_temp() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	PSOC5_DIE_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mbeep(duration int32) ([]byte,error){
	p, err := json.Marshal(SBeepPayload{
		Duration: duration,
	})
	b, err := json.Marshal(SCmd{
		Uri:	SB_BEEP_URI,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_status() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_STATUS_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_config_settings(config SSetConfigSettings) ([]byte,error){
	p, err := json.Marshal(config)
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_CONFIG_SETTINGS_URI,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_config_settings() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_CONFIG_SETTINGS_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_mode(mode string) ([]byte,error){	
	if(validOption(mode,mode_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_MODE_URI,
			Cmd:  POST,
			Payload: mode,
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
}

func mget_mode() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_MODE_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_fan(fan string) ([]byte,error){
	if(validOption(fan,fan_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_FAN_URI,
			Cmd:  POST,
			Payload: fan,
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
}

func mget_fan() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_FAN_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mread_ambient_conditions() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_REPORT_AMBIENT_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_api_test_ambient_conditions(temperature int32, humidity int32) ([]byte,error){
	p, err := json.Marshal(SSetAmbientCond{
		Temperature: temperature * 10,
		Humidity: humidity,
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_REPORT_AMBIENT_URI,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_target(t_heat float32, t_cool float32) ([]byte,error){
	p, err := json.Marshal(SSetTarget{
		THeat: int32(t_heat * 10),
		TCool: int32(t_cool * 10),
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_TARGET,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_target() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_TARGET,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func Close(){
	serial_port.Close()
}

func validJSON(b []byte) bool {
	var x struct{}
	var res bool = true
	if err := json.Unmarshal(b, &x); err != nil {
	  res = false
	}
  return res
}

func validOption(val string, options []string) bool {
	var res bool = false
	for _, item := range options {
		if item == val {
			res = true
		}
	}
	return res
}


