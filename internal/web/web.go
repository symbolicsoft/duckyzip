/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package web

import (
	"log"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        2 * time.Minute,
	}))
	router.LoadHTMLGlob(filepath.Join("internal", "web", "assets", "html", "*"))
	router.GET("/", routeHome)
	router.GET("/assets/:category/:filename", routeAssets)
	router.GET("/captcha", routeCaptcha)
	router.POST("/shorten", routeShorten)
	router.GET("/:shortURL", routeLengthen)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
