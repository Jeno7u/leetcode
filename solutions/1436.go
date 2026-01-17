package main

// создаем hash map[отправление] -> куда прибывает. Потом берем любое значение прибытия и пока есть в hash map
// места где место отправления = место прибытия получаем новое значение destination
func destCity(paths [][]string) string {
    pathsMap := map[string]string{}
    for i := range paths {
        pathsMap[paths[i][0]] = paths[i][1]
    }

    destination := paths[0][1]
	for true {
		if _, ok := pathsMap[destination]; !ok {
			break
		}
		destination = pathsMap[destination]
	}
    
	return destination
}
