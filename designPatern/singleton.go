package designpatern

// Config interface [1]
type Config interface{
	PutString(key string, value string)
	GetString(key string)(string, bool)
}