package main

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type chunkHeader struct {
	ID   [4]byte
	Size uint32
}

type getMetadataParams struct {
	filepath string
	fileName string
	wg       *sync.WaitGroup
}
type iXML struct {
	XMLName xml.Name          `xml:"BWFXML"`
	Project *string           `xml:"PROJECT"`
	Speed   *TimecodeMetadata `xml:"SPEED"`
	Scene   *string           `xml:"SCENE"`
	Take    *string           `xml:"TAKE"`
}
type TimecodeMetadata struct {
	XMLName                xml.Name `xml:"SPEED"`
	SpeedString            *string  `xml:"TIMECODE_RATE"`
	Speed                  *float64
	SampleRate             *float64 `xml:"FILE_SAMPLE_RATE"`
	SamplesSinceMidnightLo *float64 `xml:"TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO"`
	SamplesSinceMidnightHi *float64 `xml:"TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_HI"`
}

func main() {
	dirpath := "../../../shared/audio/"
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	start := time.Now()
	wg := sync.WaitGroup{}
	for _, entry := range entries {
		wg.Add(1)
		params := getMetadataParams{
			filepath: filepath.Join(dirpath, entry.Name()),
			fileName: entry.Name(),
			wg:       &wg,
		}
		go GetIxmlMetadata(params)
	}
	wg.Wait()
	fmt.Println("Time taken for Go implementation: ", time.Since(start))
	fmt.Print("Current Memory Allocation: ")
	printAlloc()
}

func GetIxmlMetadata(params getMetadataParams) {
	defer params.wg.Done()
	fd, err := os.Open(params.filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()
	_, err = fd.Seek(12, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var header chunkHeader
		err = binary.Read(fd, binary.LittleEndian, &header)
		if err != nil {
			if err == io.EOF {
				fmt.Println("No XML for ", params.filepath)
				break
			}
			fmt.Println(err)
			return
		}
		chunkID := string(header.ID[:])

		if chunkID == "iXML" {
			data := make([]byte, header.Size)
			if _, err := io.ReadFull(fd, data); err != nil {
				fmt.Println(err)
				return
			}
			data = bytes.TrimRight(data, "\x00")
			os.WriteFile(fmt.Sprintf("./outputs/%s.xml", params.fileName), data, 0o755)
			ixml, err := ParseiXML(data)
			if err != nil {
				fmt.Println("ERROR while parsing iXML for: ", params.fileName, err.Error())
			}
			tc := ixml.ConvertSamplesToTimecode()
			fmt.Println("TIMECODE: ", tc)

			break
		} else {
			skip := int64(header.Size)
			if skip%2 == 1 {
				skip++
			}
			_, err := fd.Seek(skip, io.SeekCurrent)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func ParseiXML(ixmlData []byte) (iXML, error) {
	var bwfXML iXML
	err := xml.Unmarshal(ixmlData, &bwfXML)
	if err != nil {
		return bwfXML, err
	}
	err = bwfXML.ConvertSpeed()
	if err != nil {
		return bwfXML, err
	}
	return bwfXML, nil
}

func (i iXML) ConvertSpeed() error {
	if i.Speed.SpeedString == nil {
		return fmt.Errorf("no speed data present")
	}
	splitStr := strings.Split(*i.Speed.SpeedString, "/")
	rate, err := strconv.ParseInt(splitStr[0], 0, 64)
	if err != nil {
		return err
	}
	div, err := strconv.ParseInt(splitStr[1], 0, 64)
	if err != nil {
		return err
	}
	rateI := float64(rate)
	divI := float64(div)
	speed := rateI / divI
	if divI == 0 {
		return fmt.Errorf("couldn't find speed data")
	}
	i.Speed.Speed = &speed
	return nil
}

func (i iXML) ConvertSamplesToTimecode() string {
	if i.Speed.SamplesSinceMidnightLo == nil ||
		i.Speed.SampleRate == nil {
		return "00:00:00:00"
	}
	if i.Speed.SampleRate == nil {
		fmt.Println("bad sample rate for: ", *i.Scene, *i.Take)
	}
	totalSeconds := *i.Speed.SamplesSinceMidnightLo / *i.Speed.SampleRate
	hours := int(totalSeconds / 3600)
	minutes := (int(totalSeconds) / 60) % 60
	seconds := (int(totalSeconds) % 60)
	secondsFraction := totalSeconds - (math.Floor(totalSeconds))
	frames := int(secondsFraction * *i.Speed.Speed)
	tc := fmt.Sprintf("%.2d:%.2d:%.2d:%.2d", hours, minutes, seconds, frames)
	return tc
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
