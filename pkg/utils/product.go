/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 */

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type UpdateMedata struct {
	UpdateLevel string `json:"update_level"`
	Channel     string `json:"channel"`
}

type SetupInfo struct {
	Product     string     `json:"product"`
	UpdateLevel string     `json:"update_level"`
	Channel     string     `json:"channel"`
	OS          string     `json:"os"`
	Files       []FileData `json:"files"`
}

func GetProductPath() string {
	executable, err := os.Executable()
	if err != nil {
		HandleErrorAndExit(ExecutablePathError, err, FileSystemError)
	}
	exPath := filepath.Dir(executable)
	productPath := filepath.Dir(exPath)
	return productPath
}

func GetProductName(productPath string) string {
	updatesDir := productPath + PathSeparator + UpdatesDirectory
	productFile := updatesDir + PathSeparator + ProductFile

	file, err := os.Open(productFile)
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToOpenFileMsg, productFile), err, FileSystemError)
	}
	defer CloseFile(file)

	productFileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToReadFileMsg, productFile), err, FileSystemError)
	}

	productName := string(productFileBytes)
	productName = strings.TrimSpace(productName)
	return productName
}

func GetUpdateInfo(productPath string) UpdateMedata {
	updatesDir := productPath + PathSeparator + UpdatesDirectory
	metadataFileName := updatesDir + PathSeparator + MetadataFile

	metadataFile, err := ioutil.ReadFile(metadataFileName)
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToReadFileMsg, metadataFileName), err, FileSystemError)
	}

	var metadata UpdateMedata
	err = json.Unmarshal([]byte(metadataFile), &metadata)
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(JSONParseError, metadataFileName), err, DefaultError)
	}

	return metadata
}

func GetFileData(productPath string) []FileData {
	var filesData []FileData

	files := scanRecursive(productPath)
	for _, file := range files {
		relativeFilePath := strings.Replace(file, productPath+PathSeparator, "", -1)
		fileMD5String := getMD5(file)

		fileData := FileData{
			File:    relativeFilePath,
			MD5Hash: fileMD5String,
		}
		filesData = append(filesData, fileData)
	}

	return filesData
}
