package api

import (
	"fmt"
	"strings"

	"github.com/gabrielroacueto/locc/filesystem"
	"github.com/gabrielroacueto/locc/llm"
)

const ANALYYZE_DIRECTORY_PROMPT = `Given the following files in a code repository:

%s

Please analyze these files and explain:
1. What programming language(s) appear to be used
2. How the code is organized and structured
3. The main components/modules of the codebase
4. Any notable patterns or architectural decisions you observe
5. The likely purpose of this codebase

Please provide a comprehensive analysis based on the file structure.`

// Given a directory, generate the analysis of the structure and code in there. The output will be streamed to the streaming destination as tokens
// arrive. The print callback function will be used to pass the tokens as they stream back from the LLM server.
func StreamDirectoryAnalysis(directory string, callback func(string)) error {

	contents, err := filesystem.GetDirectoryContents(directory)
	if err != nil {
		return err
	}
	prompt := fmt.Sprintf(ANALYYZE_DIRECTORY_PROMPT, strings.Join(contents, "\n"))

	llm_err := llm.GenerateStream(prompt, callback)

	if llm_err != nil {
		return err
	}

	return nil
}
