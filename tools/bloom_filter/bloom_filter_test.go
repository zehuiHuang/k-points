package bloom_filter

import "testing"

func TestEncrypt(t *testing.T) {
	enc := NewEncryptor()
	a := enc.Encrypt("test2")
	println(a)
}

func TestNewLocalBloomService(t *testing.T) {
	enc := NewEncryptor()
	localBloomService := NewLocalBloomService(32, 2, enc)
	localBloomService.Set("test")
	localBloomService.Exist("test1")
}
