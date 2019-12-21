package intent

import (
	"fmt"
	"strings"
)

const intentFormatTemplate = "mic_scheme://%s/push?%s#Intent;launchFlags=0x24000000;end"

func GenerateIntent(appPkgName string, extra map[string]string) string {
	var (
		extraStr string
	)
	extraStr = generateIntentExtra(extra)
	return fmt.Sprintf(intentFormatTemplate, appPkgName, extraStr)
}

func generateIntentExtra(extra map[string]string) string {
	var (
		extraStr string
		buf      strings.Builder
	)
	for key, value := range extra {
		buf.WriteString("&")
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(value)
	}
	extraStr = buf.String()
	extraStr = strings.Trim(extraStr, " ")
	return extraStr
}

