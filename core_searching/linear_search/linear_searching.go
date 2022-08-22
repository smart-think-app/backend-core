package linear_search

import "reflect"

type SearchKey struct {
	Field string
	Value interface{}
}
func SearchString(arr []string , value string) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		if v == value {
			index = i
			found = true
			return
		}
	}
	return
}

func SearchFloat64(arr []float64 , value float64) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		if v == value {
			index = i
			found = true
			return
		}
	}
	return
}

func SearchFloat32(arr []float32 , value float32) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		if v == value {
			index = i
			found = true
			return
		}
	}
	return
}

func SearchInt(arr []int , value int) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		if v == value {
			index = i
			found = true
			return
		}
	}
	return
}

func SearchInt64(arr []int64 , value int64) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		if v == value {
			index = i
			found = true
			return
		}
	}
	return
}

func SearchInt32(arr []int32 , value int32) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		if v == value {
			index = i
			found = true
			return
		}
	}
	return
}

func SearchByKey(arr []interface{} , keys SearchKey) (found bool , index int) {
	index = -1
	for i, v:= range arr {
		ref := reflect.ValueOf(v)
		field := ref.FieldByName(keys.Field)
		value := field.Interface()
		if value == keys.Value {
			found = true
			index = i
			return
		}
	}
	return
}

func SearchByMultiKey(arr []interface{} , keys []SearchKey) (found bool , index int) {
	index = -1
	type compare struct {
		field interface{}
		value interface{}
	}
	check := func(cpa []compare) bool {
		for _ , cp := range cpa {
			if cp.value != cp.field {
				return false
			}
		}
		return true
	}
	for i, v:= range arr {
		ref := reflect.ValueOf(v)
		arrCpa := make([]compare  , 0)
		for _ , key := range keys{
			arrCpa = append(arrCpa, compare{
				field: ref.FieldByName(key.Field).Interface(),
				value: key.Value,
			})
		}
		if check(arrCpa) {
			index = i
			found = true
			return
		}
	}
	return
}