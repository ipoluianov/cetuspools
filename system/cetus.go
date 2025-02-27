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
	bs, _ := json.MarshalIndent(c.table, "", "  ")
	return string(bs)
}

/*
Response:
{"code":200,"msg":"OK","data":{"total":24009,"lp_list":[{"symbol":"USDC-SUI","name":"Native USDC-SUI Token","decimals":6,"fee":"0.0025","tick_spacing":"60","pool_type":"","address":"0xb8d7d9e66a60c239e7a60110efcf8de6c705580ed924d0dde141f4a0e2c90105","coin_a_address":"0xdba34672e30cb065b1f93e3ab55318768fd6fef66c15942c9f7cb846e2f900e7::usdc::USDC","coin_b_address":"0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI","project_url":"","is_display_rewarder":true,"is_closed":false,"rewarder_display1":true,"rewarder_display2":true,"rewarder_display3":true,"rewarder_display4":true,"rewarder_display5":true,"labels":null,"coin_a":{"name":"Native USDC","symbol":"USDC","decimals":6,"address":"0xdba34672e30cb065b1f93e3ab55318768fd6fef66c15942c9f7cb846e2f900e7::usdc::USDC","balance":"3996088.31717","logo_url":"https://gateway.irys.xyz/EGpc2cG886CrWwLMneF2RyVpZ7D33a6znz6XE8n8nU7h","coingecko_id":"usd-coin","project_url":"","labels":[],"is_trusted":true},"coin_b":{"name":"SUI Token","symbol":"SUI","decimals":9,"address":"0x0000000000000000000000000000000000000000000000000000000000000002::sui::SUI","balance":"5848579.290635528","logo_url":"https://archive.cetus.zone/assets/image/sui/sui.png","coingecko_id":"sui","project_url":"","labels":[],"is_trusted":true},"price":"0.340530013270094244142471191467731472880556159256038595000171061","rewarder_usd":["0.04234086937873556752156993118349354355973796021114885332510546778757323448430867518059808854","0.16451561053017987329999362259079629958203721890212146157072420352174871568775877614262364896"],"rewarder_apr":["6.09235627167323186854909062797388687962579300263281459014690616%","23.6717693551068978134212157064199325863317438985168942460331252%","0%","0%","0%"],"is_forward":false,"price_range_config":null,"object":{"coin_a":4006889527953,"coin_b":5852582887224387,"tick_spacing":60,"liquidity":"1336279221810653","current_sqrt_price":"340406167575098811728","rewarder_manager":{"fields":{"rewarders":[{"fields":{"reward_coin":{"fields":{"name":"06864a6f921804860930db6ddbe2e16acdf8504495ea7481637a1c8b9a8fe54b::cetus::CETUS"}},"emissions_per_second":"6405119470038038755555555556","growth_global":"22844531770778864314"}},{"fields":{"reward_coin":{"fields":{"name":"0000000000000000000000000000000000000000000000000000000000000002::sui::SUI"}},"emissions_per_second":"1033786282464139455146666667","growth_global":"4513163272202884428"}}],"points_released":"225897112379448314101235712000000","points_growth_global":"258980132520830670","last_updated_time":1740643860}},"is_pause":false,"index":3137},"category":"","is_vaults":false,"stable_farming":{"stable_farming_pool":"","is_show_rewarder":false,"show_rewarder_1":false,"show_rewarder_2":false,"show_rewarder_3":false,"tvl":"0","apr":"0","effective_tick_lower":0,"effective_tick_upper":0,"total_wrapped_amount":"","stable_rewarder":null},"vaults":["0x8740c0c9ae8cd90b3c608ad4b574250acd40032cb639b6167eb1c61fabe1cc2d"],"pure_tvl_in_usd":"21917397.206064","vol_in_usd_24h":"45864983.305755","fee_24_h":"114662.45826438843367004068632465776226413153325","total_apr":"2.1653023985577053602947203065276264031111211560654271057473085561","apr":{"fee_apr_24h":"186.76611422899040634750172431836882084515457870539300173855082425%"},"extensions":null,"has_mining":true,"has_farming":false,"no_incentives":false}]}}
*/

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

	logger.Println("UpdateCetusPools", "body", string(body))

	var stats CetusStatsPools
	err = json.Unmarshal(body, &stats)
	if err != nil {
		logger.Println("UpdateCetusPools", "json.Unmarshal Error", err)
		return 0, err
	}
	logger.Println("UpdateCetusPools", "stats", stats)
	return stats.Data.Total, nil
}

func (c *System) UpdateCetusPools() {
	totalCount, err := c.CetusGetPoolsTotalCount()
	if err != nil {
		logger.Println("UpdateCetusPools", "CetusGetPoolsTotalCount Error", err)
		return
	}
	logger.Println("UpdateCetusPools", "totalCount", totalCount)

	var items []*CetusPoolsItem

	if totalCount > 200 {
		totalCount = 200
	}

	pageSize := 30

	for offset := 0; offset < totalCount; offset += pageSize {
		time.Sleep(500 * time.Millisecond)
		logger.Println("UpdateCetusPools", "offset", offset)
		url := "https://api-sui.cetus.zone/v2/sui/stats_pools?is_vaults=false&display_all_pools=false&has_mining=true&has_farming=true&no_incentives=true&order_by=-tvl&limit=" + fmt.Sprint(pageSize) + "&offset=" + fmt.Sprint(offset)
		resp, err := http.Get(url)
		if err != nil {
			logger.Println("UpdateCetusPools", "http.Get Error", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Println("UpdateCetusPools", "io.ReadAll Error", err)
			return
		}

		//logger.Println("UpdateCetusPools", "body", string(body))

		var stats CetusStatsPools
		err = json.Unmarshal(body, &stats)
		if err != nil {
			logger.Println("UpdateCetusPools", "json.Unmarshal Error", err)
			return
		}
		//logger.Println("UpdateCetusPools", "stats", stats)
		for _, item := range stats.Data.LpList {
			items = append(items, &CetusPoolsItem{
				Coins:        item.Symbol,
				CoinA:        item.CoinA.Name,
				CoinB:        item.CoinB.Name,
				Price:        item.Price,
				PureTvlInUsd: item.PureTvlInUsd,
				VolInUsd24H:  item.VolInUsd24H,
			})
		}
		logger.Println("UpdateCetusPools", "len(items)", len(items))
	}

	//table.Count = len(stats.Data.LpList)
	c.table = &CetusPoolsTable{
		Count: len(items),
		Items: items,
	}
}
