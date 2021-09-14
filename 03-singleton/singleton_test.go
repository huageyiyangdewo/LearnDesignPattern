package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins2 != ins1 {
		t.Fatal("instance is not equal")
	}
}


const parCount = 100

func TestParallelSingleton(t *testing.T)  {
	// 无缓冲的通道
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	instances := [parCount]*Singleton{}
	//var instances []*Singleton

	wg.Add(parCount)

	for i := 0; i < parCount; i++ {
		go func(i int) {
			// 阻塞等待，关闭通道时在执行，实现并发
			<- start
			instances[i] = GetInstance()
			// 注意，不能使用切片，因为 slice 和 map 都是线程不安全的，所以 instances 最终的数量可能不是 100
			//instances = append(instances, GetInstance())
			wg.Done()
		}(i)
	}
	//关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	//阻塞等待, 直到WaitGroup,计数器为0, 即等待上面的100个协程执行完毕后
	wg.Wait()

	fmt.Println(len(instances))
	for i := 1; i < parCount; i++ {
		if instances[i - 1] != instances[i] {
			t.Fatal("instance is not equal")
		}
	}
}