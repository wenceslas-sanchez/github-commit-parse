package commit

type NestedCounter map[string]map[string]int

func Counters(commits *[]*Commit) (*map[string]int, *NestedCounter) {
	commitCounter := make(map[string]int)
	emojiCounter := make(NestedCounter)
	pEmojiCounter := &emojiCounter

	for _, c := range *commits {
		commit := (*c).Commit
		login := (*c).Author.Login
		message := commit.Message
		// Ignore Merge messages (automatic one)
		if message.IsMergeMessage() {
			continue
		}
		commitCounter[login] += 1
		if message.ContainsEmoji() {
			countEmojis(&message, &login, pEmojiCounter)
		}
	}

	return &commitCounter, pEmojiCounter
}

func countEmojis(m *Message, login *string, counter *NestedCounter) {
	for _, emoji := range m.FindEmojis() {
		// Can't use []byte as key, so it's needed to convert it to a string
		convEmoji := string(emoji)
		if _, ok := (*counter)[*login]; !ok {
			(*counter)[*login] = map[string]int{convEmoji: 1}
			continue
		}
		(*counter)[*login][convEmoji] += 1
	}
}
