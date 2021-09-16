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
