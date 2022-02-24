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
	"strconv"
)

const (
	BLOCK_NAME = "BLOCK"

	BLOCK_MULTIPLIER_MIN = 1
	BLOCK_MULTIPLIER_MAX = 10
)

const (
	BLOCK_ALIGNMENT_DEFAULT BlockAlignment = iota
	BLOCK_ALIGNMENT_LEFT
	BLOCK_ALIGNMENT_CENTER
	BLOCK_ALIGNMENT_RIGHT
)

const (
	BLOCK_FIT_NOSHRINK BlockFit = iota
	BLOCK_FIT_SHRINK
)

type BlockAlignment int
type BlockFit int

type BlockImpl struct {
	xCoordinate     *int
	yCoordinate     *int
	width           *int
	height          *int
	fontName        *string
	rotation        *int
	xMultiplication *float64
	yMultiplication *float64
	space           *int
	alignment       *int
	fit             *int
	content         *string
	contentQuote    bool
}

type BlockBuilder interface {
	TSPLCommand
	XCoordinate(x int) BlockBuilder
	YCoordinate(y int) BlockBuilder
	Width(dots int) BlockBuilder
	Height(dots int) BlockBuilder
	FontName(name string) BlockBuilder
	Rotation(angle int) BlockBuilder
	XMultiplier(xm float64) BlockBuilder
	YMultiplier(ym float64) BlockBuilder
	Space(dots int) BlockBuilder
	Alignment(align BlockAlignment) BlockBuilder
	Fit(fit BlockFit) BlockBuilder
	Content(content string, quote bool) BlockBuilder
}

func Block() BlockBuilder {
	return BlockImpl{}
}

func (t BlockImpl) GetMessage() ([]byte, error) {
	if t.xCoordinate == nil ||
		t.yCoordinate == nil ||
		t.width == nil ||
		t.height == nil ||
		t.fontName == nil ||
		t.rotation == nil ||
		t.content == nil ||
		t.xMultiplication == nil ||
		t.yMultiplication == nil {
		return nil, errors.New("ParseError BLOCK Command: " +
			"xCoordinate, yCoordinate, width, height, fontName, rotation, xMultiplication, yMultiplication and content should be specified")
	}

	if t.rotation != nil {
		if !findIntInSlice(ROTATION_ANGLES, *t.rotation) {

			var err_str string

			for _, v := range ROTATION_ANGLES {
				err_str += strconv.Itoa(v)
				err_str += ","
			}
			return nil, errors.New("ParseError BLOCK Command: " +
				"rotation must be one of [" + err_str[:len(err_str)-1] + "]")
		}
	}

	if t.xMultiplication != nil {
		if !(*t.xMultiplication >= BLOCK_MULTIPLIER_MIN && *t.xMultiplication <= BLOCK_MULTIPLIER_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError BLOCK Command: "+
				"xMultiplication parameter must be between %d and %d", BLOCK_MULTIPLIER_MIN, BLOCK_MULTIPLIER_MAX))
		}
	}

	if t.yMultiplication != nil {
		if !(*t.yMultiplication >= BLOCK_MULTIPLIER_MIN && *t.yMultiplication <= BLOCK_MULTIPLIER_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError BLOCK Command: "+
				"yMultiplication parameter must be between %d and %d", BLOCK_MULTIPLIER_MIN, BLOCK_MULTIPLIER_MAX))
		}
	}

	buf := bytes.NewBufferString(BLOCK_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*t.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.width))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.height))
	buf.WriteString(VALUE_SEPARATOR + DOUBLE_QUOTE)
	buf.WriteString(*t.fontName)
	buf.WriteString(DOUBLE_QUOTE + VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.rotation))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(formatFloatWithUnits(*t.xMultiplication, false))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(formatFloatWithUnits(*t.yMultiplication, false))
	if t.space != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.WriteString(strconv.Itoa(*t.space))
	}
	if t.alignment != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.WriteString(strconv.Itoa(*t.alignment))
	}
	if t.fit != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.WriteString(strconv.Itoa(*t.fit))
	}

	buf.WriteString(VALUE_SEPARATOR)
	if t.contentQuote {
		buf.WriteString(DOUBLE_QUOTE)
	}
	buf.WriteString(*t.content)
	if t.contentQuote {
		buf.WriteString(DOUBLE_QUOTE)
	}
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (t BlockImpl) XCoordinate(x int) BlockBuilder {
	if t.xCoordinate == nil {
		t.xCoordinate = new(int)
	}
	*t.xCoordinate = x
	return t
}

func (t BlockImpl) YCoordinate(y int) BlockBuilder {
	if t.yCoordinate == nil {
		t.yCoordinate = new(int)
	}
	*t.yCoordinate = y
	return t
}

func (t BlockImpl) Width(dots int) BlockBuilder {
	if t.width == nil {
		t.width = new(int)
	}
	*t.width = dots
	return t
}

func (t BlockImpl) Height(dots int) BlockBuilder {
	if t.height == nil {
		t.height = new(int)
	}
	*t.height = dots
	return t
}

func (t BlockImpl) FontName(name string) BlockBuilder {
	if t.fontName == nil {
		t.fontName = new(string)
	}
	*t.fontName = name
	return t
}

func (t BlockImpl) Rotation(angle int) BlockBuilder {
	if t.rotation == nil {
		t.rotation = new(int)
	}
	*t.rotation = angle
	return t
}

func (t BlockImpl) XMultiplier(xm float64) BlockBuilder {
	if t.xMultiplication == nil {
		t.xMultiplication = new(float64)
	}
	*t.xMultiplication = xm
	return t
}

func (t BlockImpl) YMultiplier(ym float64) BlockBuilder {
	if t.yMultiplication == nil {
		t.yMultiplication = new(float64)
	}
	*t.yMultiplication = ym
	return t
}

func (t BlockImpl) Space(dots int) BlockBuilder {
	if t.space == nil {
		t.space = new(int)
	}
	*t.space = dots
	return t
}

func (t BlockImpl) Alignment(align BlockAlignment) BlockBuilder {
	if t.alignment == nil {
		t.alignment = new(int)
	}
	*t.alignment = int(align)
	return t
}

func (t BlockImpl) Fit(dots BlockFit) BlockBuilder {
	if t.fit == nil {
		t.fit = new(int)
	}
	*t.fit = int(dots)
	return t
}

func (t BlockImpl) Content(content string, quote bool) BlockBuilder {
	if t.content == nil {
		t.content = new(string)
	}
	*t.content = content
	t.contentQuote = quote
	return t
}
