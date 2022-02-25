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
	PUTBMP_NAME = "PUTBMP"
)

const (
	PUTBMP_BPP_1BIT PutbmpBpp = 1
	PUTBMP_BPP_8BIT PutbmpBpp = 8
)

type PutbmpBpp int

type PutbmpImpl struct {
	xCoordinate *int
	yCoordinate *int
	fileName    *string
	bpp         *int
	contrast    *int
}

type PutbmpBuilder interface {
	TSPLCommand
	XCoordinate(x int) PutbmpBuilder
	YCoordinate(y int) PutbmpBuilder
	FileName(filename string) PutbmpBuilder
	Bpp(bpp PutbmpBpp) PutbmpBuilder
	Contrast(contrast int) PutbmpBuilder
}

func Putbmp() PutbmpBuilder {
	return PutbmpImpl{}
}

func (t PutbmpImpl) GetMessage() ([]byte, error) {
	if t.xCoordinate == nil ||
		t.yCoordinate == nil ||
		t.fileName == nil || len(*t.fileName) == 0 {
		return nil, errors.New("ParseError PUTBMP Command: " +
			"xCoordinate, yCoordinate, width, height, fontName, rotation, xMultiplication, yMultiplication and content should be specified")
	}

	buf := bytes.NewBufferString(PUTBMP_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*t.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(*t.fileName)
	buf.WriteString(DOUBLE_QUOTE)
	if t.bpp != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.WriteString(strconv.Itoa(*t.bpp))
	}
	if t.contrast != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.WriteString(strconv.Itoa(*t.contrast))
	}
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (t PutbmpImpl) XCoordinate(x int) PutbmpBuilder {
	if t.xCoordinate == nil {
		t.xCoordinate = new(int)
	}
	*t.xCoordinate = x
	return t
}

func (t PutbmpImpl) YCoordinate(y int) PutbmpBuilder {
	if t.yCoordinate == nil {
		t.yCoordinate = new(int)
	}
	*t.yCoordinate = y
	return t
}

func (t PutbmpImpl) FileName(filename string) PutbmpBuilder {
	if t.fileName == nil {
		t.fileName = new(string)
	}
	*t.fileName = filename
	return t
}

func (t PutbmpImpl) Bpp(bpp PutbmpBpp) PutbmpBuilder {
	if t.bpp == nil {
		t.bpp = new(int)
	}
	*t.bpp = int(bpp)
	return t
}

func (t PutbmpImpl) Contrast(contrast int) PutbmpBuilder {
	if t.contrast == nil {
		t.contrast = new(int)
	}
	*t.contrast = contrast
	return t
}
