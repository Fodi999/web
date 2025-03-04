package library

// AddBackgrounds добавляет стили фонов в карту CSS-классов.
func AddBackgrounds(cssClasses map[string]string) {
	backgrounds := map[string]string{
		// Background Attachment
		"bg-fixed":   "background-attachment: fixed;",
		"bg-local":   "background-attachment: local;",
		"bg-scroll":  "background-attachment: scroll;",

		// Background Clip
		"bg-clip-border":  "background-clip: border-box;",
		"bg-clip-padding": "background-clip: padding-box;",
		"bg-clip-content": "background-clip: content-box;",
		"bg-clip-text":    "background-clip: text;",

		// Background Origin
		"bg-origin-border":  "background-origin: border-box;",
		"bg-origin-padding": "background-origin: padding-box;",
		"bg-origin-content": "background-origin: content-box;",

		// Background Position
		"bg-bottom":        "background-position: bottom;",
		"bg-center":        "background-position: center;",
		"bg-left":          "background-position: left;",
		"bg-left-bottom":   "background-position: left bottom;",
		"bg-left-top":      "background-position: left top;",
		"bg-right":         "background-position: right;",
		"bg-right-bottom":  "background-position: right bottom;",
		"bg-right-top":     "background-position: right top;",
		"bg-top":           "background-position: top;",

		// Background Repeat
		"bg-repeat":       "background-repeat: repeat;",
		"bg-no-repeat":    "background-repeat: no-repeat;",
		"bg-repeat-x":     "background-repeat: repeat-x;",
		"bg-repeat-y":     "background-repeat: repeat-y;",
		"bg-repeat-round": "background-repeat: round;",
		"bg-repeat-space": "background-repeat: space;",

		// Background Size
		"bg-auto":    "background-size: auto;",
		"bg-cover":   "background-size: cover;",
		"bg-contain": "background-size: contain;",

		// Background Image
		"bg-none":            "background-image: none;",
		"bg-gradient-to-t":   "background-image: linear-gradient(to top, var(--tw-gradient-stops));",
		"bg-gradient-to-tr":  "background-image: linear-gradient(to top right, var(--tw-gradient-stops));",
		"bg-gradient-to-r":   "background-image: linear-gradient(to right, var(--tw-gradient-stops));",
		"bg-gradient-to-br":  "background-image: linear-gradient(to bottom right, var(--tw-gradient-stops));",
		"bg-gradient-to-b":   "background-image: linear-gradient(to bottom, var(--tw-gradient-stops));",
		"bg-gradient-to-bl":  "background-image: linear-gradient(to bottom left, var(--tw-gradient-stops));",
		"bg-gradient-to-l":   "background-image: linear-gradient(to left, var(--tw-gradient-stops));",
		"bg-gradient-to-tl":  "background-image: linear-gradient(to top left, var(--tw-gradient-stops));",
	}

	for key, value := range backgrounds {
		cssClasses[key] = value
	}
}

