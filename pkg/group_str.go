package group

import (
	"bytes"
	"fmt"
	"sort"
)

// Псевдоним для типа
type str = string

// Реализация группировки строки по символам
func GroupStr(sourceStr str) str {

	// Если исходная строка пустая — возвращаем пустую строку
	if sourceStr == "" {
		return ""
	}

	// Буфер для формирования конечного результата
	var resBuffer bytes.Buffer

	// Вспомогательная функция для буфера
	resBufferWrite := func(b *bytes.Buffer, g str, n int) {
		if n == 1 {
			/* 
				Предположение / допущение:
				Если в группе 1 элемент - количество не указывается
				"aa" > "a2"
				"a" > "a"
			*/
			b.WriteString(g)
		 } else if n > 1 {
			b.WriteString(fmt.Sprintf("%s%d", g, n))
		}
	}

	// Подсчет символов в исходной строке
	mapCompact := make(map[str]int)
	for _, currRune := range sourceStr {
		currChar := string(currRune) // rune > string
		_ , ok := mapCompact[currChar]
		if ok {
			mapCompact[currChar] += 1
		} else {
			mapCompact[currChar] = 1
		}
	}
	
	// Сортировка по ключу (группа в исходной строке)
	/*
		"aa" > "a2" ("a" - группа, "2" - количество элементов)

		Предположение / допущение:
		Сортировка производиться по ключу (группа)
		"bbbaa" > "a2b3"
		В описании явно не указаны критерии сортировки
		Альтернатива — сортировка по количеству элементов в группе
		"bbbaa" > "b3a2"
	*/
	keySlice := make([]str, 0, len(mapCompact))
	for k := range mapCompact {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)

	// Формирование результата
	for _, k := range keySlice {
		resBufferWrite(&resBuffer, k, mapCompact[k])
	}
	return resBuffer.String()
}