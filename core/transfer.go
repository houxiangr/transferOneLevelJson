package core

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func TransferToOneLevel(source string) (string, error) {
	var objMap interface{}
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
	resbyte, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(resbyte), nil
}

func dealObjMap(baseKey string, obj interface{}, res *map[string]interface{}) error {
	var err error
	baseKeyBytes := []byte(baseKey)

	switch reflect.TypeOf(obj).Kind() {
	case reflect.Map:
		for k,v := range obj.(map[string]interface{}) {
			tempBaseKeyBytes := []byte{}
			if baseKey != "" {
				tempBaseKeyBytes = append(baseKeyBytes, []byte(".")...)
			}
			tempBaseKeyBytes = append(tempBaseKeyBytes, []byte(k)...)
			err = dealObjMap(string(tempBaseKeyBytes), v, res)
			if err != nil {
				return err
			}
		}
	case reflect.Slice:
		for k, v := range obj.([]interface{}) {
			tempBaseKeyBytes := append(baseKeyBytes, []byte("[")...)
			tempBaseKeyBytes = append(tempBaseKeyBytes, []byte(strconv.Itoa(k))...)
			tempBaseKeyBytes = append(tempBaseKeyBytes, []byte("]")...)
			err = dealObjMap(string(tempBaseKeyBytes), v, res)
			if err != nil {
				return err
			}
		}
	default:
		(*res)[baseKey] = obj
	}

	return nil
}
