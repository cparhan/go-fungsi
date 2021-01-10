// Package fungsi is a collection of helper functions.
//
// Copyright (c)2021, Chandra Parhan (github.com/cparhan), all rights reserved.
// License: MIT, for more details check the included LICENSE file.
package fungsi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mohae/deepcopy"
)

// ExpandToMap returns a map constructed from string keys separated by string sep.
// It accepts map init as initialization and will merge the map with the returned map.
// For example
//
//    keys := "api.server.host"
//    sep := "."
//    value := "localhost"
//
// will return a map
//
//    map[string]interface{}{
//        "api": map[string]interface{}{
//            "server": map[string]interface{}{
//                "host": "localhost",
//            },
//        },
//    }
//
// and with map init
//
//    init := map[string]interface{}{
//        "api": map[string]interface{}{
//            "server": map[string]interface{}{
//                "port": 8080,
//            },
//        },
//    }
//
// the returned map will be
//
//    map[string]interface{}{
//        "api": map[string]interface{}{
//            "server": map[string]interface{}{
//                "host": "localhost",
//                "port": 8080,
//            },
//        },
//    }
//
func ExpandToMap(keys, sep string, value interface{}, init map[string]interface{}) map[string]interface{} {
	obj := map[string]interface{}{}
	if init != nil {
		obj = deepcopy.Copy(init).(map[string]interface{})
	}
	tokens := strings.Split(keys, sep)
	key := tokens[0]
	rem := strings.Join(tokens[1:], sep)

	childInit := map[string]interface{}{}
	if v, ok := obj[key].(map[string]interface{}); ok {
		childInit = v
	}

	if rem != "" {
		obj[key] = ExpandToMap(rem, sep, value, childInit)
	} else if value != "" {
		obj[key] = value
	} else {
		obj[key] = map[string]interface{}{}
	}

	return obj
}

// FlattenMap will flatten map m by concatenating the keys
// of inner map.
// For example for map m
//
//    map[string]interface{}{
//        "api": map[string]interface{}{
//            "server": map[string]interface{}{
//                "host": "localhost",
//                "port": 8080,
//            },
//        },
//    }
//
// the returned map will be
//
//    map[string]interface{}{
//        "api.server.host": "localhost",
//        "api.server.port": 8080,
//    }
//
func FlattenMap(m interface{}) map[string]interface{} {
	if reflect.ValueOf(m).Kind() != reflect.Map {
		panic("flattenMap: m must be a map")
	}

	repeat := false
	obj := make(map[string]interface{})
	for k, v := range m.(map[string]interface{}) {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			iter := reflect.ValueOf(v).MapRange()
			for iter.Next() {
				kStr := fmt.Sprintf("%v.%v", k, iter.Key().Interface())
				if reflect.TypeOf(iter.Value().Interface()).Kind() == reflect.Map {
					repeat = true
				}
				obj[kStr] = iter.Value().Interface()
			}
		default:
			obj[k] = deepcopy.Copy(v)
		}
	}

	if repeat {
		return FlattenMap(obj)
	}
	return obj
}
