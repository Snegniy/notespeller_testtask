package service

import (
	"errors"
	"fmt"
	"github.com/Snegniy/notespeller-testtask/pkg/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

func speller(s string) error {
	temp := strings.ReplaceAll(s, " ", "+")

	url := fmt.Sprintf("https://speller.yandex.net/services/spellservice.json/checkText?text=%s", temp)
	resp, err := http.Get(url)
	if err != nil {
		logger.Warn("Couldn't check note", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("Couldn't read response body", zap.Error(err))
		return err
	}

	if len(body) > 2 {
		return errors.New("spelling error in the message")
	}
	return nil
}
