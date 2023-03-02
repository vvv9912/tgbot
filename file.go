package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
)

func savefile(outfile tgbotapi.File) (string, error) {
	nameid := uuid.New()
	pathserv := "https://api.telegram.org/file/" + "bot" + get_token() + "/" + outfile.FilePath
	w := strings.Split(outfile.FilePath, ".") //todo
	fileext := w[len(w)-1]
	//nameid + "." + fileext
	filename := fmt.Sprintf("%s.%s", &nameid, fileext)
	p1 := "database"
	p2 := "photo"
	f := path.Join(p1, p2, filename)
	resp, err := http.Get(pathserv)
	_ = err
	out, err := os.Create(f)
	defer out.Close()
	n, err := io.Copy(out, resp.Body)
	_ = n //todo
	defer resp.Body.Close()
	return f, err
}
