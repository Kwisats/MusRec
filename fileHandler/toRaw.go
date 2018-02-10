package fileHandler

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/cmplx"
	"os"
	"os/exec"
	"strconv"

	"github.com/mjibson/go-dsp/fft"
)

//ToRaw is a function which should convert initial format to .raw and use .raw to build the
func ToRaw(fileName string) error {
	fmt.Println(fileName)
	err := exec.Command("mp3ToRaw", fileName).Run()
	if err != nil {
		log.Print("ToRaw:")
		log.Fatal(err)
	}

	bytesLE, err := ioutil.ReadFile("rawData/" + fileName + ".raw")
	check(err)
	reals := make([]float64, len(bytesLE)/2)
	for i := 1; i < len(bytesLE)-1; i += 2 {
		reals[i/2] = float64(binary.LittleEndian.Uint16(bytesLE[i : i+2]))
	}

	complex := fft.FFTReal(reals)
	amplitudes := make([]float64, len(complex))
	for i, z := range complex {
		amplitudes[i] = cmplx.Abs(z)
	}

	//протестировать на простых файлах

	output, err := os.Create("forTest/output.txt")
	fmt.Println(len(amplitudes))
	for i, buff := range amplitudes {
		if i >= 1 {
			//fmt.Println(buff)
			output.WriteString(strconv.Itoa(i) + "\t" + strconv.FormatFloat(buff, 'e', 3, 64) + "\n")
		}
	}

	//построить график

	err = exec.Command("gnuPlot").Run()
	if err != nil {
		log.Print("ToRaw:")
		log.Fatal(err)
	}
	return nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
