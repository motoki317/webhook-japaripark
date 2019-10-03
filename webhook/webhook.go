package webhook

import (
	"github.com/labstack/echo"
	"github.com/motoki317/webhook-japaripark/model"
	"log"
	"net/http"
	"os/exec"
)

func MakeWebhookHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		event := c.Request().Header.Get("X-Gitea-Event")
		log.Printf("Received event %s\n", event)

		switch event {
		case "push":
			return handlePushEvent(c)
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func handlePushEvent(c echo.Context) error {
	payload := model.PushEvent{}
	if err := c.Bind(&payload); err != nil {
		log.Printf("Error occured while binding payload: %s\n", err)
		return err
	}

	if payload.Ref == "ref/heads/master" {
		// deploy
		log.Println("Deploying...")
		go func() {
			err := exec.Command("sh", "deploy.sh").Run()
			if err != nil {
				log.Printf("Error while executing deploy shell commands: %s\n", err)
			} else {
				log.Println("Deploy complete")
			}
		}()
	}

	return c.NoContent(http.StatusNoContent)
}
