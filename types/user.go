package types

// User - is implement https://core.telegram.org/bots/api#user
type User struct {
	ID                    int64   `json:"id"`                                // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifie
	IsBot                 bool    `json:"is_bot"`                            // True, if this user is a bot
	FirstName             string  `json:"first_name"`                        // User's or bot's first name
	LastName              *string `json:"last_name,omitempty"`               // User's or bot's last name
	UserName              *string `json:"username,omitempty"`                // User's or bot's username
	LanguageCode          *string `json:"language_code,omitempty"`           // IETF language tag of the user's language
	IsPremium             *bool   `json:"is_premium,omitempty"`              // True, if this user is a Telegram Premium user
	AddedToAttachmentMenu *bool   `json:"added_to_attachment_menu"`          // True, if this user added the bot to the attachment menu
	CanJoinGroups         *bool   `json:"can_join_groups,omitempty"`         // True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroups      *bool   `json:"can_read_all_groups,omitempty"`     // True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries *bool   `json:"supports_inline_queries,omitempty"` // True, if the bot supports inline queries. Returned only in getMe.
	CanConnectToBusiness  *bool   `json:"can_connect_to_business,omitempty"` // True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
}
