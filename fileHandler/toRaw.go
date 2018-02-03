package fileHandler

import (
	"fmt"
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
	fmt.Println(fft.FFTReal([]float64{1, 2, 3}))
	// Теперь считать все числа в массив(?) и применить к "аудио-числам"
	// Найти аудио-файлы для теста
	// fmt.Printf("output is %s\n", out)
	return nil
}
