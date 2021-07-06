package to

type Setting struct {
	UseDefaultEmpty bool
	CustomValue     interface{}
}

type Option func(setting *Setting)

func UseDefaultEmpty(setting *Setting) {
	setting.UseDefaultEmpty = true
}
func UseCustomValueWrapper(customValue interface{}) Option {
	return func(setting *Setting) {
		setting.CustomValue = customValue
	}
}
