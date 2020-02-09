package main

import (
	"image"

	. "github.com/gizak/termui/v3"
)

type Paragraph struct {
	Block
	Text      string
	TextStyle Style
	WrapText  bool
	start     int
	end       int
}

func NewParagraph() *Paragraph {
	return &Paragraph{
		Block:     *NewBlock(),
		TextStyle: Theme.Paragraph.Text,
		WrapText:  true,
	}
}

func (self *Paragraph) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	cells := ParseStyles(self.Text, self.TextStyle)
	if self.WrapText {
		cells = WrapCells(cells, uint(self.Inner.Dx()))
	}

	rows := SplitCells(cells, '\n')
	if self.end-self.start <= len(rows) {
		if self.end > len(rows) {
			self.end = len(rows)
			self.start = self.end - 40
		}

		if self.start <= 0 {
			self.start = 0
			self.end = 40
		}
		rows = rows[self.start:self.end]
	}
	for y, row := range rows {
		if y+self.Inner.Min.Y >= self.Inner.Max.Y {
			break
		}
		row = TrimCells(row, self.Inner.Dx())
		for _, cx := range BuildCellWithXArray(row) {
			x, cell := cx.X, cx.Cell

			buf.SetCell(cell, image.Pt(x, y).Add(self.Inner.Min))
		}
	}
}
