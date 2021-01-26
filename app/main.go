package main

import (
	"app/sb"
	"log"
	"time"
)

var (
	start = time.Now()
)

func main() {
	
	err := sb.Open()
	if err != nil {
		log.Println("Error opening serialport.")
	}
	log.Println("Serial port opened succesfully.")


	// Polling sensors 
	go func() {
		
		for {
			sensors, err := sb.Read_Sensors()
			if err != nil {
				log.Println("Error reading sensors")	
			}
			log.Printf("Read_Sensors: %s   %v   errors: %d",string(sensors), time.Since(start), sb.Err_cnt)	
		}
		
	}()

	// Polling proximity 
	go func() {
		
		for {
			proximity, err := sb.Read_Proximity()
			if err != nil {
				log.Println("Error reading proximity")	
			}
			log.Printf("Read_Proximity: %s   %v   errors: %d",string(proximity), time.Since(start), sb.Err_cnt)	
		}
		
	}()

	// Polling iaq 
	go func() {
		
		for {
			iaq, err := sb.Read_IAQ()
			if err != nil {
				log.Println("Error reading iaq")	
			}
			log.Printf("Read_IAQ: %s   %v   errors: %d",string(iaq), time.Since(start), sb.Err_cnt)	
		}
		
	}()

	// Polling status 
	go func() {
		
		for {
			status, err := sb.Get_Status()
			if err != nil {
				log.Println("Error reading status")	
			}
			log.Printf("Get_Status: %s   %v   errors: %d",string(status), time.Since(start), sb.Err_cnt)		
		}
		
	}()
	


	/* Main loop */
	for {

		log.Println("Message 2 : Set HVAC config setting")
		log.Println("Changing mode to \"off\" so config changes can be made")

		setmode, err := sb.Set_Mode("off")
		if err != nil {
			log.Println("Error reading status")	
		}
		log.Printf("Set_Mode off: %s   %v   errors: %d",string(setmode), time.Since(start), sb.Err_cnt)		

		getmode, err := sb.Get_Mode()
		if err != nil {
			log.Println("Error getting mode")	
		}
		log.Printf("Get_Mode off: %s   %v   errors: %d",string(getmode), time.Since(start), sb.Err_cnt)	

		




		//sb.Beep(1000)

		log.Printf("Set config settings to : electric, 2H, 1C, O/B for Heat")
		config := sb.SSetConfigSettings {
			HeatAuxType: "electric",
			HeatStage: "2 stage AUX heat",
			CoolStage: "1-stage",
			HeatPumpConfig: "B",
		}
		setcfg, err := sb.Set_config_settings(config)
		if err != nil {
			log.Println("Error setting config settings")	
		}
		log.Printf("Set_config_settings: %s   %v   errors: %d",string(setcfg), time.Since(start), sb.Err_cnt)
		
		





		getcfg, err := sb.Get_config_settings()
		if err != nil {
			log.Println("Error getting config settings")	
		}
		log.Printf("Get_config_settings: %s   %v   errors: %d",string(getcfg), time.Since(start), sb.Err_cnt)


		log.Println("Changing mode to \"off\" so config changes can be made")

		mode_off, err := sb.Set_Mode("off")
		if err != nil {
			log.Println("Error setting mode to off")	
		}
		log.Printf("Set_Mode off: %s   %v   errors: %d",string(mode_off), time.Since(start), sb.Err_cnt)

		read_mode_off, err := sb.Get_Mode()
		if err != nil {
			log.Println("Error setting mode to off")	
		}
		log.Printf("Get_Mode off: %s   %v   errors: %d",string(read_mode_off), time.Since(start), sb.Err_cnt)
		


		mode_heat, err := sb.Set_Mode("heat")
		if err != nil {
			log.Println("Error setting mode to heat")	
		}
		log.Printf("Set_Mode heat: %s   %v   errors: %d",string(mode_heat), time.Since(start), sb.Err_cnt)

		read_mode_heat, err := sb.Get_Mode()
		if err != nil {
			log.Println("Error setting mode to heat")	
		}
		log.Printf("Get_Mode heat: %s   %v   errors: %d",string(read_mode_heat), time.Since(start), sb.Err_cnt)

		mode_cool, err := sb.Set_Mode("cool")
		if err != nil {
			log.Println("Error setting mode to cool")	
		}
		log.Printf("Set_Mode cool: %s   %v   errors: %d",string(mode_cool), time.Since(start), sb.Err_cnt)

		read_mode_cool, err := sb.Get_Mode()
		if err != nil {
			log.Println("Error setting mode to cool")	
		}
		log.Printf("Get_Mode cool: %s   %v   errors: %d",string(read_mode_cool), time.Since(start), sb.Err_cnt)

		mode_auto, err := sb.Set_Mode("auto")
		if err != nil {
			log.Println("Error setting mode to auto")	
		}
		log.Printf("Set_Mode auto: %s   %v   errors: %d",string(mode_auto), time.Since(start), sb.Err_cnt)

		read_mode_auto, err := sb.Get_Mode()
		if err != nil {
			log.Println("Error setting mode to auto")	
		}
		log.Printf("Get_Mode auto: %s   %v   errors: %d",string(read_mode_auto), time.Since(start), sb.Err_cnt)
		


		set_auto_fan, err := sb.Set_Fan("auto")
		if err != nil {
			log.Println("Error setting fan to auto")	
		}
		log.Printf("Set_Fan: %s   %v   errors: %d",string(set_auto_fan), time.Since(start), sb.Err_cnt)

		read_auto_fan, err := sb.Get_Fan()
		if err != nil {
			log.Println("Error reading fan")	
		}
		log.Printf("Get_Fan auto: %s   %v   errors: %d",string(read_auto_fan), time.Since(start), sb.Err_cnt)	
		
		set_on_fan, err := sb.Set_Fan("on")
		if err != nil {
			log.Println("Error setting fan to on")	
		}
		log.Printf("Set_Fan: %s   %v   errors: %d",string(set_on_fan), time.Since(start), sb.Err_cnt)

		read_on_fan, err := sb.Get_Fan()
		if err != nil {
			log.Println("Error reading fan")	
		}
		log.Printf("Get_Fan on: %s   %v   errors: %d",string(read_on_fan), time.Since(start), sb.Err_cnt)	

		read_ambient, err := sb.Read_ambient_conditions()
		if err != nil {
			log.Println("Error reading ambient conditions")	
		}
		log.Printf("Read_ambient_conditions: %s   %v   errors: %d",string(read_ambient), time.Since(start), sb.Err_cnt)	

		set_target, err := sb.Set_target(70.5, 80.6)
		if err != nil {
			log.Println("Error setting target")	
		}
		log.Printf("Set_target heat target 70.5, cool target 80.6: %s   %v   errors: %d",string(set_target), time.Since(start), sb.Err_cnt)	

		read_target, err := sb.Get_target()
		if err != nil {
			log.Println("Error getting target")	
		}
		log.Printf("Get_target: %s   %v   errors: %d",string(read_target), time.Since(start), sb.Err_cnt)	

		set_time, err := sb.Set_time(16,35,"Monday")
		if err != nil {
			log.Println("Error setting time")	
		}
		log.Printf("Set_time: %s   %v   errors: %d",string(set_time), time.Since(start), sb.Err_cnt)	

		get_time, err := sb.Get_time()
		if err != nil {
			log.Println("Error getting time")	
		}
		log.Printf("Get_time: %s   %v   errors: %d",string(get_time), time.Since(start), sb.Err_cnt)
		
		get_psoc5_version, err := sb.Get_psoc5_version()
		if err != nil {
			log.Println("Error getting version")	
		}
		log.Printf("Get_psoc5_version: %s   %v   errors: %d",string(get_psoc5_version), time.Since(start), sb.Err_cnt)

		get_version, err := sb.Get_version()
		if err != nil {
			log.Println("Error getting version")	
		}
		log.Printf("Get_version: %s   %v   errors: %d",string(get_version), time.Since(start), sb.Err_cnt)	

		get_serial, err := sb.Get_serial()
		if err != nil {
			log.Println("Error getting serial")	
		}
		log.Printf("Get_serial: %s   %v   errors: %d",string(get_serial), time.Since(start), sb.Err_cnt)

		set_emer_off, err := sb.Set_heat_pump_emer("off")
		if err != nil {
			log.Println("Error setting heat pump emer off")	
		}
		log.Printf("Set_heat_pump_emer: %s   %v   errors: %d",string(set_emer_off), time.Since(start), sb.Err_cnt)

		get_emer_off, err := sb.Get_heat_pump_emer()
		if err != nil {
			log.Println("Error getting serial")	
		}
		log.Printf("Get_heat_pump_emer off: %s   %v   errors: %d",string(get_emer_off), time.Since(start), sb.Err_cnt)

		set_emer_on, err := sb.Set_heat_pump_emer("on")
		if err != nil {
			log.Println("Error setting heat pump emer on")	
		}
		log.Printf("Set_heat_pump_emer: %s   %v   errors: %d",string(set_emer_on), time.Since(start), sb.Err_cnt)

		get_emer_on, err := sb.Get_heat_pump_emer()
		if err != nil {
			log.Println("Error getting serial")	
		}
		log.Printf("Get_heat_pump_emer on: %s   %v   errors: %d",string(get_emer_on), time.Since(start), sb.Err_cnt)

		set_stage2stage_delay, err := sb.Set_stage2stage_delay(48)
		if err != nil {
			log.Println("Error setting stage to stage delay")	
		}
		log.Printf("Set_stage2stage_delay: %s   %v   errors: %d",string(set_stage2stage_delay), time.Since(start), sb.Err_cnt)

		get_stage2stage_delay, err := sb.Get_stage2stage_delay()
		if err != nil {
			log.Println("Error getting stage to stage delay")	
		}
		log.Printf("Get_stage2stage_delay: %s   %v   errors: %d",string(get_stage2stage_delay), time.Since(start), sb.Err_cnt)

		set_swing_temp, err := sb.Set_swing_temperature(2.6, "swing both ways")
		if err != nil {
			log.Println("Error setting swing temperature")	
		}
		log.Printf("Set_swing_temperature: %s   %v   errors: %d",string(set_swing_temp), time.Since(start), sb.Err_cnt)

		get_swing_temp, err := sb.Get_swing_temperature()
		if err != nil {
			log.Println("Error getting swing temperature")	
		}
		log.Printf("Get_swing_temperature: %s   %v   errors: %d",string(get_swing_temp), time.Since(start), sb.Err_cnt)

		set_temp_diff, err := sb.Set_temperature_differential(3.0, 5.0)
		if err != nil {
			log.Println("Error setting temperature differential")	
		}
		log.Printf("Set_temperature_differential: %s   %v   errors: %d",string(set_temp_diff), time.Since(start), sb.Err_cnt)

		get_temp_diff, err := sb.Get_temperature_differential()
		if err != nil {
			log.Println("Error getting temperature differential")	
		}
		log.Printf("Get_temperature_differential: %s   %v   errors: %d",string(get_temp_diff), time.Since(start), sb.Err_cnt)

		set_recvy_fast, err := sb.Set_recovery("fast")
		if err != nil {
			log.Println("Error setting recovery fast")	
		}
		log.Printf("Set_recovery fast: %s   %v   errors: %d",string(set_recvy_fast), time.Since(start), sb.Err_cnt)

		get_recvy_fast, err := sb.Get_recovery()
		if err != nil {
			log.Println("Error getting recovery fast")	
		}
		log.Printf("Get_recovery fast: %s   %v   errors: %d",string(get_recvy_fast), time.Since(start), sb.Err_cnt)

		set_recvy_economy, err := sb.Set_recovery("economy")
		if err != nil {
			log.Println("Error setting recovery economy")	
		}
		log.Printf("Set_recovery economy: %s   %v   errors: %d",string(set_recvy_economy), time.Since(start), sb.Err_cnt)

		get_recvy_economy, err := sb.Get_recovery()
		if err != nil {
			log.Println("Error getting recovery economy")	
		}
		log.Printf("Get_recovery economy: %s   %v   errors: %d",string(get_recvy_economy), time.Since(start), sb.Err_cnt)

		set_humidity_sp, err := sb.Set_humidity_setpoint(40)
		if err != nil {
			log.Println("Error setting humidity setpoint")	
		}
		log.Printf("Set_humidity_setpoint : %s   %v   errors: %d",string(set_humidity_sp), time.Since(start), sb.Err_cnt)

		get_humidity_sp, err := sb.Get_humidity_setpoint()
		if err != nil {
			log.Println("Error getting humidity setpoint")	
		}
		log.Printf("Get_humidity_setpoint: %s   %v   errors: %d",string(get_humidity_sp), time.Since(start), sb.Err_cnt)

	}
	
}
