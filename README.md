# GPT-4o Image API SDK for RunAPI

The gpt-4o image api SDK packages JavaScript, Ruby, and Go clients for GPT-4o Image on RunAPI. Use this gpt-4o image api SDK for text-to-image generation workflows that need typed installs, JSON request bodies, task polling, and consistent RunAPI errors across services.

GPT-4o Image belongs to the OpenAI catalog on RunAPI. The public model page is https://runapi.ai/models/gpt-4o-image; variant pages below carry pricing, rate-limit, and commercial-usage details. The public `gpt-4o-image-sdk` repository groups the JavaScript, Ruby, and Go packages for this model.

## Install

```bash
npm install @runapi.ai/gpt-4o-image
gem install runapi-gpt_4o_image
go get github.com/runapi-ai/gpt-4o-image-sdk/go@latest
```

## What you can build

- Build creative tools, agent pipelines, and production integrations with the gpt-4o image api SDK.
- Keep one model-specific repository while installing only the language package your app needs.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Handle authentication, validation, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

The JavaScript client exposes text to image resources, and the Ruby and Go packages mirror the same RunAPI task lifecycle.

## JavaScript quick start

```typescript
import { Gpt4oImageClient } from '@runapi.ai/gpt-4o-image';

const client = new Gpt4oImageClient();

const task = await client.textToImage.create({
  // Pass the GPT-4o Image request body documented at https://runapi.ai/docs#gpt-4o-image.
});

const status = await client.textToImage.get(task.id);
```

For short scripts, use `run` with the same JSON body to create the task and wait for completion. For web request handlers, prefer `create` plus webhook or later `get` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/gpt-4o-image`.
- `ruby/` publishes `runapi-gpt_4o_image` when RubyGems publishing resumes.
- `go/` publishes `github.com/runapi-ai/gpt-4o-image-sdk/go` and depends on `github.com/runapi-ai/core-sdk/go`.

## Public links

- Model page: https://runapi.ai/models/gpt-4o-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-4o-image
- Product docs: https://runapi.ai/docs#gpt-4o-image
- SDK repository: https://github.com/runapi-ai/gpt-4o-image-sdk
- Skill repository: https://github.com/runapi-ai/gpt-4o-image
- Provider comparison: https://runapi.ai/providers/openai
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific gpt-4o image api variant page for pricing, rate limits, and commercial usage:
- [GPT-4o Image](https://runapi.ai/models/gpt-4o-image/gpt-4o-image)

Default pricing link for the gpt-4o image api SDK: https://runapi.ai/models/gpt-4o-image/gpt-4o-image

## FAQ

### Which package should I install for gpt-4o image api work?

Install the model package for your language: `@runapi.ai/gpt-4o-image`, `runapi-gpt_4o_image`, or `github.com/runapi-ai/gpt-4o-image-sdk/go`. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary gpt-4o image api links point to https://runapi.ai/models/gpt-4o-image. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/gpt-4o-image/gpt-4o-image. Provider comparisons point to https://runapi.ai/providers/openai, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
