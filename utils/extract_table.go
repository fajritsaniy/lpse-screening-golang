package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractTableData(htmlContent string) [][]string {
	var headings, row []string
	var rows [][]string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	// Find each table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				headings = append(headings, tableheading.Text())
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})
			rows = append(rows, row)
			row = nil
		})
	})

	return rows
}
