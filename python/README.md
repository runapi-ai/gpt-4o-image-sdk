# GPT-4o Image API Python SDK for RunAPI

The gpt-4o image api Python SDK is the language-specific package for GPT-4o Image on RunAPI. Use this gpt-4o image api package for text-to-image, image editing, and creative production flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Python.

This gpt-4o image api README is the Python package guide inside the public `gpt4o-image-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/gpt-4o-image; for API reference, use https://runapi.ai/docs#gpt-4o-image; for SDK docs, use https://runapi.ai/docs#sdk-gpt-4o-image.

## Install

```bash
pip install runapi-gpt-4o-image
```

## Quick start

```python
from runapi.gpt_4o_image import Gpt4oImageClient

client = Gpt4oImageClient()  # reads RUNAPI_API_KEY, or pass api_key="sk-..."

task = client.text_to_image.create(
    model="gpt-4o-image",
    prompt="A cinematic product photo on warm paper",
    aspect_ratio="1:1",
)
status = client.text_to_image.get(task.id)
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion:

```python
result = client.text_to_image.run(
    model="gpt-4o-image",
    prompt="A futuristic cityscape at dusk",
    aspect_ratio="3:2",
)
print(result.images[0].url)
```

In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.gpt_4o_image` error classes when building image jobs or scripts. The available resource is `text_to_image`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
