package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func GetUpdatedBody(body interface{}, replacements map[string]string, removes []string, path string) interface{} {
	if len(replacements) == 0 && len(removes) == 0 {
		return body
	}
	switch bodyValue := body.(type) {
	case map[string]interface{}:
		res := make(map[string]interface{})
		for key, value := range bodyValue {
			if temp := GetUpdatedBody(value, replacements, removes, path+"."+key); temp != nil {
				if replaceKey := replacements[fmt.Sprintf("key:%s.%s", path, key)]; replaceKey != "" {
					key = replaceKey
				}
				res[key] = temp
			}
		}
		return res
	case []interface{}:
		res := make([]interface{}, 0)
		for index, value := range bodyValue {
			if temp := GetUpdatedBody(value, replacements, removes, path+"."+strconv.Itoa(index)); temp != nil {
				res = append(res, temp)
			}
		}
		return res
	case string:
		for key, replacement := range replacements {
			if key == path {
				return replacement
			}
		}
		for _, remove := range removes {
			if path == remove {
				return nil
			}
		}
	default:

	}
	return body
}

// GetIdFromResponseExample extracts id from response
func GetIdFromResponseExample(response interface{}) string {
	if response != nil {
		if responseMap, ok := response.(map[string]interface{}); ok && responseMap["body"] != nil {
			if bodyMap, ok := responseMap["body"].(map[string]interface{}); ok && bodyMap["id"] != nil {
				if id, ok := bodyMap["id"].(string); ok {
					return id
				}
			}
		}
	}
	return ""
}

// GetParentIdFromId returns the id of its parent resource from id
func GetParentIdFromId(id string) string {
	idURL, err := url.ParseRequestURI(id)
	if err != nil {
		return ""
	}
	path := idURL.Path

	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")

	components := strings.Split(path, "/")
	parentId := ""
	length := len(components) - 2
	if length-2 >= 0 && components[length-2] == "providers" {
		length -= 2
	}
	for current := 0; current <= length-2; current += 2 {
		key := components[current]
		value := components[current+1]
		parentId += "/" + key + "/" + value
	}

	return parentId
}
