package sample

import (
	
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-sample")

const (
	ivString1 = "string1"
	ivString2 = "string2"
	ivString3 = "string3"

	ovOutputString = "output_string"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}
// sample program demonstrates the concatenation of three string into one
type sample struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &sample{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *sample) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *sample) Eval(context activity.Context) (done bool, err error) {

	// Get the runtime values
	s1, _ := context.GetInput(ivString1).(string)
	s2, _ := context.GetInput(ivString2).(string)
	s3, _ := context.GetInput(ivString3).(string)
	
    s := []string{s1, s2, s3}
	temp := strings.Join(s, "");

	context.SetOutput(ovOutputString, temp)
	return true, nil
}
