# GPT-4o Image API Skill for RunAPI

Generate images with GPT-4o Image text-to-image. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate GPT-4o Image through RunAPI.

The canonical agent file is `skills/gpt-4o-image/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/gpt-4o-image -g
```

Or manually: clone this repo and copy `skills/gpt-4o-image/` into your agent's skills directory.

## Quick example

```typescript
import { Gpt4oImageClient } from '@runapi.ai/gpt-4o-image';

const client = new Gpt4oImageClient();
const result = await client.textToImage.run({
  model: 'gpt-4o-image',
  prompt: 'A minimalist logo of a mountain range at sunset',
});
const url = result.images[0].url;
```

## Routing

- Model page: https://runapi.ai/models/gpt-4o-image
- Product docs: https://runapi.ai/docs#gpt-4o-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-4o-image
- SDK repository: https://github.com/runapi-ai/gpt-4o-image-sdk
- Pricing and rate limits: https://runapi.ai/models/gpt-4o-image/gpt-4o-image
- Provider comparison: https://runapi.ai/providers/openai
- Browse all RunAPI models and skills: https://runapi.ai/models

## Variants

- [GPT-4o Image](https://runapi.ai/models/gpt-4o-image/gpt-4o-image)

## Agent rules

- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For gpt-4o image api pricing, rate-limit, and commercial-usage answers, link to the variant page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
