package main

import "testing"

func TestTrainApp(t *testing.T) {
	expected:="Output"
	actual:=trainApp();
	if actual!=expected{
		t.Errorf(format: "Test failed",expected,actual)
	}

}
