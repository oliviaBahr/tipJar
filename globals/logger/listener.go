package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
)

var stdLog = log.New(os.Stdout)

var LOG_LEVELS = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
	"fatal": log.FatalLevel,
}

func Listen(logPath string) {
	cmd := exec.Command("tail", "-f", "-n", "0", logPath)

	reader, err := cmd.StdoutPipe()
	if err != nil {
		stdLog.Fatal("failed to get stdout pipe", "error", err)
	}

	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			stdLog.Debug("line", "line", line)
			logLine(line)
			stdLog.Debug("checked line")
		}
	}()

	err = cmd.Start()
	if err != nil {
		stdLog.Fatal("failed to start tail command", "error", err)
	}

	err = cmd.Wait()
	if err != nil {
		stdLog.Fatal("tail command failed", "error", err)
	}
}

var logger = log.NewWithOptions(os.Stdout, log.Options{
	Level:           log.DebugLevel,
	ReportCaller:    false,
	ReportTimestamp: false,
})

func logLine(line string) {
	// read line json
	var record map[string]interface{}
	err := json.Unmarshal([]byte(line), &record)
	if err != nil {
		stdLog.Fatal("failed to unmarshal log line", "error", err)
	}

	level := record["level"].(string)
	message := record["msg"].(string)
	caller := record["caller"].(string)
	delete(record, "level")
	delete(record, "msg")
	delete(record, "caller")

	logLine := styleLine(message, caller, record)
	logger.Log(LOG_LEVELS[level], logLine)
}

var dfStyles = log.DefaultStyles()

func styleLine(message string, caller string, keyVals map[string]interface{}) string {
	// caller
	result := dfStyles.Key.Render(fmt.Sprintf("<%s>", caller)) + " "

	// message
	result += dfStyles.Message.Render(message) + " "

	// keyvals
	for k, v := range keyVals {
		key := dfStyles.Key.Render(fmt.Sprintf("%s=", k))
		val := dfStyles.Value.Render(fmt.Sprintf("%v", v))
		result += fmt.Sprintf("%s%s ", key, val)
	}

	return result
}
