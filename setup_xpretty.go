package xpretty

type XOpts struct {
	noColor  bool
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

func WithNoColor(b bool) XOptFunc {
	return func(o *XOpts) {
		o.noColor = b
	}
}

var ctrl = &XOpts{}

// Initialize setups
//   - color: used for `terminal vivid output`, true by default
//   - dummyLog: used in `DummyLog`, true by default
func Initialize(opts ...XOptFunc) {
	opt := XOpts{noColor: true, dummyLog: true}
	bindXOpts(&opt, opts...)
	ctrl = &opt
	SetNoColor(ctrl.noColor)
}

func InitializeWithColor(opts ...XOptFunc) {
	opt := XOpts{noColor: false, dummyLog: true}
	bindXOpts(&opt, opts...)
	ctrl = &opt

	SetNoColor(false)
}
