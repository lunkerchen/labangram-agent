---
name: labangram-agent
description: Use Labangram's public MCP and API surfaces for portfolio discovery, service fit, pricing, and human-confirmed inquiry preparation.
---

# Labangram Agent Skill

## Routes

- Read-only documentation: `https://labangram.kamera-ichi.com/api/mcp`
- Product action preparation: `https://labangram.kamera-ichi.com/api/product-mcp`
- REST contract: `https://labangram.kamera-ichi.com/openapi.json`
- Test data: `https://labangram.kamera-ichi.com/api/sandbox/v1/projects`

## Safety

Use the sandbox for validation. Treat `submit_inquiry` as a draft-only operation and require a human to confirm before any outbound action. Use live `tools/list` and `/pricing.md` instead of guessing.

