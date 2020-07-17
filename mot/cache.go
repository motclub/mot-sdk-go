package mot

import (
	"time"
)

func Has(path string) bool {
	return Client.cache.Has(path)
}

func HasGet(path string, dst interface{}) bool {
	return Client.cache.HasGet(path, dst)
}

func HasGetInt(path string) (int, bool) {
	return Client.cache.HasGetInt(path)
}

func HasGetInt8(path string) (int8, bool) {
	return Client.cache.HasGetInt8(path)
}

func HasGetInt16(path string) (int16, bool) {
	return Client.cache.HasGetInt16(path)
}

func HasGetInt32(path string) (int32, bool) {
	return Client.cache.HasGetInt32(path)
}

func HasGetInt64(path string) (int64, bool) {
	return Client.cache.HasGetInt64(path)
}

func HasGetUint(path string) (uint, bool) {
	return Client.cache.HasGetUint(path)
}

func HasGetUint8(path string) (uint8, bool) {
	return Client.cache.HasGetUint8(path)
}

func HasGetUint16(path string) (uint16, bool) {
	return Client.cache.HasGetUint16(path)
}

func HasGetUint32(path string) (uint32, bool) {
	return Client.cache.HasGetUint32(path)
}

func HasGetUint64(path string) (uint64, bool) {
	return Client.cache.HasGetUint64(path)
}

func HasGetFloat(path string) (float64, bool) {
	return Client.cache.HasGetFloat(path)
}

func HasGetFloat32(path string) (float32, bool) {
	return Client.cache.HasGetFloat32(path)
}

func HasGetFloat64(path string) (float64, bool) {
	return Client.cache.HasGetFloat64(path)
}

func HasGetString(path string) (string, bool) {
	return Client.cache.HasGetString(path)
}

func HasGetBool(path string) (bool, bool) {
	return Client.cache.HasGetBool(path)
}

func HasGetTime(path string) (time.Time, bool) {
	return Client.cache.HasGetTime(path)
}

func Get(path string, dst interface{}) {
	Client.cache.Get(path, dst)
}

func GetInt(path string) int {
	return Client.cache.GetInt(path)
}

func GetInt8(path string) int8 {
	return Client.cache.GetInt8(path)
}

func GetInt16(path string) int16 {
	return Client.cache.GetInt16(path)
}

func GetInt32(path string) int32 {
	return Client.cache.GetInt32(path)
}

func GetInt64(path string) int64 {
	return Client.cache.GetInt64(path)
}

func GetUint(path string) uint {
	return Client.cache.GetUint(path)
}

func GetUint8(path string) uint8 {
	return Client.cache.GetUint8(path)
}

func GetUint16(path string) uint16 {
	return Client.cache.GetUint16(path)
}

func GetUint32(path string) uint32 {
	return Client.cache.GetUint32(path)
}

func GetUint64(path string) uint64 {
	return Client.cache.GetUint64(path)
}

func GetFloat(path string) float64 {
	return Client.cache.GetFloat(path)
}

func GetFloat32(path string) float32 {
	return Client.cache.GetFloat32(path)
}

func GetFloat64(path string) float64 {
	return Client.cache.GetFloat64(path)
}

func GetString(path string) string {
	return Client.cache.GetString(path)
}

func GetBool(path string) bool {
	return Client.cache.GetBool(path)
}

func GetTime(path string) time.Time {
	return Client.cache.GetTime(path)
}

func DefaultGet(path string, dst interface{}, defaultValue interface{}) {
	Client.cache.DefaultGet(path, dst, defaultValue)
}

func DefaultGetInt(path string, defaultValue int) int {
	return Client.cache.DefaultGetInt(path, defaultValue)
}

func DefaultGetInt8(path string, defaultValue int8) int8 {
	return Client.cache.DefaultGetInt8(path, defaultValue)
}

func DefaultGetInt16(path string, defaultValue int16) int16 {
	return Client.cache.DefaultGetInt16(path, defaultValue)
}

func DefaultGetInt32(path string, defaultValue int32) int32 {
	return Client.cache.DefaultGetInt32(path, defaultValue)
}

func DefaultGetInt64(path string, defaultValue int64) int64 {
	return Client.cache.DefaultGetInt64(path, defaultValue)
}

func DefaultGetUint(path string, defaultValue uint) uint {
	return Client.cache.DefaultGetUint(path, defaultValue)
}

func DefaultGetUint8(path string, defaultValue uint8) uint8 {
	return Client.cache.DefaultGetUint8(path, defaultValue)
}

func DefaultGetUint16(path string, defaultValue uint16) uint16 {
	return Client.cache.DefaultGetUint16(path, defaultValue)
}

func DefaultGetUint32(path string, defaultValue uint32) uint32 {
	return Client.cache.DefaultGetUint32(path, defaultValue)
}

func DefaultGetUint64(path string, defaultValue uint64) uint64 {
	return Client.cache.DefaultGetUint64(path, defaultValue)
}

func DefaultGetFloat(path string, defaultValue float64) float64 {
	return Client.cache.DefaultGetFloat(path, defaultValue)
}

func DefaultGetFloat32(path string, defaultValue float32) float32 {
	return Client.cache.DefaultGetFloat32(path, defaultValue)
}

func DefaultGetFloat64(path string, defaultValue float64) float64 {
	return Client.cache.DefaultGetFloat64(path, defaultValue)
}

func DefaultGetString(path string, defaultValue string) string {
	return Client.cache.DefaultGetString(path, defaultValue)
}

func DefaultGetBool(path string, defaultValue bool) bool {
	return Client.cache.DefaultGetBool(path, defaultValue)
}

func DefaultGetTime(path string, defaultValue time.Time) time.Time {
	return Client.cache.DefaultGetTime(path, defaultValue)
}

func TTL(key string) (time.Duration, bool) {
	return Client.cache.TTL(key)
}

func Set(key string, value interface{}, expiration ...time.Duration) error {
	return Client.cache.Set(key, value, expiration...)
}

func HasPrefix(s string, limit ...int) (map[string]string, error) {
	return Client.cache.HasPrefix(s, limit...)
}

func HasSuffix(s string, limit ...int) (map[string]string, error) {
	return Client.cache.HasSuffix(s, limit...)
}

func Contains(s string, limit ...int) (map[string]string, error) {
	return Client.cache.Contains(s, limit...)
}

func Incr(key string) (int, error) {
	return Client.cache.Incr(key)
}

func IncrBy(key string, step int) (int, error) {
	return Client.cache.IncrBy(key, step)
}

func IncrByFloat(key string, step float64) (float64, error) {
	return Client.cache.IncrByFloat(key, step)
}

func Del(keys ...string) error {
	return Client.cache.Del(keys...)
}

func Publish(channel string, message interface{}) error {
	return Client.cache.Publish(channel, message)
}

func Subscribe(channels []string, handler func(string, string)) error {
	return Client.cache.Subscribe(channels, handler)
}

func PSubscribe(patterns []string, handler func(string, string)) error {
	return Client.cache.PSubscribe(patterns, handler)
}
