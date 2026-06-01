import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { TextToImage } from '../../src/resources/text-to-image';
import type { TextToImageResponse, TaskCreateResponse } from '../../src/types';

describe('Gpt4oImage TextToImage', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('creates a generation task with direct service params', async () => {
    const mockResponse: TaskCreateResponse = { id: 'task-gen-123', status: 'processing' };
    vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

    const textToImage = new TextToImage(mockHttp);
    const result = await textToImage.create({
      model: 'gpt-4o-image',
      prompt: 'A still life',
      aspect_ratio: '1:1',
      source_image_urls: ['https://cdn.runapi.ai/public/samples/input.png'],
      output_count: 2,
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/gpt_4o_image/text_to_image', {
      body: {
        model: 'gpt-4o-image',
        prompt: 'A still life',
        aspect_ratio: '1:1',
        source_image_urls: ['https://cdn.runapi.ai/public/samples/input.png'],
        output_count: 2,
      },
    });
    expect(result).toEqual(mockResponse);
  });

  it('gets a generation task', async () => {
    const mockResponse: TextToImageResponse = {
      id: 'task-gen-456',
      status: 'completed',
      progress: '1.00',
      images: [{ url: 'https://file.runapi.ai/result.png' }],
    };
    vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

    const textToImage = new TextToImage(mockHttp);
    const result = await textToImage.get('task-gen-456');

    expect(mockHttp.request).toHaveBeenCalledWith('GET', '/api/v1/gpt_4o_image/text_to_image/task-gen-456', {});
    expect(result).toEqual(mockResponse);
  });

});
