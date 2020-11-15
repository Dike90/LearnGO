package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type ProcStat struct {
	CPU       string
	User      int64
	Nice      int64
	System    int64
	Idle      int64
	IoWait    int64
	Irq       int64
	SoftIrq   int64
	Steal     int64
	Guest     int64
	GuestNice int64
}

var preProcStat []ProcStat
var curProcStat []ProcStat
var readed bool

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, strings.Trim(line, "\n"))
	}
	return lines, nil
}

func parseProcStat() []ProcStat {
	var porcStats []ProcStat
	lines, _ := ReadLines("/proc/stat")
	for _, line := range lines {
		if strings.HasPrefix(line, "cpu") {
			var procStat ProcStat
			_, err := fmt.Sscanf(line, "%s %d %d %d %d %d %d %d %d %d %d",
				&procStat.CPU, &procStat.User, &procStat.Nice, &procStat.System, &procStat.Idle,
				&procStat.IoWait, &procStat.Irq, &procStat.SoftIrq, &procStat.Steal, &procStat.Guest, &procStat.GuestNice)
			if err != nil {
				continue
			}
			porcStats = append(porcStats, procStat)
		}
	}
	return porcStats
}

func main() {

	curProcStat = parseProcStat()
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for { //循环
			<-ticker.C
			preProcStat = curProcStat
			curProcStat = parseProcStat()
			fmt.Println(preProcStat)
			fmt.Println(curProcStat)
			for i := 0; i < len(curProcStat); i++ {
				preIdle := preProcStat[i].Idle + preProcStat[i].IoWait
				preNonIdle := preProcStat[i].User + preProcStat[i].Nice + preProcStat[i].System +
					preProcStat[i].Irq + preProcStat[i].SoftIrq + preProcStat[i].Steal
				preTotal := preNonIdle + preIdle
				curIdle := curProcStat[i].Idle + curProcStat[i].IoWait
				curNonIdle := curProcStat[i].User + curProcStat[i].Nice + curProcStat[i].System +
					curProcStat[i].Irq + curProcStat[i].SoftIrq + curProcStat[i].Steal
				curTotal := curNonIdle + curIdle
				totalDiff := curTotal - preTotal
				idleDiff := curIdle - preIdle
				usage := (float64(curNonIdle) - float64(preNonIdle)) / float64(totalDiff)
				fmt.Printf("preIdle:%d, curIdle:%d\n", preIdle, curIdle)
				fmt.Printf("totalDiff:%d, idleDiff:%d\n", totalDiff, idleDiff)
				fmt.Printf("preTotal:%d, curTotal:%d\n", preTotal, curTotal)
				fmt.Printf("usage:%f\n", usage*100)
			}

		}
	}()
	for {
		time.Sleep(time.Second)
	}
}
