package helpers

type Context interface {
	BindJSON(obj interface{}) error
	GetHeader(key string) string
	JSON(code int, obj interface{})
	Param(key string) string
	PostForm(key string) string
	Query(key string) string
}
