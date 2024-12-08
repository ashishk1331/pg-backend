package util

import (
	"sync"
	"os"
	"path/filepath"
	"os/exec"
	"log"
	"time"
)

func RunCode(wg *sync.WaitGroup, code *CodeSample, content string) {
	defer wg.Done()

	wd, _ := os.Getwd()
	var filename string = "py-"+ UUID() + ".py"
	var filepath string = filepath.Join(wd, filename)

	WriteFile(filepath, content)
	defer DeleteFile(filepath)

	start := time.Now()
	cmd := exec.Command("python", filepath)
	elapsed := time.Since(start)

	code.TimeTaken = int(elapsed)
	code.TimeUnits = "ns"

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing: %v\n", err)
		code.Error = string(output)
		code.OK = false
		return
	}
	code.Output = string(output)
}