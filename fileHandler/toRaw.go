package fileHandler

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/mjibson/go-dsp/fft"
)

//ToRaw is a function which should convert initial format to .raw and use .raw to build the
func ToRaw(fileName string) error {
	fmt.Println(fileName)
	err := exec.Command("mp3ToRaw").Run()
	if err != nil {
		log.Print("ToRaw:")
		log.Fatal(err)
	}

	bytesLE, err := ioutil.ReadFile("rawData/ostan.raw")
	check(err)
	fmt.Println(len(bytesLE))
	digits := make([]uint16, len(bytesLE)/2)
	for i := 0; i < len(bytesLE); i += 2 {
		digits[i/2] = binary.LittleEndian.Uint16(bytesLE[i : i+2])
	}
	fmt.Println(digits)
	//type convertion
	fmt.Println(fft.FFTReal([]float64{1, 2, 3}))
	// Теперь считать все числа в массив(?) и применить к "аудио-числам"
	// Найти аудио-файлы для теста
	// fmt.Printf("output is %s\n", out)
	return nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
