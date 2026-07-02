import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { TextToImage } from './resources/text-to-image';

/**
 * GPT-4o Image generation and editing API client.
 *
 * Supports pure text-to-image generation, image editing with source images,
 * and inpainting with an optional mask. At least one of `prompt` or
 * `source_image_urls` must be provided.
 *
 * @example
 * ```typescript
 * const client = new Gpt4oImageClient({ apiKey: 'your-api-key' });
 *
 * // Pure generation
 * const result = await client.textToImage.run({
 *   model: 'gpt-4o-image',
 *   prompt: 'A watercolor painting of a sunset',
 *   aspect_ratio: '3:2',
 * });
 *
 * // Edit with source images
 * const edited = await client.textToImage.run({
 *   model: 'gpt-4o-image',
 *   prompt: 'Add a rainbow',
 *   aspect_ratio: '3:2',
 *   source_image_urls: ['https://cdn.runapi.ai/public/samples/image.jpg'],
 * });
 * ```
 */
export class Gpt4oImageClient extends BaseClient {
  /** Image generation, editing, and inpainting operations. */
  public readonly textToImage: TextToImage;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.textToImage = new TextToImage(this.http);
  }
}
