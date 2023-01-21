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
	"testing"
)

func TestProcessLine_NonTagListLineReturnedAsIs(t *testing.T) {
	line := "<dc:date>2022-10-09T10:00:24+00:00</dc:date>"
	expectation := line
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_BlankTagListLineReturnedAsIs(t *testing.T) {
	line := "<dc:subject></dc:subject>"
	expectation := line
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_SingleWordTagLineReturnedAsIs(t *testing.T) {
	line := "<dc:subject>test</dc:subject>"
	expectation := line
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_TagsShouldBeTitleCase(t *testing.T) {
	line := "<dc:subject>one two-three</dc:subject>"
	expectation := "<dc:subject>One,Two Three</dc:subject>"
	titleCaseTags = true
	actual := process(line)
	titleCaseTags = false
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_HyphenatedTagShouldBeSplitBySpaces(t *testing.T) {
	line := "<dc:subject>test-one</dc:subject>"
	expectation := "<dc:subject>test one</dc:subject>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_HyphenatedTagShouldBeSplitBySpacesToBeOneTag(t *testing.T) {
	line := "<dc:subject>test-one</dc:subject>"
	expectation := "<dc:subject>test one</dc:subject>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
func TestProcessLine_SpaceTagShouldUseCommas(t *testing.T) {
	line := "<dc:subject>one two three</dc:subject>"
	expectation := "<dc:subject>one,two,three</dc:subject>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_SupportSpaceAndHypenatedTags(t *testing.T) {
	line := "<dc:subject>one a-b two bb-ccc-dddd three</dc:subject>"
	expectation := "<dc:subject>one,a b,two,bb ccc dddd,three</dc:subject>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessLine_SupportSubjectAndCreatorOnSameXmlLine(t *testing.T) {
	line := "  <dc:creator>user</dc:creator><dc:subject>one a-b</dc:subject>"
	expectation := "  <dc:creator>user</dc:creator><dc:subject>one,a b</dc:subject>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessTitleLine_PreserveTitleWithoutPrivMarker(t *testing.T) {
	line := "  <title>Some title</title>"
	expectation := "  <title>Some title</title>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestProcessTitleLine_StripPrivMarkerFromTitle(t *testing.T) {
	line := "  <title>[priv] Some title</title>"
	expectation := "  <title>Some title</title>"
	actual := process(line)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

