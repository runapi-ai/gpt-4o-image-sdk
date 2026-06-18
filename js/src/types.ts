import type { AsyncTaskStatus } from '@runapi.ai/core';

export type Gpt4oImageModel = 'gpt-4o-image';
/** Required output aspect ratio. */
export type AspectRatio = '1:1' | '3:2' | '2:3';
/** Batch output size: 1, 2, or 4 images per request. */
export type OutputCount = 1 | 2 | 4;

/**
 * Parameters for GPT-4o image generation and editing.
 *
 * - For pure generation, provide `prompt` (required) and `aspect_ratio`.
 * - For editing, provide `source_image_urls` (up to 5); `prompt` becomes
 *   optional since the model can generate variants without one.
 * - For inpainting, also provide `mask_url` to indicate which regions of
 *   the source image to regenerate.
 */
export interface TextToImageParams {
  model: Gpt4oImageModel;
  /** Text prompt; required when source_image_urls is not provided. */
  prompt?: string;
  aspect_ratio: AspectRatio;
  /** Up to 5 source image URLs for edits or variants. */
  source_image_urls?: string[];
  /** Mask image URL marking regions to regenerate (inpainting); only meaningful with source_image_urls. */
  mask_url?: string;
  output_count?: OutputCount;
  callback_url?: string;
  /** Let the model expand short prompts with additional detail. */
  enable_prompt_expansion?: boolean;
}

export interface TaskCreateResponse {
  id: string;
  status: string;
}

/** A generated image with its CDN URL. */
export interface Image {
  url: string;
}

export interface TextToImageResponse {
  id: string;
  status: AsyncTaskStatus;
  images?: Image[];
  /** Intermediate status string while the task is running. */
  progress?: string;
  error?: string;
  [key: string]: unknown;
}
