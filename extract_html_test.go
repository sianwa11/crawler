package main

import "testing"

func TestGetH1FromHTMLBasic(t *testing.T) {
	testCases := []struct {
		name          string
		inputBody     string
		expected      string
		errorContains string
	}{
		{
			name:      "test 1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "test 2",
			inputBody: "<html><body><main><h1>Test Title 2</h1></main></body></html>",
			expected:  "Test Title 2",
		},
		{
			name:          "test 3",
			inputBody:     "<html><body><main>Test Title 3</h1></main></body></html>",
			expected:      "",
			errorContains: "couldn't get h1",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %d Failed, expected %q, got %q", i, tc.expected, actual)

			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLFallback(t *testing.T) {
	inputBody := `<html><body>
		<p>First paragraph outside main.</p>
		<p>Second paragraph outside main.</p>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "First paragraph outside main."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetParagraphFromHTMLEmpty(t *testing.T) {
	inputBody := `<html><body><p></p></body></html>`
	actual := getH1FromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH1FromHTMLEmpty(t *testing.T) {
	inputBody := `<html><body><p>No h1 here</p></body></html>`
	actual := getH1FromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
