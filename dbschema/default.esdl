module default {
	type Shortening {
		required property identifier -> str {
			constraint exclusive;
		};
		link created_by -> User;
		required property original_url -> str;
		required property visits -> int64 {
			default := 0;
		};
		required property created_at -> datetime {
			default := datetime_of_transaction();
		};
		property updated_at -> datetime;
		index on (.identifier);
	}

	type User {
		required property is_active -> bool {
			default := true;
		};
		required property gh_login -> str {
			constraint exclusive;
		};
		property gh_access_key -> str;
		required property created_at -> datetime {
			default := datetime_of_transaction();
		};
	}
}
