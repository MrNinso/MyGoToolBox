package tests

import (
	"github.com/MrNinso/MyGoToolBox/objects/workPool"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func TestSimplePool(t *testing.T) {
	workesCalls := 0
	resultCalls := 0
	key := time.Now().UnixNano()

	p := workPool.NewSimplePool(runtime.NumCPU(), func(workerId int, job interface{}) (result interface{}) {
		assert.Equal(t, key, job)
		time.Sleep(time.Second)
		workesCalls++
		return job
	})

	p.SetOnResult(func(p *workPool.SimplePool, result interface{}) {
		assert.Equal(t, key, result)
		resultCalls++
		if resultCalls == runtime.NumCPU() {
			p.Close()
		}
	})

	for i := 0; i < runtime.NumCPU(); i++ {
		p.AddJob(key)
	}

	p.StartWorkes(true)

	assert.Equal(t, runtime.NumCPU(), workesCalls)
	assert.Equal(t, runtime.NumCPU(), resultCalls)
}
