package confhandler

import (
	"os"
	"encoding/json"
)

func SetConf(confData interface{}, fileName string) {
	//Return configuration struct from json file
	file, _ := os.Open("src/app/conf/" + fileName)
	jsonDecoder := json.NewDecoder(file)

	//Unpack json to struct
	//configuration := ConfigurationApiAmazon{}
	jsonDecoder.Decode(confData)
}
