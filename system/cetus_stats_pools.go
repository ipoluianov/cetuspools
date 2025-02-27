package system

type CetusStatsPools struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total  int `json:"total"`
		LpList []struct {
			Symbol            string      `json:"symbol"`
			Name              string      `json:"name"`
			Decimals          int         `json:"decimals"`
			Fee               string      `json:"fee"`
			TickSpacing       string      `json:"tick_spacing"`
			PoolType          string      `json:"pool_type"`
			Address           string      `json:"address"`
			CoinAAddress      string      `json:"coin_a_address"`
			CoinBAddress      string      `json:"coin_b_address"`
			ProjectUrl        string      `json:"project_url"`
			IsDisplayRewarder bool        `json:"is_display_rewarder"`
			IsClosed          bool        `json:"is_closed"`
			RewarderDisplay1  bool        `json:"rewarder_display1"`
			RewarderDisplay2  bool        `json:"rewarder_display2"`
			RewarderDisplay3  bool        `json:"rewarder_display3"`
			RewarderDisplay4  bool        `json:"rewarder_display4"`
			RewarderDisplay5  bool        `json:"rewarder_display5"`
			Labels            interface{} `json:"labels"`
			CoinA             struct {
				Name        string        `json:"name"`
				Symbol      string        `json:"symbol"`
				Decimals    int           `json:"decimals"`
				Address     string        `json:"address"`
				Balance     string        `json:"balance"`
				LogoUrl     string        `json:"logo_url"`
				CoingeckoId string        `json:"coingecko_id"`
				ProjectUrl  string        `json:"project_url"`
				Labels      []interface{} `json:"labels"`
				IsTrusted   bool          `json:"is_trusted"`
			} `json:"coin_a"`
			CoinB struct {
				Name        string        `json:"name"`
				Symbol      string        `json:"symbol"`
				Decimals    int           `json:"decimals"`
				Address     string        `json:"address"`
				Balance     string        `json:"balance"`
				LogoUrl     string        `json:"logo_url"`
				CoingeckoId string        `json:"coingecko_id"`
				ProjectUrl  string        `json:"project_url"`
				Labels      []interface{} `json:"labels"`
				IsTrusted   bool          `json:"is_trusted"`
			} `json:"coin_b"`
			Price            string      `json:"price"`
			RewarderUsd      []string    `json:"rewarder_usd"`
			RewarderApr      []string    `json:"rewarder_apr"`
			IsForward        bool        `json:"is_forward"`
			PriceRangeConfig interface{} `json:"price_range_config"`
			Object           struct {
				CoinA            int    `json:"coin_a"`
				CoinB            int    `json:"coin_b"`
				TickSpacing      int    `json:"tick_spacing"`
				Liquidity        string `json:"liquidity"`
				CurrentSqrtPrice string `json:"current_sqrt_price"`
				RewarderManager  struct {
					Fields struct {
						Rewarders []struct {
							Fields struct {
								RewardCoin struct {
									Fields struct {
										Name string `json:"name"`
									} `json:"fields"`
								} `json:"reward_coin"`
								EmissionsPerSecond string `json:"emissions_per_second"`
								GrowthGlobal       string `json:"growth_global"`
							} `json:"fields"`
						} `json:"rewarders"`
						PointsReleased     string `json:"points_released"`
						PointsGrowthGlobal string `json:"points_growth_global"`
						LastUpdatedTime    int    `json:"last_updated_time"`
					} `json:"fields"`
				} `json:"rewarder_manager"`
				IsPause bool `json:"is_pause"`
				Index   int  `json:"index"`
			} `json:"object"`
			Category      string `json:"category"`
			IsVaults      bool   `json:"is_vaults"`
			StableFarming struct {
				StableFarmingPool  string      `json:"stable_farming_pool"`
				IsShowRewarder     bool        `json:"is_show_rewarder"`
				ShowRewarder1      bool        `json:"show_rewarder_1"`
				ShowRewarder2      bool        `json:"show_rewarder_2"`
				ShowRewarder3      bool        `json:"show_rewarder_3"`
				Tvl                string      `json:"tvl"`
				Apr                string      `json:"apr"`
				EffectiveTickLower int         `json:"effective_tick_lower"`
				EffectiveTickUpper int         `json:"effective_tick_upper"`
				TotalWrappedAmount string      `json:"total_wrapped_amount"`
				StableRewarder     interface{} `json:"stable_rewarder"`
			} `json:"stable_farming"`
			Vaults       []string `json:"vaults"`
			PureTvlInUsd string   `json:"pure_tvl_in_usd"`
			VolInUsd24H  string   `json:"vol_in_usd_24h"`
			Fee24H       string   `json:"fee_24_h"`
			TotalApr     string   `json:"total_apr"`
			Apr          struct {
				FeeApr24H string `json:"fee_apr_24h"`
			} `json:"apr"`
			Extensions   interface{} `json:"extensions"`
			HasMining    bool        `json:"has_mining"`
			HasFarming   bool        `json:"has_farming"`
			NoIncentives bool        `json:"no_incentives"`
		} `json:"lp_list"`
	} `json:"data"`
}
