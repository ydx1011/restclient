package reflection

import "reflect"

func IsNil(o interface{}) bool {
	if o == nil {
		return true
	}

	return reflect.ValueOf(o).IsNil()
}
