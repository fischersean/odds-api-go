package oddsapi

import (
	"testing"
)

func TestRegionTypes(t *testing.T) {
	if RegionUS.String() != "us" {
		t.Error("US region string does not match expected value")
	}
	if RegionEU.String() != "eu" {
		t.Error("EU region string does not match expected value")
	}
	if RegionUK.String() != "uk" {
		t.Error("UK region string does not match expected value")
	}
	if RegionAU.String() != "au" {
		t.Error("AU region string does not match expected value")
	}
}

func TestMarketTypes(t *testing.T) {
	if MarketMoneyline.String() != "h2h" {
		t.Error("Moneyline market string does not match expected value")
	}
	if MarketMoneylineLay.String() != "h2h_lay" {
		t.Error("Moneyline Lay market string does not match expected value")
	}
	if MarketSpread.String() != "spreads" {
		t.Error("Spread market string does not match expected value")
	}
	if MarketTotal.String() != "total" {
		t.Error("Total market string does not match expected value")
	}
	if MarketOutright.String() != "outrights" {
		t.Error("Outright market string does not match expected value")
	}
}

func TestOddsFormateTypes(t *testing.T) {
	if OddsFormatAmerican.String() != "american" {
		t.Error("American odds format string does not match expected value")
	}
	if OddsFormatDecimal.String() != "decimal" {
		t.Error("Decimal odds format string does not match expected value")
	}
}
