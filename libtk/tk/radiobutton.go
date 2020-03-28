// Copyright 2018 visualfc. All rights reserved.

package tk

import "fmt"

// radio button
type RadioButton struct {
	BaseWidget
	command *Command
}

func NewRadioButton(parent Widget, text string, attributes ...*WidgetAttr) *RadioButton {
	theme := checkInitUseTheme(attributes)
	iid := makeNamedWidgetId(parent, "atk_radiobutton")
	attributes = append(attributes, &WidgetAttr{"text", text})
	info := CreateWidgetInfo(iid, WidgetTypeRadioButton, theme, attributes)
	if info == nil {
		return nil
	}
	w := &RadioButton{}
	w.id = iid
	w.info = info
	RegisterWidget(w)
	return w
}

func (w *RadioButton) Attach(id string) error {
	info, err := CheckWidgetInfo(id, WidgetTypeRadioButton)
	if err != nil {
		return err
	}
	w.id = id
	w.info = info
	RegisterWidget(w)
	return nil
}

func (w *RadioButton) SetText(text string) error {
	setObjText("atk_tmp_text", text)
	return eval(fmt.Sprintf("%v configure -text $atk_tmp_text", w.id))
}

func (w *RadioButton) Text() string {
	r, _ := evalAsString(fmt.Sprintf("%v cget -text", w.id))
	return r
}

func (w *RadioButton) SetWidth(width int) error {
	return eval(fmt.Sprintf("%v configure -width {%v}", w.id, width))
}

func (w *RadioButton) Width() int {
	r, _ := evalAsInt(fmt.Sprintf("%v cget -width", w.id))
	return r
}

func (w *RadioButton) SetImage(image *Image) error {
	if image == nil {
		return ErrInvalid
	}
	return eval(fmt.Sprintf("%v configure -image {%v}", w.id, image.Id()))
}

func (w *RadioButton) Image() *Image {
	r, err := evalAsString(fmt.Sprintf("%v cget -image", w.id))
	return parserImageResult(r, err)
}

func (w *RadioButton) SetCompound(compound Compound) error {
	return eval(fmt.Sprintf("%v configure -compound {%v}", w.id, compound))
}

func (w *RadioButton) Compound() Compound {
	r, err := evalAsString(fmt.Sprintf("%v cget -compound", w.id))
	return parserCompoundResult(r, err)
}

func (w *RadioButton) SetPaddingN(padx int, pady int) error {
	if w.info.IsTtk {
		return eval(fmt.Sprintf("%v configure -padding {%v %v}", w.id, padx, pady))
	}
	return eval(fmt.Sprintf("%v configure -padx {%v} -pady {%v}", w.id, padx, pady))
}

func (w *RadioButton) PaddingN() (int, int) {
	var r string
	var err error
	if w.info.IsTtk {
		r, err = evalAsString(fmt.Sprintf("%v cget -padding", w.id))
	} else {
		r1, _ := evalAsString(fmt.Sprintf("%v cget -padx", w.id))
		r2, _ := evalAsString(fmt.Sprintf("%v cget -pady", w.id))
		r = r1 + " " + r2
	}
	return parserPaddingResult(r, err)
}

func (w *RadioButton) SetPadding(pad Pad) error {
	return w.SetPaddingN(pad.X, pad.Y)
}

func (w *RadioButton) Padding() Pad {
	x, y := w.PaddingN()
	return Pad{x, y}
}

func (w *RadioButton) SetState(state State) error {
	return eval(fmt.Sprintf("%v configure -state {%v}", w.id, state))
}

func (w *RadioButton) State() State {
	r, err := evalAsString(fmt.Sprintf("%v cget -state", w.id))
	return parserStateResult(r, err)
}

func (w *RadioButton) SetTakeFocus(takefocus bool) error {
	return eval(fmt.Sprintf("%v configure -takefocus {%v}", w.id, boolToInt(takefocus)))
}

func (w *RadioButton) IsTakeFocus() bool {
	r, _ := evalAsBool(fmt.Sprintf("%v cget -takefocus", w.id))
	return r
}

func (w *RadioButton) SetChecked(check bool) *RadioButton {
	if check {
		eval(fmt.Sprintf("set [%v cget -variable] [%v cget -value]", w.id, w.id))
	} else {
		eval(fmt.Sprintf("set [%v cget -variable] {}", w.id))
	}
	return w
}

func (w *RadioButton) IsChecked() bool {
	r, _ := evalAsBool(fmt.Sprintf("expr $[%v cget -variable]=={[%v cget -value]}", w.id, w.id))
	return r
}

func (w *RadioButton) OnCommand(fn func()) error {
	if fn == nil {
		return ErrInvalid
	}
	if w.command == nil {
		w.command = &Command{}
		bindCommand(w.id, "command", w.command.Invoke)
	}
	w.command.Bind(fn)
	return nil
}

func (w *RadioButton) Invoke() {
	eval(fmt.Sprintf("%v invoke", w.id))
}

func RadioButtonAttrText(text string) *WidgetAttr {
	return &WidgetAttr{"text", text}
}

func RadioButtonAttrWidth(width int) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func RadioButtonAttrImage(image *Image) *WidgetAttr {
	if image == nil {
		return nil
	}
	return &WidgetAttr{"image", image.Id()}
}

func RadioButtonAttrCompound(compound Compound) *WidgetAttr {
	return &WidgetAttr{"compound", compound}
}

func RadioButtonAttrPadding(padding Pad) *WidgetAttr {
	return &WidgetAttr{"padding", padding}
}

func RadioButtonAttrState(state State) *WidgetAttr {
	return &WidgetAttr{"state", state}
}

func RadioButtonAttrTakeFocus(takefocus bool) *WidgetAttr {
	return &WidgetAttr{"takefocus", boolToInt(takefocus)}
}
