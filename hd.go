//HD rrelated calculations
package hd

import (
	"fmt"

	"github.com/RulezKT/structs"
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

	longitude = convert_to_0_360_DEG(longitude)
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
