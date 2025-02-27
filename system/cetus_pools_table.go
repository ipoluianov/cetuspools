package system

type CetusPoolsTable struct {
	Count int               `json:"count"`
	Items []*CetusPoolsItem `json:"items"`
}

type CetusPoolsItem struct {
	Coins        string `json:"coins"`
	CoinA        string `json:"coin_a"`
	CoinB        string `json:"coin_b"`
	Price        string `json:"price"`
	PureTvlInUsd string `json:"pure_tvl_in_usd"`
	VolInUsd24H  string `json:"vol_in_usd_24h"`
}

func (c *CetusPoolsTable) Add(item *CetusPoolsItem) {
	c.Items = append(c.Items, item)
}
