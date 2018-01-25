package log

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"logger/fmtlogger"
	"sync"
	"testing"
)

func BenchmarkFmtSingleGroutine(b *testing.B) {
	// (fmtlogger.LogConfig)
	var conf fmtlogger.LogConfig
	if _, err := toml.DecodeFile("logconfig.toml", &conf); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmtlogger.InitConfig(&conf, "debug")
	t := T{123, "test"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmtlogger.Debug("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
		fmtlogger.Error("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
	}
}

func BenchmarkFmtMultiGroutine(b *testing.B) {
	// (fmtlogger.LogConfig)
	var conf fmtlogger.LogConfig
	if _, err := toml.DecodeFile("logconfig.toml", &conf); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmtlogger.InitConfig(&conf, "debug")
	t := T{123, "test"}
	var wg sync.WaitGroup
	var beginWg sync.WaitGroup

	f := func() {
		defer wg.Done()
		beginWg.Wait()
		for i := 0; i < b.N; i++ {
			fmtlogger.Debug("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
			fmtlogger.Error("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
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
