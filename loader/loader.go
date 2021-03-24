package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type EmissionYear struct {
	IsoCode           string `json:"iso_code"`
	Country           string `json:"country"`
	Year              string `json:"year"`
	Co2               string `json:"co2"`
	Co2Growth         string `json:"co2_growth_prct"`
	Co2GrowthAbs      string `json:"co2_growth_abs"`
	ConsCo2           string `json:"consumption_co2"`
	TradeCo2          string `json:"trade_co2"`
	TradeCo2Share     string `json:"trade_co2_share"`
	Co2PC             string `json:"co2_per_capita"`
	ConsCo2PC         string `json:"consumption_co2_per_capita"`
	ShareGbCo2        string `json:"share_global_co2"`
	ShareGbCumCo2     string `json:"share_global_cumulative_co2"`
	CumCo2            string `json:"cumulative_co2"`
	Co2PGpd           string `json:"co2_per_gdp"`
	ConsCo2PGpd       string `json:"consumption_co2_per_gdp"`
	Co2PUEnergy       string `json:"co2_per_unit_energy"`
	CementCo2         string `json:"cement_co2"`
	CoalCo2           string `json:"coal_co2"`
	FlaringCo2        string `json:"flaring_co2"`
	GasCo2            string `json:"gas_co2"`
	OilCo2            string `json:"oil_co2"`
	OtherCo2          string `json:"other_industry_co2"`
	CementCo2PC       string `json:"cement_co2_per_capita"`
	CoalCo2PC         string `json:"coal_co2_per_capita"`
	FlaringCo2PC      string `json:"flaring_co2_per_capita"`
	GasCo2PC          string `json:"gas_co2_per_capita"`
	OilCo2PC          string `json:"oil_co2_per_capita"`
	OtherCo2PC        string `json:"other_co2_per_capita"`
	ShareGbCoalCo2    string `json:"share_global_coal_co2"`
	ShareGbOilCo2     string `json:"share_global_oil_co2"`
	ShareGbGasCo2     string `json:"share_global_gas_co2"`
	ShareGbFlaringCo2 string `json:"share_global_flaring_co2"`
	ShareGbCementCo2  string `json:"share_global_cement_co2"`
	CumCoalCo2        string `json:"cumulative_coal_co2"`
	CumOilCo2         string `json:"cumulative_oil_co2"`
	CumGasCo2         string `json:"cumulative_gas_co2"`
	CumFlaringCo2     string `json:"cumulative_flaring_co2"`
	CumCementCo2      string `json:"cumulative_cement_co2"`
	ShareGbCumCoal    string `json:"share_global_cumulative_coal_co2"`
	ShareGbCumOil     string `json:"share_global_cumulative_oil_co2"`
	ShareGbCumGas     string `json:"share_global_cumulative_gas_co2"`
	ShareGbCumFlaring string `json:"share_global_cumulative_flaring_co2"`
	ShareGbCumCement  string `json:"share_global_cumulative_cement_co2"`
	TotalGhg          string `json:"total_ghg"`
	GhgPC             string `json:"ghg_per_capita"`
	Methane           string `json:"methane"`
	MethanePC         string `json:"methane_per_capita"`
	NitrousOxide      string `json:"nitrous_oxide"`
	NitrousOxidePC    string `json:"nitrous_oxide_per_capita"`
	PrimaryEnergyCons string `json:"primary_energy_consumption"`
	EnergyPC          string `json:"energy_per_capita"`
	EnergyPGpd        string `json:"energy_per_gdp"`
	Population        string `json:"population"`
	Gpd               string `json:"gdp"`
}

func LoadData(r io.Reader) *[]*EmissionYear {
	reader := csv.NewReader(r)

	ret := make([]*EmissionYear, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		year := &EmissionYear{
			IsoCode:           row[0],
			Country:           row[1],
			Year:              row[2],
			Co2:               row[3],
			Co2Growth:         row[4],
			Co2GrowthAbs:      row[5],
			ConsCo2:           row[6],
			TradeCo2:          row[7],
			TradeCo2Share:     row[8],
			Co2PC:             row[9],
			ConsCo2PC:         row[10],
			ShareGbCo2:        row[11],
			ShareGbCumCo2:     row[12],
			CumCo2:            row[13],
			Co2PGpd:           row[14],
			ConsCo2PGpd:       row[15],
			Co2PUEnergy:       row[16],
			CementCo2:         row[17],
			CoalCo2:           row[18],
			FlaringCo2:        row[19],
			GasCo2:            row[20],
			OilCo2:            row[21],
			OtherCo2:          row[22],
			CementCo2PC:       row[23],
			CoalCo2PC:         row[24],
			FlaringCo2PC:      row[25],
			GasCo2PC:          row[26],
			OilCo2PC:          row[27],
			OtherCo2PC:        row[28],
			ShareGbCoalCo2:    row[29],
			ShareGbOilCo2:     row[30],
			ShareGbGasCo2:     row[31],
			ShareGbFlaringCo2: row[32],
			ShareGbCementCo2:  row[33],
			CumCoalCo2:        row[34],
			CumOilCo2:         row[35],
			CumGasCo2:         row[36],
			CumFlaringCo2:     row[37],
			CumCementCo2:      row[38],
			ShareGbCumCoal:    row[39],
			ShareGbCumOil:     row[40],
			ShareGbCumGas:     row[41],
			ShareGbCumFlaring: row[42],
			ShareGbCumCement:  row[43],
			TotalGhg:          row[44],
			GhgPC:             row[45],
			Methane:           row[46],
			MethanePC:         row[47],
			NitrousOxide:      row[48],
			NitrousOxidePC:    row[49],
			PrimaryEnergyCons: row[50],
			EnergyPC:          row[51],
			EnergyPGpd:        row[52],
			Population:        row[53],
			Gpd:               row[54],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, year)
	}

	return &ret
}
