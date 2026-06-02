package util

func MapMap[
	InputKey comparable, InputValue any, OutputKey comparable, OutputValue any,
](
	value map[InputKey]InputValue,
	converter func(key InputKey, value InputValue) (OutputKey, OutputValue),
) map[OutputKey]OutputValue {
	result := make(map[OutputKey]OutputValue)
	for key, value := range value {
		newKey, newValue := converter(key, value)
		result[newKey] = newValue
	}
	return result
}

func GetMapKeys[InputKey comparable, OutputKey any](m map[InputKey]OutputKey) []InputKey {
	keys := make([]InputKey, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetMapVals[InputKey comparable, OutputKey any](m map[InputKey]OutputKey) []OutputKey {
	values := make([]OutputKey, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
