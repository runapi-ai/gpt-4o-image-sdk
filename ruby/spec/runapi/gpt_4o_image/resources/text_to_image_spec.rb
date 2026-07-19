# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::Gpt4oImage::Resources::TextToImage do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:text_to_image) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gpt_4o_image/text_to_image" }

  describe "#create" do
    it "POSTs to the correct endpoint with direct service params" do
      params = {
        model: "gpt-4o-image",
        prompt: "a still life",
        aspect_ratio: "1:1",
        source_image_urls: ["https://cdn.runapi.ai/public/samples/input.png"],
        output_count: 2
      }
      expect(http).to receive(:request).with(:post, endpoint, body: params).and_return("id" => "task-1")

      result = text_to_image.create(**params)
      expect(result).to be_a(RunApi::Gpt4oImage::Types::TextToImageResponse)
      expect(result.id).to eq("task-1")
    end

    it "raises ValidationError when model is missing" do
      expect { text_to_image.create(aspect_ratio: "1:1") }
        .to raise_error(RunApi::Core::ValidationError, /model must be one of: gpt-4o-image/)
    end
  end

  describe "#get" do
    it "GETs the correct endpoint with task id" do
      expect(http).to receive(:request).with(:get, "#{endpoint}/task-123")
        .and_return("id" => "task-123", "status" => "completed", "progress" => "1.00", "images" => [{"url" => "https://file.runapi.ai/result.png"}])

      result = text_to_image.get("task-123")
      expect(result.id).to eq("task-123")
      expect(result.images.first.url).to eq("https://file.runapi.ai/result.png")
    end
  end
end
