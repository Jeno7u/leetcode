package main


// делаем из требований к курсу ориентированный граф. Идем от каждого значения курса
// к последнему (проверяем на цикл). Если встречаем цикл, то поднимаем этот результат
// выше по стеку вызовов и возвращаем false. Иначе true.
func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := map[int][]int{}
    for _, prerequisite := range prerequisites {
        adj[prerequisite[0]] = append(adj[prerequisite[0]], prerequisite[1])
    }

    visited := map[int]struct{}{}
    var dfs func(course int) bool
    dfs = func(course int) bool {
        _, ok := visited[course]
        if ok {
            return false
        }
        if len(adj[course]) == 0 {
            return true
        }

        visited[course] = struct{}{}
        for i := range adj[course] {
            if !dfs(adj[course][i]) {
                return false
            }
        }
        delete(visited, course)
        adj[course] = adj[course][:0]
        return true
    }

    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return false
        }
    }
    return true
}