package xpretty

type XOpts struct {
	color    bool
	dummyLog bool
}

type XOptFunc func(o *XOpts)

func bindXOpts(opt *XOpts, opts ...XOptFunc) {
	for _, f := range opts {
		f(opt)
	}
}

func WithDummyLog(b bool) XOptFunc {
	return func(o *XOpts) {
		o.dummyLog = b
	}
}

func WithColor(b bool) XOptFunc {
	return func(o *XOpts) {
		o.color = b
	}
}

var ctrl = &XOpts{}

// Initialize setups
//   - color: used in `DummyLog`
//   - dummyLog: used in `DummyLog`
func Initialize(opts ...XOptFunc) {
	opt := XOpts{}
	bindXOpts(&opt, opts...)
	ctrl = &opt
	ToggleColor(ctrl.color)
}
