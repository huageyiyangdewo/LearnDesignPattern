package singleton

import "sync"

// Singleton 是单例模式
type Singleton struct {}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
