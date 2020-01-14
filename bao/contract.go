package bao

type IContract interface {
	Invoke(accessToken, fn string, args []interface{}) (string, string, error)
	Query(accessToken, fn string, args []interface{}) (string, error)
}
