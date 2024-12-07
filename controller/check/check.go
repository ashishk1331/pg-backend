package check

import (
	"github.com/gin-gonic/gin"
	"pg-backend/util"
	"pg-backend/template"
	"sync"
)

type CodeSample = util.CodeSample

func Post(c *gin.Context) {
	var wg sync.WaitGroup

	var questionId int = util.Stoi(c.Query("question_id"))
	var testId int = util.Stoi(c.Query("test_id"))
	var code CodeSample = CodeSample{QuestionId: questionId, TestId: testId, OK: true}

	c.ShouldBind(&code)

	var content string = template.GenerateBase(&code)

	wg.Add(1)
	go util.RunCode(&wg, &code, content)

	wg.Wait()
	c.JSON(200, code)
}