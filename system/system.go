package system

import (
	"sync"
	"time"

	"github.com/ipoluianov/cetuspools/logger"
)

var system *System

func init() {
	logger.Println("system init")
	system = NewSystem()
	system.Start()
}

func Get() *System {
	return system
}

type System struct {
	mtx      sync.Mutex
	stopping bool
	Name     string
}

func NewSystem() *System {
	var c System
	c.Name = "123123123"
	return &c
}

func (c *System) Start() {
	logger.Println("System start")
	go c.ThWork()
}

func (c *System) Stop() {
	logger.Println("System stop")
}

func (c *System) ThWork() {
	for !c.stopping {
		logger.Println("System working")
		time.Sleep(1 * time.Second)
	}
}
