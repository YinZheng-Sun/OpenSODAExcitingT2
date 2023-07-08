package service

import (
	"exciting-opendigger/utils"
	"fmt"
	"testing"
)

func TestSingleDownloadService(t *testing.T) {
	fmt.Println("TestSingleDownloadService：")
	downloadService := &SingleDownloadService{}
	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["metricOne"] = testData1
	testDataMap["metricTwo"] = testData2

	testDataMap["bus_factor_detail"] = nil
	testDataMap["new_contributors_detail"] = nil
	testDataMap["activity_details"] = nil

	myMap := make(map[string]([][]string))

	// 添加值到 map 中
	myMap["2020-08"] = [][]string{{"sunshinemingo1", "5"}, {"frank-zsy1", "12"}}
	myMap["2020-09"] = [][]string{{"sunshinemingo2", "50"}, {"frank-zsy2", "6"}}
	myMap["2020-10"] = [][]string{{"sunshinemingo3", "500"}, {"frank-zsy3", "120"}}

	myMap2 := make(map[string]([]string))

	// 添加值到 map 中
	myMap2["2020-08"] = []string{"sunshinemingo1new", "frank-zsy1new"}
	myMap2["2020-09"] = []string{"sunshinemingo2new", "frank-zsy2new"}
	myMap2["2020-10"] = []string{"sunshinemingo3new", "frank-zsy3new"}

	testSpecial := &utils.SpecialDataStructure{BusFactorDetail: myMap, ActivityDetails: myMap, NewContributorsDetail: myMap2}

	ret := RepoInfo{
		RepoName:    "opendigger",
		RepoUrl:     "www.github.com/X-lab2017/open-digger",
		Month:       "",
		Dates:       testDates,
		Data:        testDataMap,
		SpecialData: *testSpecial,
	}

	err := downloadService.SetData(ret, "html_output")
	if err != nil {
		return
	}
	err2 := downloadService.Download()
	if err2 != nil {
		return
	}
}

func TestBatchDownloadService(t *testing.T) {
	fmt.Println("TestBatchDownloadService：")
	downloadService := &BatchDownloadService{}

	var rets []RepoInfo

	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["metricOne"] = testData1
	testDataMap["metricTwo"] = testData2
	ret := RepoInfo{
		RepoName: "opendigger",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	rets = append(rets, ret)

	ret2 := RepoInfo{
		RepoName: "opendigger2",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	rets = append(rets, ret2)

	ret3 := RepoInfo{
		RepoName: "opendigger3",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	rets = append(rets, ret3)

	err := downloadService.SetData(rets, "metricOne", "csv_output")
	if err != nil {
		return
	}
	err2 := downloadService.Download()
	if err2 != nil {
		return
	}
}

func TestCompareDownloadService(t *testing.T) {
	fmt.Println("TestCompareDownloadService：")
	downloadService := &CompareDownloadService{}
	testDates := []string{"2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-10-raw", "2021-11", "2021-12", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"}
	testData1 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testData2 := map[string]interface{}{"2020-08": 4.5, "2020-09": 4.91, "2020-10": 5.59, "2020-11": 6.31, "2020-12": 9.96, "2021-01": 10.61, "2021-02": 6.28, "2021-03": 4.14, "2021-04": 4.44, "2021-05": 4.26, "2021-06": 6.46, "2021-07": 4.84, "2021-08": 3.93, "2021-09": 3.34, "2021-10": 3, "2021-11": 2.89, "2021-12": 3.33, "2022-01": 4.71, "2022-02": 4.87, "2022-03": 6.06, "2022-04": 3.76, "2022-05": 4.14, "2022-06": 7.67, "2022-07": 9.17, "2022-08": 8.53, "2022-09": 9.96, "2022-10": 11.84, "2022-11": 14.65, "2022-12": 19.36, "2023-01": 19.9, "2023-02": 40.48, "2023-03": 22.05, "2023-04": 18.79, "2023-05": 18.42, "2021-10-raw": 2.84}
	testDataMap := make(map[string](map[string]interface{}))
	testDataMap["metricOne"] = testData1
	testDataMap["metricTwo"] = testData2
	ret1 := RepoInfo{
		RepoName: "opendigger1",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	ret2 := RepoInfo{
		RepoName: "opendigger2",
		RepoUrl:  "www.github.com/X-lab2017/open-digger",
		Month:    "",
		Dates:    testDates,
		Data:     testDataMap,
	}

	err := downloadService.SetData(ret1, ret2, "html_output_compare")
	if err != nil {
		return
	}
	err2 := downloadService.Download()
	if err2 != nil {
		return
	}
}
