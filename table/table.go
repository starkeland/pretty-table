package table

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/starkeland/pretty-table/style"
)

type Option func(*Table)

func ShowSeqColumn() Option {
	return func(t *Table) {
		t.showSeqColumn = true
	}
}

func ShowHr() Option {
	return func(t *Table) {
		t.showHr = true
	}
}

func WithStyle(tableStyle *style.TableStyle) Option {
	return func(t *Table) {
		t.style = tableStyle
	}
}

type Table struct {
	showSeqColumn bool
	showHr        bool
	style         *style.TableStyle

	caption string
	header  []*Cell
	rows    [][]*Cell
	footer  []*Cell

	printer     strings.Builder
	columnCount int   // the count num of columns
	tableWidth  int   // the width of table
	columnWidth []int // the width of each column
}

func (t *Table) SetCaption(caption string) {
	t.caption = caption
}

func (t *Table) SetHeader(header []string) {
	for _, h := range header {
		t.header = append(t.header, NewCell(h))
	}
}

func (t *Table) SetFooter(footer []string) {
	for _, f := range footer {
		t.footer = append(t.footer, NewCell(f))
	}
}

func (t *Table) AddRow(row []string) {
	var cells []*Cell
	for _, r := range row {
		cells = append(cells, NewCell(r))
	}
	t.rows = append(t.rows, cells)
}

func (t *Table) SetStyle(s *style.TableStyle) {
	t.style = s
}

func (t *Table) Render() {
	t.addSeqColumn()
	t.initWidth()
	t.initCellStyle()

	t.renderCaption()
	t.renderTopBorder()
	t.renderHeader()
	t.renderRows()
	t.renderFooter()
	t.renderBottomBorder()

	fmt.Println(t.printer.String())
	t.printer.Reset()
}

func (t *Table) Style() *style.TableStyle {
	if t.style == nil {
		return style.DefaultTableStyle
	}
	return t.style
}

func (t *Table) addSeqColumn() {
	if !t.showSeqColumn {
		return
	}

	t.header = append([]*Cell{NewCell("#")}, t.header...)
	for i, cells := range t.rows {
		t.rows[i] = append([]*Cell{NewCell(strconv.Itoa(i + 1))}, cells...)
	}

	t.footer = append([]*Cell{NewCell("")}, t.footer...)
}

func (t *Table) initWidth() {
	t.columnCount = len(t.header)
	for _, row := range t.rows {
		if len(row) > t.columnCount {
			t.columnCount = len(row)
		}

		for i, cell := range row {
			if i >= len(t.columnWidth) {
				t.columnWidth = append(t.columnWidth, cell.Width())
			} else if cell.Width() > t.columnWidth[i] {
				t.columnWidth[i] = cell.Width()
			}
		}
	}

	if t.header != nil {
		for i, cell := range t.header {
			if i >= len(t.columnWidth) {
				t.columnWidth = append(t.columnWidth, cell.Width())
			} else if cell.Width() > t.columnWidth[i] {
				t.columnWidth[i] = cell.Width()
			}
		}
	}

	if t.footer != nil {
		for i, cell := range t.footer {
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

func (t *Table) initCellStyle() {
	for i, cellStyle := range t.style.Columns { // column style
		for j := 0; j < len(t.rows); j++ { // every row
			if i < len(t.rows[j]) {
				t.rows[j][i].style = cellStyle
			}
		}
	}
	for i, cellStyle := range t.style.Rows {
		if i < len(t.rows) {
			for _, cell := range t.rows[i] {
				cell.style = cellStyle
			}
		}
	}
}

func (t *Table) renderCaption() {
	if t.caption == "" {
		return
	}

	text := strings.Repeat(" ", (t.tableWidth+t.columnCount+1-len(t.caption))/2) + t.caption
	if captionStyle := t.Style().Caption; captionStyle != nil {
		text = captionStyle.Apply(text)
	}

	t.printer.WriteString(text)
	t.printer.WriteString("\n")
}

func (t *Table) renderTopBorder() {
	if t.header == nil && len(t.rows) == 0 && t.footer == nil {
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
	if len(t.header) == 0 || (t.showSeqColumn && len(t.header) == 1) {
		return
	}

	t.printer.WriteString(t.Style().Border.RowSeparator)
	for i, col := range t.header {
		t.printer.WriteString(col.Render(t.columnWidth[i]))
		t.printer.WriteString(t.Style().Border.RowSeparator)
	}
	t.printer.WriteString("\n")

	t.printer.WriteString(t.Style().Border.HrLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.Style().Border.HeaderBottom, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.Style().Border.HrSeparator)
		} else {
			t.printer.WriteString(t.Style().Border.HrRight)
		}
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderHr() {
	if t.showHr == false {
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
	for i, row := range t.rows {
		t.printer.WriteString(t.Style().Border.RowSeparator)
		for j, col := range row {
			t.printer.WriteString(col.Render(t.columnWidth[j]))
			t.printer.WriteString(t.Style().Border.RowSeparator)
		}
		t.printer.WriteString("\n")
		if i < len(t.rows)-1 {
			t.renderHr()
		}
	}
}

func (t *Table) renderFooter() {
	if len(t.footer) == 0 || (t.showSeqColumn && len(t.footer) == 1) {
		return
	}

	t.printer.WriteString(t.Style().Border.RowSeparator)
	for i, cell := range t.footer {
		t.printer.WriteString(cell.Render(t.columnWidth[i]))
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
