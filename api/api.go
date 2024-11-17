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

const ANALYZE_DIRECTORY_WITH_USER_CONTEXT_PROMPT = `Given the following files in a code repository:

%s

and the following information about the project:

%s


Please provide a comprehensive analysis based on the file structure. Do NOT go over specific files, but instead provide an overview of what the project does.`

type RepoContext struct {
	contents    string
	userContext string
}

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

func StreamDirectoryAnalysisWithAdditionalContext(directory string, callback func(string), userContext string) error {

	contents, err := filesystem.GetDirectoryContents(directory)
	if err != nil {
		return err
	}
	prompt := fmt.Sprintf(ANALYZE_DIRECTORY_WITH_USER_CONTEXT_PROMPT, strings.Join(contents, "\n"), userContext)

	llm_err := llm.GenerateStream(prompt, callback)

	if llm_err != nil {
		return err
	}

	return nil
}

func GenerateRepoDocumentation(ctx RepoContext) (string, error) {
	if ctx.contents == "" {
		panic("We need repository contents to actually analyze.")
	}

	if ctx.userContext == "" {
		panic("User context is required to generate documentation. File structure is not sufficient.")
	}

	var prompt string

	context, err := llm.Generate(prompt)

	if err != nil {
		return "", err
	}

	return context, nil

}
