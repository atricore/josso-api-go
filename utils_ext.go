package jossoappi

func IPtrString(s interface{}) *string {
	if _, ok := s.(string); ok {
		return PtrString(s.(string))
	}
	return nil
}

func IPtrBool(s interface{}) *bool {
	if _, ok := s.(bool); ok {
		return PtrBool(s.(bool))
	}
	return nil
}

func AsBool(i interface{}, d bool) bool {
	if i == nil {
		return d
	}

	switch v := i.(type) {
	case bool:
		return v
	case *bool:
		return *v

	}

	return d
}

func AsInt32(i interface{}, d int32) int32 {
	if i == nil {
		return d
	}

	switch v := i.(type) {
	case int32:
		return v
	case int:
	case int64:
		return int32(v)
	}

	return d

}

func AsString(i interface{}, d string) string {
	if i == nil {
		return d
	}

	return i.(string)
}
