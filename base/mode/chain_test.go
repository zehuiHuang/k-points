package mode

import (
	"context"
	"fmt"
	"go-learn/base/mode/common"
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

func TestWrapToolCall(t *testing.T) {
	f := wrapToolCall(&Weather{}, []common.InvokableToolMiddleware{
		func(next common.InvokableToolEndpoint) common.InvokableToolEndpoint {
			return func(ctx context.Context, input *string) (*string, error) {
				fmt.Println("11111111111111 - before")
				result, err := next(ctx, input)
				fmt.Println("11111111111111 - after")
				return result, err
			}
		},
		func(next common.InvokableToolEndpoint) common.InvokableToolEndpoint {
			return func(ctx context.Context, input *string) (*string, error) {
				fmt.Println("22222222222222 - before")
				result, err := next(ctx, input)
				fmt.Println("22222222222222 - after")
				return result, err
			}
		},
		func(next common.InvokableToolEndpoint) common.InvokableToolEndpoint {
			return func(ctx context.Context, input *string) (*string, error) {
				fmt.Println("33333333333333 - before")
				result, err := next(ctx, input)
				fmt.Println("33333333333333 - after")
				return result, err
			}
		},
	})

	data := "testdata"
	result, err := f(context.Background(), &data)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", *result)
	// 打印应该如下:
	/**
	11111111111111 - before
	22222222222222 - before
	33333333333333 - before
	---------------call tool---------------------
	---------------call Weather tool---------------------
	33333333333333 - after
	22222222222222 - after
	11111111111111 - after
	*/
}
