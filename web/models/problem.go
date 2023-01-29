package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	redisdao "ahutoj/web/dao/redisDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/mapping"
	"ahutoj/web/utils"
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
)

// 判断题目是否存在
func IsProblemExistByPID(ctx context.Context, problem *dao.Problem) bool {
	count, err := mysqldao.SelectProblemCountByPID(ctx, problem.PID)
	if err != nil {
		return false
	}
	return count > 0
}

// 创建题目
func CreateProblem(ctx context.Context, problem *dao.Problem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertProblemTable(ctx, *problem)
	if err != nil {
		logger.Errorf("call InsertProblemTable failed,problem= %+v, err=%s", utils.Sdump(problem), err.Error())
	}
	return err
}

// 编辑题目
func EditProblem(ctx context.Context, problem *dao.Problem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.EditProblemTable(ctx, *problem)
	if err != nil {
		logger.Errorf("call EditProblemTable failed,problem= %+v, err=%s", utils.Sdump(problem), err.Error())
	}
	return err
}

func DeleteProblem(ctx context.Context, PID string) error {
	logger := utils.GetLogInstance()
	err := mysqldao.DeleteProblem(ctx, PID)
	if err != nil {
		logger.Errorf("call DeleteProblem failed,problem= %d, err=%s", PID, err.Error())
	}
	return err
}

// 前期的话 先用 mysql 后期针对活跃数据会引入redis
func GetProblemByPID(ctx context.Context, PID string) (dao.Problem, error) {
	logger := utils.GetLogInstance()
	problem := dao.Problem{}
	problem.PID = PID
	err := mysqldao.SelectProblemByPID(ctx, &problem)
	if err != nil {
		logger.Errorf("call SelectProblemByPID failed,PID=%d,err=%s", PID, err.Error())
		return problem, err
	}
	return problem, err
}

func GetProblemCount(ctx context.Context, problem dao.Problem) (int64, error) {
	return mysqldao.SelectProblemCount(ctx, problem)
}

func GetProblems(ctx context.Context, PIDs []string) ([]dao.Problem, error) {
	problems := make([]dao.Problem, len(PIDs))
	logger := utils.GetLogInstance()
	for idx, PID := range PIDs {
		problem, err := GetProblemByPID(ctx, PID)
		if err != nil {
			logger.Errorf("call GetProblemByPID failed,err=%s", err.Error())
			return nil, err
		}
		problems[idx] = problem
	}
	return problems, nil
}

func GetProblemList(ctx context.Context, offset, size int, problem dao.Problem) ([]dao.Problem, error) {
	return mysqldao.SelectListProblem(ctx, offset, size, problem)
}

func ChekckProblemType(ctx context.Context, PType constanct.ProblemType) bool {
	if PType == "" {
		return true
	}
	if PType == constanct.LOCALTYPE {
		return true
	}
	if PType == constanct.ATCODERTYPE {
		return true
	}
	if PType == constanct.CODEFORCESTYPE {
		return true
	}
	return false
}

func GetNextProblemPID(ctx context.Context) (string, error) {
	logger := utils.GetLogInstance()
	PID, err := redisdao.GetLastANDPID(ctx)
	if err != nil || PID == 0 {
		PID, err = mysqldao.SelectProblemLastPID(ctx)
		logger.Debugf("PID:%v", PID)

		if err != nil {
			logger.Errorf("call SelectProblemLastPID failed,err:%v", err.Error())
			return "", err
		}
	}
	return fmt.Sprintf("%v", PID+1), nil
}

// 解析json
func ParseJsonToProblem(ctx context.Context, fileText string) (mapping.JsonProblems, error) {
	jsonData := []byte(fileText)
	var jproblem mapping.JsonProblems
	err := json.Unmarshal(jsonData, &jproblem)
	if err != nil {
		logger.Errorf("call ParseJsonToDB failed, err:%v", err.Error())
		return jproblem, err
	}
	return jproblem, nil
}

// 解析xml
func ParseXmlToproblem(ctx context.Context, fileText string) (mapping.JsonProblems, error) {
	logger := utils.GetLogInstance()
	xmlData := []byte(fileText)
	var xmlProblems mapping.XMLProblems
	var problems mapping.JsonProblems
	err := xml.Unmarshal(xmlData, &xmlProblems)
	problems = mapping.XMLToJsonProblem(xmlProblems)
	if err != nil {
		logger.Errorf("call ParseXmlToDB failed, err:%v", err.Error())
		return problems, err
	}
	return problems, err
}

// 下载json数据
func ParseProblemToJsonProblem(ctx context.Context, PID string) (mapping.JsonProblem, error) {
	var result mapping.JsonProblem
	problem, err := GetProblemByPID(ctx, PID)
	if err != nil {
		logger.Errorf("call ParseDBToJson failed, err:%v", err.Error())
		return result, err
	}
	result = mapping.ProblemToJsonProblem(problem)
	getProblemSampleFile(ctx, &result, PID)
	result.ImgList = GetProblemImageByProblem(ctx, problem)
	return result, nil
}

// 获取样例文件
func getProblemSampleFile(ctx context.Context, problem *mapping.JsonProblem, PID string) error {
	filepath := utils.GetConfInstance().DataPath + "/" + PID
	ok, err := utils.CheckPathExists(filepath)
	if err != nil {
		logger.Errorf("call CheckPathExists faile,filepath=%s err=%s", filepath, err.Error())
		return err
	}
	if !ok {
		return nil
	}

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		logger.Errorf("call ReadDir faile,filepath=%s err=%s", filepath, err.Error())
		return err
	}
	problem.DataSize = len(files)
	problem.Data = make([]mapping.DataFile, 0)
	for _, finfo := range files {
		datafile := mapping.DataFile{}
		if utils.ChekfileHashSuffix(finfo.Name(), "in", "out") {
			datafile.FileName = finfo.Name()
			filename := filepath + "/" + finfo.Name()
			content, err := os.ReadFile(filename)
			if err != nil {
				logger.Errorf("call ReadFile faile,fileName=%s, err=%s", filename, err.Error())
				return err
			}
			datafile.Data = string(content)
			problem.Data = append(problem.Data, datafile)
		}
	}
	return nil
}

// 生成json代码
func ParseJsonProblemToJson(ctx context.Context, problems mapping.JsonProblems) (string, error) {
	buf, err := json.Marshal(problems) //结构体对象生成json字符串
	if err != nil {
		logger.Errorf("call ParseDBToJson failed, err:%v", err.Error())
		return "", err
	}
	var str bytes.Buffer //格式化json字符串
	err = json.Indent(&str, buf, "", "\t")
	if err != nil {
		logger.Errorf("call ParseDBToJson failed, err:%v", err.Error())
		return "", err
	}
	return str.String(), nil
}

// 创建测试文件
func CreateTestFile(ctx context.Context, PID string, datas []mapping.DataFile) error {
	var passSuffix = []string{"in", "out", "zip"}
	path := utils.GetConfInstance().DataPath + "/" + PID
	if len(datas) == 0 {
		return nil
	}
	os.Mkdir(path, 0777)
	for _, data := range datas {
		if !utils.ChekfileHashSuffix(data.FileName, passSuffix...) {
			return fmt.Errorf("file type %v is not supported", utils.GetFileSuffix(data.FileName))
		}
		ioutil.WriteFile(path+"/"+data.FileName, []byte(data.Data), 0666)
	}
	return nil
}
func GetProblemImageByProblem(ctx context.Context, problem dao.Problem) []mapping.ImgItem {
	ret := GetProblemImgByProblemByHTML(ctx, problem)
	retMarkDown := GetProblemImgByProblemByMarkDown(ctx, problem)
	ret = append(ret, retMarkDown...)
	return ret
}

// 获得题目的图片HTML
func GetProblemImgByProblemByHTML(ctx context.Context, problem dao.Problem) []mapping.ImgItem {
	re := regexp.MustCompile("<[iI][mM][gG][^<>]+[sS][rR][cC]=\"?([^ \">]+)/?>")
	descprtionImg := re.FindAllStringSubmatch(problem.Description, -1)
	inputImg := re.FindAllStringSubmatch(problem.Input, -1)
	outputImg := re.FindAllStringSubmatch(problem.Output, -1)
	hitImg := re.FindAllStringSubmatch(problem.Hit, -1)

	result := make([]mapping.ImgItem, 0)
	result = append(result, ImageStringToImgItem(ctx, descprtionImg)...)
	result = append(result, ImageStringToImgItem(ctx, inputImg)...)
	result = append(result, ImageStringToImgItem(ctx, outputImg)...)
	result = append(result, ImageStringToImgItem(ctx, hitImg)...)
	return result
}

// 获得题目的图片MarkDown
func GetProblemImgByProblemByMarkDown(ctx context.Context, problem dao.Problem) []mapping.ImgItem {
	re := regexp.MustCompile(`!\[\]\((.*?)\)`)
	descprtionImg := re.FindAllStringSubmatch(problem.Description, -1)
	inputImg := re.FindAllStringSubmatch(problem.Input, -1)
	outputImg := re.FindAllStringSubmatch(problem.Output, -1)
	hitImg := re.FindAllStringSubmatch(problem.Hit, -1)

	result := make([]mapping.ImgItem, 0)
	result = append(result, ImageStringToImgItem(ctx, descprtionImg)...)
	result = append(result, ImageStringToImgItem(ctx, inputImg)...)
	result = append(result, ImageStringToImgItem(ctx, outputImg)...)
	result = append(result, ImageStringToImgItem(ctx, hitImg)...)
	return result
}
func ImageStringToImgItem(ctx context.Context, imgSrc [][]string) []mapping.ImgItem {
	ret := make([]mapping.ImgItem, 0)
	for _, img := range imgSrc {
		temp := mapping.ImgItem{
			ImageSrc: img[1],
		}
		Strs := strings.Split(temp.ImageSrc, "/")
		temp.ImageName = Strs[len(Strs)-1]
		imageData, _ := os.ReadFile(utils.GetConfInstance().ImagePath + temp.ImageName)
		temp.ImageData = utils.EncodeBase64FromByte(imageData)
		ret = append(ret, temp)
	}
	return ret
}
