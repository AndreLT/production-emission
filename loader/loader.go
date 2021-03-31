package loader

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type EmissionYear struct {
	IsoCode           string  `json:"iso_code"`
	Country           string  `json:"country"`
	Year              int     `json:"year"`
	Co2               float64 `json:"co2"`
	Co2Growth         float64 `json:"co2_growth_prct"`
	Co2GrowthAbs      float64 `json:"co2_growth_abs"`
	ConsCo2           float64 `json:"consumption_co2"`
	TradeCo2          float64 `json:"trade_co2"`
	TradeCo2Share     float64 `json:"trade_co2_share"`
	Co2PC             float64 `json:"co2_per_capita"`
	ConsCo2PC         float64 `json:"consumption_co2_per_capita"`
	ShareGbCo2        float64 `json:"share_global_co2"`
	ShareGbCumCo2     float64 `json:"share_global_cumulative_co2"`
	CumCo2            float64 `json:"cumulative_co2"`
	Co2PGpd           float64 `json:"co2_per_gdp"`
	ConsCo2PGpd       float64 `json:"consumption_co2_per_gdp"`
	Co2PUEnergy       float64 `json:"co2_per_unit_energy"`
	CementCo2         float64 `json:"cement_co2"`
	CoalCo2           float64 `json:"coal_co2"`
	FlaringCo2        float64 `json:"flaring_co2"`
	GasCo2            float64 `json:"gas_co2"`
	OilCo2            float64 `json:"oil_co2"`
	OtherCo2          float64 `json:"other_industry_co2"`
	CementCo2PC       float64 `json:"cement_co2_per_capita"`
	CoalCo2PC         float64 `json:"coal_co2_per_capita"`
	FlaringCo2PC      float64 `json:"flaring_co2_per_capita"`
	GasCo2PC          float64 `json:"gas_co2_per_capita"`
	OilCo2PC          float64 `json:"oil_co2_per_capita"`
	OtherCo2PC        float64 `json:"other_co2_per_capita"`
	ShareGbCoalCo2    float64 `json:"share_global_coal_co2"`
	ShareGbOilCo2     float64 `json:"share_global_oil_co2"`
	ShareGbGasCo2     float64 `json:"share_global_gas_co2"`
	ShareGbFlaringCo2 float64 `json:"share_global_flaring_co2"`
	ShareGbCementCo2  float64 `json:"share_global_cement_co2"`
	CumCoalCo2        float64 `json:"cumulative_coal_co2"`
	CumOilCo2         float64 `json:"cumulative_oil_co2"`
	CumGasCo2         float64 `json:"cumulative_gas_co2"`
	CumFlaringCo2     float64 `json:"cumulative_flaring_co2"`
	CumCementCo2      float64 `json:"cumulative_cement_co2"`
	ShareGbCumCoal    float64 `json:"share_global_cumulative_coal_co2"`
	ShareGbCumOil     float64 `json:"share_global_cumulative_oil_co2"`
	ShareGbCumGas     float64 `json:"share_global_cumulative_gas_co2"`
	ShareGbCumFlaring float64 `json:"share_global_cumulative_flaring_co2"`
	ShareGbCumCement  float64 `json:"share_global_cumulative_cement_co2"`
	TotalGhg          float64 `json:"total_ghg"`
	GhgPC             float64 `json:"ghg_per_capita"`
	Methane           float64 `json:"methane"`
	MethanePC         float64 `json:"methane_per_capita"`
	NitrousOxide      float64 `json:"nitrous_oxide"`
	NitrousOxidePC    float64 `json:"nitrous_oxide_per_capita"`
	PrimaryEnergyCons float64 `json:"primary_energy_consumption"`
	EnergyPC          float64 `json:"energy_per_capita"`
	EnergyPGpd        float64 `json:"energy_per_gdp"`
	Population        float64 `json:"population"`
	Gpd               float64 `json:"gdp"`
}

func ConvertToInt(s string) int {
	if s == "" {
		return -1
	} else {
		converted, _ := strconv.Atoi(s)
		return converted
	}
}

func ConvertToFloat(s string) float64 {
	if s == "" {
		return -1
	} else {
		converted, _ := strconv.ParseFloat(s, 64)
		return converted
	}
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
			Year:              ConvertToInt(row[2]),
			Co2:               ConvertToFloat(row[3]),
			Co2Growth:         ConvertToFloat(row[4]),
			Co2GrowthAbs:      ConvertToFloat(row[5]),
			ConsCo2:           ConvertToFloat(row[6]),
			TradeCo2:          ConvertToFloat(row[7]),
			TradeCo2Share:     ConvertToFloat(row[8]),
			Co2PC:             ConvertToFloat(row[9]),
			ConsCo2PC:         ConvertToFloat(row[10]),
			ShareGbCo2:        ConvertToFloat(row[11]),
			ShareGbCumCo2:     ConvertToFloat(row[12]),
			CumCo2:            ConvertToFloat(row[13]),
			Co2PGpd:           ConvertToFloat(row[14]),
			ConsCo2PGpd:       ConvertToFloat(row[15]),
			Co2PUEnergy:       ConvertToFloat(row[16]),
			CementCo2:         ConvertToFloat(row[17]),
			CoalCo2:           ConvertToFloat(row[18]),
			FlaringCo2:        ConvertToFloat(row[19]),
			GasCo2:            ConvertToFloat(row[20]),
			OilCo2:            ConvertToFloat(row[21]),
			OtherCo2:          ConvertToFloat(row[22]),
			CementCo2PC:       ConvertToFloat(row[23]),
			CoalCo2PC:         ConvertToFloat(row[24]),
			FlaringCo2PC:      ConvertToFloat(row[25]),
			GasCo2PC:          ConvertToFloat(row[26]),
			OilCo2PC:          ConvertToFloat(row[27]),
			OtherCo2PC:        ConvertToFloat(row[28]),
			ShareGbCoalCo2:    ConvertToFloat(row[29]),
			ShareGbOilCo2:     ConvertToFloat(row[30]),
			ShareGbGasCo2:     ConvertToFloat(row[31]),
			ShareGbFlaringCo2: ConvertToFloat(row[32]),
			ShareGbCementCo2:  ConvertToFloat(row[33]),
			CumCoalCo2:        ConvertToFloat(row[34]),
			CumOilCo2:         ConvertToFloat(row[35]),
			CumGasCo2:         ConvertToFloat(row[36]),
			CumFlaringCo2:     ConvertToFloat(row[37]),
			CumCementCo2:      ConvertToFloat(row[38]),
			ShareGbCumCoal:    ConvertToFloat(row[39]),
			ShareGbCumOil:     ConvertToFloat(row[40]),
			ShareGbCumGas:     ConvertToFloat(row[41]),
			ShareGbCumFlaring: ConvertToFloat(row[42]),
			ShareGbCumCement:  ConvertToFloat(row[43]),
			TotalGhg:          ConvertToFloat(row[44]),
			GhgPC:             ConvertToFloat(row[45]),
			Methane:           ConvertToFloat(row[46]),
			MethanePC:         ConvertToFloat(row[47]),
			NitrousOxide:      ConvertToFloat(row[48]),
			NitrousOxidePC:    ConvertToFloat(row[49]),
			PrimaryEnergyCons: ConvertToFloat(row[50]),
			EnergyPC:          ConvertToFloat(row[51]),
			EnergyPGpd:        ConvertToFloat(row[52]),
			Population:        ConvertToFloat(row[53]),
			Gpd:               ConvertToFloat(row[54]),
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, year)
	}

	return &ret
}
