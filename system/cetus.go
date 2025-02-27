package system

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ipoluianov/cetuspools/logger"
)

func (c *System) CetusGetName() string {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.Name
}

func (c *System) PoolsAsTable() string {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return ""
}

func (c *System) CetusGetPoolsTotalCount() (int, error) {
	url := "https://api-sui.cetus.zone/v2/sui/stats_pools?is_vaults=false&display_all_pools=false&has_mining=true&has_farming=true&no_incentives=true&order_by=-tvl&limit=1&offset=0"
	resp, err := http.Get(url)
	if err != nil {
		logger.Println("UpdateCetusPools", "http.Get Error", err)
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println("UpdateCetusPools", "io.ReadAll Error", err)
		return 0, err
	}
	var stats CetusStatsPools
	err = json.Unmarshal(body, &stats)
	if err != nil {
		logger.Println("UpdateCetusPools", "json.Unmarshal Error", err)
		return 0, err
	}
	return stats.Data.Total, nil
}

func (c *System) UpdateCetusPools() (*CetusStatsPools, error) {
	result := &CetusStatsPools{}

	totalCount, err := c.CetusGetPoolsTotalCount()
	if err != nil {
		logger.Println("UpdateCetusPools", "CetusGetPoolsTotalCount Error", err)
		return nil, err
	}
	logger.Println("UpdateCetusPools", "totalCount", totalCount)

	if totalCount > 200 {
		totalCount = 200
	}

	result.Data.Total = totalCount

	pageSize := 30

	for offset := 0; offset < totalCount; offset += pageSize {
		time.Sleep(500 * time.Millisecond)
		logger.Println("UpdateCetusPools", "offset", offset)
		url := "https://api-sui.cetus.zone/v2/sui/stats_pools?is_vaults=false&display_all_pools=false&has_mining=true&has_farming=true&no_incentives=true&order_by=-tvl&limit=" + fmt.Sprint(pageSize) + "&offset=" + fmt.Sprint(offset)
		resp, err := http.Get(url)
		if err != nil {
			logger.Println("UpdateCetusPools", "http.Get Error", err)
			return nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Println("UpdateCetusPools", "io.ReadAll Error", err)
			return nil, err
		}

		var stats CetusStatsPools
		err = json.Unmarshal(body, &stats)
		if err != nil {
			logger.Println("UpdateCetusPools", "json.Unmarshal Error", err)
			return nil, err
		}
		result.Data.LpList = append(result.Data.LpList, stats.Data.LpList...)
		logger.Println("UpdateCetusPools", "len(items)", len(result.Data.LpList))
	}
	return result, nil
}
