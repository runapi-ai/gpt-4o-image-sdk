---
name: gpt-4o-image
description: Generate images with GPT-4o Image text-to-image through RunAPI.ai using the @runapi.ai/gpt-4o-image Node/TypeScript SDK. Use when the user asks for GPT-4o image generation, or writes against @runapi.ai/gpt-4o-image. Triggers on "gpt-4o-image", "GPT-4o Image", "gpt4o image", "@runapi.ai/gpt-4o-image".
documentation: https://runapi.ai/models/gpt-4o-image
provider_page: https://runapi.ai/providers/openai
catalog: https://runapi.ai/models
---
# @runapi.ai/gpt-4o-image -- RunAPI.ai GPT-4o Image generation

Build Node / TypeScript integrations that generate images with GPT-4o Image through RunAPI.ai.

## Setup

Requires **Node 18+** (global `fetch`).

```bash
npm install @runapi.ai/gpt-4o-image
```

```dotenv
# .env
RUNAPI_API_KEY=runapi_xxx   # get one at https://runapi.ai/settings/api_keys
```

```ts
import { Gpt4oImageClient } from '@runapi.ai/gpt-4o-image';

const client = new Gpt4oImageClient();
```

Pass `{ apiKey }` explicitly if you manage secrets differently. `baseUrl` defaults to `https://runapi.ai`; override only for local development.

## Resource

`client.textToImage` uses the async task contract:

```ts
const { id } = await client.textToImage.create({ ... });
const status = await client.textToImage.get(id);
const result = await client.textToImage.run({ ... });
```

## Text to image

```ts
const result = await client.textToImage.run({
  model: 'gpt-4o-image',
  prompt: 'A minimalist logo of a mountain range at sunset',
});

const url = result.images[0].url;
```

## Errors

All errors are re-exported from `@runapi.ai/core`. Use `instanceof` checks instead of string-matching messages. For long-running tasks, prefer `create()` plus webhook or `get(id)` in request handlers, and reserve `run()` for jobs / CLI.

## RunAPI public routing

gpt-4o image api public links use the API-379 catalog route map. The main gpt-4o image api page is https://runapi.ai/models/gpt-4o-image. SDK docs live at https://runapi.ai/docs#sdk-gpt-4o-image and product docs live at https://runapi.ai/docs#gpt-4o-image.

Pricing, rate limits, and commercial usage for gpt-4o image api should point to the most specific variant page:
- [GPT-4o Image](https://runapi.ai/models/gpt-4o-image/gpt-4o-image)

Compare GPT-4o Image with other OpenAI models at https://runapi.ai/providers/openai. Browse every RunAPI model and skill at https://runapi.ai/models. SDK repository: https://github.com/runapi-ai/gpt-4o-image-sdk. Skill repository: https://github.com/runapi-ai/gpt-4o-image.
