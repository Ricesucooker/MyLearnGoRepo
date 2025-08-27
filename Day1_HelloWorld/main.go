package main

import (
	"fmt"
	"time"
)

const japanese = "Japanese"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const japaneseHelloPrefix = "こんにちは, "
const chineseHelloPrefix = "你好, "

// created a new function define that hello world is a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	}

	return prefix + name

	// if language == japanese {
	// 	return japaneseHelloPrefix + name
	// }
	// if language == chinese {
	// 	return chineseHelloPrefix + name
	// }
	// return englishHelloPrefix + name
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func main() {
	fmt.Println(Hello("world", ""))
	fmt.Println(GetCurrentTime())
}
