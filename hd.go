//HD rrelated calculations
package hd

import (
	"fmt"

	"github.com/RulezKT/structs"
)

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
