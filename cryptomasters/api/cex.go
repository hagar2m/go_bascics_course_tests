package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"test/go/ccryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) < 3 {
		return nil, fmt.Errorf("3 characters required")
	}
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
	if err != nil {
		return nil, err
	}
	
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var response CEXResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

		//string(bodyBytes)
		// fmt.Println(json)
		rate := datatypes.Rate{
			Currency: currency,
			Price: response.Bid,
		}
		return &rate, nil
 	}
	return nil, fmt.Errorf("status code received: %v", res.StatusCode)
}

