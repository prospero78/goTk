// Copyright 2018 visualfc. All rights reserved.

package tk

import (
	"strings"
	"testing"
)

func init() {
	registerTest("TestWidgetId", testWidgetId)
	registerTest("TestWidgetParent", testWidgetParent)
}

type TestWidget struct {
	BaseWidget
}

func (w *TestWidget) Type() WidgetType {
	return WidgetTypeLast + 1
}

func (w *TestWidget) Attach(id string) error {
	w.id = id
	w.info = &WidgetInfo{WidgetTypeCanvas, "TestWidget", false, nil}
	RegisterWidget(w)
	return nil
}

func NewTestWidget(parent Widget, id string) *TestWidget {
	w := &TestWidget{}
	w.Attach(makeNamedWidgetId(parent, id))
	return w
}

func testWidgetId(t *testing.T) {
	var id string
	parent := NewTestWidget(nil, "base")
	id = makeNamedWidgetId(nil, "atk_wtest")
	if !strings.HasPrefix(id, ".atk_wtest") {
		t.Fatal(id)
	}
	id = makeNamedWidgetId(parent, "atk_wchild")
	if !strings.HasPrefix(id, ".base1.atk_wchild1") {
		t.Fatal(id)
	}
	id = makeNamedWidgetId(nil, "idtest")
	if id != ".idtest1" {
		t.Fatal(id)
	}
	id = makeNamedWidgetId(parent, "idtest")
	if id != ".base1.idtest1" {
		t.Fatal(id)
	}
	id = makeNamedWidgetId(parent, "idtest")
	if id != ".base1.idtest2" {
		t.Fatal(id)
	}
	DestroyWidget(parent)
}

func findOfList(w Widget, list []Widget) bool {
	for _, v := range list {
		if v == w {
			return true
		}
	}
	return false
}

func testWidgetParent(t *testing.T) {
	a1 := NewTestWidget(nil, "a1")
	a2 := NewTestWidget(nil, "a2")
	defer a1.Destroy()
	defer a2.Destroy()
	a1_b1 := NewTestWidget(a1, "b1")
	a1_b1_c1 := NewTestWidget(a1_b1, "c1")
	a1_b1_c2 := NewTestWidget(a1_b1, "c2")
	a1_b1_c3 := NewTestWidget(a1_b1, "c3")
	a2_b1 := NewTestWidget(a2, "b1")
	a2_b1_c1 := NewTestWidget(a2_b1, "c1")
	a2_b1_c2 := NewTestWidget(a2_b1, "c2")
	a2_b1_c3 := NewTestWidget(a2_b1, "c3")

	if p := ParentOfWidget(a1); p != RootWindow() {
		t.Fatal("ParentWidget", p)
	}
	if p := ParentOfWidget(a1_b1); p != a1 {
		t.Fatal("ParentWidget", p)
	}
	if p := ParentOfWidget(a1_b1_c1); p != a1_b1 {
		t.Fatal("ParentWidget", p)
	}
	list := ChildrenOfWidget(rootWindow)
	if !findOfList(a1, list) || !findOfList(a2, list) {
		t.Fatal("ChildrenOfWidget", list)
	}
	list = ChildrenOfWidget(a1_b1)
	if len(list) != 3 || !findOfList(a1_b1_c1, list) || !findOfList(a1_b1_c2, list) || !findOfList(a1_b1_c3, list) {
		t.Fatal("ChildrenOfWidget", list)
	}
	DestroyWidget(a1_b1_c3)
	list = ChildrenOfWidget(a1_b1)
	if len(list) != 2 {
		t.Fatal("DestroyWidget", list)
	}
	DestroyWidget(a1)
	list = ChildrenOfWidget(rootWindow)
	if findOfList(a1, list) {
		t.Fatal("DestroyWidget", list)
	}
	if IsValidWidget(a1_b1_c1) {
		t.Fatal("IsValidWidget", a1_b1_c1)
	}

	list = a2_b1.Children()
	if len(list) != 3 {
		t.Fatal("Children", list)
	}
	a2_b1_c3.Destroy()
	if a2_b1_c1.Parent() != a2_b1 || a2_b1_c2.Parent() != a2_b1 {
		t.Fatal("Destroy")
	}
	a2_b1.DestroyChildren()
	if a2_b1_c2.IsValid() {
		t.Fatal("DestroyChildren", a2_b1_c2)
	}
}
