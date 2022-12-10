package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPUState struct {
	Cycle          int
	SpriteWideness int
	XValue         int
}

type CRTState struct {
	Index     int
	Wideness  int
	Largeness int
	Drawing   string
}

type CommDeviceConfig struct {
	CPUWantedCycle int
	CPUIncrement   int
	CPUSpriteWide  int
	CRTWideness    int
	CRTLargeness   int
}

func main() {
	deviceConfig := CommDeviceConfig{CPUWantedCycle: 20, CPUIncrement: 40, CPUSpriteWide: 3, CRTWideness: 40, CRTLargeness: 6}
	fmt.Println(Challenge("./chall_input.txt", deviceConfig))
}

func Challenge(fileName string, deviceConfig CommDeviceConfig) int {

	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	CPUPower := 0

	CPUState := CPUState{Cycle: 1, SpriteWideness: deviceConfig.CPUSpriteWide, XValue: 1}
	CRTState := CRTState{Index: 1, Wideness: deviceConfig.CRTWideness, Largeness: deviceConfig.CRTLargeness, Drawing: ""}
	CRTState = updateCRTDrawing(CRTState, CPUState)

	for fileScanner.Scan() {
		inpt := strings.Split(fileScanner.Text(), " ")
		switch inpt[0] {
		case "noop":
			for stt := range noop(CPUState) {
				CPUState = stt
				CRTState = updateCRTDrawing(CRTState, CPUState)
				if CPUState.Cycle == deviceConfig.CPUWantedCycle {
					CPUPower += CPUState.Cycle * CPUState.XValue
					deviceConfig.CPUWantedCycle += deviceConfig.CPUIncrement
				}
			}
		case "addx":
			xVal, err := strconv.Atoi(inpt[1])
			if err != nil {
				fmt.Printf("%s is not a number. Terminating.\n", xVal)
				return -1
			}
			for stt := range addx(CPUState, xVal) {
				CPUState = stt
				CRTState = updateCRTDrawing(CRTState, CPUState)
				if CPUState.Cycle == deviceConfig.CPUWantedCycle {
					CPUPower += CPUState.Cycle * CPUState.XValue
					deviceConfig.CPUWantedCycle += deviceConfig.CPUIncrement
				}
			}
		}
	}

	PrintCrtDraw(CRTState)

	return CPUPower
}

func updateCRTDrawing(crtState CRTState, cpuState CPUState) CRTState {

	intervalCPULow := cpuState.XValue + 1 - (cpuState.SpriteWideness / 2)
	if intervalCPULow < 1 {
		intervalCPULow = 1
	}
	intervalCPUHigh := cpuState.XValue + 1 + (cpuState.SpriteWideness / 2)
	fmt.Printf("low interval is %d, high interval is %d, middle is %d, index is %d\n", intervalCPULow, intervalCPUHigh, cpuState.XValue+1, crtState.Index)

	if crtState.Index >= intervalCPULow && crtState.Index <= intervalCPUHigh {
		fmt.Println("Drawing #")
		crtState.Drawing += "#"
	} else {
		fmt.Println("Drawing .")
		crtState.Drawing += "."
	}

	if (len(crtState.Drawing))%crtState.Wideness == 0 {
		fmt.Printf("%d is len opf crt so we go back\n", len(crtState.Drawing))
		crtState.Index = 0
	}
	crtState.Index += 1

	return crtState
}

func PrintCrtDraw(state CRTState) {
	tempStr := ""
	for _, rne := range state.Drawing {
		if len(tempStr)%state.Wideness == 0 {
			fmt.Println(tempStr)
			tempStr = ""
		}

		tempStr += string(rne)
	}

	fmt.Println(tempStr)
}

func noop(state CPUState) <-chan (CPUState) {
	chnl := make(chan CPUState)
	go func() {
		chnl <- CPUState{Cycle: state.Cycle + 1, XValue: state.XValue, SpriteWideness: state.SpriteWideness}
		close(chnl)
	}()
	return chnl
}

func addx(state CPUState, addxValue int) <-chan (CPUState) {
	chnl := make(chan CPUState)
	go func() {
		chnl <- CPUState{Cycle: state.Cycle + 1, XValue: state.XValue, SpriteWideness: state.SpriteWideness}
		chnl <- CPUState{Cycle: state.Cycle + 2, XValue: state.XValue + addxValue, SpriteWideness: state.SpriteWideness}
		close(chnl)
	}()
	return chnl
}
