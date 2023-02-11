package helpers

type Context interface {
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	Param(key string) string
}
