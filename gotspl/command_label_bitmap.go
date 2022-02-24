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
	BITMAP_NAME = "BITMAP"
)

const (
	BITMAP_MODE_OVERWRITE BitmapMode = iota
	BITMAP_MODE_OR
	BITMAP_MODE_XOR
)

type BitmapMode int

type BitmapImpl struct {
	xCoordinate *int
	yCoordinate *int
	width       *int
	height      *int
	mode        *int
	bitmapData  *[]byte
}

type BitmapBuilder interface {
	TSPLCommand
	XCoordinate(x int) BitmapBuilder
	YCoordinate(y int) BitmapBuilder
	Width(dots int) BitmapBuilder
	Height(dots int) BitmapBuilder
	Mode(mode BitmapMode) BitmapBuilder
	BitmapData(bitmapdata []byte) BitmapBuilder
}

func Bitmap() BitmapBuilder {
	return BitmapImpl{}
}

func (t BitmapImpl) GetMessage() ([]byte, error) {
	if t.xCoordinate == nil ||
		t.yCoordinate == nil ||
		t.width == nil ||
		t.height == nil ||
		t.mode == nil ||
		t.bitmapData == nil {
		return nil, errors.New("ParseError BLOCK Command: " +
			"xCoordinate, yCoordinate, width, height, fontName, rotation, xMultiplication, yMultiplication and content should be specified")
	}

	buf := bytes.NewBufferString(BITMAP_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*t.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.width))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.height))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.mode))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(*t.bitmapData)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (t BitmapImpl) XCoordinate(x int) BitmapBuilder {
	if t.xCoordinate == nil {
		t.xCoordinate = new(int)
	}
	*t.xCoordinate = x
	return t
}

func (t BitmapImpl) YCoordinate(y int) BitmapBuilder {
	if t.yCoordinate == nil {
		t.yCoordinate = new(int)
	}
	*t.yCoordinate = y
	return t
}

func (t BitmapImpl) Width(bytes int) BitmapBuilder {
	if t.width == nil {
		t.width = new(int)
	}
	*t.width = bytes
	return t
}

func (t BitmapImpl) Height(dots int) BitmapBuilder {
	if t.height == nil {
		t.height = new(int)
	}
	*t.height = dots
	return t
}

func (t BitmapImpl) Mode(mode BitmapMode) BitmapBuilder {
	if t.mode == nil {
		t.mode = new(int)
	}
	*t.mode = int(mode)
	return t
}

func (t BitmapImpl) BitmapData(bitmapdata []byte) BitmapBuilder {
	if t.bitmapData == nil {
		t.bitmapData = new([]byte)
	}
	*t.bitmapData = bitmapdata
	return t
}
