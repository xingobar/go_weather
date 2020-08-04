package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", "Taipei", os.Getenv("KEY"))

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("get weather data error")
	}

	// 轉換成字串
	//fmt.Println(string(body))

	var data map[string] interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error: " , err)
	}

	for k, v := range data {
		var tmp map[string] interface{}
		if k == "main" {
			// 只印出主要的資訊
			fmt.Println("======= main ========")
			// 將字串轉換成 json object
			// 回傳值回 []bytes
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("convert error")
				break
			}
			// 將 json object 轉換成 map 資料型態
			err = json.Unmarshal(b, &tmp)
			if err != nil {
				fmt.Println("json convert error")
				break
			}

			for tmpk, tmpv := range tmp {
				fmt.Println(tmpk, " : ", tmpv)
			}
		}
	}
}