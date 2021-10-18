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

type DateFormatType int

const (
	DateFormatISO DateFormatType = iota
	DateFormatUnix
)

func (df DateFormatType) String() string {
	return []string{"iso", "unix"}[df]
}

type OddsFormatType int

const (
	OddsFormatAmerican OddsFormatType = iota
	OddsFormatDecimal
)

func (fo OddsFormatType) String() string {
	return []string{"american", "decimal"}[fo]
}
