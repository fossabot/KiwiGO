package command

type AllowedChatType int

const (
	All AllowedChatType = iota
	GuildOnly
	DMOnly
)

type Command struct {
	usage_doc string
	short_doc string
	long_doc  string

	name     string
	aliases  []string
	category string // *Category ?
	// bot_perms  []*Permission
	// user_perms []*Permission
	min_args int
	// flags      []*FlagOptions
	AllowedChatType AllowedChatType
	nsfw            bool
	hidden          bool
	disabled        bool
	// events     ???
}

type SubCommand struct {
	usage_doc string
	short_doc string
	long_doc  string

	name    string
	aliases []string
	// bot_perms  []*Permission
	// user_perms []*Permission
	min_args int
	nsfw     bool
	hidden   bool
	disabled bool
}

// TODO: check funcs
// TODO: errors
