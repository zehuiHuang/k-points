package mode

import "testing"

func Test_adapter(t *testing.T) {
	// 创建一个华为手机实例
	huaWeiPhone := NewHuaWeiPhone()

	// 使用华为手机充电器进行充电
	HuaWeiCharger := NewHuaWeiCharger()
	huaWeiPhone.Charge(HuaWeiCharger)

	// 使用适配器转换后的 macbook 充电器进行充电
	macBookCharger := NewMacBookCharger()
	macBookChargerAdapter := NewMacBookChargerAdapter(macBookCharger)
	huaWeiPhone.Charge(macBookChargerAdapter)
}
