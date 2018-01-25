package log

import (
	"fmt"
	logger "github.com/YoungPioneers/blog4go"
	"sync"
	"testing"
)

// MyHook .
type MyHook struct {
	something string
}

// Fire .
func (hook *MyHook) Fire(level logger.LevelType, args ...interface{}) {
	fmt.Println(args...)
}

type T struct {
	A int
	B string
}

func BenchmarkBlog4goSingleGoroutine(b *testing.B) {
	err := logger.NewWriterFromConfigAsFile("blog4go_config.xml")
	defer logger.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	// hook := new(MyHook)
	// logger.SetHook(hook)
	// logger.SetHookLevel(logger.INFO)
	t := T{123, "test"}
	logger.SetColored(true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
		logger.Errorf("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
	}
}

func BenchmarkBlog4goMultiGoroutine(b *testing.B) {
	err := logger.NewWriterFromConfigAsFile("blog4go_config.xml")
	defer logger.Close()
	if nil != err {
		fmt.Println(err.Error())
	}

	t := T{123, "testing"}
	var wg sync.WaitGroup
	var beginWg sync.WaitGroup

	f := func() {
		defer wg.Done()
		beginWg.Wait()
		for i := 0; i < b.N; i++ {
			logger.Infof("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
			logger.Errorf("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
		}
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		beginWg.Add(1)
	}

	b.ResetTimer()
	for i := 0; i < 100; i++ {
		go f()
		beginWg.Done()
	}

	wg.Wait()
}
