import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { TextToImage } from './resources/text-to-image';

/**
 * GPT-4o Image text-to-image API client.
 */
export class Gpt4oImageClient {
  public readonly textToImage: TextToImage;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.textToImage = new TextToImage(http);
  }
}
