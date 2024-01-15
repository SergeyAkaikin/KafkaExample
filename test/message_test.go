package test

import (
	"task/internal/message"
	"testing"
)

func TestMessage(t *testing.T) {

	number1 := 20
	value1 := message.Sync{Number: number1, Roman: "XX", Description: "twenty"}
	value2 := message.New(number1)

	if value1 != value2 {
		t.Errorf("value1 != value2, value1=%v, value2=%v\n", value1, value2)
	}

	number2 := 191
	value1 = message.Sync{Number: number2, Roman: "CXCI", Description: "one hundred ninety-one"}
	value2 = message.New(number2)

	if value1 != value2 {
		t.Errorf("value1 != value2, value1=%v, value2=%v\n", value1, value2)
	}

}
