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

	/*
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


	// Main loop
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

		die_temp, err := sb.Get_psoc5_die_temp()
		if err != nil {
			log.Println("Error getting die temperature")	
		}
		log.Printf("Get_psoc5_die_temp: %s   %v   errors: %d",string(die_temp), time.Since(start), sb.Err_cnt)	

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

		set_dehumidifier_sp, err := sb.Set_dehumidifier_setpoint(50)
		if err != nil {
			log.Println("Error setting dehumidifier setpoint")	
		}
		log.Printf("Set_dehumidifier_setpoint : %s   %v   errors: %d",string(set_dehumidifier_sp), time.Since(start), sb.Err_cnt)

		get_dehumidifier_sp, err := sb.Get_dehumidifier_setpoint()
		if err != nil {
			log.Println("Error getting dehumidifier setpoint")	
		}
		log.Printf("Get_dehumidifier_setpoint: %s   %v   errors: %d",string(get_dehumidifier_sp), time.Since(start), sb.Err_cnt)

		set_ext_dehumidifier_sp, err := sb.Set_external_dehumidifier_setpoint(35)
		if err != nil {
			log.Println("Error setting external dehumidifier setpoint")	
		}
		log.Printf("Set_external_dehumidifier_setpoint : %s   %v   errors: %d",string(set_ext_dehumidifier_sp), time.Since(start), sb.Err_cnt)

		get_ext_dehumidifier_sp, err := sb.Get_external_dehumidifier_setpoint()
		if err != nil {
			log.Println("Error getting external dehumidifier setpoint")	
		}
		log.Printf("Get_external_dehumidifier_setpoint : %s   %v   errors: %d",string(get_ext_dehumidifier_sp), time.Since(start), sb.Err_cnt)

		set_humidifier_mode_off, err := sb.Set_humidifier_mode("off")
		if err != nil {
			log.Println("Error setting humidifier mode off")	
		}
		log.Printf("Set_humidifier_mode off : %s   %v   errors: %d",string(set_humidifier_mode_off), time.Since(start), sb.Err_cnt)

		get_humidifier_mode_off, err := sb.Get_humidifier_mode()
		if err != nil {
			log.Println("Error getting humidifier mode off")	
		}
		log.Printf("Get_humidifier_mode off : %s   %v   errors: %d",string(get_humidifier_mode_off), time.Since(start), sb.Err_cnt)

		set_humidifier_mode_runonly, err := sb.Set_humidifier_mode("run only with heat")
		if err != nil {
			log.Println("Error setting humidifier mode runonly")	
		}
		log.Printf("Set_humidifier_mode runonly : %s   %v   errors: %d",string(set_humidifier_mode_runonly), time.Since(start), sb.Err_cnt)

		get_humidifier_mode_runonly, err := sb.Get_humidifier_mode()
		if err != nil {
			log.Println("Error getting humidifier mode runonly")	
		}
		log.Printf("Get_humidifier_mode runonly : %s   %v   errors: %d",string(get_humidifier_mode_runonly), time.Since(start), sb.Err_cnt)

		set_humidifier_mode_anytime, err := sb.Set_humidifier_mode("humidify anytime")
		if err != nil {
			log.Println("Error setting humidifier mode anytime")	
		}
		log.Printf("Set_humidifier_mode anytime : %s   %v   errors: %d",string(set_humidifier_mode_anytime), time.Since(start), sb.Err_cnt)

		get_humidifier_mode_anytime, err := sb.Get_humidifier_mode()
		if err != nil {
			log.Println("Error getting humidifier mode anytime")	
		}
		log.Printf("Get_humidifier_mode anytime : %s   %v   errors: %d",string(get_humidifier_mode_anytime), time.Since(start), sb.Err_cnt)

		set_dehumidifier_mode_off, err := sb.Set_dehumidifier_mode("off")
		if err != nil {
			log.Println("Error setting dehumidifier mode off")	
		}
		log.Printf("Set_dehumidifier_mode off : %s   %v   errors: %d",string(set_dehumidifier_mode_off), time.Since(start), sb.Err_cnt)

		get_dehumidifier_mode_off, err := sb.Get_dehumidifier_mode()
		if err != nil {
			log.Println("Error getting dehumidifier mode off")	
		}
		log.Printf("Get_dehumidifier_mode off: %s   %v   errors: %d",string(get_dehumidifier_mode_off), time.Since(start), sb.Err_cnt)

		set_dehumidifier_mode_on, err := sb.Set_dehumidifier_mode("on")
		if err != nil {
			log.Println("Error setting dehumidifier mode on")	
		}
		log.Printf("Set_dehumidifier_mode on : %s   %v   errors: %d",string(set_dehumidifier_mode_on), time.Since(start), sb.Err_cnt)

		get_dehumidifier_mode_on, err := sb.Get_dehumidifier_mode()
		if err != nil {
			log.Println("Error getting dehumidifier mode on")	
		}
		log.Printf("Get_dehumidifier_mode on: %s   %v   errors: %d",string(get_dehumidifier_mode_on), time.Since(start), sb.Err_cnt)

		set_ext_dehumidifier_mode_off, err := sb.Set_external_dehumidifier_mode("off")
		if err != nil {
			log.Println("Error setting external dehumidifier mode off")	
		}
		log.Printf("Set_external_dehumidifier_mode off : %s   %v   errors: %d",string(set_ext_dehumidifier_mode_off), time.Since(start), sb.Err_cnt)

		get_ext_dehumidifier_mode_off, err := sb.Get_external_dehumidifier_mode()
		if err != nil {
			log.Println("Error getting external dehumidifier mode off")	
		}
		log.Printf("Get_external_dehumidifier_mode off : %s   %v   errors: %d",string(get_ext_dehumidifier_mode_off), time.Since(start), sb.Err_cnt)

		set_ext_dehumidifier_mode_wfan, err := sb.Set_external_dehumidifier_mode("on with fan")
		if err != nil {
			log.Println("Error setting external dehumidifier mode wfan")	
		}
		log.Printf("Set_external_dehumidifier_mode wfan : %s   %v   errors: %d",string(set_ext_dehumidifier_mode_wfan), time.Since(start), sb.Err_cnt)

		get_ext_dehumidifier_mode_wfan, err := sb.Get_external_dehumidifier_mode()
		if err != nil {
			log.Println("Error getting external dehumidifier mode wfan")	
		}
		log.Printf("Get_external_dehumidifier_mode wfan : %s   %v   errors: %d",string(get_ext_dehumidifier_mode_wfan), time.Since(start), sb.Err_cnt)

		set_ext_dehumidifier_mode_wofan, err := sb.Set_external_dehumidifier_mode("on without fan")
		if err != nil {
			log.Println("Error setting external dehumidifier mode wofan")	
		}
		log.Printf("Set_external_dehumidifier_mode wofan : %s   %v   errors: %d",string(set_ext_dehumidifier_mode_wofan), time.Since(start), sb.Err_cnt)

		get_ext_dehumidifier_mode_wofan, err := sb.Get_external_dehumidifier_mode()
		if err != nil {
			log.Println("Error getting external dehumidifier mode wofan")	
		}
		log.Printf("Get_external_dehumidifier_mode wofan : %s   %v   errors: %d",string(get_ext_dehumidifier_mode_wofan), time.Since(start), sb.Err_cnt)



		// Relays control
		log.Println("Setting user mode off")
		sb.Set_Mode("off")
		log.Println("Setting fan mode to auto")
		sb.Set_Fan("auto")
		log.Println("Set Humidifier Mode off")
		sb.Set_humidifier_mode("off")
		log.Println("Set De-Humidifier Mode off")
		sb.Set_dehumidifier_mode("off")
		log.Println("Set External De-Humidifier mode off")
		sb.Set_external_dehumidifier_mode("off")
		log.Println("Setting user mode to manual")
		sb.Set_Mode("manual")
		log.Println("Message Type 21 : W, W2, W3, Y, Y2, G, OB, H, DH, EX ...All HIGH")

		relays := sb.SSRelays { 
			W :1,  
			W2 :1, 
			W3 :1, 
			Y :1,  
			Y2 :1,
			G :1,  
			OB :1,
			H :1, 
			DH :1, 
			EX :1,
		}
		set_relays_one, err := sb.Set_control_relay(relays)
		if err != nil {
			log.Println("Error setting control relays to 1")	
		}
		log.Printf("Set_control_relay to 1: %s   %v   errors: %d",string(set_relays_one), time.Since(start), sb.Err_cnt)

		relays_zero := sb.SSRelays { 
			W :0,  
			W2 :0, 
			W3 :0, 
			Y :0,  
			Y2 :0,
			G :0,  
			OB :0,
			H :0, 
			DH :0, 
			EX :0,
		}
		set_relays_zero, err := sb.Set_control_relay(relays_zero)
		if err != nil {
			log.Println("Error setting control relays to 0")	
		}
		log.Printf("Set_control_relay to 0: %s   %v   errors: %d",string(set_relays_zero), time.Since(start), sb.Err_cnt)


		sb.Set_Mode("off")
		sb.Set_Fan("auto")

		log.Printf("Close EX relay")
		sb.Set_ex_relay("close")
		log.Printf("Open EX relay")
		sb.Set_ex_relay("open")
	}
	*/

	sb.PIC18_upload_firmware("ota_test_files/pic18/ota_v0.21.bin")

	pic18_version, err := sb.Get_version()
	if err != nil {
		log.Println("Error getting pic18_version")	
	}
	log.Printf("pic18_version: %s",string(pic18_version))	


	
}
