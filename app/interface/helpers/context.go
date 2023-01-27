package helpers

type Context interface {
	JSON(code int, obj interface{})
}
