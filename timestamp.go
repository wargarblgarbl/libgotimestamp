package libgotimestamp

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//TimeStamp stores a timestamp as a struct of ints. 
type TimeStamp struct {
	Hours    int
	Minutes  int
	Seconds  int
	Decimals int
	Frame    int
}

//PaddedStamp stores a timestamp as a struct of strings with padding applied
type PaddedStamp struct {
	Hours    string
	Minutes  string
	Seconds  string
	Decimals string
	Frame    string
}

func intit(in string) (out int) {
	if in == "" {
		out = 0
	} else {
		outa, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		out = outa
	}
	return
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func roundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}

func padStamp(in int) (out string) {
	proc := strconv.Itoa(in)
	if len(proc) == 1 {
		out = "0" + proc
	} else if len(proc) < 1 {
		out = "00"
	} else {
		out = proc
	}
	return
}

func padDec(in int) (out string) {
	proc := strconv.Itoa(in)
	if in < 100 {
		out = "0" + proc
	} else {
		out = proc
	}
	return
}

//MakeTimeStamp converts a frame int, and an fps value to a timestamp and padded timestamp struct
func MakeTimeStamp(fps float64, fpspos int) (timestamp *TimeStamp, paddedstamp *PaddedStamp) {
	var hours int64
	hours = 0
	secpos := float64(fpspos) / fps
	minutes := secpos / 60
	if minutes > 60 {
		// check how many hours we have and round down
		hours = int64(minutes / 60)
		minutes = minutes - 60*float64(hours)
	}
	minstamp := int32(minutes)
	seconds := int32((minutes - float64(minstamp)) * 60)
	decimal := ((minutes - float64(minstamp)) * 60) - float64(seconds)
	decimals := int32(decimal * 1000)
	timestamp = &TimeStamp{
		Hours:    int(hours),
		Minutes:  int(minutes),
		Seconds:  int(seconds),
		Decimals: int(decimals),
		Frame:    fpspos,
	}

	paddedstamp = &PaddedStamp{
		Hours:    padStamp(timestamp.Hours),
		Minutes:  padStamp(timestamp.Minutes),
		Seconds:  padStamp(timestamp.Seconds),
		Decimals: padDec(timestamp.Decimals),
		Frame:    strconv.Itoa(timestamp.Frame),
	}
	return
}


//MakeFrame converts a frame and an fps value to a TimeStamp and PaddedStamp struct
func MakeFrame(fps float64, timestamp string) (outtimestamp *TimeStamp, outpaddedstamp *PaddedStamp) {
	splits := strings.Split(timestamp, ":")

	//	hours := splits[0]
	//	minutes := splits[1]
	ssplits := strings.Split(splits[2], ".")
	//	seconds := ssplits[0]
	//	decimals := ssplits[1]
	if len(ssplits[1]) < 3 {
		ssplits[1] = ssplits[1] + "0"
	}
	timestring := splits[0] + "h" + splits[1] + "m" + ssplits[0] + "s" + ssplits[1] + "ms"
	dur, err := time.ParseDuration(timestring)
	if err != nil {
		fmt.Println(err)
	}

	ts := roundPlus((dur.Seconds() * fps), 0)
	//return int64(ts)
	outtimestamp = &TimeStamp{
		Hours:    intit(splits[0]),
		Minutes:  intit(splits[1]),
		Seconds:  intit(ssplits[0]),
		Decimals: intit(ssplits[1]),
		Frame:    int(ts),
	}
	outpaddedstamp = &PaddedStamp{
		Hours:    splits[0],
		Minutes:  splits[1],
		Seconds:  ssplits[0],
		Decimals: ssplits[1],
		Frame:    strconv.Itoa(int(ts)),
	}
	return
}
