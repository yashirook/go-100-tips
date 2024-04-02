package redis

import "fmt"

func init() {
	fmt.Println("init in redis package")
}

func Store(key, value string) error {
	fmt.Printf("dummy store! key: %s, value: %s\n", key, value)
	return nil
}
