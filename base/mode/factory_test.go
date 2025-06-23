package mode

import "testing"

//func TestBaidu(t *testing.T) {
//	// 尝尝佳农品牌的水果
//	baiduChatFactory := BaiduFactory{}
//	baiduChat := baiduChatFactory.CreateChat()
//	baiduChat.a()
//}

func TestBaidu(t *testing.T) {
	//
	baiduFactory := BaiduFactory{}
	baiduChat := baiduFactory.CreateChat()
	baiduChat.a()
}
