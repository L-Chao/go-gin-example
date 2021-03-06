package cron

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/L-Chao/go-gin-example/models"
)

func CronMain() {
	log.Println("starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.DeleteAllArticle...")
		models.DeleteAllArticle()
	})

	c.Start()

	t1 := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-t1.C:
			t1.Reset(10 * time.Second)
		}
	}
}
