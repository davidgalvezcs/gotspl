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
	"strconv"
)

const (
	PUTPCX_NAME = "PUTPCX"
)

type PutpcxImpl struct {
	xCoordinate *int
	yCoordinate *int
	fileName    *string
	bpp         *int
	contrast    *int
}

type PutpcxBuilder interface {
	TSPLCommand
	XCoordinate(x int) PutpcxBuilder
	YCoordinate(y int) PutpcxBuilder
	FileName(filename string) PutpcxBuilder
}

func Putpcx() PutpcxBuilder {
	return PutpcxImpl{}
}

func (t PutpcxImpl) GetMessage() ([]byte, error) {
	if t.xCoordinate == nil ||
		t.yCoordinate == nil ||
		t.fileName == nil || len(*t.fileName) == 0 {
		return nil, errors.New("ParseError PUTPCX Command: " +
			"xCoordinate, yCoordinate, width, height, fontName, rotation, xMultiplication, yMultiplication and content should be specified")
	}

	buf := bytes.NewBufferString(PUTPCX_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*t.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(*t.fileName)
	buf.WriteString(DOUBLE_QUOTE)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (t PutpcxImpl) XCoordinate(x int) PutpcxBuilder {
	if t.xCoordinate == nil {
		t.xCoordinate = new(int)
	}
	*t.xCoordinate = x
	return t
}

func (t PutpcxImpl) YCoordinate(y int) PutpcxBuilder {
	if t.yCoordinate == nil {
		t.yCoordinate = new(int)
	}
	*t.yCoordinate = y
	return t
}

func (t PutpcxImpl) FileName(filename string) PutpcxBuilder {
	if t.fileName == nil {
		t.fileName = new(string)
	}
	*t.fileName = filename
	return t
}
