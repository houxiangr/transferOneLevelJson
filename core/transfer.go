package core

import (
	"encoding/json"
	"reflect"
)

func TransferToOneLevel(source string) (string, error) {
	objMap := make(map[string]interface{})
	res := make(map[string]interface{})
	var err error
	err = json.Unmarshal([]byte(source), &objMap)
	if err != nil {
		return "", err
	}
	err = dealObjMap("", objMap, &res)
	if err != nil {
		return "", err
	}
	resbyte,err := json.Marshal(res)
	if err != nil {
		return "",err
	}

	return string(resbyte), nil
}

func dealObjMap(baseKey string, objMap map[string]interface{}, res *map[string]interface{}) error {
	var err error
	baseKeyBytes := []byte(baseKey)

	for k, v := range objMap {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			tempBaseKeyBytes := append(baseKeyBytes,[]byte(k)...)
			tempBaseKeyBytes = append(tempBaseKeyBytes,[]byte(".")...)
			err = dealObjMap(string(tempBaseKeyBytes),v.(map[string]interface{}),res)
			if err != nil {
				return err
			}
		default:
			tempBaseKeyBytes := append(baseKeyBytes,[]byte(k)...)
			(*res)[string(tempBaseKeyBytes)] = v
		}
	}

	return nil
}
