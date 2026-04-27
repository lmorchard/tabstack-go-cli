# Agent

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AutomateEventUnion">AutomateEventUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ResearchEventUnion">ResearchEventUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentAutomateInputResponse">AgentAutomateInputResponse</a>

Methods:

- <code title="post /automate">client.Agent.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentService.Automate">Automate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentAutomateParams">AgentAutomateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AutomateEventUnion">AutomateEventUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /automate/{requestID}/input">client.Agent.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentService.AutomateInput">AutomateInput</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentAutomateInputParams">AgentAutomateInputParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentAutomateInputResponse">AgentAutomateInputResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /research">client.Agent.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentService.Research">Research</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#AgentResearchParams">AgentResearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ResearchEventUnion">ResearchEventUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Extract

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractJsonResponse">ExtractJsonResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractMarkdownResponse">ExtractMarkdownResponse</a>

Methods:

- <code title="post /extract/json">client.Extract.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractService.Json">Json</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractJsonParams">ExtractJsonParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractJsonResponse">ExtractJsonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /extract/markdown">client.Extract.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractService.Markdown">Markdown</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractMarkdownParams">ExtractMarkdownParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#ExtractMarkdownResponse">ExtractMarkdownResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Generate

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#GenerateJsonResponse">GenerateJsonResponse</a>

Methods:

- <code title="post /generate/json">client.Generate.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#GenerateService.Json">Json</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#GenerateJsonParams">GenerateJsonParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go">tabstack</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/tabstack-go#GenerateJsonResponse">GenerateJsonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
