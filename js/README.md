# GPT-4o Image API JavaScript SDK for RunAPI

The gpt-4o image api JavaScript SDK is the language-specific package for GPT-4o Image on RunAPI. Use this gpt-4o image api package for text-to-image, image editing, and creative production flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in JavaScript.

This gpt-4o image api README is the JavaScript package guide inside the public `gpt4o-image-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/gpt-4o-image; for API reference, use https://runapi.ai/docs#gpt-4o-image; for SDK docs, use https://runapi.ai/docs#sdk-gpt-4o-image.

## Install

```bash
npm install @runapi.ai/gpt4o-image
```

## Quick start

```typescript
import { Gpt4oImageClient } from '@runapi.ai/gpt4o-image';

const client = new Gpt4oImageClient();
const task = await client.generations.create({
  // Pass the GPT-4o Image JSON request body from https://runapi.ai/docs#gpt-4o-image.
});
const status = await client.generations.get(task.id);
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

## Language notes

Use the TypeScript types in `src/types.ts` and the resource classes under `src/resources` when building image applications. The available resources include generations. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
