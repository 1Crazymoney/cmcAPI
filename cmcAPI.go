package cmcAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	baseURL = "https://api.coinmarketcap.com/v1"
	url     string
	l       string
)

//GetMarketData - Get information about the global market data of the cryptocurrencies
func GetMarketData() (GlobalMarketData, error) {
	url = fmt.Sprintf(baseURL + "/global/")

	resp, err := makeReq(url)

	var data GlobalMarketData
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

//GetCoinInfo - Get information about single crypto currency
func GetCoinInfo(coin string) (Coin, error) {
	url = fmt.Sprintf("%s/ticker/%s", baseURL, coin)
	resp, err := makeReq(url)

	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data[0], nil
}

//GetAllCoinInfo - Get information about all coins listed in Coin Market Cap. If you want to limit the search to top 10 coins pass 10 as int, if you want all - pass 0 == No Limit
func GetAllCoinInfo(limit int) ([]Coin, error) {
	if limit > 0 {
		l = fmt.Sprintf("?limit=%v", limit)
	}
	url = fmt.Sprintf("%s/ticker/%s", baseURL, l)

	resp, err := makeReq(url)

	var data []Coin
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

//Client
func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := doReq(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp, err
}
