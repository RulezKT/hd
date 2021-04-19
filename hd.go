//HD rrelated calculations
package hd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/RulezKT/mathfn"
	"github.com/RulezKT/structs"
)

const (

	//Astronomical Unit
	AU = 0.1495978707e9 // km 149597870.7

	// Здесь мы как начальное значение ставим eps = 23°26'21,448" градуса согласно теории вогнутой Земли
	// double const RAD_TO_DEG = 5.7295779513082320877e1;
	// Obliquity of the ecliptic  = 23°26'21,448"  - на 1 января 2000 года = 23.43929111111111
	// 23.43929111111111/5.7295779513082320877e1 = 0.4090928042223289
	MED_EPS = 0.4090928042223289

	SSB       = 0
	MERCURY   = 1 // 7,01° (относительно эклиптики)
	VENUS     = 2 // 3,39458° (относительно эклиптики)
	EARTH     = 3
	MARS      = 4 // 1,85061° (относительно эклиптики)
	JUPITER   = 5 // 1,304° (относительно эклиптики)
	SATURN    = 6 // 2,485 240° (относительно эклиптики)
	URANUS    = 7 // 0,772556° (относительно эклиптики)
	NEPTUNE   = 8 // 1,767975° (относительно эклиптики)
	PLUTO     = 9 // 17°,14 (относительно эклиптики)
	SUN       = 10
	MOON      = 11 // 5,14° (относительно эклиптики)
	NORTHNODE = 12
	SOUTHNODE = 13
	HIRON     = 14

	HEAD   = 0
	AJNA   = 1
	THROAT = 2
	G      = 3
	SACRAL = 4
	ROOT   = 5
	EGO    = 6
	SPLEEN = 7
	EMO    = 8
)

const (
	// размер одной линии в десятичных градусах
	oneLineInDec float64 = 0.9375

	// размер одного цвета в десятичных градусах
	oneColorInDec float64 = 0.15625

	// размер одного тона в десятичных градусах
	oneToneInDec float64 = 0.026041666666666668

	// размер одной базы в десятичных градусах
	oneBaseInDec float64 = 0.005208333333333334
)

//гексаграммы с  диапазонами градусов на космограмме, первое значение включая, второе не включая
type HexRangeRAD struct {
	startDegree float64
	endDegree   float64
}

func CalcHexLineColorToneBase(longitude float64) structs.HdStructure {

	var hexSortByDeg = map[int]HexRangeRAD{

		// from 3.875 to 9.49
		17: {3.875, 9.5},

		21: {9.5, 15.125},

		51: {15.125, 20.75},

		42: {20.75, 26.375},

		3: {26.375, 32.0},

		27: {32.0, 37.625},

		24: {37.625, 43.25},

		2: {43.25, 48.875},

		23: {48.875, 54.5},

		8: {54.5, 60.125},

		20: {60.125, 65.75},

		16: {65.75, 71.375},

		35: {71.375, 77.0},

		45: {77.0, 82.625},

		12: {82.625, 88.255},

		15: {88.25, 93.875},

		52: {93.875, 99.5},

		39: {99.5, 105.125},

		53: {105.125, 110.75},

		62: {110.75, 116.375},

		56: {116.375, 122.0},

		31: {122.0, 127.625},

		33: {127.625, 133.25},

		7: {133.25, 138.875},

		4: {138.875, 144.5},

		29: {144.5, 150.125},

		59: {150.125, 155.75},

		40: {155.75, 161.375},

		64: {161.375, 167.0},

		47: {167.0, 172.625},

		6: {172.625, 178.25},

		46: {178.25, 183.875},

		18: {183.875, 189.5},

		48: {189.5, 195.125},

		57: {195.125, 200.75},

		32: {200.75, 206.375},

		50: {206.375, 212.0},

		28: {212.0, 217.625},

		44: {217.625, 223.25},

		1: {223.25, 228.875},

		43: {228.875, 234.5},

		14: {234.5, 240.125},

		34: {240.125, 245.75},

		9: {245.75, 251.375},

		5: {251.375, 257.0},

		26: {257.0, 262.625},

		11: {262.625, 268.25},

		10: {268.25, 273.875},

		58: {273.875, 279.5},

		38: {279.5, 285.125},

		54: {285.125, 290.75},

		61: {290.75, 296.375},

		60: {296.375, 302.0},

		41: {302.0, 307.625},

		19: {307.625, 313.25},

		13: {313.25, 318.875},

		49: {318.875, 324.5},

		30: {324.5, 330.125},

		55: {330.125, 335.75},

		37: {335.75, 341.375},

		63: {341.375, 347.0},

		22: {347.0, 352.625},

		36: {352.625, 358.25},

		25: {358.25, 3.875},
	}

	var hex int

	var line float64

	var color float64

	var tone float64

	var base float64

	var number_of_passed_degrees float64

	longitude *= mathfn.RAD_TO_DEG

	longitude = mathfn.Convert_to_0_360_DEG(longitude)
	// console.log(`longitude = ${longitude}`);

	for key, entry := range hexSortByDeg {

		//25 - последняя гексаграмма и там начальное значение больше конечного, круг замыкается
		if (longitude >= entry.startDegree) &&
			(longitude < entry.endDegree || key == 25) {

			hex = key
			//System.out.println("hex = " + hex + "===" + longitude);
			number_of_passed_degrees = longitude - entry.startDegree

			line = number_of_passed_degrees / oneLineInDec

			color = (number_of_passed_degrees - float64(int(line))*oneLineInDec) / oneColorInDec
			// console.log(`Цвет   = ${math.ceil(color)}, цвет завершена на ${(color - math.trunc(color))*100}  процентов`);

			tone = (number_of_passed_degrees - float64(int(line))*oneLineInDec - float64(int(color))*oneColorInDec) / oneToneInDec
			// console.log(`Тон  =  ${math.ceil(tone)},тон завершен на ${(tone - math.trunc(tone))*100} процентов`);

			base = (number_of_passed_degrees - float64(int(line))*oneLineInDec - float64(int(color))*oneColorInDec - float64(int(tone))*oneToneInDec) / oneBaseInDec

			break

		} else if longitude >= 0 && longitude < 3.875 {
			// здесь будут находиться все значения в градусах от 0 до 3.875 (не включительно) [25, [358.25, 3.875]]
			// так как круг в 360 градусов завершается и начинается новый

			hex = 25

			number_of_passed_degrees = longitude - 358.25 + 360

			line = number_of_passed_degrees / oneLineInDec
			// console.log(`Линия  = ${math.ceil(line)}, линия завершена на ${(line - math.trunc(line))*100}  процентов`);

			color = (number_of_passed_degrees - float64(int(line))*oneLineInDec) / oneColorInDec
			// console.log(`Цвет   = ${math.ceil(color)}, цвет завершена на ${(color - math.trunc(color))*100}  процентов`);

			tone = (number_of_passed_degrees - float64(int(line))*oneLineInDec - float64(int(color))*oneColorInDec) / oneToneInDec
			// console.log(`Тон  =  ${math.ceil(tone)},тон завершен на ${(tone - math.trunc(tone))*100} процентов`);

			base = (number_of_passed_degrees - float64(int(line))*oneLineInDec - float64(int(color))*oneColorInDec - float64(int(tone))*oneToneInDec) / oneBaseInDec

			break
		}
	}

	if line > 6 || color > 6 || tone > 6 || base > 5 {
		fmt.Println("error in calc_hex_line_color_tone_base")
	}

	return structs.HdStructure{Hex: hex, Line: line, Color: color, Tone: tone, Base: base, NumberOfPassedDegrees: number_of_passed_degrees}
}

func GatesChannelsCenters(info *structs.CdInfo) {

	// this.formula.personality[key].hex] и this.formula.design[key].hex показывают какие ворота определены
	// в channels[number] = [red, black] определяем канал и как он образован red/black/both для двух ворот
	// number - по таблице

	// инициализируем ворота, отсчет от 1
	//	var gates [65]string

	//	var channels [37]structs.Channel

	// инициализируем центры
	//	var centers structs.Centers

	//0 не берем, это SSB
	for i := 1; i < len(info.HdInfo.Design.Planets.Planet); i++ {

		info.HdInfo.Gates[info.HdInfo.Design.Planets.Planet[i].Hex].Des++
		info.HdInfo.Gates[info.HdInfo.Design.Planets.Planet[i].Hex].Defined = true

	}

	for i := 1; i < len(info.HdInfo.Personality.Planets.Planet); i++ {

		info.HdInfo.Gates[info.HdInfo.Personality.Planets.Planet[i].Hex].Pers++
		info.HdInfo.Gates[info.HdInfo.Personality.Planets.Planet[i].Hex].Defined = true

	}

	gates := &info.HdInfo.Gates
	//?  info.HdInfo.Centers.Init()
	centers := &info.HdInfo.Centers

	channels := &info.HdInfo.Channels

	// HEAD && AJNA

	// 1 - 64-47
	if gates[64].Defined && gates[47].Defined {
		centers.Center["Head"] = true
		centers.Center["Ajna"] = true

		channels[1].FirstGate = gates[64]
		channels[1].SecondGate = gates[47]
		channels[1].Defined = true
	}

	// 2 - 61-24
	if gates[61].Defined && gates[24].Defined {
		centers.Center["Head"] = true
		centers.Center["Ajna"] = true

		channels[2].FirstGate = gates[61]
		channels[2].SecondGate = gates[24]
		channels[2].Defined = true
	}

	// 3 - 63-4
	if gates[63].Defined && gates[4].Defined {
		centers.Center["Head"] = true
		centers.Center["Ajna"] = true

		channels[3].FirstGate = gates[63]
		channels[3].SecondGate = gates[4]
		channels[3].Defined = true
	}

	// AJNA && THROAT

	// 4 - 17-62
	if gates[17].Defined && gates[62].Defined {
		centers.Center["Throat"] = true
		centers.Center["Ajna"] = true

		channels[4].FirstGate = gates[17]
		channels[4].SecondGate = gates[62]
		channels[4].Defined = true
	}

	// 5 - 43-23
	if gates[43].Defined && gates[23].Defined {
		centers.Center["Throat"] = true
		centers.Center["Ajna"] = true

		channels[5].FirstGate = gates[43]
		channels[5].SecondGate = gates[23]
		channels[5].Defined = true
	}

	// 6 - 11-56
	if gates[11].Defined && gates[56].Defined {
		centers.Center["Throat"] = true
		centers.Center["Ajna"] = true

		channels[6].FirstGate = gates[11]
		channels[6].SecondGate = gates[56]
		channels[6].Defined = true
	}

	// THROAT && G

	// 14- 7-31
	if gates[31].Defined && gates[7].Defined {
		centers.Center["Throat"] = true
		centers.Center["G"] = true

		channels[14].FirstGate = gates[7]
		channels[14].SecondGate = gates[31]
		channels[14].Defined = true
	}

	// 15- 1-8
	if gates[8].Defined && gates[1].Defined {
		centers.Center["Throat"] = true
		centers.Center["G"] = true

		channels[15].FirstGate = gates[1]
		channels[15].SecondGate = gates[8]
		channels[15].Defined = true
	}

	// 16- 13-33
	if gates[33].Defined && gates[13].Defined {
		centers.Center["Throat"] = true
		centers.Center["G"] = true

		channels[16].FirstGate = gates[13]
		channels[16].SecondGate = gates[33]
		channels[16].Defined = true
	}

	// G && SACRAL

	// 20- 5-15
	if gates[15].Defined && gates[5].Defined {
		centers.Center["G"] = true
		centers.Center["Sacral"] = true

		channels[20].FirstGate = gates[5]
		channels[20].SecondGate = gates[15]
		channels[20].Defined = true
	}

	// 21- 14-2
	if gates[2].Defined && gates[14].Defined {
		centers.Center["G"] = true
		centers.Center["Sacral"] = true

		channels[21].FirstGate = gates[14]
		channels[21].SecondGate = gates[2]
		channels[21].Defined = true
	}

	// 22- 29-46
	if gates[46].Defined && gates[29].Defined {
		centers.Center["G"] = true
		centers.Center["Sacral"] = true

		channels[22].FirstGate = gates[29]
		channels[22].SecondGate = gates[46]
	}

	// SACRAL && ROOT

	// 31- 53-42
	if gates[42].Defined && gates[53].Defined {
		centers.Center["Sacral"] = true
		centers.Center["Root"] = true

		channels[31].FirstGate = gates[53]
		channels[31].SecondGate = gates[42]
		channels[31].Defined = true
	}

	// 32- 60-3
	if gates[3].Defined && gates[60].Defined {
		centers.Center["Sacral"] = true
		centers.Center["Root"] = true

		channels[32].FirstGate = gates[60]
		channels[32].SecondGate = gates[3]
		channels[32].Defined = true
	}

	// 33- 52-9
	if gates[9].Defined && gates[52].Defined {
		centers.Center["Sacral"] = true
		centers.Center["Root"] = true

		channels[33].FirstGate = gates[52]
		channels[33].SecondGate = gates[9]
		channels[33].Defined = true
	}

	// ROOT && EMO

	// 34- 19-49
	if gates[19].Defined && gates[49].Defined {
		centers.Center["Root"] = true
		centers.Center["Emo"] = true

		channels[34].FirstGate = gates[19]
		channels[34].SecondGate = gates[49]
		channels[34].Defined = true
	}

	// 35- 39-55
	if gates[39].Defined && gates[55].Defined {
		centers.Center["Root"] = true
		centers.Center["Emo"] = true

		channels[35].FirstGate = gates[39]
		channels[35].SecondGate = gates[55]
		channels[35].Defined = true
	}

	// 36- 41-30
	if gates[41].Defined && gates[30].Defined {
		centers.Center["Root"] = true
		centers.Center["Emo"] = true

		channels[36].FirstGate = gates[41]
		channels[36].SecondGate = gates[30]
		channels[36].Defined = true
	}

	// ROOT && SPLEEN
	// 30- 58-18
	if gates[18].Defined && gates[58].Defined {
		centers.Center["Root"] = true
		centers.Center["Spleen"] = true

		channels[30].FirstGate = gates[58]
		channels[30].SecondGate = gates[18]
		channels[30].Defined = true
	}

	// 29- 38-28
	if gates[28].Defined && gates[38].Defined {
		centers.Center["Root"] = true
		centers.Center["Spleen"] = true

		channels[29].FirstGate = gates[38]
		channels[29].SecondGate = gates[28]
		channels[29].Defined = true
	}

	// 28- 54-32
	if gates[32].Defined && gates[54].Defined {
		centers.Center["Root"] = true
		centers.Center["Spleen"] = true

		channels[28].FirstGate = gates[54]
		channels[28].SecondGate = gates[32]
		channels[28].Defined = true
	}

	// EMO && SACRAL, EGO, THROAT

	// 26- 59-6
	if gates[59].Defined && gates[6].Defined {
		centers.Center["Emo"] = true
		centers.Center["Sacral"] = true

		channels[26].FirstGate = gates[59]
		channels[26].SecondGate = gates[6]
		channels[26].Defined = true
	}

	// 27- 37-40
	if gates[37].Defined && gates[40].Defined {
		centers.Center["Emo"] = true
		centers.Center["Ego"] = true

		channels[27].FirstGate = gates[37]
		channels[27].SecondGate = gates[40]
		channels[27].Defined = true
	}

	// 18- 22-12
	if gates[22].Defined && gates[12].Defined {
		centers.Center["Emo"] = true
		centers.Center["Throat"] = true

		channels[18].FirstGate = gates[22]
		channels[18].SecondGate = gates[12]
		channels[18].Defined = true
	}

	// 19- 36-35
	if gates[35].Defined && gates[36].Defined {
		centers.Center["Emo"] = true
		centers.Center["Throat"] = true

		channels[19].FirstGate = gates[36]
		channels[19].SecondGate = gates[35]
		channels[19].Defined = true
	}

	// EGO && SPLEEN, G, THROAT

	// 24- 44-26
	if gates[44].Defined && gates[26].Defined {
		centers.Center["Ego"] = true
		centers.Center["Spleen"] = true

		channels[24].FirstGate = gates[44]
		channels[24].SecondGate = gates[26]
		channels[24].Defined = true
	}

	// 23- 51-25
	if gates[51].Defined && gates[25].Defined {
		centers.Center["Ego"] = true
		centers.Center["G"] = true

		channels[23].FirstGate = gates[51]
		channels[23].SecondGate = gates[25]
		channels[23].Defined = true
	}

	// 17- 21-45
	if gates[21].Defined && gates[45].Defined {
		centers.Center["Ego"] = true
		centers.Center["Throat"] = true

		channels[17].FirstGate = gates[21]
		channels[17].SecondGate = gates[45]
		channels[17].Defined = true
	}

	// SACRAL && SPLEEN

	// 25- 27-50
	if gates[27].Defined && gates[50].Defined {
		centers.Center["Spleen"] = true
		centers.Center["Sacral"] = true

		channels[25].FirstGate = gates[27]
		channels[25].SecondGate = gates[50]
		channels[25].Defined = true
	}

	// THROAT && SPLEEN

	// 7 - 48-16
	if gates[48].Defined && gates[16].Defined {
		centers.Center["Spleen"] = true
		centers.Center["Throat"] = true

		channels[7].FirstGate = gates[48]
		channels[7].SecondGate = gates[16]
		channels[7].Defined = true
	}

	// INTEGRATION

	if gates[20].Defined || gates[57].Defined || gates[10].Defined || gates[34].Defined {
		// 8 - 57-20
		if gates[20].Defined && gates[57].Defined {
			centers.Center["Spleen"] = true
			centers.Center["Throat"] = true

			channels[8].FirstGate = gates[57]
			channels[8].SecondGate = gates[20]
			channels[8].Defined = true
		}

		// 10- 10-20
		if gates[20].Defined && gates[10].Defined {
			centers.Center["G"] = true
			centers.Center["Throat"] = true

			channels[10].FirstGate = gates[10]
			channels[10].SecondGate = gates[20]
			channels[10].Defined = true
		}

		// 9 - 34-20
		if gates[20].Defined && gates[34].Defined {
			centers.Center["Sacral"] = true
			centers.Center["Throat"] = true

			channels[9].FirstGate = gates[34]
			channels[9].SecondGate = gates[20]
			channels[9].Defined = true
		}

		// 11- 57-10
		if gates[10].Defined && gates[57].Defined {
			centers.Center["Spleen"] = true
			centers.Center["G"] = true

			channels[11].FirstGate = gates[57]
			channels[11].SecondGate = gates[10]
			channels[11].Defined = true
		}

		// 12- 57-34
		if gates[34].Defined && gates[57].Defined {
			centers.Center["Spleen"] = true
			centers.Center["Sacral"] = true

			channels[12].FirstGate = gates[57]
			channels[12].SecondGate = gates[34]
			channels[12].Defined = true
		}

		// 13- 34-10
		if gates[34].Defined && gates[10].Defined {
			centers.Center["G"] = true
			centers.Center["Sacral"] = true

			channels[13].FirstGate = gates[34]
			channels[13].SecondGate = gates[10]
			channels[13].Defined = true
		}
	}

}

func Profile(info *structs.CdInfo) {

	info.HdInfo.Profile = strconv.Itoa(int(math.Ceil(info.HdInfo.Personality.Planets.Planet[SUN].Line))) + "/" + strconv.Itoa(int(math.Ceil(info.HdInfo.Design.Planets.Planet[SUN].Line)))

}

func Authority(info *structs.CdInfo) {

	var authority string

	if info.HdInfo.Centers.Center["Emo"] {
		authority = "Emo"
	} else if info.HdInfo.Centers.Center["Sacral"] {
		authority = "Sacral"
	} else if info.HdInfo.Centers.Center["Spleen"] {
		authority = "Spleen"
	} else if info.HdInfo.Centers.Center["Ego"] {
		authority = "Ego"
	} else if info.HdInfo.Centers.Center["G"] {
		authority = "Self projected"
	} else if info.HdInfo.Centers.Center["Throat"] || info.HdInfo.Centers.Center["Ajna"] || info.HdInfo.Centers.Center["Head"] {
		authority = "No inner authority"
	} else {
		authority = "Moon"
	}

	info.HdInfo.Authority = authority

}

func Variable(info *structs.CdInfo) {

	var first string
	if info.HdInfo.Personality.Planets.Planet[SUN].Tone > 3 {
		first = "R"
	} else {

		first = "L"
	}

	var second string
	if info.HdInfo.Personality.Planets.Planet[NORTHNODE].Tone > 3 {
		second = "R"
	} else {

		second = "L"
	}

	var third string
	if info.HdInfo.Design.Planets.Planet[SUN].Tone > 3 {
		third = "R"
	} else {

		third = "L"
	}

	var forth string
	if info.HdInfo.Design.Planets.Planet[NORTHNODE].Tone > 3 {
		forth = "R"
	} else {

		forth = "L"
	}

	info.HdInfo.Variable = "P" + first + second + "D" + third + forth

}

func Cross(info *structs.CdInfo) {

	info.HdInfo.Cross.First = info.HdInfo.Personality.Planets.Planet[SUN].Hex
	info.HdInfo.Cross.Second = info.HdInfo.Personality.Planets.Planet[EARTH].Hex
	info.HdInfo.Cross.Third = info.HdInfo.Design.Planets.Planet[SUN].Hex
	info.HdInfo.Cross.Forth = info.HdInfo.Design.Planets.Planet[EARTH].Hex

}

// TYPE
func HdType(info *structs.CdInfo) string {

	centers := info.HdInfo.Centers.Center
	channels := info.HdInfo.Channels

	hdType := "Reflector"

	// если хоть один центр определен - тип как минимум Проектор, если нет Рефлектор
	for _, value := range centers {
		if value {
			hdType = "Projector"
			break
		}
	}

	if hdType == "Reflector" {
		return hdType
	}

	//checking direct connections from motors to throat

	// 9 - 34-20

	if channels[9].Defined {
		hdType = "Manifesting Generator"
		return hdType
	}

	// 18- 22-12
	// 19- 36-35
	// 17- 21-45
	if channels[17].Defined || channels[18].Defined || channels[19].Defined {
		if centers["Sacral"] {
			hdType = "Manifesting Generator"
			return hdType
		}

		hdType = "Manifestor"
		return hdType
	}

	if !centers["Throat"] {

		if centers["Sacral"] {
			hdType = " Generator"
			return hdType
		}

		hdType = "Projector"
		return hdType
	}

	// 14- 7-31
	// 15- 1-8
	// 16- 13-33
	// 10- 10-20
	var gToThroat bool = channels[14].Defined || channels[15].Defined || channels[16].Defined || channels[10].Defined

	// 8 - 57-20
	// 7 - 48-16
	var spleenToThroat bool = channels[8].Defined || channels[7].Defined

	// 30- 58-18
	// 29- 38-28
	// 28- 54-32
	var rootToSpleen bool = channels[30].Defined || channels[29].Defined || channels[28].Defined

	// 11- 57-10
	var spleenToG bool = channels[11].Defined

	// 23- 51-25
	var egoToG bool = channels[23].Defined

	// 24- 44-26
	var egoToSpleen bool = channels[24].Defined

	// 12- 57-34
	// 25- 27-50
	var sacralToSpleen bool = channels[12].Defined || channels[25].Defined

	//sacral through G center
	// 20- 5-15
	// 21- 14-2
	// 22- 29-46
	// 13- 34-10
	var sacralToG bool = channels[20].Defined || channels[21].Defined || channels[22].Defined || channels[13].Defined

	// Generator or Manifesting Generator
	if centers["Sacral"] {
		hdType = "Generator"

		if sacralToG && (gToThroat || (spleenToG && spleenToThroat)) {

			hdType = "Manifesting Generator"
			return hdType
		}

		if sacralToSpleen && (spleenToThroat || (spleenToG && gToThroat)) {

			hdType = "Manifesting Generator"
			return hdType

		}

		if egoToG && (gToThroat || (spleenToG && spleenToThroat)) {
			hdType = "Manifesting Generator"
			return hdType

		}

		if egoToSpleen && (spleenToThroat || (spleenToG && gToThroat)) {

			hdType = "Manifesting Generator"
			return hdType

		}

		if rootToSpleen && (spleenToThroat || (spleenToG && gToThroat)) {

			hdType = "Manifesting Generator"
			return hdType

		}

		return hdType
	}

	if centers["Ego"] || centers["Emo"] || centers["Root"] {

		if egoToG && (gToThroat || (spleenToG && spleenToThroat)) {

			hdType = "Manifestor"
			return hdType

		}

		if egoToSpleen && (spleenToThroat || (spleenToG && gToThroat)) {
			hdType = "Manifestor"
			return hdType
		}

		if rootToSpleen && (spleenToThroat || (spleenToG && gToThroat)) {
			hdType = "Manifestor"
			return hdType
		}

	}

	return hdType

}

func CentersConnections(info *structs.CdInfo) [][]string {

	channels := info.HdInfo.Channels
	centers := info.HdInfo.Centers.Center

	var headToAjna bool = channels[1].Defined || channels[2].Defined || channels[3].Defined

	var ajnaToThroat bool = channels[4].Defined || channels[5].Defined || channels[6].Defined
	var gToThroat bool = channels[14].Defined || channels[15].Defined || channels[16].Defined || channels[10].Defined
	var spleenToThroat bool = channels[8].Defined || channels[7].Defined
	var egoToThroat bool = channels[17].Defined
	var emoToThroat bool = channels[18].Defined || channels[19].Defined
	var sacralToThroat bool = channels[9].Defined

	var egoToG bool = channels[23].Defined
	var spleenToG bool = channels[11].Defined
	var sacralToG bool = channels[20].Defined || channels[21].Defined || channels[22].Defined || channels[13].Defined

	var sacralToRoot bool = channels[31].Defined || channels[32].Defined || channels[33].Defined
	var sacralToSpleen bool = channels[12].Defined || channels[25].Defined
	var sacralToEmo bool = channels[26].Defined

	var rootToSpleen bool = channels[30].Defined || channels[29].Defined || channels[28].Defined
	var rootToEmo = channels[34].Defined || channels[35].Defined || channels[36].Defined

	var egoToSpleen bool = channels[24].Defined

	var emoToEgo bool = channels[27].Defined

	var connArray [][]string
	var conn []string

	if centers["Head"] {
		if headToAjna {
			conn = []string{"Head", "Ajna"}
			connArray = append(connArray, conn)
		}
	}

	if centers["Ajna"] {
		if ajnaToThroat {
			conn = []string{"Throat", "Ajna"}
			connArray = append(connArray, conn)
		}
	}

	if centers["Throat"] {
		conn = []string{"Throat"}
		if gToThroat {
			conn = append(conn, "G")

		}
		if spleenToThroat {
			conn = append(conn, "Spleen")

		}

		if egoToThroat {
			conn = append(conn, "Ego")

		}

		if emoToThroat {
			conn = append(conn, "Emo")

		}

		if sacralToThroat {
			conn = append(conn, "Sacral")

		}

		connArray = append(connArray, conn)
	}

	if centers["G"] {

		conn = []string{"G"}
		if gToThroat {
			conn = append(conn, "Throat")

		}

		if spleenToG {
			conn = append(conn, "Spleen")

		}

		if sacralToG {
			conn = append(conn, "Sacral")

		}

		if egoToG {
			conn = append(conn, "Ego")

		}

		connArray = append(connArray, conn)

	}

	if centers["Sacral"] {

		conn = []string{"Sacral"}

		if sacralToRoot {
			conn = append(conn, "Root")
		}

		if sacralToSpleen {
			conn = append(conn, "Spleen")
		}

		if sacralToEmo {
			conn = append(conn, "Emo")
		}

		if sacralToG {
			conn = append(conn, "G")
		}

		if sacralToThroat {
			conn = append(conn, "Throat")
		}

		connArray = append(connArray, conn)

	}

	if centers["Root"] {
		conn = []string{"Root"}

		if rootToSpleen {
			conn = append(conn, "Spleen")
		}

		if rootToEmo {
			conn = append(conn, "Emo")
		}

		if sacralToRoot {
			conn = append(conn, "Sacral")
		}

		connArray = append(connArray, conn)
	}

	if centers["Spleen"] {
		conn = []string{"Spleen"}

		if rootToSpleen {
			conn = append(conn, "Root")
		}

		if sacralToSpleen {
			conn = append(conn, "Sacral")
		}

		if spleenToG {
			conn = append(conn, "G")
		}

		if spleenToThroat {
			conn = append(conn, "Throat")
		}

		if egoToSpleen {
			conn = append(conn, "Ego")
		}

		connArray = append(connArray, conn)
	}

	if centers["Emo"] {
		conn = []string{"Emo"}

		if rootToEmo {
			conn = append(conn, "Root")
		}

		if sacralToEmo {
			conn = append(conn, "Sacral")
		}

		if emoToThroat {
			conn = append(conn, "Throat")
		}

		if emoToEgo {
			conn = append(conn, "Ego")
		}

		connArray = append(connArray, conn)
	}

	if centers["Ego"] {
		conn = []string{"Ego"}

		if emoToEgo {
			conn = append(conn, "Emo")
		}

		if rootToEmo {
			conn = append(conn, "Root")
		}

		if sacralToEmo {
			conn = append(conn, "Sacral")
		}

		if emoToThroat {
			conn = append(conn, "Throat")
		}

		connArray = append(connArray, conn)
	}

	return connArray
}

func Definition(info *structs.CdInfo) string {

	cenConn := CentersConnections(info)

	fmt.Println("cenConn at start == ", cenConn, "cenConn lenth at start == ", len(cenConn))
	finalArr := make([][]string, 0)
	//fmt.Println("cenConn[0]) ==  ", cenConn[0])
	var toDelete []int
	index := 0
	for {
		finalArr = append(finalArr, []string{})
		//fmt.Println("finalArr[index] ==  ", finalArr[index])
		finalArr[index] = append(finalArr[index], cenConn[0]...)
		toDelete = append(toDelete, 0)
		//fmt.Println("toDelete == ", toDelete)

		var finalArrTemp []string
		finalArrTemp = append(finalArrTemp, finalArr[index]...)
		//fmt.Println("finalArrTemp == ", finalArrTemp)

		//fmt.Println(v)

		for {

			found := false

			for key, value := range cenConn {
				//if key == 0 {
				//	continue
				//}

				for _, v := range value {

					for i := 0; i < len(finalArr[index]); i++ {

						if finalArr[index][i] == v {
							//fmt.Println("v ===", v)

							var finalArrTemp2 []string
							isIn := false
							for _, temp_v := range value {
								for i := 0; i < len(finalArrTemp); i++ {

									if finalArrTemp[i] == temp_v {

										isIn = true
									}
								}

								if !isIn {
									finalArrTemp2 = append(finalArrTemp2, temp_v)
								}
							}

							finalArrTemp = append(finalArrTemp, finalArrTemp2...)
							//finalArrTemp = append(finalArrTemp, value...)

							found = true //начинаем с начала
							//fmt.Println("finalArrTemp ===", finalArrTemp)
							//found = true
							//fmt.Println(len(toDelete))
							alreadyExists := false
							for i := 0; i < len(toDelete); i++ {
								if toDelete[i] == key {
									alreadyExists = true
									break
								}
							}
							if !alreadyExists {
								toDelete = append(toDelete, key)
							}

							break
						}
					}

					if found {
						break
					}

				}
				if found {
					cenConn = DelElements(cenConn, toDelete)
					finalArr[index] = finalArrTemp
					toDelete = nil
					break
				}
				//fmt.Println(key, "   ", value)
			}

			if !found {
				break
			}

		}

		//fmt.Println(" befor copy")
		//fmt.Println("finalArrTemp ===", finalArrTemp)
		//fmt.Println("finalArr[index] ===", finalArr[index])
		//copy(finalArr[index], finalArrTemp)
		//finalArr[index] = finalArrTemp
		//fmt.Println(" after copy")
		//fmt.Println("finalArrTemp ===", finalArrTemp)
		//fmt.Println("finalArr[index] ===", finalArr[index])
		index++

		fmt.Println(finalArr)
		fmt.Println("--------")

		//fmt.Println("toDelete  == ", toDelete, "len(toDelete) == ", len(toDelete))
		//fmt.Println("cenConn  == ", cenConn, "cenConn lenth  == ", len(cenConn))
		//cenConn = utils.DelElements(cenConn, toDelete)
		//fmt.Println("cenConn  after Del == ", cenConn, "cenConn lenth  == ", len(cenConn))

		//toDelete = nil

		if len(cenConn) == 0 {
			//fmt.Println("Quitting")
			//fmt.Println("cenConn  == ", cenConn, "cenConn lenth  == ", len(cenConn))
			break
		}
	}

	fmt.Println(finalArr)

	retValue := ""
	switch len(finalArr) {
	case 0:
		retValue = "No Definition"
	case 1:
		retValue = "Single Definition"
	case 2:
		retValue = "Split Definition"
	case 3:
		retValue = "Triple Split Definition"
	case 4:
		retValue = "Quadruple Split Definition"
	default:
		retValue = "???"

	}

	//fmt.Println(cenConn)
	return retValue

}

func DelElements(s [][]string, intArr []int) [][]string {

	for i := 0; i < len(intArr); i++ {
		s[intArr[i]] = nil
	}

	newSlice := make([][]string, 0)
	for _, item := range s {
		if item != nil {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice

}
