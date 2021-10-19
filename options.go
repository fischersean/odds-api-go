package oddsapi

type RegionType int

const (
	RegionUS RegionType = iota
	RegionEU
	RegionUK
	RegionAU
)

func (r RegionType) String() string {
	return []string{"us", "eu", "uk", "au"}[r]
}

type MarketType int

const (
	MarketMoneyline MarketType = iota
	MarketMoneylineLay
	MarketSpread
	MarketTotal
	MarketOutright
)

func (m MarketType) String() string {
	return []string{"h2h", "h2h_lay", "spreads", "total", "outrights"}[m]
}

type OddsFormatType int

const (
	OddsFormatDecimal OddsFormatType = iota
	OddsFormatAmerican
)

func (fo OddsFormatType) String() string {
	return []string{"decimal", "american"}[fo]
}
