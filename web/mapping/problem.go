package mapping

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/request"
	"ahutoj/web/utils"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type JsonProblems []JsonProblem
type JsonProblem struct {
	PType        constanct.ProblemType `json:"PType"`
	Title        string                `json:"Title"`
	Description  string                `json:"Description"`
	ImgList      []ImgItem             `json:"ImgList"`
	Input        string                `json:"Input"`
	Output       string                `json:"Output" `
	SampleInput  string                `json:"SampleInput"`
	SampleOutput string                `json:"SampleOutput"`
	LimitTime    int64                 `json:"LimitTime"`
	LimitMemory  int64                 `json:"LimitMemory"`
	Hit          string                `json:"Hit"`
	Origin       int64                 `json:"Origin"`
	OriginPID    string                `json:"OriginPID"`
	ContentType  int64                 `json:"ContentType"`
	SpjJudge     string                `json:"SpjJudge"`
	SpjSource    string                `json:"SpjSource"`
	Source       string                `json:"Source"`
	DataSize     int                   `json:"DataSize"`
	Data         []DataFile            `json:"Data"`
}
type DataFile struct {
	FileName string `json:"FileName"`
	Data     string `json:"Data"`
}
type XMLProblems struct {
	Items []XMLItem `xml:"item"`
}
type XMLItem struct {
	XMLName         xml.Name       `xml:"item"`
	Title           string         `xml:"title"`
	TimeLimitItem   XMLTimeLimit   `xml:"time_limit"`
	MemoryLimitItem XMLMemoryLimit `xml:"memory_limit"`
	ImgList         []ImgItem      `xml:"img"`
	Description     string         `xml:"description"`
	Input           string         `xml:"input"`
	OutPut          string         `xml:"output"`
	Source          string         `xml:"source"`
	SampleInput     string         `xml:"sample_input"`
	SampleOutput    string         `xml:"sample_output"`
	SpjSource       string         `xml:"spj"`
	TpjSource       string         `xml:"tpj"`
	Hit             string         `xml:"hit"`
	TestInPut       []string       `xml:"test_input"`
	TestOutPut      []string       `xml:"test_output"`
}
type XMLTimeLimit struct {
	xml.Name `xml:"time_limit"`
	Data     string `xml:",innerxml"`
	TimeUnit string `xml:"unit,attr"`
}
type XMLMemoryLimit struct {
	xml.Name   `xml:"memory_limit"`
	Data       string `xml:",innerxml"`
	MemoryUnit string `xml:"unit,attr"`
}
type ImgItem struct {
	ImageName string `json:"ImageName"`
	ImageSrc  string `json:"ImageSrc" xml:"src"`
	ImageData string `json:"ImageData" xml:"base64"`
}

func ProblemReqToDao(req request.Problem) dao.Problem {
	problem := dao.Problem{}
	if req.PID != nil {
		problem.PID = *req.PID
	}
	// 	LimitMemory:   req.LimitMemory,
	if req.Title != nil {
		problem.Title = *req.Title
	}
	if req.Description != nil {
		problem.Description = *req.Description
	}
	if req.Input != nil {
		problem.Input = *req.Input
	}
	if req.Output != nil {
		problem.Output = *req.Output
	}
	if req.Sample_input != nil {
		problem.SampleInput = *req.Sample_input
	}
	if req.Sample_output != nil {
		problem.SampleOutput = *req.Sample_output
	}
	if req.Hit != nil {
		problem.Hit = *req.Hit
	}
	if req.Label != nil {
		problem.Label = *req.Label
	}
	if req.LimitTime != nil {
		problem.LimitTime = *req.LimitTime
	}
	if req.LimitMemory != nil {
		problem.LimitMemory = *req.LimitMemory
	}
	if req.Origin != nil {
		problem.Origin = *req.Origin
	}
	if req.OriginPID != nil {
		problem.OriginPID = *req.OriginPID
	}
	problem.PType = req.PType
	problem.ContentType = req.ContentType
	problem.Visible = req.Visible
	if req.SpjJudge != nil {
		problem.SpjJudge = *req.SpjJudge
	}
	if req.Source != nil {
		problem.Source = *req.Source
	}
	return problem
}

func ProblemToJsonProblem(problem dao.Problem) JsonProblem {
	result := JsonProblem{}
	result.PType = problem.PType
	result.Title = problem.Title
	result.Description = problem.Description
	result.Input = problem.Input
	result.Output = problem.Output
	result.SampleInput = problem.SampleInput
	result.SampleOutput = problem.SampleOutput
	result.LimitTime = problem.LimitTime
	result.LimitMemory = problem.LimitMemory
	result.Hit = problem.Hit
	result.SpjJudge = problem.SpjJudge
	result.Origin = problem.Origin
	result.OriginPID = problem.OriginPID
	result.ContentType = problem.ContentType
	return result
}

func JsonProblemToProblem(problem JsonProblem) dao.Problem {
	result := dao.Problem{}
	result.PType = problem.PType
	result.Title = problem.Title
	result.Description = problem.Description
	result.Input = problem.Input
	result.Output = problem.Output
	result.SampleInput = problem.SampleInput
	result.SampleOutput = problem.SampleOutput
	result.LimitTime = problem.LimitTime
	result.LimitMemory = problem.LimitMemory
	result.Hit = problem.Hit
	result.SpjJudge = problem.SpjJudge
	result.Origin = problem.Origin
	result.OriginPID = problem.OriginPID
	result.ContentType = problem.ContentType
	result.Visible = 1
	return result
}

func XMLToJsonProblem(problems XMLProblems) JsonProblems {
	result := make(JsonProblems, 0)
	for _, problem := range problems.Items {
		jproblem := JsonProblem{
			PType:       constanct.LOCALTYPE,
			Title:       problem.Title,
			SpjSource:   problem.SpjSource,
			Source:      problem.Source,
			Origin:      -1,
			ContentType: 1,
			DataSize:    len(problem.TestInPut),
		}
		TestData := fmt.Sprintf("### 样例输入\n\n```\n%v\n```\n\n### 样例输出\n\n```\n%v\n```\n\n",
			problem.SampleInput, problem.SampleOutput)
		jproblem.Description = fmt.Sprintf("## 题目描述\n\n%v\n\n## 输入格式\n\n%v\n\n## 输出格式\n\n%v\n\n## 题目样例\n\n%v\n\n## 提示\n\n%v",
			problem.Description,
			problem.Input,
			problem.OutPut,
			TestData,
			problem.Hit,
		)
		if problem.SpjSource != "" {
			jproblem.SpjJudge = "Y"
			jproblem.SpjSource = problem.SpjSource
		}

		if problem.TpjSource != "" {
			jproblem.SpjJudge = "Y"
			jproblem.SpjSource = problem.TpjSource
		}

		jproblem.Data = make([]DataFile, 0)
		for idx, input := range problem.TestInPut {
			filename := fmt.Sprintf("%d.in", idx+1)
			var datafile DataFile
			datafile.FileName = filename
			datafile.Data = input
			jproblem.Data = append(jproblem.Data, datafile)
		}

		for idx, output := range problem.TestOutPut {
			filename := fmt.Sprintf("%d.out", idx+1)
			var datafile DataFile
			datafile.FileName = filename
			datafile.Data = output
			jproblem.Data = append(jproblem.Data, datafile)
		}
		times := GetCdataValue(problem.TimeLimitItem.Data)
		time, _ := strconv.ParseFloat(times, 64)
		memorys := GetCdataValue(problem.MemoryLimitItem.Data)
		memory, _ := strconv.ParseFloat(memorys, 64)
		switch problem.TimeLimitItem.TimeUnit {
		case "s":
			{
				jproblem.LimitTime = int64(time * 1000)
				break
			}
		case "ms":
			{
				jproblem.LimitTime = int64(time)
				break
			}
		default:
			{
				break
			}
		}
		switch problem.MemoryLimitItem.MemoryUnit {
		case "mb":
			{
				jproblem.LimitMemory = int64(memory)
				break
			}
		case "kb":
			{
				jproblem.LimitMemory = int64(memory / 1024)
				break
			}
		default:
			{
				break
			}
		}
		jproblem.ImgList = problem.ImgList
		for idx, img := range jproblem.ImgList {
			names := strings.Split(img.ImageSrc, "/")
			jproblem.ImgList[idx].ImageName = names[len(names)-1]
		}
		result = append(result, jproblem)
	}
	return result
}

func GetCdataValue(data string) string {
	return data[9 : len(data)-3]
}

func FixProblemURL(oldFilename, newFilename string, problem *JsonProblem) {
	problem.Description = strings.ReplaceAll(problem.Description, oldFilename, newFilename)
	problem.Input = strings.ReplaceAll(problem.Input, oldFilename, newFilename)
	problem.Output = strings.ReplaceAll(problem.Output, oldFilename, newFilename)
	problem.Hit = strings.ReplaceAll(problem.Hit, oldFilename, newFilename)
}

func SaveJproblemImage(jproblem *JsonProblem) error {
	for _, image := range jproblem.ImgList {
		data, err := utils.DeCodeBase64ToByte(image.ImageData)
		if err != nil {
			return err
		}
		filename := utils.GetFileName(image.ImageName, data)
		err = os.WriteFile(utils.GetConfInstance().ImagePath+filename, data, 0666)
		if err != nil {
			return err
		}
		FixProblemURL(image.ImageSrc, "image/"+filename, jproblem)
		if err != nil {
			return err
		}
	}
	return nil
}
