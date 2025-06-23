package mode

import (
	"fmt"
)

//适配器模式

type PhoneCharger interface {
	Output5V()
}

type HuaWeiCharger struct {
}

func NewHuaWeiCharger() *HuaWeiCharger {
	return &HuaWeiCharger{}
}
func (h *HuaWeiCharger) Output5V() {
	fmt.Println("华为手机充电器输出 5V 电压...")
}

type XiaoMiCharger struct {
}

func NewXiaoMiCharger() *XiaoMiCharger {
	return &XiaoMiCharger{}
}
func (x *XiaoMiCharger) Output5V() {
	fmt.Println("小米手机充电器输出 5V 电压...")
}

type MacBookCharger struct {
}

func NewMacBookCharger() *MacBookCharger {
	return &MacBookCharger{}
}
func (m *MacBookCharger) Output28V() {
	fmt.Println("苹果笔记本充电器输出 28V 电压...")
}

// -----------------------------Adapter

type MacBookChargerAdapter struct {
	core *MacBookCharger
}

func NewMacBookChargerAdapter(m *MacBookCharger) *MacBookChargerAdapter {
	return &MacBookChargerAdapter{
		core: m,
	}
}

func (m *MacBookChargerAdapter) Output5V() {
	m.core.Output28V()
	fmt.Println("适配器将输出电压调整为 5V...")
}

//------------------phone

type Phone interface {
	Charge(phoneCharger PhoneCharger)
}
type HuaWeiPhone struct {
}

func NewHuaWeiPhone() Phone {
	return &HuaWeiPhone{}
}

func (h *HuaWeiPhone) Charge(phoneCharger PhoneCharger) {
	fmt.Println("华为手机准备开始充电...")
	phoneCharger.Output5V()
}

func skip() {

}
