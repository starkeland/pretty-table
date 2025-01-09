package table

import (
	"fmt"
	"strings"

	"github.com/starkeland/pretty-table/style"
)

type Option func(*Table)

func ShowSeqColumn() Option {
	return func(t *Table) {
		t.showSeqColumn = true
	}
}

func WithTableStyle(tableStyle *style.TableStyle) Option {
	return func(t *Table) {
		if tableStyle != nil {
			t.tableStyle = tableStyle
		}
	}
}

type Table struct {
	showSeqColumn bool
	tableStyle    *style.TableStyle

	caption *Cell
	header  []*Cell
	rows    []*Row
	footer  []*Cell

	printer     strings.Builder
	columnCount int   // the count num of columns
	tableWidth  int   // the width of table
	columnWidth []int // the width of each column
}

func (t *Table) SetCaption(caption string) {
	if caption != "" {
		t.caption = NewCell(caption)
	}
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

func (t *Table) AddStringRow(stringRow []string) {
	if len(stringRow) == 0 {
		return
	}
	var row []*Cell
	for _, cell := range stringRow {
		row = append(row, NewCell(cell))
	}
	t.rows = append(t.rows, NewRow(row))
}

func (t *Table) AddStringRows(stringRows ...[]string) {
	for _, row := range stringRows {
		t.AddStringRow(row)
	}
}

func (t *Table) AddRow(row *Row) {
	if row != nil {
		t.rows = append(t.rows, row)
	}
}

func (t *Table) AddRows(rows ...*Row) {
	for _, row := range rows {
		t.AddRow(row)
	}
}

func (t *Table) SetTableStyle(s *style.TableStyle) {
	if s == nil {
		t.tableStyle = style.DefaultTableStyle
	} else {
		t.tableStyle = s
	}
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

func (t *Table) addSeqColumn() {
	if !t.showSeqColumn {
		return
	}

	t.header = append([]*Cell{NewCell("#")}, t.header...)
	for i, row := range t.rows {
		row.InsertSeqColumn(i + 1)
	}
	t.footer = append([]*Cell{NewCell("")}, t.footer...)
}

func (t *Table) initWidth() {
	t.columnCount = len(t.header)
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

// cell style priority: column < row < cell
func (t *Table) initCellStyle() {
	for i := 0; i < len(t.rows); i++ { // row
		for j := 0; j < len(t.rows[i].cells); j++ { // column
			if t.showSeqColumn && j == 0 {
				continue
			}
			colStyle := t.tableStyle.Columns[j]
			if t.showSeqColumn {
				colStyle = t.tableStyle.Columns[j-1]
			}
			t.rows[i].cells[j].style = style.MergeCellStyles(colStyle, t.rows[i].style, t.rows[i].cells[j].style)
		}
	}
}

func (t *Table) renderCaption() {
	if t.caption == nil {
		return
	}

	if t.tableStyle.Caption != nil {
		t.caption.style = t.tableStyle.Caption
	}

	if t.tableStyle.HideBorder {
		t.printer.WriteString(t.caption.Render(t.tableWidth))
	} else {
		t.printer.WriteString(t.caption.Render(t.tableWidth + t.columnCount + 1))
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderTopBorder() {
	if t.tableStyle.HideBorder {
		return
	}
	if t.header == nil && len(t.rows) == 0 && t.footer == nil {
		return
	}

	t.printer.WriteString(t.tableStyle.Border.TopLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.tableStyle.Border.Top, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.tableStyle.Border.TopSeparator)
		}
	}
	t.printer.WriteString(t.tableStyle.Border.TopRight)
	t.printer.WriteString("\n")
}

func (t *Table) renderHeader() {
	if len(t.header) == 0 || (t.showSeqColumn && len(t.header) == 1) {
		return
	}

	if !t.tableStyle.HideBorder {
		t.printer.WriteString(t.tableStyle.Border.ColumnSeparator)
	}

	for i, cell := range t.header {
		cell.style = t.tableStyle.Header
		t.printer.WriteString(cell.Render(t.columnWidth[i]))
		if !t.tableStyle.HideBorder {
			t.printer.WriteString(t.tableStyle.Border.ColumnSeparator)
		}
	}
	t.printer.WriteString("\n")

	if t.tableStyle.HideBorder {
		return
	}

	t.printer.WriteString(t.tableStyle.Border.HrLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.tableStyle.Border.HeaderBottom, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.tableStyle.Border.HrSeparator)
		} else {
			t.printer.WriteString(t.tableStyle.Border.HrRight)
		}
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderHr() {
	if t.tableStyle.HideBorder {
		return
	}
	t.printer.WriteString(t.tableStyle.Border.HrLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.tableStyle.Border.Hr, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.tableStyle.Border.HrSeparator)
		} else {
			t.printer.WriteString(t.tableStyle.Border.HrRight)
		}
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderRows() {
	for i, row := range t.rows {
		if !t.tableStyle.HideBorder {
			t.printer.WriteString(t.tableStyle.Border.ColumnSeparator)
		}
		for j, cell := range row.cells {
			t.printer.WriteString(cell.Render(t.columnWidth[j]))
			if !t.tableStyle.HideBorder {
				t.printer.WriteString(t.tableStyle.Border.ColumnSeparator)
			}
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

	if !t.tableStyle.HideBorder {
		t.printer.WriteString(t.tableStyle.Border.ColumnSeparator)
	}

	for i, cell := range t.footer {
		t.printer.WriteString(cell.Render(t.columnWidth[i]))
		if !t.tableStyle.HideBorder {
			t.printer.WriteString(t.tableStyle.Border.ColumnSeparator)
		}
	}
	t.printer.WriteString("\n")
}

func (t *Table) renderBottomBorder() {
	if t.tableStyle.HideBorder {
		return
	}

	t.printer.WriteString(t.tableStyle.Border.BottomLeft)
	for i := 0; i < t.columnCount; i++ {
		t.printer.WriteString(strings.Repeat(t.tableStyle.Border.Bottom, t.columnWidth[i]))
		if i < t.columnCount-1 {
			t.printer.WriteString(t.tableStyle.Border.BottomSeparator)
		}
	}
	t.printer.WriteString(t.tableStyle.Border.BottomRight)
	t.printer.WriteString("\n")
}
