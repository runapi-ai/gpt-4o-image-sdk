# GPT-4o Image API Ruby SDK for RunAPI

The gpt-4o image api Ruby SDK is the language-specific package for GPT-4o Image on RunAPI. Use this gpt-4o image api package for text-to-image, image editing, and creative production flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Ruby.

This gpt-4o image api README is the Ruby package guide inside the public `gpt4o-image-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/gpt-4o-image; for API reference, use https://runapi.ai/docs#gpt-4o-image; for SDK docs, use https://runapi.ai/docs#sdk-gpt-4o-image.

## Install

```bash
gem install runapi-gpt_4o_image
```

## Quick start

```ruby
require "runapi-gpt_4o_image"

client = RunApi::Gpt4oImage::Client.new
task = client.generations.create(
  # Pass the GPT-4o Image JSON request body from https://runapi.ai/docs#gpt-4o-image.
)
status = client.generations.get(task.id)
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Use Ruby keyword arguments and the `RunApi::Gpt4oImage` error classes when building image jobs, Rails workers, or scripts. The available resources include generations. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
