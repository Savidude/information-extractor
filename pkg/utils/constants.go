/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 */

package utils

import "os"

const ToolName string = "Information Extractor"

const SetupInfoFilename string = "setup-info.json"

const PathSeparator = string(os.PathSeparator)

const UpdatesDirectory string = "updates"
const ProductFile string = "product.txt"
const MetadataFile string = "metadata.json"

const DefaultErrorTemplate = "{{.Error}}"

// Exit Codes
const DefaultError = 1
const FileSystemError = 6
