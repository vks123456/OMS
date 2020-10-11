package cache

//
//import (
//	"time"
//
//	"github.com/go-redis/redis"
//	"github.com/prometheus/client_golang/prometheus"
//)
//
//// Repository represent the Caching Interface
//type Repository interface {
//	Set(key string, value interface{}, exp time.Duration) error
//	Get(key string) (string, error)
//}
//
//// repository represent the repository model
//type repository struct {
//	Client *redis.Client
//}
//
//// NewRedisRepository will create an object that represent the Repository interface
//func NewRedisRepository(Client *redis.Client) Repository {
//	return &repository{Client}
//}
//
//// Set attaches the redis repository and set the data
//func (r *repository) Set(key string, value interface{}, exp time.Duration) error {
//	timer := prometheus.NewTimer(metric.GetPrometheusInstance().MethodRequestLatency().WithLabelValues("RedisSet", "SetInCache"))
//	defer timer.ObserveDuration()
//	metric.GetPrometheusInstance().MethodRequestCounter().WithLabelValues("RedisSet", "SetInCache").Inc()
//
//	err := r.Client.Set(key, value, exp).Err()
//
//	if err != nil {
//		metric.GetPrometheusInstance().MethodExceptionCounter().WithLabelValues("RedisSet", "SetInCache", err.Error()).Inc()
//	}
//
//	return err
//}
//
//// Get attaches the redis repository and get the data
//func (r *repository) Get(key string) (string, error) {
//	timer := prometheus.NewTimer(metric.GetPrometheusInstance().MethodRequestLatency().WithLabelValues("RedisGet", "GetFromCache"))
//	defer timer.ObserveDuration()
//	metric.GetPrometheusInstance().MethodRequestCounter().WithLabelValues("RedisGet", "GetFromCache").Inc()
//
//	get := r.Client.Get(key)
//	s, err := get.Result()
//
//	if err != nil {
//		metric.GetPrometheusInstance().MethodExceptionCounter().WithLabelValues("RedisGet", "GetFromCache", err.Error()).Inc()
//	}
//
//	return s, err
//}
