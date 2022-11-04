module default {
	type Shortening {
		required property identifier -> str {
			constraint exclusive;
		};
		required property original_url -> str;
		required property created_at -> datetime;
		index on (.identifier);
	}
}
