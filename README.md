# Labangram Agent Services

Public agent integration metadata for [Labangram](https://labangram.kamera-ichi.com), a Taiwan-based photography and digital product studio.

## Machine-readable surfaces

- Product MCP: <https://labangram.kamera-ichi.com/api/product-mcp>
- Documentation MCP: <https://labangram.kamera-ichi.com/api/mcp>
- MCP server card: <https://labangram.kamera-ichi.com/.well-known/mcp/server-card.json>
- MCP Registry identity: `io.github.lunkerchen/labangram`
- REST API and OpenAPI: <https://labangram.kamera-ichi.com/openapi.json>
- No-auth sandbox: <https://labangram.kamera-ichi.com/api/sandbox/v1/projects>
- Go SDK: <https://github.com/lunkerchen/labangram-agent/tree/main/sdk/go>

The product MCP exposes portfolio search, pricing and services, project lookup, and `submit_inquiry`. Inquiry preparation always returns `requires_human_confirmation`; it never sends or confirms a booking.

The documentation MCP exposes read-only portfolio, pricing, project, resource, and prompt guidance. Agents should use the documentation surface for explanation and the product surface for action preparation.

## Source and policy

Use live JSON-RPC responses and the public pricing/API documents as the source of truth. Do not invent project facts, prices, identities, ratings, or availability. Treat third-party directory listings as discovery metadata and link back to the canonical website.
