package monkey

import (
	"fmt"
	"math/big"
	"strconv"
)

type Operator string

const (
	MULT Operator = "MULT"
	ADD  Operator = "ADD"
	DIV  Operator = "DIV"
	SUB  Operator = "SUB"
)

type StressOperationDescriptor struct {
	Operation Operator
	Value     string
}

type ThrowOperationDescriptor struct {
	Operation     Operator
	Value         *big.Int
	MonkeyIfTrue  int
	MonkeyIfFalse int
}

type Monkey struct {
	StartingItems       []*big.Int
	StressOperation     StressOperationDescriptor
	ThrowOperation      ThrowOperationDescriptor
	InspectedItemsValue int
}

func PrintMonkey(monk Monkey) {
	fmt.Println("{")
	fmt.Printf(" - Starting Items: [ ")
	for _, items := range monk.StartingItems {
		fmt.Printf("%d ", items)
	}
	fmt.Printf("]\n")
	fmt.Printf(" - Stress Operation: operator %s and value %s\n", monk.StressOperation.Operation, monk.StressOperation.Value)
	fmt.Printf(" - Throw Operation: operator %s, value %d, monkey if true %d, monkey if false %d\n",
		monk.ThrowOperation.Operation, monk.ThrowOperation.Value, monk.ThrowOperation.MonkeyIfTrue, monk.ThrowOperation.MonkeyIfFalse)
	fmt.Printf(" - Monkey inspected items %d times\n", monk.InspectedItemsValue)
	fmt.Println("}")
}

func ApplyMonkeyOp(op StressOperationDescriptor, itemvAlue *big.Int) *big.Int {
	returnVal := big.NewInt(-1)
	switch op.Operation {
	case "MULT":
		intVal, err := strconv.Atoi(op.Value)
		if err != nil {
			return returnVal.Mul(itemvAlue, itemvAlue)
		}
		return returnVal.Mul(itemvAlue, big.NewInt(int64(intVal)))

	case "DIV":
		intVal, err := strconv.Atoi(op.Value)
		if err != nil {
			return returnVal.Div(itemvAlue, itemvAlue)
		}
		return returnVal.Div(itemvAlue, big.NewInt(int64(intVal)))

	case "ADD":
		intVal, err := strconv.Atoi(op.Value)
		if err != nil {
			return returnVal.Add(itemvAlue, itemvAlue)
		}
		return returnVal.Add(itemvAlue, big.NewInt(int64(intVal)))
	case "SUB":
		intVal, err := strconv.Atoi(op.Value)
		if err != nil {
			return returnVal.Sub(itemvAlue, itemvAlue)
		}
		return returnVal.Sub(itemvAlue, big.NewInt(int64(intVal)))
	}

	fmt.Printf("This shouldnt happen, OP %s unrecognized.\n", op.Operation)
	return returnVal
}

func ApplyMonkeyThrowOp(op ThrowOperationDescriptor, itemvAlue *big.Int) bool {

	switch op.Operation {
	case "DIV":
		modulus := new(big.Int)
		return modulus.Mod(itemvAlue, op.Value).BitLen() == 0
	}

	fmt.Printf("This shouldnt happen, OP %s unrecognized.\n", op.Operation)
	return false
}
