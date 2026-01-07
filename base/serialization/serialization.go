package serialization

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
)

var mp = map[string][]byte{}

func seria(key string) {

	//序列化
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(&serialization{
		Data:                "data-----",
		InterruptID2Address: map[string]string{"1": "1"},
		EnableStreaming:     true,
	})
	if err != nil {
		fmt.Println(err)
		panic("序列号报错")
	}
	mp[key] = buf.Bytes()

	//反序列化
	data := mp[key]
	s := &serialization{}
	err = gob.NewDecoder(bytes.NewReader(data)).Decode(s)
	if err != nil {
		fmt.Println(err)
		panic("反序列号报错")
	}
	fmt.Println(s.strings())
}

type serialization struct {
	Data                string
	EnableStreaming     bool
	InterruptID2Address map[string]string
}

func (s *serialization) strings() string {
	fmt.Println(s.Data)
	fmt.Println(s.InterruptID2Address)
	return s.Data
}
func Set(ctx context.Context, key string, value []byte) error {
	mp[key] = value
	return nil
}
