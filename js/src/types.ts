import type { AsyncTaskStatus } from '@runapi.ai/core';

export type Gpt4oImageModel = 'gpt-4o-image';
export type AspectRatio = '1:1' | '3:2' | '2:3';
export type OutputCount = 1 | 2 | 4;

export interface TextToImageParams {
  model: Gpt4oImageModel;
  prompt?: string;
  aspect_ratio: AspectRatio;
  source_image_urls?: string[];
  mask_url?: string;
  output_count?: OutputCount;
  callback_url?: string;
  enable_prompt_expansion?: boolean;
}

export interface TaskCreateResponse {
  id: string;
  status: string;
}

export interface Image {
  url: string;
}

export interface TextToImageResponse {
  id: string;
  status: AsyncTaskStatus;
  images?: Image[];
  progress?: string;
  error?: string;
  [key: string]: unknown;
}
