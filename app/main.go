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

	}
	
}
