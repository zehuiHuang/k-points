package mode

import (
	"context"
	"errors"
	"fmt"
)

type ChainRule interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	//Next() ChainRule
}

type baseChainRule struct {
	next ChainRule
}

func (b *baseChainRule) Apply(ctx context.Context, params map[string]interface{}) error {
	panic("not implement")
}

func (b *baseChainRule) Next() ChainRule {
	return b.next
}

func (b *baseChainRule) applyNext(ctx context.Context, params map[string]interface{}) error {
	if b.next != nil {
		b.next.Apply(ctx, params)
	}
	return nil
}

type A struct {
	baseChainRule
}

func NewA(chainRule ChainRule) *A {
	return &A{
		baseChainRule{
			next: chainRule,
		},
	}
}

func (i *A) Apply(ctx context.Context, params map[string]interface{}) error {
	age := params["age"].(int)
	if age < 18 {
		return errors.New("age not condition")
	}
	if err := i.applyNext(ctx, params); err != nil {
		return errors.New("after process not condition")
	}
	fmt.Println("A success")
	return nil
}

type B struct {
	baseChainRule
}

func NewB(chainRule ChainRule) *B {
	return &B{
		baseChainRule{
			next: chainRule,
		},
	}
}

func (i *B) Apply(ctx context.Context, params map[string]interface{}) error {
	name := params["name"].(string)
	if name != "hzh" {
		return errors.New("name not condition")
	}
	if err := i.applyNext(ctx, params); err != nil {
		return errors.New("other not condition")
	}
	fmt.Println("B success")
	return nil
}
