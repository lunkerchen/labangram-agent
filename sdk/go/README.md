# Labangram Go SDK

Dependency-free Go client for Labangram's public portfolio, services, pricing, and product MCP endpoints.

```go
client := labangram.New(nil)
projects, err := client.GetProjects(context.Background(), "camera", "digital")
```

Install with `go get github.com/lunkerchen/labangram-agent/sdk/go`. The client only uses public read APIs and the draft-safe product MCP surface; inquiry preparation still requires human confirmation.
