package common

import "encoding/json"

func SwapTo(req, product interface{}) error {
	dataByte, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, product)
}
