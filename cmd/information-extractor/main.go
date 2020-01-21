/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 */

package main

import (
	"fmt"
	"github.com/wso2/information-extractor/pkg/utils"
	"runtime"
)

var version string
var buildDate string

func printHello() {
	fmt.Printf("Running %v tool", utils.ToolName)
	fmt.Printf("Version: %v", version)
	fmt.Printf("Release Date: %v", buildDate)
}

func main() {
	printHello()
	productPath := utils.GetProductPath()
	updateInfo := utils.GetUpdateInfo(productPath)

	setupInfo := utils.SetupInfo{
		Product:     utils.GetProductName(productPath),
		UpdateLevel: updateInfo.UpdateLevel,
		Channel:     updateInfo.Channel,
		OS:          runtime.GOOS,
		Files:       utils.GetFileData(productPath),
	}

	utils.WriteToFile(setupInfo)
}
