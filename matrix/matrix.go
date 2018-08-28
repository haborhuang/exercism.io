package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type matrix struct {
	colNum int
	data   []int
}

var (
	errScanInput = errors.New("Scan input error")
	errEmptyLine = errors.New("Empty line is invalid")
	errParseInt  = errors.New("Parse integer error")
	errUnevenRow = errors.New("Rows are uneven")
)

func New(input string) (*matrix, error) {
	m := matrix{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			return nil, errEmptyLine
		}

		tokens := strings.Split(line, " ")

		if m.colNum == 0 {
			m.colNum = len(tokens)
		} else if m.colNum != len(tokens) {
			return nil, errUnevenRow
		}
		for _, s := range tokens {
			i, err := strconv.ParseInt(s, 10, 0)
			if nil != err {
				return nil, errParseInt
			}
			m.data = append(m.data, int(i))
		}

	}

	return &m, nil
}

func (m *matrix) Rows() [][]int {
	res := make([][]int, m.rowNum())
	for i := range res {
		res[i] = make([]int, m.colNum)
		for j := 0; j < m.colNum; j++ {
			res[i][j] = m.data[m.getIndex(i, j)]
		}
	}

	return res
}

func (m *matrix) Cols() [][]int {
	res := make([][]int, m.colNum)
	for j := range res {
		rowNum := m.rowNum()
		res[j] = make([]int, rowNum)
		for i := 0; i < rowNum; i++ {
			res[j][i] = m.data[m.getIndex(i, j)]
		}
	}

	return res
}

func (m *matrix) Set(i, j, val int) (ok bool) {
	defer func() {
		if ok {
			m.data[m.getIndex(i, j)] = val
		}
	}()
	return i >= 0 && i < m.rowNum() &&
		j >= 0 && j < m.colNum
}

func (m *matrix) rowNum() int {
	return len(m.data) / m.colNum
}

func (m *matrix) getIndex(i, j int) int {
	return i*m.colNum + j
}
