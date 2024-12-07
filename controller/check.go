package check

import (
	"github.com/gin-gonic/gin"
	"pg-backend/util"
	"pg-backend/template"
	"sync"
	"os"
	"path/filepath"
	"os/exec"
	"fmt"
	"time"
)

type CodeSample = util.CodeSample

// routine to write content to a file
func runCode(wg *sync.WaitGroup, code *CodeSample) {
	defer wg.Done()

	wd, _ := os.Getwd()
	var filename string = filepath.Join(wd, "script.py")

	var content string = template.GenerateBase(code)

	util.WriteFile(content)
	defer util.DeleteFile(filename)

	start := time.Now()
	cmd := exec.Command("python", filename)
	elapsed := time.Since(start)

	code.TimeTaken = int(elapsed)
	code.TimeUnits = "ns"

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing: %v\n", err)
		code.OK = false
		return
	}
	code.Output = string(output)
}

func Get(c *gin.Context) {
	var wg sync.WaitGroup

	var questionId int = util.Stoi(c.Query("question_id"))
	var testId int = util.Stoi(c.Query("test_id"))
	var code CodeSample = CodeSample{QuestionId: questionId, TestId: testId, OK: true}

	c.ShouldBind(&code)

	wg.Add(1)
	go runCode(&wg, &code)

	wg.Wait()
	c.JSON(200, code)
}