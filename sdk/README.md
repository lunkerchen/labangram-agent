# @lunkerchen/labangram-agent

Fetch-based JavaScript SDK for the public [Labangram](https://labangram.kamera-ichi.com) REST and product MCP surfaces.

```sh
npm install @lunkerchen/labangram-agent
```

```js
import { LabangramClient } from '@lunkerchen/labangram-agent';

const labangram = new LabangramClient();
const { projects } = await labangram.getProjects({ category: 'commercial' });
const pricing = await labangram.getPricing();
const draft = labangram.prepareInquiry({
  name: 'Ada',
  email: 'ada@example.com',
  message: 'Commercial product photography',
});
```

`prepareInquiry` is local and returns `requires_human_confirmation`; it never sends a booking. `callProductTool` calls the public MCP endpoint for read-only tools or an explicitly confirmed action flow.

Source of truth: [OpenAPI](https://labangram.kamera-ichi.com/openapi.json), [product MCP](https://labangram.kamera-ichi.com/api/product-mcp), and the [public agent repository](https://github.com/lunkerchen/labangram-agent).
