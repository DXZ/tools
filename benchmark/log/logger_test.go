package log

import (
	"logger/logger"
	"sync"
	"testing"
)

func BenchmarkLoggerSingle(b *testing.B) {
	logger.SetRollingDaily("./", "logger.log")
	logger.SetLevel(logger.INFO)
	logger.SetConsole(false)
	log := logger.GetLogger()
	log.SetRollingDaily("./", "logger.log")
	log.SetLevel(logger.INFO)
	log.SetConsole(false)
	t := T{123, "test"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
		log.Error("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
	}
}

func BenchmarkLoggerMultiGroutine(b *testing.B) {
	logger.SetRollingDaily("./", "logger.log")
	logger.SetLevel(logger.INFO)
	logger.SetConsole(false)
	log := logger.GetLogger()
	log.SetRollingDaily("./", "logger.log")
	log.SetLevel(logger.INFO)
	log.SetConsole(false)
	t := T{123, "test"}
	var wg sync.WaitGroup
	var beginWg sync.WaitGroup

	f := func() {
		defer wg.Done()
		beginWg.Wait()
		for i := 0; i < b.N; i++ {
			logger.Info("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
			log.Error("haha %s. en\\en, always %d and %f, %t, %+v", "eddie", 18, 3.1415, true, t)
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
