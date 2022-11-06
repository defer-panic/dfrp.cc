module default {
	type Shortening {
		required property identifier -> str {
			constraint exclusive;
		};
		required property original_url -> str;
		property visits -> int64;
		required property created_at -> datetime;
		property updated_at -> datetime;
		index on (.identifier);
	}
}
