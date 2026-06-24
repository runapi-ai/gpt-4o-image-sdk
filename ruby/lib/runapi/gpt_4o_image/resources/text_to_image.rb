# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    module Resources
      # GPT-4o Image generation and editing resource.
      # Generate images from text, edit with source images, or inpaint with a mask.
      class TextToImage
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gpt_4o_image/text_to_image"

        RESPONSE_CLASS = Types::TextToImageResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedTextToImageResponse

        def initialize(http)
          @http = http
        end

        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        def create(**params)
          params = compact_params(params)
          validate_contract!(CONTRACT["text-to-image"], params)
          request(:post, ENDPOINT, body: params)
        end

        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end
      end
    end
  end
end
