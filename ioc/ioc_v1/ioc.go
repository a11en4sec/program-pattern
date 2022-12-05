package ioc

import "errors"

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	set.data[x] = true
}

func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

type UndoableIntSet struct {
	IntSet
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{
		NewIntSet(),
		nil,
	}
}

func (set *UndoableIntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() {
			set.Delete(x)
		})
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() {
			// 调用嵌套的IntSet的方法
			set.Add(x)
		})
	} else {
		set.functions = append(set.functions, nil)
	}
}

// Undo 执行undo操作，上一个操作的逆向操作函数，会添加到functions切片的末尾
func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // 方便垃圾回收
	}

	set.functions = set.functions[:index]
	return nil
}
