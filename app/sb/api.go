package sb



/* Enable TEST_MODE in Support Board */
func Enable_api_test_mode() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = menable_api_test_mode()
	}
	sb_mutex.Unlock()
	return response, err
}

/* Disable TEST_MODE in Support Board */
func Disable_api_test_mode() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mdisable_api_test_mode()
	}
	sb_mutex.Unlock()
	return response, err
}

/* GET Support Board sensors data */
func Read_Sensors() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mread_sensors()
	}
	sb_mutex.Unlock()
	return response, err
}

// GET Support Board Poximity sensor data
func Read_Proximity() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mread_proximity()
	}
	sb_mutex.Unlock()
	return response, err	
}

// GET Support Board IAQ sensor data
func Read_IAQ() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mread_iaq()
	}
	sb_mutex.Unlock()
	return response, err	
}

/* 
	Set Support Board OFFSET
	offset temperature in Fahrenheit e.g 11.4
	use example : sb.Set_Offset(12.6)
*/ 
func Set_Offset(offset int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_offset(offset)
	}
	sb_mutex.Unlock()
	return response, err	
}


/* Get Support Board OFFSET X 10 */
func Get_Offset() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_offset()
	}
	sb_mutex.Unlock()
	return response, err	
}


/* RESET Cypress PSoC5 MCU */
func PSOC5_Reset() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mpsoc5_reset()
	}
	sb_mutex.Unlock()
	return response, err	
}

/* Get Cypress PSoC5 Firmware Version */
func Get_psoc5_version() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_psoc5_version()
	}
	sb_mutex.Unlock()
	return response, err	
}

/* Get Cypress PSoC5 Die Temperature */
func Get_psoc5_die_temp() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_psoc5_die_temp()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Set Support Board BEEP
	duration in milliseconds
	use example : sb.Beep(1000)  #Beep for 1 second = 1000 ms
*/
func Beep(duration int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mbeep(duration)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 1
	Report HVAC status
*/
func Get_Status() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_status()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 2
	Set HVAC config setting
	This function receives an object with 
	the next signature e.g 
	config = {
		'heat_aux_type': 'gas/oil',
		'heat_stage': '1 stage AUX heat',
		'cool_stage': '1-stage',   				 				
		'heat_pump_config': 'B'
	}

	Before and after call this method a Set_Mode("off") has to be called.
	--------------------------------------------------------------------
*/
func Set_config_settings(config SSetConfigSettings) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_config_settings(config)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 2
	Get HVAC config setting
*/
func Get_config_settings() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_config_settings()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 3
	Set Mode of operation
*/
func Set_Mode(mode string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_mode(mode)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 3
	Get Mode of operation
*/
func Get_Mode() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_mode()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 4
	Set Fan setting
*/
func Set_Fan(fan string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_fan(fan)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 4
	Get Fan setting
*/
func Get_Fan() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_fan()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 5
	Read ambient conditions
*/
func Read_ambient_conditions() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mread_ambient_conditions()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 5
	Set ambient conditions
*/
func Set_api_test_ambient_conditions(temperature int32, humidity int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_api_test_ambient_conditions(temperature, humidity)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 6
	Set heat / cool targets
	t_heat range [35-94F]
	t_cool range [35-95F]
*/
func Set_target(t_heat float32, t_cool float32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_target(t_heat, t_cool)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 6
	Set heat / cool targets
*/
func Get_target() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_target()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 7
	Set time & weekday
*/
func Set_time(hour int32, minutes int32, weekday string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_time(hour, minutes, weekday)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 7
	Get time & weekday
*/
func Get_time() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_time()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 8
	Read HVAC board firmware version
*/
func Get_version() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_version()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 9
	Read HVAC board Serial Number
*/
func Get_serial() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_serial()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 10
	Set Heat Pump EMER
*/
func Set_heat_pump_emer(fan string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_heat_pump_emer(fan)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 10
	Get Heat Pump EMER
*/
func Get_heat_pump_emer() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_heat_pump_emer()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 11
	Set Stage-to-Stage Delay
	delay in minutes [0-60]
*/
func Set_stage2stage_delay(delay int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_stage2stage_delay(delay)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 11
	Get Stage-to-Stage Delay
*/
func Get_stage2stage_delay() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_stage2stage_delay()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 12
	Set Swing Temperature
 	swing => [0.0 to 3.0F]
	swing_mode => ['swing one way','swing both ways']
*/
func Set_swing_temperature(swing float32, swing_mode string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_swing_temperature(swing, swing_mode)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 12
	Get Swing Temperature
*/
func Get_swing_temperature() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = get_swing_temperature()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 13
	Set Temperature Differential
	heat_diff, cool_diff range [2.0 - 6.0]
*/
func Set_temperature_differential(heat_diff float32, cool_diff float32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_temperature_differential(heat_diff, cool_diff)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 13
	Get Temperature Differential
*/
func Get_temperature_differential() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_temperature_differential()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 14
	Set HVAC Recovery Mode
*/
func Set_recovery(recovery string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_recovery(recovery)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 14
	Get HVAC Recovery Mode
*/
func Get_recovery() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_recovery()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 15
	Set Humidity Setpoint
	value from 0% to 100%
*/
func Set_humidity_setpoint(humidity int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_humidity_setpoint(humidity)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 15
	Set Humidity Setpoint
*/
func Get_humidity_setpoint() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_humidity_setpoint()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 16
	Set De-Humidifier Setpoint
	value from 20% to 90%
*/
func Set_dehumidifier_setpoint(dehumidifier int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_dehumidifier_setpoint(dehumidifier)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 16
	Set De-Humidifier Setpoint
*/
func Get_dehumidifier_setpoint() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_dehumidifier_setpoint()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 17
	Set External De-Humidifier Setpoint
	value from 20% to 90%
*/
func Set_external_dehumidifier_setpoint(extdehumi int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_external_dehumidifier_setpoint(extdehumi)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 17
	Set External De-Humidifier Setpoint
*/
func Get_external_dehumidifier_setpoint() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_external_dehumidifier_setpoint()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 18
	Set Humidifier Mode
*/
func Set_humidifier_mode(mode string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_humidifier_mode(mode)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 18
	Set Humidifier Mode
*/
func Get_humidifier_mode() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_humidifier_mode()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 19
	Set De-Humidifier Mode
*/
func Set_dehumidifier_mode(mode string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_dehumidifier_mode(mode)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 19
	Get De-Humidifier Mode
*/
func Get_dehumidifier_mode() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_dehumidifier_mode()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 20
	Set External De-Humidifier mode
*/
func Set_external_dehumidifier_mode(mode string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_external_dehumidifier_mode(mode)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 20
	Get External De-Humidifier mode
*/
func Get_external_dehumidifier_mode() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mget_external_dehumidifier_mode()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 21
	Control relay independently
	Operating condition: User mode is OFF, Fan mode is Auto, No humidity control

	Before and after call this method a Set_Mode("off") has to be called.
	---------------------------------------------------------------------
 	Set Thermostat Control relay independently
             	7  6   5   4    3    2    1    0
 -----------------------------------------------
 	1st byte	  W  W2  W3  Y    Y2   G    O/B  N/A
 	2nd byte	  H  DH  EX  N/A  N/A  N/A  N/A  N/A

 This function receives an object with the next signature e.g 
	relays = {
		'W' : 1,
		'W2' : 0,
		'W3' : 1,
		'Y' : 1,
		'Y2' : 0,
		'G' : 1,
		'OB' : 0,
		'H' : 1,
		'DH' : 0,
		'EX' : 0
	}
*/
func Set_control_relay(relays SSRelays) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_control_relay(relays)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 22
	Set Thermostat EX Relay

	Hub CAN control when user mode is OFF, HEAT, COOL, AUTO;
	Hub CANNOT control when user mode is MANUAL;
	EX relay will be Open after reset.
*/
func Set_ex_relay(ex string) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = mset_ex_relay(ex)
	}
	sb_mutex.Unlock()
	return response, err	
}




