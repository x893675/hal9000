package util

import "os"

func GetEnvironment(key, value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return value
	}
	return val
}

func GetEnvironmentToInt(key, value string) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return S(value).DefaultInt(0)
	}
	return S(val).DefaultInt(0)
}

func GetEnvironmentToBool(key, value string) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return S(value).DefaultBool(false)
	}
	return S(val).DefaultBool(false)
}
