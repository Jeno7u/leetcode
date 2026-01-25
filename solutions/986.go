package main

// очевидное решение, оно такое же по скорости вроде, но требует намного больше условий и проверок,
// но все еще корректное и хорошее решение
// func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
//     result := [][]int{}
// 	p1, p2 := 0, 0
// 	for p1 < len(firstList) && p2 < len(secondList) {
// 		fmt.Println(p1, p2, result)
// 		firstStart := firstList[p1][0]
// 		firstEnd := firstList[p1][1]
// 		secondStart := secondList[p2][0]
// 		secondEnd := secondList[p2][1]

// 		if firstEnd < secondStart {
// 			p1++
// 			continue
// 		} else if secondEnd < firstStart {
// 			p2++
// 			continue
// 		}
// 		// p2 starts last
// 		if firstStart <= secondStart {
// 			// p2 ends first
// 			if firstEnd >= secondEnd{
// 				result = append(result, secondList[p2])
// 				p2++
// 			// p1 ends first
// 			} else {
// 				result = append(result, []int{secondStart, firstEnd})
// 				p1++
// 			}
// 		// p1 starts last
// 		} else {
// 			// p2 ends first
// 			if firstEnd >= secondEnd {
// 				result = append(result, []int{firstStart, secondEnd})
// 				p2++
// 			// p1 ends first
// 			} else {
// 				result = append(result, firstList[p1])
// 				p1++
// 			}
// 		}
// 	}

// 	return result
// }

// второе решение более элегантное. По сути завязано на двух поинтерах. Если есть хоть какое-то пересечение, то добавляем в result интервал через min и max.
// А если нету пересечений, то двигаем поинтер
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    result := [][]int{}
	p1, p2 := 0, 0
	for p1 < len(firstList) && p2 < len(secondList) {
		firstStart, firstEnd := firstList[p1][0], firstList[p1][1]
		secondStart, secondEnd := secondList[p2][0], secondList[p2][1]
        if firstStart <= secondEnd && secondStart <= firstEnd {
            result = append(result, []int{max(firstStart, secondStart), min(firstEnd, secondEnd)})
        }

        if firstEnd <= secondEnd {
            p1++
        } else {
            p2++
        }

	}

	return result
}

