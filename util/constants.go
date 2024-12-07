package util

type CodeSample struct {
	QuestionId   int    `json:"question_id"`
	TestId       int    `json:"test_id"`
	Content      string `json:"content"`
	Output       string `json:"output"`
	OK           bool   `json:"ok"`
	TimeTaken    int    `json:"time_taken"`
	TimeUnits    string `json:"time_units"`
	Error        string `json:"error"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
