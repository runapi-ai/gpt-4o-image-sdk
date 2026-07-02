# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    # GPT-4o Image generation and editing API client.
    #
    # Supports pure text-to-image generation, image editing with source images,
    # and inpainting with an optional mask.
    #
    # @example
    #   client = RunApi::Gpt4oImage::Client.new(api_key: "your-api-key")
    #
    #   # Pure generation
    #   result = client.text_to_image.run(
    #     model: "gpt-4o-image", prompt: "A mountain lake at dawn",
    #     aspect_ratio: "3:2"
    #   )
    #
    #   # Edit with source images
    #   edited = client.text_to_image.run(
    #     model: "gpt-4o-image", prompt: "Add a rainbow",
    #     aspect_ratio: "3:2",
    #     source_image_urls: ["https://cdn.runapi.ai/public/samples/image.jpg"]
    #   )
    class Client < RunApi::Core::Client
      # @return [Resources::TextToImage] Image generation and editing operations.
      attr_reader :text_to_image

      def initialize(api_key: nil, **options)
        super
        @text_to_image = Resources::TextToImage.new(http)
      end
    end
  end
end
