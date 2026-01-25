package main

import "strings"

func simplifyPath(path string) string {
    stack := []string{}

    i, wordStart := 0, -1

    for i < len(path) {
        // пропустить все последовательные '/'
        if path[i] == '/' {
            i++
            continue
        }
        if path[i] == '.' {
            // находим размер
            wordStart = i
            haveLetters := false
            for i + 1 < len(path) && path[i+1] != '/' {
                if path[i + 1] != '.' {
                    haveLetters = true
                }
                i++
            }

            // содержит буквы или больше двух точек - обрабатываем как слово
            if (i - wordStart) >= 2 || haveLetters {
                stack = append(stack, path[wordStart:i+1])
            // '..' - возвращаемся назад, то есть удаляем из стека последний элемент
            } else if (i - wordStart) == 1 {
                if len(stack) != 0 {
                    stack = stack[:len(stack) - 1]
                }
            }
        } else {
            wordStart = i
            for i + 1 < len(path) && path[i+1] != '/' {
                i++
            }
            stack = append(stack, path[wordStart:i+1])
        }
        i++
    }
    return "/" + strings.Join(stack, "/")
}