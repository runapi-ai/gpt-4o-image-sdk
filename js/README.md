# GPT-4o Image JavaScript SDK for RunAPI

The GPT-4o Image JavaScript SDK is the language-specific package for GPT-4o Image on RunAPI. Use this package for image generation, image editing, and creative production workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in JavaScript.

This README is the JavaScript package guide inside the public `gpt-4o-image-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/gpt-4o-image; for API reference, use https://runapi.ai/docs#gpt-4o-image; for SDK docs, use https://runapi.ai/docs#sdk-gpt-4o-image.

## Install

```bash
npm install @runapi.ai/gpt-4o-image
```

## Quick start

```typescript
import { Gpt4oImageClient } from '@runapi.ai/gpt-4o-image';

const client = new Gpt4oImageClient();
const task = await client.textToImage.create({
  // Pass the GPT-4o Image JSON request body from https://runapi.ai/docs#gpt-4o-image.
});
const status = await client.textToImage.get(task.id);
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Use the TypeScript types in `src/types.ts` and the resource classes under `src/resources` when building image applications. The available resources are `textToImage`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/gpt-4o-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-4o-image
- Product docs: https://runapi.ai/docs#gpt-4o-image
- Pricing and rate limits: https://runapi.ai/models/gpt-4o-image
- Provider comparison: https://runapi.ai/providers/openai
- Full catalog: https://runapi.ai/models
- Repository: https://github.com/runapi-ai/gpt4o-image-sdk

## License

Licensed under the Apache License, Version 2.0.
