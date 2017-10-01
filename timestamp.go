package libgotimestamp
import (
	"strconv"
	"fmt"
	"math"
//	"time"
)


type TimeStamp struct {
	Hours int
	Minutes int
	Seconds int
	Decimals int
}

type PaddedStamp struct {
	Hours string
	Minutes string
	Seconds string
	Decimals string
}


func round(f float64) float64 {
	return math.Floor(f + .5)
}
 
func roundPlus(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return round(f * shift) / shift;	
}


func padStamp(in int)(out string){
	proc := strconv.Itoa(in)
	if len(proc) == 1 {
		out = "0"+proc
	} else if len(proc) < 1 {
		out = "00"
	} else {
		out = proc
	}
	return
}

//MakeTimeStamp converts a frame int, and an fps value to a timestamp and padded timestamp struct
func MakeTimeStamp(fps float32, fpspos int) (timestamp *TimeStamp, paddedstamp *PaddedStamp){
	var hours int64
	hours = 0
	secpos := float32(fpspos) / fps
	fmt.Println(roundPlus(float64(secpos), 3))
	fmt.Println(secpos)
	minutes := secpos/60
	if minutes > 60 {
		// check how many hours we have and round down
		hours = int64(minutes / 60)
		minutes = minutes - 60*float32(hours)
	}
	minstamp := int32(minutes)
	seconds := int32((minutes - float32(minstamp))*60)
	decimal := ((minutes-float32(minstamp))*60) - float32(seconds)
	fmt.Println(roundPlus(float64(decimal*10), 0))
	decimals := int32(decimal*1000)
	timestamp = &TimeStamp{
		Hours: int(hours),
		Minutes: int(minutes),
		Seconds: int(seconds),
		Decimals: int(decimals),
	}

	paddedstamp = &PaddedStamp{
		Hours: padStamp(timestamp.Hours),
		Minutes: padStamp(timestamp.Minutes),
		Seconds: padStamp(timestamp.Seconds),
		Decimals: strconv.Itoa(timestamp.Decimals),
	}
	return
}
