CONTRACT = {
    "text-to-image": {
        "models": ["gpt-4o-image"],
        "fields_by_model": {
            "gpt-4o-image": {
                "aspect_ratio": {
                    "enum": ["1:1", "3:2", "2:3"],
                    "required": True
                },
                "output_count": {
                    "enum": [1, 2, 4],
                    "type": "integer"
                }
            }
        }
    }
}
