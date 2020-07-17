package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/toms1441/urlsh/internal/config"
	"github.com/toms1441/urlsh/internal/repo/json"
	"github.com/toms1441/urlsh/internal/shortener"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("invalid config: %v", err)
	}

	srepo, err := json.NewShortenerRepository("./db")
	if err != nil {
		log.Fatalf("couldn't create json database: %v", err)
	}

	sserv, err := shortener.NewService(srepo, conf.Shortener)
	if err != nil {
		log.Fatalf("couldn't create new service, either config is invalid or database is invalid: %v", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*.html")

	r.GET("/redirect/:id", func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == conf.Shortener.Length {
			urlstring, err := sserv.GetShortener(id)
			if err != nil {
				c.JSON(404, "shortener record does not exist")
				return
			}

			c.Redirect(307, urlstring)
		} else {
			c.JSON(400, "length is invalid")
			return
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/post", func(c *gin.Context) {
		urlstring := c.PostForm("url")

		if len(urlstring) > 0 {
			model, err := sserv.NewShortener(urlstring)
			if err == nil {
				c.String(200, model.ID)
				return
			}
		}

		c.String(400, "invalid value")
		return
	})

	r.Static("/static/", "./web/assets/")

	r.Run(":8080")

}
