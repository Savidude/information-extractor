/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileData struct {
	File    string `json:"file"`
	MD5Hash string `json:"md5hash"`
}

func WriteToFile(setupInfo SetupInfo) {
	setupInfoJson, err := json.MarshalIndent(setupInfo, "", "\t")
	if err != nil {
		HandleErrorAndExit(JSONMarshalError, err, DefaultError)
	}
	err = ioutil.WriteFile(SetupInfoFilename, setupInfoJson, 0644)
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToWriteFileMsg, SetupInfoFilename), err, FileSystemError)
	}
	fmt.Printf("Information on deployment written to file %v.", SetupInfoFilename)
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToCloseFileMsg, file.Name()), err, FileSystemError)
	}
}

func scanRecursive(dirPath string) []string {
	var files []string

	err := filepath.Walk(dirPath, func(path string, fileStatus os.FileInfo, err error) error {
		fileStatus, err = os.Stat(path)
		if err != nil {
			HandleErrorAndExit(fmt.Sprintf(UnableToCheckPathStatus, path), err, FileSystemError)
		}

		fileMode := fileStatus.Mode()
		if fileMode.IsRegular() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToScanDirectory, dirPath), err, FileSystemError)
	}
	return files
}

func getMD5(filePath string) string {
	var fileMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToOpenFileMsg, filePath), err, FileSystemError)
	}
	defer CloseFile(file)

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		HandleErrorAndExit(fmt.Sprintf(UnableToGetMD5ofFile, filePath), err, FileSystemError)
	}
	hashInBytes := hash.Sum(nil)[:16]
	fileMD5String = hex.EncodeToString(hashInBytes)
	return fileMD5String
}
