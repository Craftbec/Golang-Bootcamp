package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var str = []string{
	"http://google.com",
	"http://google.com",
	"http://google.com",
	"http://google.com",
	"http://google.com",
	"http://ya.ru",
	"http://ya.ru",
	"http://ya.ru",
	"http://ya.ru",
	"http://ya.ru",
	"http://google.com",
	"http://ya.ru",
	"http://google.com",
	"http://ya.ru",
	"https://edu.21-school.ru/",
	"https://go.dev/",
	"https://ru.wikipedia.org/wiki/Go",
	"https://github.com/golang/go",
	"https://tproger.ru/translations/golang-basics/",
	"https://dev-gang.ru/article/golang-osnovnoi-sintaksis-pua4fd0n18/",
	"https://dev-gang.ru/article/golang-making-http-requests-3/",
	"https://translate.yandex.ru/?source_lang=en&target_lang=ru",
	"https://repos.21-school.ru/users/sign_in",
	"https://dzen.ru/?yredirect=true&utm_referer=www.google.com",
	"https://edu.21-school.ru/",
	"https://go.dev/",
	"https://ru.wikipedia.org/wiki/Go",
	"https://github.com/golang/go",
	"https://tproger.ru/translations/golang-basics/",
	"https://dev-gang.ru/article/golang-osnovnoi-sintaksis-pua4fd0n18/",
	"https://dev-gang.ru/article/golang-making-http-requests-3/",
	"https://translate.yandex.ru/?source_lang=en&target_lang=ru",
	"https://repos.21-school.ru/users/sign_in",
	"https://dzen.ru/?yredirect=true&utm_referer=www.google.com",
	"https://edu.21-school.ru/",
	"https://go.dev/",
	"https://ru.wikipedia.org/wiki/Go",
	"https://github.com/golang/go",
	"https://tproger.ru/translations/golang-basics/",
	"https://dev-gang.ru/article/golang-osnovnoi-sintaksis-pua4fd0n18/",
	"https://dev-gang.ru/article/golang-making-http-requests-3/",
	"https://translate.yandex.ru/?source_lang=en&target_lang=ru",
	"https://repos.21-school.ru/users/sign_in",
	"https://dzen.ru/?yredirect=true&utm_referer=www.google.com",
}

func stop(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	sig := <-c
	switch sig {
	case os.Interrupt:
		return errors.New("ctrl+c")
	default:
		return nil
	}
}

func crawlWeb(ctx context.Context, c chan string) chan *string {
	res := make(chan *string)
	ctxx, cancel := context.WithCancel(ctx)
	chek := make(chan struct{}, 8)
	var wg sync.WaitGroup
	go func() {
		for {
			err := stop(ctxx)
			if err != nil {
				cancel()
			}
		}
	}()
	go func() {
		for val := range c {
			select {
			case <-ctxx.Done():
				go func() {
					for {
						wg.Add(-1)
						wg.Wait()
					}
				}()
				close(res)
				return
			default:
				wg.Add(1)
				chek <- struct{}{}
				go func(val string) {
					resp, err := http.Get(val)
					if err != nil {
						log.Fatalln(err)
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Fatalln(err)
					}
					str := string(body)
					res <- &str
					wg.Done()
					<-chek
				}(val)
			}
		}
		wg.Wait()
		close(res)
	}()
	return res
}

func main() {
	c := make(chan string)
	go func() {
		for i := 0; i < len(str); i++ {
			c <- str[i]
		}
		close(c)
	}()
	out := crawlWeb(context.Background(), c)
	for in := range out {
		fmt.Println((*in)[:300])
	}
}
