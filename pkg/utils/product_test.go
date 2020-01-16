package utils

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var sampleProductsPath = "../../test/testdata/sample-products"
var sampleFileData = "../../test/testdata/file-data.json"

type SampleProduct struct {
	name        string
	updateLevel string
	channel     string
}

var sampleProducts = []SampleProduct{
	{"wso2am-3.0.0", "3.0.0.0", "full"},
	{"wso2ei-6.5.0", "6.5.0.0", "security"},
	{"wso2is-5.9.0", "5.9.0.0", "private_user"},
}

func TestGetProductName(t *testing.T) {
	for _, product := range sampleProducts {
		productName := GetProductName(sampleProductsPath + PathSeparator + product.name)
		if productName != product.name {
			t.Errorf("GetProductName() : FAILED, expected '%v' but got value '%v'", product.name, productName)
		} else {
			t.Logf("GetProductName() : PASSED, expected '%v' and got value '%v'", product.name, productName)
		}
	}
}

func TestGetUpdateInfo(t *testing.T) {
	for _, product := range sampleProducts {
		updateInfo := GetUpdateInfo(sampleProductsPath + PathSeparator + product.name)
		t.Logf("Testing update info for %v", product.name)

		if updateInfo.UpdateLevel != product.updateLevel {
			t.Errorf("GetUpdateInfo() : FAILED, expected update level '%v' but got value '%v'",
				product.updateLevel, updateInfo.UpdateLevel)
		} else {
			t.Logf("GetUpdateInfo() : PASSED, expected update level '%v' and got value '%v'",
				product.updateLevel, updateInfo.UpdateLevel)
		}

		if updateInfo.Channel != product.channel {
			t.Errorf("GetUpdateInfo() : FAILED, expected channel '%v' but got value '%v'",
				product.channel, updateInfo.Channel)
		} else {
			t.Logf("GetUpdateInfo() : PASSED, expected channel '%v' and got value '%v'",
				product.channel, updateInfo.Channel)
		}
	}
}

func TestGetFileData(t *testing.T) {
	fileData := GetFileData(sampleProductsPath)

	sampleFiledataFile, _ := ioutil.ReadFile(sampleFileData)
	var sampleFileData []FileData
	err := json.Unmarshal([]byte(sampleFiledataFile), &sampleFileData)
	if err != nil {
		panic(err)
	}

	for _, file := range fileData {
		for _, sampleFile := range sampleFileData {
			if file.File == sampleFile.File {
				if file.MD5Hash != sampleFile.MD5Hash {
					t.Errorf("getFileData() : FAILED, md5 hash for file '%v' does not match. Expected %v but got %v.",
						file.File, sampleFile.MD5Hash, file.MD5Hash)
				} else {
					t.Logf("getFileData() : PASSED, md5 hash for file '%v' matched. Expected %v and got %v.",
						file.File, sampleFile.MD5Hash, file.MD5Hash)
				}
				break
			}
		}
	}
}
