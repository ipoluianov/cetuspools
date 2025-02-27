package system

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/ipoluianov/cetuspools/logger"
)

var system *System

func init() {
	logger.Println("system init")
	system = NewSystem()
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
		resp, err := c.UpdateCetusPools()
		if err != nil {
			logger.Println("System working", "UpdateCetusPools Error", err)
			time.Sleep(60 * time.Second)
			continue
		}
		//logger.Println("System working", "resp", resp)
		c.WriteHistory(resp)
		time.Sleep(60 * time.Second)
	}
}

func (c *System) CreateZipWithJSON(jsonData []byte) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	fileWriter, err := zipWriter.Create("data.json")
	if err != nil {
		return nil, err
	}

	_, err = fileWriter.Write(jsonData)
	if err != nil {
		return nil, err
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *System) WriteHistory(data *CetusStatsPools) {
	bsJson, _ := json.MarshalIndent(data, "", "  ")

	zipFileBytes, err := c.CreateZipWithJSON(bsJson)
	if err != nil {
		logger.Println("WriteHistory", "CreateZipWithJSON Error", err)
		return
	}

	os.MkdirAll("./data", 0777)
	f, err := os.Create("./data/" + time.Now().UTC().Format("2006-01-02-15-04-05") + ".zip")
	if err != nil {
		logger.Println("WriteHistory", "os.Create Error", err)
		return
	}
	defer f.Close()
	f.Write(zipFileBytes)
}
