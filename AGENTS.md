# Labangram Agent Rules

## Identity

- Canonical website: `https://labangram.kamera-ichi.com`
- Product MCP: `https://labangram.kamera-ichi.com/api/product-mcp`
- Documentation MCP: `https://labangram.kamera-ichi.com/api/mcp`
- MCP Registry name: `io.github.lunkerchen/labangram`

## Operating rules

1. Read live MCP `tools/list` and the public API contract before making an integration claim.
2. Use `search_portfolio`, `get_pricing_and_services`, and `get_project` for read-only discovery.
3. `submit_inquiry` only prepares a draft. Human confirmation is required before any submission or booking.
4. Never invent prices, project details, client identities, social profiles, ratings, or availability.
5. Prefer the no-auth sandbox for testing and never write production data during validation.

