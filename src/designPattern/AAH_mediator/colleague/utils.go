package colleague

import (
	"reflect"
	"strings"
)

func GetTypeName(ob interface{}) string {
	v := reflect.ValueOf(ob)
	vs := strings.Split(v.Type().String(), ".")
	return vs[len(vs)-1]
}
