package commands

import (
	"strings"

	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix      string
	cmdInstance []Command
	cmdMap      map[string]Command
	middlewares []Middleware
	OnError     func(err error, ctx *Context)
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:      prefix,
		cmdInstance: make([]Command, 0),
		cmdMap:      make(map[string]Command),
		middlewares: make([]Middleware, 0),
		OnError:     func(error, *Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstance = append(c.cmdInstance, cmd)
	for _, v := range cmd.Invokes() {
		c.cmdMap[v] = cmd
	}
}

func (c *CommandHandler) RegisterMiddleware(mw Middleware) {
	c.middlewares = append(c.middlewares, mw)
}

func (c *CommandHandler) HandleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	logger.Debug("Handling message: %s", e.Content)
	if e.Author.ID == s.State.User.ID || e.Author.Bot || !strings.HasPrefix(e.Content, c.prefix) {
		return
	}

	split := strings.Split(e.Content[len(c.prefix):], " ")
	logger.Debug("Split: %s", split)

	if len(split) < 1 {
		return
	}

	invoke := strings.ToLower(split[0])
	args := split[1:]

	logger.Debug("Invoke: %s", invoke)
	logger.Debug("Args: %v", args)

	cmd, ok := c.cmdMap[invoke]

	if !ok || cmd == nil {
		return
	}

	ctx := &Context{
		Session: s,
		Args:    args,
		Handler: c,
		Message: e.Message,
	}

	for _, mw := range c.middlewares {
		next, err := mw.Exec(ctx, cmd)

		if err != nil {
			c.OnError(err, ctx)
			return
		}

		if !next {
			return
		}

	}

	if err := cmd.Exec(ctx); err != nil {
		c.OnError(err, ctx)
	}

}
