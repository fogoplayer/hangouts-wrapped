package model

import "zarinloosli.com/hangouts-wrapped/browserApis"

var ChatDataDirectoryChannel chan browserApis.DirectoryHandle = make(chan browserApis.DirectoryHandle)
