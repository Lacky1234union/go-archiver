package variable

import (
	"bufio"
	"fmt"
	"os"
)

var RuneToNumbers = map[rune][]int{
	'a': {0, 1, 0, 1},
}

func bitsToBytes(bits []int) []byte {
	var bytes []byte
	var byteValue byte

	// Проходим по всем битам
	for i, bit := range bits {
		// Добавляем бит в текущий байт
		byteValue = byteValue<<1 | byte(bit)

		// Если набрали 8 бит, сохраняем байт
		if (i+1)%8 == 0 {
			bytes = append(bytes, byteValue)
			byteValue = 0 // Сбрасываем текущий байт
		}
	}

	// Если остались неполные биты, добавляем последний байт
	if len(bits)%8 != 0 {
		bytes = append(bytes, byteValue<<(8-len(bits)%8))
	}

	return bytes
}

func compress(line string) ([]byte, error) {
	var bits []int
	for _, s := range line {

		result := RuneToNumbers[s]
		for _, n := range result {
			bits = append(bits, n)
		}
	} //
	bytes := bitsToBytes(bits)
	return bytes, nil
}

func CompressToFile(fileorign, fileansw string) error {
	origin, err := os.Open(fileorign)
	if err != nil {
		return err
	}
	defer origin.Close()
	scanner := bufio.NewScanner(origin)
	sours, err := os.Create(fileansw)
	if err != nil {
		return err
	}
	defer sours.Close()
	for scanner.Scan() {
		data, err := compress(scanner.Text())
		if err != nil {
			return err
		}
		n, err := sours.Write(data)
		if err != nil {
			return err
		}
		fmt.Println("Compress to ", n, " bytes")
	}
	return nil
}
