package mode

import (
	"context"
	"errors"
	"fmt"
)

type Weather struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func (w Weather) InvokableRun(ctx context.Context, argumentsInJSON string) (string, error) {
	fmt.Println("---------------call Weather tool---------------------")
	return "{\"temperature\": 25.0, \"humidity\": 60.0}", nil
}

type InvokableTool interface {
	InvokableRun(ctx context.Context, argumentsInJSON string) (string, error)
}

// InvokableToolEndpoint 定义 函数类型
type InvokableToolEndpoint func(ctx context.Context, input *string) (*string, error)

// InvokableToolMiddleware 定义函数类型
type InvokableToolMiddleware func(InvokableToolEndpoint) InvokableToolEndpoint

// 通过包装进行链式调用,对实现了InvokableTool接口的方法进行包装
// 包装函数是一个链式结构调用:middleware1 -> middleware2 ->middleware3
func wrapToolCall(it InvokableTool, middlewares []InvokableToolMiddleware) InvokableToolEndpoint {
	middleware := func(next InvokableToolEndpoint) InvokableToolEndpoint {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
	return middleware(func(ctx context.Context, input *string) (*string, error) {
		fmt.Println("---------------call tool---------------------")
		result, err := it.InvokableRun(ctx, *input)
		if err != nil {
			return nil, err
		}
		return &result, nil
	})
}

/////////////////------------------------------------------------------------------------------------------------------------------------
//责任链模式,网关的白名单校验、审核等操作是否可以加入责任链模式？或者是入参校验、频度控制、llm参数校验、权益验证、审核

type RuleChain interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	Next() RuleChain
}

type baseRuleChain struct {
	next RuleChain
}

func (b *baseRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	panic("not implement")
}

func (b *baseRuleChain) Next() RuleChain {
	return b.next
}

func (b *baseRuleChain) applyNext(ctx context.Context, params map[string]interface{}) error {
	if b.Next() != nil {
		return b.Next().Apply(ctx, params)
	}
	return nil
}

/////////////////////执行逻辑
//token校验

type CheckTokenRule struct {
	baseRuleChain
}

func NewCheckTokenRule(next RuleChain) RuleChain {
	return &CheckTokenRule{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}
func (c *CheckTokenRule) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验 token 是否合法
	token, _ := params["token"].(string)
	if token != "myToken" {
		return fmt.Errorf("invalid token: %s", token)
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check token rule err post process...")
		return err
	}

	fmt.Println("check token rule common post process...")
	return nil
}

//age校验

type CheckAgeRule struct {
	baseRuleChain
}

func NewCheckAgeRule(next RuleChain) RuleChain {
	return &CheckAgeRule{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}
func (c *CheckAgeRule) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验 age 是否合法
	age, _ := params["age"].(int)
	if age < 18 {
		return fmt.Errorf("invalid age: %d", age)
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check age rule err post process...")
		return err
	}

	fmt.Println("check age rule common post process...")
	return nil
}

//AuthorizedStatus校验

type CheckAuthorizedStatus struct {
	baseRuleChain
}

func NewCheckAuthorizedStatus(next RuleChain) RuleChain {
	return &CheckAuthorizedStatus{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAuthorizedStatus) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验是否已认证
	if authorized, _ := params["authorized"].(bool); !authorized {
		return errors.New("not authorized yet")
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check authorized status rule err post process...")
		return err
	}

	fmt.Println("check authorized statuse rule common post process...")
	return nil
}
