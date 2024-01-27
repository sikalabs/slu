package editorconfig

import "os"

func CreateEditorconfig(
	withGo bool,
	withPython bool,
) {
	content := `root = true
[*]
indent_style = space
indent_size = 2
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true
end_of_line = lf
max_line_length = off
[Makefile]
indent_style = tab
`
	if withGo {
		content += `[*.go]
indent_style = tab
`
	}
	if withPython {
		content += `[*.py]
indent_size = 4
`
	}
	err := os.WriteFile(".editorconfig", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
