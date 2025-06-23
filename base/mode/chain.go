package mode

import (
	"context"
	"errors"
	"fmt"
)

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
