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
	This function receives an object (dict) with 
	the next signature e.g 
	config = {
		'heat_aux_type': 'gas/oil',
		'heat_stage': '1 stage AUX heat',
		'cool_stage': '1-stage',   				 				
		'heat_pump_config': 'B'
	}
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