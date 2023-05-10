package structs

import "reflect"

var nilable = map[reflect.Kind]bool{
	reflect.Chan:      true,
	reflect.Func:      true,
	reflect.Interface: true,
	reflect.Map:       true,
	reflect.Pointer:   true,
	reflect.Slice:     true,
}

// ApplyDefaults will set every field in value that is set to a zero or nil
// value to the value in the same field in defaults. Only exported fields will
// be set to a default value.
func ApplyDefaults[T any](value, defaults *T) {
	var x T
	tt := reflect.TypeOf(x)
	vv := reflect.ValueOf(value).Elem()
	dv := reflect.ValueOf(defaults).Elem()

	for i := 0; i < tt.NumField(); i++ {
		vf := vv.Field(i)
		tf := tt.Field(i)
		if nilable[tf.Type.Kind()] && !vf.IsNil() {
			continue
		} else if vf.IsValid() && !vf.IsZero() {
			continue
		}

		vv.Field(i).Set(dv.Field(i))
	}
}
