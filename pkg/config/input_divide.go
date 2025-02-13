// SPDX-FileCopyrightText: © 2023 OneEyeFPV oneeyefpv@gmail.com
// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-License-Identifier: FS-0.9-or-later

package config

import "github.com/kaack/elrs-joystick-control/pkg/util"

type DivideT struct {
	RightDefault *util.RawValue `json:"right_default"`

	Left  *IOHolder    `json:"left"`
	Right *[]*IOHolder `json:"right"`
}

// InputDivide *** Divide Operation ***
type InputDivide struct {
	Id    string        `json:"id"`
	Value util.RawValue `json:"value"`
	IsNaN bool          `json:"-"`

	Type   string    `json:"type"`
	Divide DivideT   `json:"divide" input:"true"`
	Holder *IOHolder `json:"-"`
}

func DivideOperation(left util.RawValue, right util.RawValue) util.RawValue {
	return left / right
}

func (i *InputDivide) Eval(c *Config) (src IOType, out util.RawValue, ch util.ChannelNumber, nan bool) {
	src, out, ch, nan = EvalOperation(c, i.Divide.Left, i.Divide.Right, i.Divide.RightDefault, DivideOperation)
	i.Value = out
	i.IsNaN = nan

	return src, out, ch, nan
}

func (i *InputDivide) InputType() string {
	return i.Type
}

func (i *InputDivide) InputValue() *util.RawValue {
	if i.IsNaN {
		return nil
	}
	return &i.Value
}

func (i *InputDivide) InputId() string {
	return i.Id
}

func (i *InputDivide) Children() (out *[]*IOHolder) {
	return GetChildren(i.Divide.Left, i.Divide.Right)
}
