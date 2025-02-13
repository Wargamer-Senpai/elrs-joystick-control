// SPDX-FileCopyrightText: © 2023 OneEyeFPV oneeyefpv@gmail.com
// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-License-Identifier: FS-0.9-or-later

package config

import "github.com/kaack/elrs-joystick-control/pkg/util"

type MultiplyT struct {
	RightDefault *util.RawValue `json:"right_default"`

	Left  *IOHolder    `json:"left"`
	Right *[]*IOHolder `json:"right"`
}

// InputMultiply *** Multiply Operation ***
type InputMultiply struct {
	Id    string        `json:"id"`
	Value util.RawValue `json:"value"`
	IsNaN bool          `json:"-"`

	Type     string      `json:"type"`
	Multiply MultiplyT   `json:"multiply" input:"true"`
	Holder   *IOHolder   `json:"-"`
}

func MultiplyOperation(left util.RawValue, right util.RawValue) util.RawValue {
	return left * right
}

func (i *InputMultiply) Eval(c *Config) (src IOType, out util.RawValue, ch util.ChannelNumber, nan bool) {
	src, out, ch, nan = EvalOperation(c, i.Multiply.Left, i.Multiply.Right, i.Multiply.RightDefault, MultiplyOperation)
	i.Value = out
	i.IsNaN = nan

	return src, out, ch, nan
}

func (i *InputMultiply) InputType() string {
	return i.Type
}

func (i *InputMultiply) InputValue() *util.RawValue {
	if i.IsNaN {
		return nil
	}
	return &i.Value
}

func (i *InputMultiply) InputId() string {
	return i.Id
}

func (i *InputMultiply) Children() (out *[]*IOHolder) {
	return GetChildren(i.Multiply.Left, i.Multiply.Right)
}
