/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package web

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ducky.zip/m/v2/internal/contract"
	"ducky.zip/m/v2/internal/link"
	"ducky.zip/m/v2/internal/sanitize"
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

func routeLink(c *gin.Context) {
	c.Request.Close = true
	time.Sleep(time.Second * 1)
	payload := c.PostForm("payload")
	captchaID := c.PostForm("captchaID")
	captchaResponse := c.PostForm("captchaResponse")
	if !captcha.VerifyString(captchaID, captchaResponse) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Captcha",
		})
		return
	}
	if len(payload) > (1024 * 10) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Payload Too Large",
		})
		return
	}
	if !sanitize.CheckPayload(payload) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Payload",
		})
		return
	}
	shortID, dbEntry, err := link.GenShortID(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"shortID":   shortID,
		"vrfValue0": dbEntry.VRFValue0,
		"vrfProof0": dbEntry.VRFProof0,
		"vrfValue1": dbEntry.VRFValue1,
		"vrfProof1": dbEntry.VRFProof1,
	})
}

func routePayload(c *gin.Context) {
	c.Request.Close = true
	time.Sleep(time.Second * 1)
	shortID := c.Param("shortID")
	if !sanitize.CheckShortID(shortID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Short ID",
		})
		return
	}
	dbEntry, err := link.GetPayload(shortID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "redirect.html", gin.H{
		"delay":     10,
		"shortID":   shortID,
		"payload":   dbEntry.Payload,
		"vrfValue0": dbEntry.VRFValue0,
		"vrfProof0": dbEntry.VRFProof0,
		"vrfValue1": dbEntry.VRFValue1,
		"vrfProof1": dbEntry.VRFProof1,
	})
}

func routeInfo(c *gin.Context) {
	c.Request.Close = true
	shortID := c.Param("shortID")
	if !sanitize.CheckShortID(shortID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Short ID",
		})
		return
	}
	dbEntry, err := link.GetPayload(shortID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shortID":   shortID,
		"payload":   dbEntry.Payload,
		"vrfValue0": dbEntry.VRFValue0,
		"vrfProof0": dbEntry.VRFProof0,
		"vrfValue1": dbEntry.VRFValue1,
		"vrfProof1": dbEntry.VRFProof1,
	})
}

func routeContract(c *gin.Context) {
	c.Request.Close = true
	time.Sleep(time.Second * 1)
	shortID := c.Param("shortID")
	if !sanitize.CheckShortID(shortID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Short ID",
		})
		return
	}
	dbEntry, err := link.GetPayload(shortID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	contractValue, err := contract.Read(strings.Join([]string{
		dbEntry.VRFValue0, dbEntry.VRFProof0,
	}, ""))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": err.Error(),
		})
		return
	}
	if len(contractValue) != 256 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": "Invalid Contract Value",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shortID":   shortID,
		"payload":   dbEntry.Payload,
		"vrfValue0": dbEntry.VRFValue0,
		"vrfProof0": dbEntry.VRFProof0,
		"vrfValue1": contractValue[0:64],
		"vrfProof1": contractValue[64:256],
	})
}
