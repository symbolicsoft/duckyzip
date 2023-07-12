/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package web

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"ducky.zip/m/v2/internal/sanitize"
	"ducky.zip/m/v2/internal/shorten"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func routeHome(c *gin.Context) {
	c.Request.Close = true
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func routeAssets(c *gin.Context) {
	c.Request.Close = true
	category := c.Param("category")
	if !sanitize.CheckAssetCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Asset File Category",
		})
		return
	}
	fileName := c.Param("filename")
	if len(fileName) > 32 || !sanitize.CheckAssetName(fileName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Asset File Name",
		})
		return
	}
	assetsPath := filepath.Join("internal", "web", "assets")
	filePath := filepath.Join(assetsPath, category, fileName)
	_, err := os.Stat(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Asset File Name",
		})
		return
	}
	c.File(filePath)
}

func routeCaptcha(c *gin.Context) {
	c.Request.Close = true
	time.Sleep(time.Second * 1)
	captchaID := captcha.New()
	var captchaImgBuffer bytes.Buffer
	err := captcha.WriteImage(&captchaImgBuffer, captchaID, 300, 100)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Captcha Generation Failed",
		})
		return
	}
	captchaImgB64 := base64.StdEncoding.EncodeToString(captchaImgBuffer.Bytes())
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"captchaID":  captchaID,
		"captchaImg": captchaImgB64,
	})
}

func routeShorten(c *gin.Context) {
	c.Request.Close = true
	time.Sleep(time.Second * 1)
	longURL := c.PostForm("longURL")
	captchaID := c.PostForm("captchaID")
	captchaResponse := c.PostForm("captchaResponse")
	if !captcha.VerifyString(captchaID, captchaResponse) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Captcha",
		})
		return
	}
	if !sanitize.CheckLongURL(longURL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid URL",
		})
		return
	}
	shortURL, err := shorten.GenShortURL(longURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": shortURL,
	})
}

func routeLengthen(c *gin.Context) {
	c.Request.Close = true
	time.Sleep(time.Second * 1)
	shortURL := c.Param("shortURL")
	if !sanitize.CheckShortURL(shortURL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid URL",
		})
		return
	}
	longURL, err := shorten.GetLongURL(shortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"delay":   0,
		"longURL": longURL,
	})
}
