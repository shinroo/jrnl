package usage

var (
	Commands = map[string]string{
		"init":       Init,
		"title":      Title,
		"post":       Post,
		"edit":       Edit,
		"rm":         Rm,
		"ls":         Ls,
		"remote":     Remote,
		"remote-ls":  RemoteLs,
		"remote-set": RemoteSet,
		"remote-rm":  RemoteRm,
		"publish":    Publish,
		"theme":      Theme,
		"theme-ls":   ThemeLs,
		"theme-save": ThemeSave,
		"theme-use":  ThemeUse,
		"theme-rm":   ThemeRm,
	}

	Jrnl = `jrnl - A simple static site generator

Usage:

  jrnl [command] [options...]

Commands:

  init     Initialize a new journal
  title    Set the journal's title
  post     Create a post
  edit     Edit a post
  rm       Remove a post
  ls       List all journal posts
  remote   Manage the journal's remotes
  publish  Publish the journal's posts
  theme    Manage the journal's themes

Options:

  --help  Display this usage message

For more information on a command run 'jrnl [command] -- help'`
)
