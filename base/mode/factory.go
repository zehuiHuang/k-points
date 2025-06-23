package mode

import (
	"fmt"
	_ "go.uber.org/dig"
)

// 简单工厂模式、工厂方法模式、抽象工程模式、容器工厂模式
// https://zhuanlan.zhihu.com/p/636171303

type Fruit interface {
	Eat()
}

//--------------------------------抽象工厂模式案例
/*
 抽象出固定不变或不常变的
*/

// Strawberry 草莓
type Strawberry interface {
	SweetAttack()
}

// Lemon 柠檬
type Lemon interface {
	AcidAttack()
}

type FruitFactory interface {
	CreateStrawberry() Strawberry
	CreateLemon() Lemon
}

type GoodFarmerStrawberry struct {
	brand string
	Strawberry
}

func (g *GoodFarmerStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s, ", g.brand)
}

type GoodFarmerLemon struct {
	brand string
	Lemon
}

func (g *GoodFarmerLemon) AcidAttack() {
	fmt.Printf("acid attack from %s, ", g.brand)
}

type GoodFarmerFactory struct{}

func (g *GoodFarmerFactory) myAspect() {
	fmt.Println("good farmer aspect...")
}
func (g *GoodFarmerFactory) CreateStrawberry() Strawberry {
	// 同一个产品族可以插入一个切面
	g.myAspect()
	defer g.myAspect()
	return &GoodFarmerStrawberry{
		brand: "goodfarmer",
	}
}

func (g *GoodFarmerFactory) CreateLemon() Lemon {
	// 同一个产品族可以插入一个切面
	g.myAspect()
	defer g.myAspect()
	return &GoodFarmerLemon{
		brand: "goodfarmer",
	}
}

// ----------------------------------------扩展一种品牌水果

type ZespriStrawberry struct {
	brand string
	Strawberry
}

func (z *ZespriStrawberry) SweetAttack() {
	fmt.Printf("sweet attack from %s, ", z.brand)
}

type ZespriLemon struct {
	brand string
	Lemon
}

func (z *ZespriLemon) AcidAttack() {
	fmt.Printf("acid attack from %s, ", z.brand)
}

type ZespriFactory struct{}

func (z *ZespriFactory) myAspect() {
	fmt.Println("dole aspect...")
}

func (z *ZespriFactory) CreateStrawberry() Strawberry {
	// 同一个产品族可以插入一个切面
	z.myAspect()
	defer z.myAspect()
	return &ZespriStrawberry{
		brand: "zespri",
	}
}
func (z *ZespriFactory) CreateLemon() Lemon {
	// 同一个产品族可以插入一个切面
	z.myAspect()
	defer z.myAspect()
	return &ZespriLemon{
		brand: "zespri",
	}
}

/**
总结：
1、不需要或不频繁变更的定义为产品等级，即芒果和草莓，后续不会有新的水果种类进来
2、经常容易变更的，这里定义为产品族，即品牌，支持水平扩展
3、符合开闭原则
4、网关是否需要考虑，
*/

//---------------------------------- 容器工厂模式

//type Factory struct {
//	container *dig.Container
//}
//
//func (f *Factory) Inject(constructor interface{}) error {
//	return f.container.Provide(constructor)
//}
//
//func (f *Factory) Invoke(invoker interface{}) error {
//	return f.container.Invoke(invoker)
//}
//
//var (
//	once    sync.Once
//	factory *Factory
//)
//
//func GetFactory() *Factory {
//	once.Do(func() {
//		factory = newFactory(dig.New())
//	})
//	return factory
//}
//
//func init() {
//	f := GetFactory()
//	f.Inject(NewComponentX)
//}
//
//type ComponentX struct{}
//
//func NewComponentX() *ComponentX {
//	return &ComponentX{}
//}
//func GetComponentX() (*ComponentX, error) {
//	f := GetFactory()
//	var componentX *ComponentX
//	return componentX, f.Invoke(func(_x *ComponentX) {
//		componentX = _x
//	})
//}

// Api 网关逻辑处理
type Api interface {
	a()
	b()
}
type Chat interface {
	Api
}

type ChatCompletions interface {
	Api
}

//type Embedding interface {
//	Api
//}

// LLMFactory BaiduFactory实现了该接口
type LLMFactory interface {
	CreateChat() Chat
	CreateChatCompletions() ChatCompletions
	//CreateChatEmbedding() Embedding
	//CreateImage() Image
	//CreateMultimodal() Multimodal
}

type BaiduChat struct {
	name string
	Chat
}
type BaiduChatCompletions struct {
	name string
	ChatCompletions
}

func (g *BaiduChat) a() {
	fmt.Printf("a chat from %s, ", g.name)
}
func (g *BaiduChat) b() {
	fmt.Printf("b chat from %s, ", g.name)
}

func (g *BaiduChatCompletions) a() {
	fmt.Printf("a ChatCompletions from %s, ", g.name)
}

func (g *BaiduChatCompletions) b() {
	fmt.Printf("b ChatCompletions from %s, ", g.name)
}

type BaiduFactory struct{}

func (g *BaiduFactory) myAspect() {
	fmt.Println("baidu aspect...")
}

func (g *BaiduFactory) CreateChat() Chat {
	// 同一个产品族可以插入一个切面
	g.myAspect()
	defer g.myAspect()
	return &BaiduChat{
		name: "baidu-chat",
	}
}

func (g *BaiduFactory) CreateChatCompletions() ChatCompletions {
	// 同一个产品族可以插入一个切面
	g.myAspect()
	defer g.myAspect()
	return &BaiduChatCompletions{
		name: "baidu-chatCompletions",
	}
}

func init() {
	//注册到worker中

}
