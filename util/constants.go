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

type Schema struct {
	HistogramTechniques HistogramTechniques `json:"histogram_techniques"`
	GraphicsAlgorithms  GraphicsAlgorithms  `json:"graphics_algorithms"`
}

type GraphicsAlgorithms struct {
	FlipVertical QuestionSchema `json:"flip_vertical"`
	Rotate90     QuestionSchema `json:"rotate_90"`
}

type HistogramTechniques struct {
	Brightness             QuestionSchema `json:"brightness"`
	Contrast               QuestionSchema `json:"contrast"`
	Histogram              QuestionSchema `json:"histogram"`
	HistogramEqualization  QuestionSchema `json:"histogram_equalization"`
	HistogramSpecification QuestionSchema `json:"QuestionSchema"`
}

type QuestionSchema struct {
	Id        string   `json:"id"`
	FnName    string   `json:"fn_name"`
	Signature []string `json:"signature"`
	Input     []string `json:"input"`
	Output    []string `json:"output"`
}
