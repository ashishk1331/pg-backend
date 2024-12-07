package run

import (
	"github.com/gin-gonic/gin"
	"pg-backend/util"
	"sync"
)

type CodeSample = util.CodeSample

func Post(c *gin.Context) {
	var wg sync.WaitGroup

	var code CodeSample = CodeSample{OK: true}

	c.ShouldBind(&code)

	wg.Add(1)
	go util.RunCode(&wg, &code, code.Content)

	wg.Wait()
	c.JSON(200, gin.H{
		"output": code.Output,
		"ok": code.OK,
		"time_taken": code.TimeTaken,
		"time_units": code.TimeUnits,
		"error": code.Error,
	})
}