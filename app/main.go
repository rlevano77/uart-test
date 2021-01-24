package main

import (
	"github.com/rlevano77/uart/sb"
	"log"
)

func main() {
	log.Println("GOTEST APP")

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
			log.Printf(sensors)	
		}
	}()

	// Polling iaq
	go func() {
		for {
			iaq, err := sb.Read_IAQ()
			if err != nil {
				log.Println("Error reading iaq")	
			}
			log.Printf(iaq)	
		}
	}()

	// Polling status
	go func() {
		for {
			status, err := sb.Get_Status()
			if err != nil {
				log.Println("Error reading status")	
			}
			log.Printf(status)	
		}
	}()

	// Polling proximity
	go func() {
		for {
			proximity, err := sb.Read_Proximity()
		if err != nil {
			log.Println("Error reading proximity")	
		}
		log.Printf(proximity)
		}
	}()

	// main loop
	for {
	
		setmode, err := sb.Set_Mode("off")
		if err != nil {
			log.Println("Error setting mode")	
		}
		log.Printf(setmode)

		getmode, err := sb.Get_Mode()
		if err != nil {
			log.Println("Error getting mode")	
		}
		log.Printf(getmode)

		config := sb.SSetConfigSettings {
			HeatAuxType: "gas/oil",
			HeatStage: "1 stage AUX heat",
			CoolStage: "1-stage",
			HeatPumpConfig: "B",
		}
		setcfg, err := sb.Set_config_settings(config)
		if err != nil {
			log.Println("Error setting config settings")	
		}
		log.Printf(setcfg)
		

		getcfg, err := sb.Get_config_settings()
		if err != nil {
			log.Println("Error getting config settings")	
		}
		log.Printf(getcfg)

		auto, err := sb.Set_Mode("auto")
		if err != nil {
			log.Println("Error setting mode to auto")	
		}
		log.Printf(auto)

	}
	
}
