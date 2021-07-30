package database

type RedisHandler interface {
	Set(string, string) (string, error)
	Get(string) (string, error)
}
