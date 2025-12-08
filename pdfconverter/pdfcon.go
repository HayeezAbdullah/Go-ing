package main

import (
	"github.com/jung-kurt/gofpdf"
)

func GenerateEduSpherePDF(filename string, content string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	// Times font supports ASCII perfectly
	pdf.SetFont("Times", "", 12)

	// Line spacing 1.5
	lineHeight := 7.5
	pdf.SetAutoPageBreak(true, 15)

	// Justify text
	pdf.MultiCell(0, lineHeight, content, "", "J", false)

	return pdf.OutputFileAndClose(filename)
}

func main() {
	content := `I'm diving into Go, wide awake through the night,
Chasing bugs, fixing loops, getting logic just right.
Packages, pointers yeah, I'll conquer the rest,
But somehow your smile runs the code in my chest.

My functions may fail, my brain may freeze,
Yet thoughts of you compile with effortless ease.
So while I'm learning Go's mysterious art,
You're the subtle import that rewrites my heart.

MWUAHH!!!
`

	if err := GenerateEduSpherePDF("poem.pdf", content); err != nil {
		panic(err)
	}
}
