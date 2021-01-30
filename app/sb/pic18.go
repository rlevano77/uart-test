package sb

import (
	"os"
	"io/ioutil"
	"time"
	"encoding/json"
	b64 "encoding/base64"
	"log"
)

/*
	Message 100
	PIC18 Bootloader - Erase specified page in flash / Set page address point
	adddress range 0 ~ 399
*/
func PIC18boot_erase_set_addr(mode string, addr int32) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_erase_set_page(mode, addr)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 101
	PIC18 Bootloader
	Set CRC code of full OTA image
*/
func PIC18boot_set_crc_full_ota(crc_H byte, crc_L byte) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_set_crc_full_ota(crc_H, crc_L)
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 101
	PIC18 Bootloader
	Get CRC code of full OTA image
*/
func PIC18boot_get_crc_full_ota() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_get_crc_full_ota()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 102
	PIC18 Bootloader
	GET CRC value of specified page in flash
*/
func PIC18boot_get_crc_page() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_get_crc_page()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Message 103
	PIC18 Bootloader
	Require HVAC board check OTA image in flash
*/
func PIC18boot_check_ota_image() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_check_ota_image()
	}
	sb_mutex.Unlock()
	return response, err
}

/*
	Message 104
	PIC18 Bootloader
	Require HVAC board upgrade firmware from OTA image in flash
*/
func PIC18boot_upgrade_fw_from_ota() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_upgrade_fw_from_ota()
	}
	sb_mutex.Unlock()
	return response, err
}

/*
	Message 105
	PIC18 Bootloader
	GET result of HVAC board firmware upgrading
*/
func PIC18boot_get_fw_upgrade_result() ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_get_fw_upgrade_result()
	}
	sb_mutex.Unlock()
	return response, err	
}

/*
	Calculates the 16-bit CCITT CRC checksum of input byte adding it to the
	ongoing global checksum stored in CRCHreg and CRCLreg.
*/
func CRC(arr []byte) S_CRC {
	var CRC_H byte= 0
	var CRC_L byte = 0

	for _, x := range arr {
		x = CRC_H ^ x
		CRC_H = CRC_L
		CRC_L = x
		x = ((x >> 4) & 0x0F)
		CRC_L = x ^ CRC_L
		x = ((CRC_L << 4) & 0xF0)
		CRC_H = x ^ CRC_H
		x = ((CRC_L >> 4) & 0x0F) | ((CRC_L << 4) & 0xF0)

		if x > 127 {
			x = x + x + 1
		} else {
			x = x + x
		}

		CRC_H = x ^ CRC_H
		x = x & 0xE0
		CRC_H = x ^ CRC_H
		CRC_L = x ^ CRC_L	
	}

	return S_CRC{
			CRCH: CRC_H,
			CRCL: CRC_L,
		}
}


func Divmod(x int32, y int32) S_QMod {
	return S_QMod{
		Q: int32(x/y),
		Mod: int32(x%y),
	}
}

/*
	Calculates an Array of pages to be send in the PIC18 Bootload process
*/
func PIC18_getpages(bin_file_path string) ([]S_Page, error) {
	block_size := 256

	// File stats
	info, err := os.Stat(bin_file_path)
	if err != nil {
		log.Printf("Error getting file stat: %v \n", err)
		return []S_Page{}, err
	}
	file_size := info.Size()
	log.Printf("file size = %d", file_size)

	s := Divmod(int32(file_size), int32(block_size))
	q := s.Q
	mod := s.Mod
	log.Printf("Number of 256 bytes blocks q = %d", q)
	log.Printf("Remainder number of blocks mod = %d", mod)

	// "arr" array will storage all bytes read from file (This assumes remaider = 0)
	arr, err := ioutil.ReadFile(bin_file_path)
	if err != nil {
		log.Printf("Error reading file: %v \n", err)
		return []S_Page{}, err
	}
	log.Printf("len(arr) = %d", len(arr))

	// "block_list" will be a slice of []byte. Each []byte will have 256 elements
	block_list := make([][]byte, 0)
	for i := 0; i < int(q); i++ {
		start := block_size*i
		stop := block_size*(i+1)
		arr_slice := arr[start : stop]
		block_list = append(block_list, arr_slice)   // slicing 256 elements and add them to the blocks_list
		/*
		log.Printf("index = %d", i)
		log.Printf("start = %d", start)
		log.Printf("stop = %d", stop)
		log.Printf("len(arr_slice) = %d", len(arr_slice))
		*/
	}

	log.Printf("len(block_list) = %d", len(block_list))

	// "b64_arr" will hold the []byte converted to base64 strings
	b64_arr := make([]string, 0)           
	for i := 0; i < int(q); i++ {
		base64_message := b64.StdEncoding.EncodeToString(block_list[i])
		b64_arr = append(b64_arr, base64_message)
		/*
		log.Printf("len(base64_message) = %d", len(base64_message))
		log.Printf("base64_message = %s", base64_message)
		*/
	}

	// "pages" will hold the S_Page objects generated using b64_arr
	pages := make([]S_Page, 0)
	for idx := 0; idx < int(q); idx++ {
		page := S_Page{
			Idx: int32(idx),
			Data: b64_arr[idx],
		}
		pages = append(pages, page)
		/*
		log.Printf("idx : %d", page.Idx)
		log.Printf("data : %s", page.Data)
		*/
	}           

	return pages, err
}

/*
	Erase and Write a Page in one call
*/
func PIC18boot_erase_and_write_flash(page S_Page) ([]byte,error){
	var err error = nil
	var response []byte
	sb_mutex.Lock()
	for ((err != nil) || (len(response) == 0) || (validJSON(response) == false)) {
		response, err = hvac_erase_and_write_flash(page)
	}
	sb_mutex.Unlock()
	return response, err	
}

/* 
	Upload Firmware (.bin file) to the HVAC Board PIC18
*/
func PIC18_upload_firmware(bin_file_path string) {

	start := time.Now() // To measure execution time

	// Convert the .bin file to Bootload page objects
	pages, err := PIC18_getpages(bin_file_path)
	if err != nil {
    log.Printf("Error PIC18_getpages: %v \n", err)
	}
	log.Printf("len(pages) : %d", len(pages))

	// Send page by page data and compare page's CRC
	for i := 0; i < len(pages); i++ {
		erase_write, err := PIC18boot_erase_and_write_flash(pages[i])
		if err != nil {
			log.Println("Error erase_write")	
		}
		log.Printf("Sending page %d : %s", i, string(erase_write))
	}

	// Calculate CRC of entire OTA image
	log.Printf("Calculate CRC of entire OTA image")
	bytes_read, err := ioutil.ReadFile(bin_file_path)
	if err != nil {
		log.Printf("Error reading file: %v \n", err)
	}
	oat_crc := CRC(bytes_read)
	log.Printf("CRC_H : %d",oat_crc.CRCH)
	log.Printf("CRC_L : %d",oat_crc.CRCL)

	log.Printf("Set CRC of entire OTA image")
	set_crc_full_ota, err := PIC18boot_set_crc_full_ota(oat_crc.CRCH, oat_crc.CRCL)
	if err != nil {
		log.Printf("Error PIC18boot_set_crc_full_ota: %v \n", err)
	}
	log.Printf("PIC18boot_set_crc_full_ota: %s",string(set_crc_full_ota))

	log.Printf("Reading back CRC of entire OTA image")
	crc_page_read, err := PIC18boot_get_crc_full_ota()
	if err != nil {
		log.Println("Error CRC of entire OTA image")	
	}
	log.Printf("CRC of entire OTA image : %s", string(crc_page_read))

	log.Println("Require HVAC board check OTA image in flash")
	check_ota_image, err := PIC18boot_check_ota_image()
	if err != nil {
		log.Println("Error PIC18boot_check_ota_image()")	
	}
	log.Printf("PIC18boot_check_ota_image : %s", string(check_ota_image))

	//log.Println("Delay 15 seconds")
	//time.Sleep(15 * time.Second)

	success := 0
	var s S_OTAStatus
	// Wait for PIC18 MCU to upgrade application image
	for success != 136 {
		status, err := PIC18boot_upgrade_fw_from_ota()
		if err != nil {
			log.Println("Error PIC18boot_upgrade_fw_from_ota()")	
		}
		log.Printf("status : %s", string(status))
		json.Unmarshal(status, &s)
		success = s.OtaImageStatus
	}

	log.Println("Delay 10 seconds")
	time.Sleep(10 * time.Second)

	log.Println(time.Since(start))
}
