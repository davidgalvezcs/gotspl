/*
 * Copyright 2020 Anton Globa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package gotspl

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	DOWNLOAD_DATANAME                                         = "DOWNLOAD"
	DOWNLOAD_DATASTORGAE_DRAM             DownloadDataStorage = ""
	DOWNLOAD_DATASTORAGE_FLASH            DownloadDataStorage = "F"
	DOWNLOAD_DATASTORAGE_EXPANSION_MODULE DownloadDataStorage = "E"
)

type DownloadDataStorage string

type DownloadDataImpl struct {
	storage *string
	name    *string
	data    *[]byte
}

type DownloadDataBuilder interface {
	TSPLCommand
	Storage(storage DownloadDataStorage) DownloadDataBuilder
	Name(name string) DownloadDataBuilder
	Data(data []byte) DownloadDataBuilder
	// DataString(data string) DownloadDataBuilder
}

func DownloadDataCmd() DownloadDataBuilder {
	return DownloadDataImpl{}
}

func (d DownloadDataImpl) GetMessage() ([]byte, error) {

	if (d.name == nil || len(*d.name) == 0) ||
		(
		// (d.dataString == nil || len(*d.dataString) == 0) &&
		d.data == nil) {
		return nil, errors.New("ParseError DOWNLOAD Command: name should be specified")
	}

	buf := bytes.NewBufferString(DOWNLOAD_DATANAME)
	buf.WriteString(EMPTY_SPACE)
	if d.storage != nil {
		buf.WriteString(*d.storage)
		buf.WriteString(VALUE_SEPARATOR)
	}
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(*d.name)
	buf.WriteString(DOUBLE_QUOTE)

	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(fmt.Sprintf("%d", len(*d.data)))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(*d.data)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (d DownloadDataImpl) Storage(storage DownloadDataStorage) DownloadDataBuilder {
	if d.storage == nil {
		d.storage = new(string)
	}
	*d.storage = string(storage)
	return d
}

func (d DownloadDataImpl) Name(name string) DownloadDataBuilder {
	if d.name == nil {
		d.name = new(string)
	}
	*d.name = name
	return d
}

func (d DownloadDataImpl) Data(data []byte) DownloadDataBuilder {
	if d.data == nil {
		d.data = new([]byte)
	}
	*d.data = data
	return d
}
