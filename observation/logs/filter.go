package logs

const fuzzyStr = "***"

// FilterOption is filter option.
type FilterOption func(*Filter)

// Filter is a logger filter.
type Filter struct {
	logger Logger
	level  Level
	key    map[interface{}]struct{}
	value  map[interface{}]struct{}
	filter func(level Level, kvs ...interface{}) bool
}

// NewFilter new a logger filter.
func NewFilter(l Logger, opts ...FilterOption) *Filter {
	f := &Filter{}
	f.Init()
	f.logger = l
	for _, o := range opts {
		o(f)
	}
	return f
}

// Log Print log by level and key values.
func (f *Filter) Log(level Level, kvs ...interface{}) error {
	if level < f.level {
		return nil
	}
	// prefixkv contains the slice of arguments defined as prefixes during the log initialization
	var prefixkv []interface{}
	l, ok := f.logger.(*logger)
	if ok && len(l.prefix) > 0 {
		prefixkv = make([]interface{}, 0, len(l.prefix))
		prefixkv = append(prefixkv, l.prefix...)
	}

	if f.filter != nil && (f.filter(level, prefixkv...) || f.filter(level, kvs...)) {
		return nil
	}

	if len(f.key) > 0 || len(f.value) > 0 {
		for i := 0; i < len(kvs); i += 2 {
			v := i + 1
			if v >= len(kvs) {
				continue
			}
			if _, ok := f.key[kvs[i]]; ok {
				kvs[v] = fuzzyStr
			}
			if _, ok := f.value[kvs[v]]; ok {
				kvs[v] = fuzzyStr
			}
		}
	}
	return f.logger.Log(level, kvs...)
}

func (f *Filter) Init() {
	f.key = make(map[interface{}]struct{})
	f.value = make(map[interface{}]struct{})
}

// FilterLevel with filter level.
func FilterLevel(level Level) FilterOption {
	return func(opts *Filter) {
		opts.level = level
	}
}

// FilterKey with filter key.
func FilterKey(key ...string) FilterOption {
	return func(o *Filter) {
		for _, v := range key {
			o.key[v] = struct{}{}
		}
	}
}

// FilterValue with filter value.
func FilterValue(value ...string) FilterOption {
	return func(o *Filter) {
		for _, v := range value {
			o.value[v] = struct{}{}
		}
	}
}

// FilterFunc with filter func.
func FilterFunc(f func(level Level, kvs ...interface{}) bool) FilterOption {
	return func(o *Filter) {
		o.filter = f
	}
}
