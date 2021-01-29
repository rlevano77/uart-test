package sb


/********************************
	PSOC5 Cypress Bootloader Load
********************************/
func Cyboot_load() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = psoc5_cyboot_load()
	}
	sb_mutex.Unlock()
	return response, err	
}