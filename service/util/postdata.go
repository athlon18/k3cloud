package util

import "time"

func CreateLoginPostData(parameters ...interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	var rid = make(map[string]interface{})
	data = map[string]interface{}{
		"format":     1,
		"useragent":  "ApiClient",
		"rid":        rid,
		"parameters": parameters,
		"timestamp":  time.Now().Format("2006-01-02"),
		"v":          "1.0",
	}
	return data
}

func CreateBusinessPostData(formId string, parameters interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	data = map[string]interface{}{
		"formid": formId,
		"data":   parameters,
	}
	return data
}
