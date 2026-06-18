<p align="center">
  <a href="https://github.com/runapi-ai/gpt-4o-image">
    <h3 align="center">GPT-4o Image API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill, inspect GPT-4o Image fields, then run jobs through the RunAPI CLI.
</p>

<p align="center">
  <a href="https://runapi.ai/models/gpt-4o-image"><strong>Model Reference</strong></a> · <a href="https://github.com/runapi-ai/cli"><strong>CLI</strong></a> · <a href="https://github.com/runapi-ai/gpt-4o-image-sdk"><strong>SDK</strong></a>
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/gpt-4o-image)](https://www.skills.sh/runapi-ai/gpt-4o-image/gpt-4o-image)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--gpt--4o--image-111827)](https://clawhub.ai/runapi-ai/runapi-gpt-4o-image)
[![License](https://img.shields.io/github/license/runapi-ai/gpt-4o-image)](https://github.com/runapi-ai/gpt-4o-image/blob/main/LICENSE)

</div>
<br/>

Generate images with GPT-4o Image text-to-image. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate GPT-4o Image through RunAPI.

The canonical agent file is `skills/gpt-4o-image/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/gpt-4o-image -g
```

Or paste this prompt to your AI agent:

```text
Install the gpt-4o-image skill for me:

1. Clone https://github.com/runapi-ai/gpt-4o-image
2. Copy the skills/gpt-4o-image/ directory into your
   user-level skills directory (e.g. ~/.claude/skills/
   for Claude Code, ~/.codex/skills/ for Codex).
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

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
- Pricing and rate limits: https://runapi.ai/models/gpt-4o-image
- Provider comparison: https://runapi.ai/providers/openai
- Browse all RunAPI models and skills: https://runapi.ai/models

## Agent rules

- Integration work uses the target language SDK; one-off generation, manual smoke tests, debugging, or user-requested CLI runs use the RunAPI CLI skill: https://github.com/runapi-ai/cli-skill
- RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.
- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For gpt-4o image api pricing, rate-limit, and commercial-usage answers, link to the model page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
