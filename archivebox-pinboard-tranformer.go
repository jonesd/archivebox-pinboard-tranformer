/*
Copyright (c) 2022 David G Jones

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
 */

package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
)

var inputFilename string
var stripPriv bool
var titleCaseTags bool

func main() {
    flag.StringVar(&inputFilename, "f", "pinboard.in.rss", "pinboard rss file")
    flag.BoolVar(&stripPriv, "p", true, "Strip [Priv] marker on title line")
    flag.BoolVar(&titleCaseTags, "t", true, "Title case tags")
    flag.Parse()

    f, err := os.Open(inputFilename)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := process(scanner.Text())
        fmt.Printf("%s\n", line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func process(line string) string {
    line = processTags(line)
    line = processTitle(line)
    return line
}

func processTags(line string) string {
    const subjectOpenXml = "<dc:subject>"
    const subjectCloseXml = "</dc:subject>"

    openIndex := strings.Index(line, subjectOpenXml)
    closeIndex := strings.Index(line, subjectCloseXml)
    if openIndex >= 0 && closeIndex >= 0 {
        openTagsIndex := openIndex + len(subjectOpenXml)
        tagsString := line[openTagsIndex : closeIndex]
        updatedTagsString := migrateTagList(tagsString)
        if titleCaseTags == true {
            updatedTagsString = strings.Title(updatedTagsString)
        }
        return line[0:openTagsIndex] + updatedTagsString + line[closeIndex:len(line)]
    } else {
        return line
    }
}

func migrateTagList(tagsLine string) string  {
    spaceToComma, _ := regexp.Compile(` `)
    hyphenToSpace, _ := regexp.Compile(`-`)
    tagsLineWithCommas := spaceToComma.ReplaceAllString(tagsLine, ",")
    archiveboxTagsLine := hyphenToSpace.ReplaceAllString(tagsLineWithCommas, " ")
    return archiveboxTagsLine
}

func processTitle(line string) string {
    const titleOpenXml = "<title>"
    const titleCloseXml = "</title>"
    const privateMarker = "[priv] "

    openIndex := strings.Index(line, titleOpenXml)
    closeIndex := strings.Index(line, titleCloseXml)
    if openIndex >= 0 && closeIndex >= 0 {
        openTagsIndex := openIndex + len(titleOpenXml)
        titleString := line[openTagsIndex : closeIndex]
        privateMarkerIndex := strings.Index(titleString, privateMarker)
        if privateMarkerIndex != -1 {
            titleString = titleString[len(privateMarker):len(titleString)]
        }
        return line[0:openTagsIndex] + titleString + line[closeIndex:len(line)]
    }
    return line
}
