package main

import (
	"fmt"
	"strconv"

	"github.com/mitchellh/go-z3"
)

func smtsolver(initCells []int) (cells []int, err error) {
	ccc = NewContext()
	defer ccc.Close()
	x := IntArrayVar("x", 81)
	for i := 0; i < 81; i++ {
		Assert(x[i].Ge(IntVal(1)).And(x[i].Lt(IntVal(10))))
	}
	for i := 0; i < 9; i++ {
		Assert(x[i*9].Distinct(x[i*9+1], x[i*9+2], x[i*9+3], x[i*9+4], x[i*9+5], x[i*9+6], x[i*9+7], x[i*9+8]))
	}
	for j := 0; j < 9; j++ {
		Assert(x[j].Distinct(x[j+9], x[j+18], x[j+27], x[j+36], x[j+45], x[j+54], x[j+63], x[j+72]))
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			Assert(x[i*9+j].Distinct(x[i*9+j+1], x[i*9+j+2], x[(i+1)*9+j], x[(i+1)*9+j+1], x[(i+1)*9+j+2], x[(i+2)*9+j], x[(i+2)*9+j+1], x[(i+2)*9+j+2]))
		}
	}
	for idx, n := range initCells {
		if n > 0 {
			knownValue := IntVal(n)
			Assert(x[idx].Eq(knownValue))
		}
	}
	cells = make([]int, len(initCells))
	var values map[string]int
	values, err = SolveIntValues("x")
	if err != nil {
		return
	}
	for i := 0; i < len(initCells); i++ {
		name := fmt.Sprintf("x[%d]", i)
		cells[i] = values[name]
	}
	return
}

// SolveIntValues は制約を解決する変数の値を表示する関数
func SolveIntValues(names ...string) (map[string]int, error) {
	return ccc.SolveIntValues(names...)
}

// SolveIntValues は制約を解決する変数の値を表示する関数
func (c Context) SolveIntValues(names ...string) (r map[string]int, err error) {
	// 解決可能かどうかを調べる
	if v := c.solver.Check(); v != z3.True {
		err = fmt.Errorf("unsolvable")
		return
	}

	// 制約を満たす値の取得
	m := c.solver.Model()
	values := m.Assignments()
	m.Close()

	r = map[string]int{}

	// 可変引数で指定された変数名の値を表示
	for _, name := range names {
		//fmt.Println("name =", name)
		if c.vars[name] {
			//fmt.Printf("%s = %s\n", name, values[name].String())
			r[name], _ = strconv.Atoi(values[name].String())
		} else {
			// 配列の可能性
			i := 0
			for {
				idxName := fmt.Sprintf("%s[%d]", name, i)
				if c.vars[idxName] {
					//fmt.Printf("%s = %s\n", idxName, values[idxName].String())
					r[idxName], _ = strconv.Atoi(values[idxName].String())
				} else {
					break
				}
				i++
			}
		}
		//fmt.Printf("%s = %s.\n", name, values[name].FString2())
	}
	return
}
