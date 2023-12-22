package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/windows"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := windows.RegisterEventSource(nil, syscall.StringToUTF16Ptr(programName))
	if err != nil {
		log.Fatal(err)
	}
	defer windows.DeregisterEventSource(sysLog)

	defer windows.DeregisterEventSource()

	message := "Logging in GO"

	err = windows.ReportEvent(windows.EVENTLOG_INFORMATION_TYPE, 0, 0, 0, nil, []string{message})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Event logged sucessfully.")
}
