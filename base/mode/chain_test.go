package mode

import (
	"context"
	"testing"
)

func TestHzh(t *testing.T) {
	a := NewA(nil)
	b := NewB(a)
	if err := b.Apply(context.Background(), map[string]interface{}{
		"name": "hzh",
		"age":  19,
	}); err != nil {
		// 校验未通过，终止发奖流程
		t.Error(err)
		return
	}
}

func Test_RuleChainV2(t *testing.T) {
	checkAuthorizedRule := NewCheckAuthorizedStatus(nil)
	checkAgeRule := NewCheckAgeRule(checkAuthorizedRule)
	checkTokenRule := NewCheckTokenRule(checkAgeRule)

	if err := checkTokenRule.Apply(context.Background(), map[string]interface{}{
		"token":      "myToken",
		"age":        18,
		"authorized": true,
	}); err != nil {
		// 校验未通过，终止发奖流程
		t.Error(err)
		return
	}

	// 通过前置校验流程，进行奖励发放
	//sendReward(ctx, params)
}
