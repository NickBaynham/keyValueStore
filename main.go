package main

var store = make(map[string]string)

func main() {
	var _ = put("key", "value")
}

func put(key string, value string) error {
	store[key] = value
	return nil
}