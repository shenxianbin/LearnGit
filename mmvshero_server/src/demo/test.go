package main

import (
	"Gameserver/logic"
	"fmt"
	"reflect"
)

func CallMethod(object interface{}, methodName string, params ...interface{}) []interface{} {
	var method = reflect.ValueOf(object).MethodByName(methodName)
	if method.IsValid() {
		args := make([]reflect.Value, len(params))

		for k, v := range params {
			args[k] = reflect.ValueOf(v)
		}

		var ret = method.Call(args)
		var realRet = make([]interface{}, len(ret))
		if len(ret) > 0 {
			for k, v := range ret {
				realRet[k] = v.Interface()
			}
		}
		return realRet

	} else {
		panic(fmt.Sprintf("Not found '%s' method", methodName))
	}
}

func main() {
	var user logic.User
	for i := 0; i <= 50000; i++ {
		CallMethod(&user, "Getlv", 11, "moro")
		//user.Getlv(10, "xxxx")
	}
}
