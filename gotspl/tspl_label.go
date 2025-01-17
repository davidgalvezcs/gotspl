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

import "bytes"

type TSPLLabel struct {
	commandList []TSPLCommand
}

type TSPLLabelBuilder interface {
	TSPLCommandSequence
	Cmd(command TSPLCommand) TSPLLabelBuilder
	GetLabelData() ([]byte, error)
}

func NewTSPLLabel() TSPLLabelBuilder {
	return TSPLLabel{}
}

func (T TSPLLabel) getTSPLCode() ([]byte, error) {
	var buf bytes.Buffer
	for _, c := range T.commandList {
		msg, err := c.GetMessage()
		if err != nil {
			return nil, err
		}

		buf.Write(msg)
	}

	return buf.Bytes(), nil
}

func (T TSPLLabel) Cmd(command TSPLCommand) TSPLLabelBuilder {
	if command != nil {
		T.commandList = append(T.commandList, command)
	}

	return T
}

func (T TSPLLabel) GetLabelData() ([]byte, error) {

	return T.getTSPLCode()
}
