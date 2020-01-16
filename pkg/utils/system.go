/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 */

package utils

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

type ErrorTemplate struct {
	Error string
}

var inputTemplate string
var isDefaultTemplateEnabled bool

var printlnOnTemplateFunc = doNothinglnFuncOnTemplate
var printfFunc = printfFunction

func HandleErrorAndExit(msg string, err error, exitCode int) {
	if inputTemplate != "" {
		inputTemplate = DefaultErrorTemplate
	}
	if err == nil {
		t := ErrorTemplate{fmt.Sprintf(ToolName + ": %v", msg)}
		printTemplate(isDefaultTemplateEnabled, DefaultErrorTemplate, inputTemplate, t)
		printf(os.Stderr, ToolName + ": %v\n", msg)
	} else {
		t := ErrorTemplate{fmt.Sprintf(ToolName + ": %v reason: %v\n", msg, err.Error())}
		printTemplate(isDefaultTemplateEnabled, DefaultErrorTemplate, inputTemplate, t)
		printf(os.Stderr, ToolName + ": %v reason: %v\n", msg, err.Error())
	}
	os.Exit(exitCode)
}

func printTemplate(isDefaultTemplateEnabled bool, defaultTemplate, inputTemplate string, tmpl interface{}) {
	if isDefaultTemplateEnabled {
		inputTemplate = defaultTemplate
	}
	t, err := template.New("Template").Parse(inputTemplate)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(DefaultError)
	}
	err = t.Execute(os.Stdout, tmpl)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(DefaultError)
	}
	printLnOnTemplate(os.Stdout)
}

func printf(writer io.Writer, format string, a ...interface{}) {
	printfFunc(writer, format, a...)
}

func printLnOnTemplate(writer io.Writer, a ...interface{}) {
	printlnOnTemplateFunc(writer, a...)
}

func printfFunction(writer io.Writer, format string, a ...interface{}) {
	fmt.Fprintf(writer, format, a...)
}

func doNothinglnFuncOnTemplate(writer io.Writer, v ...interface{}) {
}
