export const contract = {
  "text-to-image": {
    "models": [
      "gpt-4o-image"
    ],
    "fields_by_model": {
      "gpt-4o-image": {
        "aspect_ratio": {
          "enum": [
            "1:1",
            "3:2",
            "2:3"
          ],
          "required": true
        },
        "output_count": {
          "enum": [
            1,
            2,
            4
          ],
          "type": "integer"
        },
        "source_image_urls": {
          "max_items": 5
        }
      }
    }
  }
} as const;
