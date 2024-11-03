package main

import (
	"bufio"
	"context"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if err := run(); err != nil {
		println(err.Error())

		os.Exit(1)
	}
}

func run() error {
	var ctx = context.Background()

	// открываем файл
	f, err := os.Open("links_list.txt")
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	// читаем файл построчно
	var scan = bufio.NewScanner(f)
	for scan.Scan() {
		var url = strings.TrimSpace(scan.Text())

		if ok, fetchErr := fetchLink(ctx, http.MethodGet, url); fetchErr != nil {
			return fetchErr
		} else {
			if ok {
				println("OK", url)
			} else {
				println("Not OK", url)
			}
		}
	}

	// проверяем сканер на наличие ошибок
	if err = scan.Err(); err != nil {
		return err
	}

	return nil
}

// объявляем HTTP клиент для переиспользования
var httpClient = http.Client{Timeout: time.Second * 5}

func fetchLink(ctx context.Context, method, url string) (bool, error) {
	// создаём объект запроса
	var req, err = http.NewRequestWithContext(ctx, method, url, http.NoBody)
	if err != nil {
		return false, err
	}

	// выполняем его
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}

	// валидируем статус код
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
