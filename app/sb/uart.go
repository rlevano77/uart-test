package sb

import (
	"github.com/tarm/serial"
	"encoding/json"
	"bufio"
	"sync"
	"bytes"
	"strings"
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

func mset_time(hour int32, minutes int32, weekday string) ([]byte,error){
	p, err := json.Marshal(SSetTime{
		Hour: hour,
		Minutes: minutes,
		Weekday: strings.ToLower(weekday),
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_TIME_WEEKDAY,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_time() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_TIME_WEEKDAY,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_version() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_FW_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_serial() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_SERIAL_URI,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_heat_pump_emer(emer string) ([]byte,error){
	if(validOption(emer,heat_pump_emer_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_HEAT_PUMP_EMER,
			Cmd:  POST,
			Payload: emer,
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
}

func mget_heat_pump_emer() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_HEAT_PUMP_EMER,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_stage2stage_delay(delay int32) ([]byte,error){
	p, err := json.Marshal(SStage2StageDelay{
		Delay: delay,
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_STAGE_TO_STAGE_DELAY,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_stage2stage_delay() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_STAGE_TO_STAGE_DELAY,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_swing_temperature(swing float32, swing_mode string) ([]byte,error){
	if(validOption(swing_mode,swing_options)){
		p, err := json.Marshal(SSwingTemperature{
			Swing: int32(swing * 10),
			SwingMode: swing_mode,
		})
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_SWING_TEMP,
			Cmd:  POST,
			Payload: string(p),
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
}

func get_swing_temperature() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_SWING_TEMP,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_temperature_differential(heat_diff float32, cool_diff float32) ([]byte,error){
	p, err := json.Marshal(SSetTempDiff{
		HeatDiff: int32(heat_diff * 10),
		CoolDiff: int32(cool_diff * 10),
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_TEMP_DIFFERENTIAL,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_temperature_differential() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_TEMP_DIFFERENTIAL,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_recovery(recovery string) ([]byte,error){
	if(validOption(recovery,recovery_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_RECOVERY_MODE,
			Cmd:  POST,
			Payload: recovery,
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
}

func mget_recovery() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_RECOVERY_MODE,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_humidity_setpoint(humidity int32) ([]byte,error){
	p, err := json.Marshal(SSetHumidity{
		Humidity: humidity,
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_HUMIDITY_SETPOINT,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_humidity_setpoint() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_HUMIDITY_SETPOINT,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_dehumidifier_setpoint(dehumidifier int32) ([]byte,error){
	p, err := json.Marshal(SSetDehumidifier{
		Dehumidifier: dehumidifier,
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_DEHUMIDIFIER_SETPOINT,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_dehumidifier_setpoint() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_DEHUMIDIFIER_SETPOINT,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_external_dehumidifier_setpoint(extdehumi int32) ([]byte,error){
	p, err := json.Marshal(SSetExtDehumidifier{
		Extdehumi: extdehumi,
	})
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_EXT_DEHUMIDIFIER_SETPOINT,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mget_external_dehumidifier_setpoint() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_EXT_DEHUMIDIFIER_SETPOINT,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_humidifier_mode(mode string) ([]byte,error){
	if(validOption(mode,humidifier_mode_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_HUMIDIFIER_MODE,
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

func mget_humidifier_mode() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_HUMIDIFIER_MODE,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_dehumidifier_mode(mode string) ([]byte,error){
	if(validOption(mode,dehumidifier_mode_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_DEHUMIDIFIER_MODE,
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

func mget_dehumidifier_mode() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_DEHUMIDIFIER_MODE,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_external_dehumidifier_mode(mode string) ([]byte,error){
	if(validOption(mode,ext_dehumidifier_mode_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_EXT_DEHUMIDIFIER_MODE,
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

func mget_external_dehumidifier_mode() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_EXT_DEHUMIDIFIER_MODE,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_control_relay(relays SSRelays) ([]byte,error){
	p, err := json.Marshal(relays)
	b, err := json.Marshal(SCmd{
		Uri:	TSTAT_CONTROL_RELAY,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

func mset_ex_relay(ex string) ([]byte,error){
	if(validOption(ex,ex_relay_options)){
		b, err := json.Marshal(SCmd{
			Uri:	TSTAT_EX_RELAY,
			Cmd:  POST,
			Payload: ex,
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
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

/********************************
	PSOC5 Bootloader methods
********************************/

func psoc5_cyboot_load() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	CYBOOT_LOAD_URI,
		Cmd:  POST,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

/********************************
	PIC18 Bootloader methods
********************************/

/*
	PIC18 Bootloader - Erase specified page in flash / Set page address point
	adddress range 0 ~ 399
*/
func hvac_erase_set_page(mode string, addr int32) ([]byte,error){
	if(validOption(mode,page_mode_options)){
		p, err := json.Marshal(SEraseSet{
			Mode: mode,
			Address: addr,
		})
		b, err := json.Marshal(SCmd{
			Uri:	PICBOOT_ERASE_SET_PAGE,
			Cmd:  POST,
			Payload: string(p),
		})
		if err != nil {
			log.Printf("Error Marshalling: %v \n", err)
		}
		return request(b)
	} else {
		panic(errors.New("Wrong option string"))
	}
	
}

/*
	PIC18 Bootloader - Set CRC code of full OTA image
*/
func hvac_set_crc_full_ota(crc_H int32, crc_L int32) ([]byte,error){
	p, err := json.Marshal(SCRCotaFull{
		CRCH: crc_H,
		CRCL: crc_L,
	})
	b, err := json.Marshal(SCmd{
		Uri:	PICBOOT_SET_CRC_FULL_OTA,
		Cmd:  POST,
		Payload: string(p),
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

/*
	PIC18 Bootloader - Get CRC code of full OTA image
*/
func hvac_get_crc_full_ota() ([]byte,error){
	b, err := json.Marshal(SCmd{
		Uri:	PICBOOT_SET_CRC_FULL_OTA,
		Cmd:  GET,
		Payload: "",
	})
  if err != nil {
    log.Printf("Error Marshalling: %v \n", err)
	}
	return request(b)
}

