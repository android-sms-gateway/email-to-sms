package start

import (
	"fmt"

	"github.com/capcom6/go-project-template/internal/bot/handler"
	"github.com/go-core-fx/telegofx"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) handler.Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Register(router *telegofx.Router) {
	router.Handle(
		h.handleStart,
		th.CommandEqual("start"),
		th.AnyMessageWithFrom(),
	)
}

func (h *Handler) handleStart(ctx *th.Context, update telego.Update) error {
	if update.Message == nil || update.Message.From == nil {
		h.logger.Warn("received /start update without message sender")
		return nil
	}

	return h.reply(ctx, update.Message.Chat.ID, "Hello, "+update.Message.From.Username)
}

func (h *Handler) reply(ctx *th.Context, chatID int64, text string) error {
	_, err := ctx.Bot().SendMessage(ctx, &telego.SendMessageParams{
		ChatID: telego.ChatID{ID: chatID, Username: ""},
		Text:   text,
	})
	if err != nil {
		return fmt.Errorf("send telegram message: %w", err)
	}

	return nil
}
