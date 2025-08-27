package main

import (
	"fmt"
	"time"
)

const (
	japanese            = "Japanese"
	chinese             = "Chinese"
	englishHelloPrefix  = "Hello, "
	japaneseHelloPrefix = "こんにちは, "
	chineseHelloPrefix  = "你好, "
)

// created a new function define that hello world is a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func main() {
	fmt.Println(Hello("world", ""))
	fmt.Println(GetCurrentTime())
	fmt.Println(Hello("月球", "Chinese"))
	fmt.Println(Hello("美月", "Japanese"))
}
