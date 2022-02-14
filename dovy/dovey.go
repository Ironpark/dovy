package dovy

import (
	"context"
	"dovey/dovy/twitch"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	gort "runtime"
	"sync"
)

const TwitchClientId = "xxt39lhzhqa7z2wwfdzjkdrtq4cj21"

type Dovy struct {
	appCtx        context.Context
	accessToken   string
	scope         []string
	lock          sync.Mutex
	tokenReceiver *twitch.TokenReceiver
	cm            *twitch.ConnectionManager
}

func NewDovey(ctx context.Context) (*Dovy, error) {
	cm, err := twitch.NewConnectionManager()
	if err != nil {
		return nil, err
	}
	dovy := &Dovy{
		appCtx:      ctx,
		accessToken: "",
		scope: []string{
			"channel:moderate",
			"chat:edit",
			"chat:read",
			"whispers:read",
			"whispers:edit",
			"moderator:manage:chat_settings",
			"channel:edit:commercial",
			"channel:manage:broadcast",
			"user:manage:blocked_users",
			"user:read:blocked_users",
			"moderator:manage:banned_users",
			"moderator:read:blocked_terms",
		},
		lock:          sync.Mutex{},
		tokenReceiver: twitch.NewTokenReceiver(),
		cm:            cm,
	}

	dovy.tokenReceiver.SetTokenRecvCallback(func(token string) {
		cm.Initialize(token)
		cm.ConnectOrGet("woowakgood")
	})
	cm.OnMessage(func(message twitch.Message) {
		runtime.EventsEmit(ctx, "chat.stream", message)
	})
	go dovy.tokenReceiver.Serve()
	return dovy, nil
}

func (dov Dovy) OpenAuthorization() {
	authUrl := twitch.GetAuthorizationURL(dov.scope, true)
	fmt.Println(authUrl)
	switch gort.GOOS {
	case "darwin":
		exec.Command("open", authUrl).Run()
	case "windows":
		exec.Command("start", "/max", authUrl).Run()
	}
}

func (dov *Dovy) Connect(channelName string) {
	dov.cm.ConnectOrGet(channelName)
}

func (dov *Dovy) SendChatMessage(msg string) {
	fmt.Println("SendChatMessage", "woowakgood", msg)
}
