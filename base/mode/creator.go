package mode

//创建者模式
//options

type BigClass struct {
	LogOptions
}

type LogOptions struct {
	A string
	B string
	C string
}

// //////////////////////
type option interface {
	apply(*LogOptions)
}

type funcOptions struct {
	f func(*LogOptions)
}

func (fos funcOptions) apply(options *LogOptions) {
	fos.f(options)
}

func newFuncOptions(f func(*LogOptions)) funcOptions {
	return funcOptions{
		f: f,
	}
}

// ---------------------------------------------------------------------------------------------------另一种方式
// 定义一个方法类型
//原理其实就是利用闭包的特性,将外部参数加入到函数中,后续参数以多个函数的方式传入,并遍历函数进行

type Option func(opts *LogOptions)

func WithAA(a string) Option {
	return func(opts *LogOptions) {
		opts.A = a
	}
}

func WithBB(b string) Option {
	return func(opts *LogOptions) {
		opts.B = b
	}
}

func WithCC(c string) Option {
	return func(opts *LogOptions) {
		opts.C = c
	}
}

func repair(opts *LogOptions) {
	if opts.C == "" {
		opts.C = "c"
	}
}

func NewBigClass(opts ...Option) *BigClass {
	bigClass := BigClass{}
	for _, opt := range opts {
		opt(&bigClass.LogOptions)
	}
	repair(&bigClass.LogOptions)
	return &bigClass

}
