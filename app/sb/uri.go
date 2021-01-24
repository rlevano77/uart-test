package sb

var (
// Methods
  GET = "GET"
  POST = "POST"

  // Sensor Board Endpoints
  SB_SENSORS_URI = "/sb/sensors"
  PROXIMITY_URI = "/sb/proximity"
  IAQ_URI = "/sb/iaq"
  OFFSET_URI = "/sb/offset"
  RESET_URI = "/sb/reset"
  PSOC5_FW_URI = "/sb/psoc5_fw"
  PSOC5_DIE_URI = "/sb/psoc5_die"
  SB_BEEP_URI = "/sb/buzzer/beep"

  // Cypress Bootloader Endpoints
  CYBOOT_LOAD_URI = "/sb/cyboot_load"

  // PIC18 Bootloader Endpoints
  PICBOOT_ERASE_SET_PAGE = "/hvac/erase_page"
  PICBOOT_SET_CRC_FULL_OTA = "/hvac/crc_full_ota"
  PICBOOT_GET_CRC_PAGE = "/hvac/crc_page"
  PICBOOT_CHECK_OTA_IMAGE = "/hvac/check_ota_image"
  PICBOOT_UPGRADE_FW_FROM_OTA = "/hvac/upgrade_fw_ota"
  PICBOOT_GET_FW_UPGRADE_RESULT = "/hvac/upgrade_result"
  PICBOOT_WRITE_FLASH = "/hvac/write_flash"
  PICBOOT_ERASE_WRITE_PAGE = "/hvac/erase_write_page"

  // HVAC Endpoints
  TSTAT_STATUS_URI = "/tstat/status"
  TSTAT_CONFIG_SETTINGS_URI	= "/tstat/config" 
  TSTAT_MODE_URI = "/tstat/mode"
  TSTAT_FAN_URI = "/tstat/fan"
  TSTAT_REPORT_AMBIENT_URI = "/tstat/report_ambient" 
  TSTAT_TARGET  = "/tstat/target"
  TSTAT_TIME_WEEKDAY = "/tstat/time"
  TSTAT_FW_URI = "/tstat/fw"
  TSTAT_SERIAL_URI = "/tstat/serial"
  TSTAT_HEAT_PUMP_EMER = "/tstat/heatpump_emer"
  TSTAT_STAGE_TO_STAGE_DELAY = "/tstat/stage_to_stage_delay"
  TSTAT_SWING_TEMP = "/tstat/swing_temp"
  TSTAT_TEMP_DIFFERENTIAL = "/tstat/temp_diff"
  TSTAT_RECOVERY_MODE = "/tstat/recovery"
  TSTAT_HUMIDITY_SETPOINT = "/tstat/humidity_setpoint"
  TSTAT_DEHUMIDIFIER_SETPOINT = "/tstat/dehumidifier_setpoint"
  TSTAT_EXT_DEHUMIDIFIER_SETPOINT = "/tstat/ext_dehumidifier_setpoint"
  TSTAT_HUMIDIFIER_MODE = "/tstat/humidifier_mode"
  TSTAT_DEHUMIDIFIER_MODE = "/tstat/dehumidifier_mode"
  TSTAT_EXT_DEHUMIDIFIER_MODE = "/tstat/ext_dehumidifier_mode"
  TSTAT_CONTROL_RELAY = "/tstat/control_relay"
  TSTAT_EX_RELAY = "/tstat/ex_relay"

  // TEST Mode Endpoints
  TSTAT_ENABLE_API_TEST_MODE = "/tstat/api_test/enable"
  TSTAT_DISABLE_API_TEST_MODE = "/tstat/api_test/disable"

  // API options
  mode_options = []string{
    "off",
    "heat",
    "cool",
    "auto",
    "manual",
  }

  fan_options = []string{
    "auto",
    "on",
  }
  heat_aux_type_options = []string{
    "gas/oil",
    "electric",
  }
  heat_stage_options = []string{
    "off",
    "1 stage AUX heat",
    "2 stage AUX heat",
    "3 stage AUX heat",
    "1 stage heat pump",
    "1 stage heat pump + 1 AUX",
    "1 stage heat pump + 2 AUX",
    "1 stage heat pump + 3 AUX",
    "2 stage heat pump",
    "2 stage heat pump + 1 AUX",
    "2 stage heat pump + 2 AUX",
    "2 stage heat pump + 3 AUX",
  }
  cool_stage_options = []string{
    "off",
    "1-stage",
    "2-stage",
  }		
  humidity_control_options = []string{
    "no humidity control",
    "humidifier mode",
    "dehumidifier mode",
    "external dehumidifier mode",
  }	

  // O-active for Cool, B-active for Heat
  heat_pump_config_options = []string{
    "O",
    "B",
  }
  heat_pump_emer_options = []string{
    "off",
    "on",
  }
  swing_options = []string{
    "swing one way",
    "swing both ways",
  }
  recovery_options = []string{
    "fast",
    "economy",
  }
  humidifier_mode_options = []string{
    "off",
    "run only with heat",
    "humidify anytime",
  }
  dehumidifier_mode_options = []string{
    "off",
    "on",
  }
  ext_dehumidifier_mode_options = []string{
    "off",
    "on with fan",
    "on without fan",
  }
  ex_relay_options = []string{
    "open",
    "close",
  }

  // PIC18 Bootloader options
  page_mode_options = []string{
    "erase",
    "set",
  }

)