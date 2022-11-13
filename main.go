/*
Copyright 2022 Brian Pursley.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/chai2010/gettext-go/po"
	"github.com/spf13/cobra"
)

type Options struct {
	Output   string
	Template string
}

var options Options

func main() {
	cmd := cobra.Command{
		Use:   "gettext-report [.po files...]",
		Short: "Generate a report from GNU GetText PO files",
		Example: `
# Generate a report for two specific .po files
gettext-report french.po german.po

# Generate a report for two specific .po files, including a template file
gettext-report french.po german.po -t template.pot

# (Bash) Generate a report for all .po files under the current directory, including a template file
gettext-report $(find . -name *.po) -t template.pot`,
		RunE: run,
	}

	cmd.Flags().StringVarP(&options.Output, "output", "o", "table", "Output format: table or json")
	cmd.Flags().StringVarP(&options.Template, "template", "t", "", "Template (.pot) file")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

type ReportDataItem struct {
	Count   int
	Diff    int
	Percent float32
}

type FileReportDataItem struct {
	File string
	ReportDataItem
}

func run(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("no files specified")
	}

	// If a .pot file is specified, count the number of messages it contains and
	// use that as a lower bound for each file's total.
	var minTotal int
	if len(options.Template) > 0 {
		potData, err := po.LoadFile(options.Template)
		if err != nil {
			return nil
		}
		minTotal = len(potData.Messages)
	} else {
		minTotal = 0
	}

	var data []*FileReportDataItem
	total := &ReportDataItem{}

	sort.Strings(args)
	for _, poFile := range args {
		d, err := analyzeFile(poFile, minTotal)
		if err != nil {
			return err
		}
		data = append(data, d)
		total.Count += d.Count
		total.Diff += d.Diff
	}

	if total.Count > 0 {
		total.Percent = float32(total.Diff) / float32(total.Count) * 100
	} else {
		total.Percent = 0
	}

	switch options.Output {
	case "table":
		return printTable(data, total)
	case "json":
		return printJSON(data, total)
	default:
		return errors.New("unknown output type")
	}
}

func analyzeFile(path string, minTotal int) (*FileReportDataItem, error) {
	poData, err := po.LoadFile(path)
	if err != nil {
		return nil, err
	}

	d := &FileReportDataItem{File: path}
	if len(poData.Messages) > minTotal {
		d.Count = len(poData.Messages)
	} else {
		d.Count = minTotal
	}

	for _, m := range poData.Messages {
		if m.MsgId != m.MsgStr {
			d.Diff++
		}
	}

	if d.Count > 0 {
		d.Percent = float32(d.Diff) / float32(d.Count) * 100
	} else {
		d.Percent = 0
	}

	return d, nil
}

func printTable(files []*FileReportDataItem, total *ReportDataItem) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	if _, err := fmt.Fprintln(w, "File\tTranslated\tTotal\tPercent"); err != nil {
		return err
	}

	for _, rdi := range files {
		if _, err := fmt.Fprintf(w, "%s\t%d\t%d\t%.1f %%\n", rdi.File, rdi.Diff, rdi.Count, rdi.Percent); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "TOTAL\t%d\t%d\t%.1f %%\n", total.Diff, total.Count, total.Percent); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func printJSON(files []*FileReportDataItem, total *ReportDataItem) error {
	data := struct {
		ReportDataItem
		Files []*FileReportDataItem
	}{
		Files: files,
	}
	data.Count = total.Count
	data.Diff = total.Diff
	data.Percent = total.Percent

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
