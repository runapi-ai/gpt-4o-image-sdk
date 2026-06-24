<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center">
  <a href="https://github.com/runapi-ai/gpt-4o-image-sdk">GPT-4o Image API SDK for RunAPI</a>
</h3>

<p align="center">
  GPT-4o Image API SDKs for JavaScript, Python, Ruby, Go, and Java on RunAPI.
</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/gpt-4o-image)](https://www.npmjs.com/package/@runapi.ai/gpt-4o-image)
[![PyPI](https://img.shields.io/pypi/v/runapi-gpt-4o-image)](https://pypi.org/project/runapi-gpt-4o-image/)
[![RubyGems](https://img.shields.io/gem/v/runapi-gpt_4o_image)](https://rubygems.org/gems/runapi-gpt_4o_image)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/gpt-4o-image-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/gpt-4o-image-sdk/go)
[![Maven Central](https://img.shields.io/maven-central/v/ai.runapi/runapi-gpt-4o-image)](https://central.sonatype.com/artifact/ai.runapi/runapi-gpt-4o-image)
[![License](https://img.shields.io/github/license/runapi-ai/gpt-4o-image-sdk)](https://github.com/runapi-ai/gpt-4o-image-sdk/blob/main/LICENSE)

</div>
<br/>

The GPT-4o Image API SDK packages JavaScript, Python, Ruby, Go, and Java clients for GPT-4o Image on RunAPI. Use it for text-to-image generation workflows when your app needs typed request builders, predictable task polling, file upload helpers, account helpers, and consistent RunAPI errors.

GPT-4o Image is listed in the RunAPI model catalog at https://runapi.ai/models/gpt-4o-image. Variant pages below carry pricing, rate-limit, and commercial-usage details. The public `gpt-4o-image-sdk` repository groups the language packages, examples, CI, and release tags for this model.

## Install

```bash
npm install @runapi.ai/gpt-4o-image
pip install runapi-gpt-4o-image
gem install runapi-gpt_4o_image
go get github.com/runapi-ai/gpt-4o-image-sdk/go@latest
```

Gradle:

```kotlin
dependencies {
  implementation("ai.runapi:runapi-gpt-4o-image:0.1.0")
}
```

Maven:

```xml
<dependency>
  <groupId>ai.runapi</groupId>
  <artifactId>runapi-gpt-4o-image</artifactId>
  <version>0.1.0</version>
</dependency>
```

Use the Java BOM when installing multiple RunAPI Java modules:

```kotlin
dependencies {
  implementation(platform("ai.runapi:runapi-bom:0.1.0"))
  implementation("ai.runapi:runapi-gpt-4o-image")
}
```

## What you can build

- Build apps, agent workflows, batch jobs, and production services around GPT-4o Image requests.
- Install only the language package your app needs while keeping one model-specific repository for docs and releases.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Upload local files, URL files, or base64 files through shared RunAPI file helpers.
- Handle validation, authentication, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

## Java quick start

```java
import ai.runapi.gpt4oimage.Gpt4oImageClient;
import ai.runapi.gpt4oimage.types.TextToImageParams;
import ai.runapi.gpt4oimage.types.CompletedTextToImageResponse;
import ai.runapi.gpt4oimage.types.TextToImageModel;

Gpt4oImageClient client = Gpt4oImageClient.builder()
    .apiKey(System.getenv("RUNAPI_API_KEY"))
    .build();

CompletedTextToImageResponse result = client.textToImage().run(
    TextToImageParams.builder()
        .model(TextToImageModel.GPT_4O_IMAGE)
        .aspectRatio("1:1")
        .prompt("A friendly robot reading a Java manual")
        .build()
);
```

Java packages target Java 8 bytecode and are tested on Java 8, 11, 17, and 21. Each model artifact depends on `ai.runapi:runapi-core`, so application code normally installs only `ai.runapi:runapi-gpt-4o-image`.

## Task lifecycle

Most media endpoints are asynchronous. `create()` submits a task and returns its id, `get(id)` fetches the latest task state, and `run(params)` creates the task and polls until it reaches a terminal state. In web request handlers, prefer `create()` plus webhook or later `get()` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/gpt-4o-image`.
- `python/` publishes `runapi-gpt-4o-image`.
- `ruby/` publishes `runapi-gpt_4o_image` when RubyGems publishing resumes.
- `go/` publishes `github.com/runapi-ai/gpt-4o-image-sdk/go` and depends on `github.com/runapi-ai/core-sdk/go`.
- `java/` publishes `ai.runapi:runapi-gpt-4o-image` and depends on `ai.runapi:runapi-core`.

## Public links

- Model page: https://runapi.ai/models/gpt-4o-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-4o-image
- Product docs: https://runapi.ai/docs#gpt-4o-image
- SDK repository: https://github.com/runapi-ai/gpt-4o-image-sdk
- Skill repository: https://github.com/runapi-ai/gpt-4o-image
- Provider comparison: https://runapi.ai/providers/openai
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific GPT-4o Image variant page for pricing, rate limits, and commercial usage:
- [GPT-4o Image](https://runapi.ai/models/gpt-4o-image)

Default pricing link for the GPT-4o Image SDK: https://runapi.ai/models/gpt-4o-image

## File storage

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## FAQ

### Which package should I install for GPT-4o Image work?

Install the model package for your language: `@runapi.ai/gpt-4o-image` on npm, `runapi-gpt-4o-image` on PyPI, `runapi-gpt_4o_image` on RubyGems, `github.com/runapi-ai/gpt-4o-image-sdk/go`, or `ai.runapi:runapi-gpt-4o-image`. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary GPT-4o Image links point to https://runapi.ai/models/gpt-4o-image. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/gpt-4o-image. Provider comparisons point to https://runapi.ai/providers/openai, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
