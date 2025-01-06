package table

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/starkeland/pretty-table/style"
)

type Option func(*Table)

func ShowSeqColumn() func(*Table) {
	return func(t *Table) {
		t.options.ShowSeqColumn = true
	}
}

func ShowHR() func(*Table) {
	return func(t *Table) {
		t.options.ShowHR = true
	}
}

type options struct {
	ShowSeqColumn bool
	ShowHR        bool
}

type Table struct {
	options options
	style   *style.TableStyle

	caption string
	header  Row
	rows    []Row
	footer  Row

	printer     strings.Builder
	columnCount int   // the count num of columns
	tableWidth  int   // the width of table
	columnWidth []int // the width of each column
}

func (t *Table) SetCaption(caption string) {
	t.caption = caption
}

func (t *Table) SetHeader(header []string) {
	t.header = Row{}
	for _, h := range header {
		t.header.cells = append(t.header.cells, NewCell(h))
	}
}

func (t *Table) SetFooter(footer []string) {
	t.footer = Row{}
	for _, f := range footer {
		t.footer.cells = append(t.footer.cells, NewCell(f))
	}
}

func (t *Table) AddRows(rows ...[]string) {
	for _, r := range rows {
		t.rows = append(t.rows, Row{})
		for _, c := range r {
			t.rows[len(t.rows)-1].cells = append(t.rows[len(t.rows)-1].cells, NewCell(c))
		}
	}
}

func (t *Table) Render() {
	t.addSeqColumn()
	t.initWidth()

	t.renderCaption()
	t.renderTopBorder()
	t.renderHeader()
	t.renderRows()
	t.renderFooter()
	t.renderBottomBorder()

	fmt.Println(t.printer.String())
}

func (t *Table) Style() *style.TableStyle {
	if t.style == nil {
		return style.DefaultTableStyle
	}
	return t.style
}

func (t *Table) addSeqColumn() {
	if !t.options.ShowSeqColumn {
		return
	}

	t.header.cells = append([]Cell{NewCell("#")}, t.header.cells...)
	for i, row := range t.rows {
		t.rows[i].cells = append([]Cell{NewCell(strconv.Itoa(i + 1))}, row.cells...)
	}

	t.footer.cells = append([]Cell{NewCell("")}, t.footer.cells...)
}

func (t *Table) initWidth() {
	t.columnCount = len(t.header.cells)
	for _, row := range t.rows {
		if len(row.cells) > t.columnCount {
			t.columnCount = len(row.cells)
		}

		for i, cell := range row.cells {
			if i >= len(t.columnWidth) {
				t.columnWidth = append(t.columnWidth, cell.Width())
			} else if cell.Width() > t.columnWidth[i] {
				t.columnWidth[i] = cell.Width()
			}
		}
	}

	if t.header.cells != nil {
		for i, cell := range t.header.cells {
			if i >= len(t.columnWidth) {
				t.columnWidth = append(t.columnWidth, cell.Width())
			} else if cell.Width() > t.columnWidth[i] {
				t.columnWidth[i] = cell.Width()
			}
		}
	}

	if t.footer.cells != nil {
		for i, cell := range t.footer.cells {
			if i >= len(t.columnWidth) {
				t.columnWidth = append(t.columnWidth, cell.Width())
			} else if cell.Width() > t.columnWidth[i] {
				t.columnWidth[i] = cell.Width()
			}
		}
	}

	for _, w := range t.columnWidth {
		t.tableWidth += w
	}
}

func (t *Table) renderCaption() {
	if t.caption != "" {
		t.printer.WriteString(strings.Repeat(" ", (t.tableWidth+t.columnCount+1-len(t.caption))/2) + t.caption)
		t.printer.WriteString("\n")
	}
}

func (t *Table) renderTopBorder() {
	if t.header.cells == nil && len(t.rows) == 0 && t.footer.cells == nil {
		return
	}

	t.printer.WriteString(t.Style().Border.TopLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.Style().Border.Top, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.Style().Border.TopSeparator)
		}
	}
	t.printer.WriteString(t.Style().Border.TopRight)
	t.printer.WriteString("\n")
}

func (t *Table) renderHeader() {
	if t.header.cells == nil {
		return
	}

	t.printer.WriteString(t.Style().Border.RowSeparator)
	for i, col := range t.header.cells {
		t.printer.WriteString(col.Text(t.columnWidth[i]))
		t.printer.WriteString(t.Style().Border.RowSeparator)
	}
	t.printer.WriteString("\n")

	t.renderHr()
}

func (t *Table) renderHr() {
	if t.options.ShowHR == false {
		return
	}
	t.printer.WriteString(t.Style().Border.HrLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.Style().Border.Hr, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.Style().Border.HrSeparator)
		} else {
			t.printer.WriteString(t.Style().Border.HrRight)
		}
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderRows() {
	for _, row := range t.rows {
		t.printer.WriteString(t.Style().Border.RowSeparator)
		for i, col := range row.cells {
			t.printer.WriteString(col.Text(t.columnWidth[i]))
			t.printer.WriteString(t.Style().Border.RowSeparator)
		}
		t.printer.WriteString("\n")
		t.renderHr()
	}
}

func (t *Table) renderFooter() {
	if t.footer.cells == nil {
		return
	}

	t.printer.WriteString(t.Style().Border.RowSeparator)
	for i, col := range t.footer.cells {
		t.printer.WriteString(col.Text(t.columnWidth[i]))
		t.printer.WriteString(t.Style().Border.RowSeparator)
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderBottomBorder() {
	t.printer.WriteString(t.Style().Border.BottomLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.Style().Border.Bottom, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.Style().Border.BottomSeparator)
		}
	}
	t.printer.WriteString(t.Style().Border.BottomRight)
	t.printer.WriteString("\n")
}
