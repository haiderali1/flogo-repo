package sample

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalSubtract(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	fmt.Println("#######   Testing sample")

	//test1
	tc.SetInput("string1", "")
	tc.SetInput("string2", "")
	tc.SetInput("string3", "")
	act.Eval(tc)

	if tc.GetOutput("output_string") != "" {
		t.Fail()
	}
	res := tc.GetOutput("output_string")
	fmt.Println("Result should be: empty  is:", res)
	//test2
	tc.SetInput("string1", "a")
	tc.SetInput("string2", "b")
	tc.SetInput("string3", "c")
	act.Eval(tc)

	if tc.GetOutput("output_string") != "abc" {
		t.Fail()
	}
	res = tc.GetOutput("output_string")
	fmt.Println("Result should be:abc  is:", res)

}
