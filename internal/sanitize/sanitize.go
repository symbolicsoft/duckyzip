/* SPDX-FileCopyrightText: © 2023 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: GPL-3.0-only */

package sanitize

import "regexp"

var regexpMap = map[string]*regexp.Regexp{
	"shortURL":      regexp.MustCompile(`^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{8,13}$`),
	"longURL":       regexp.MustCompile(`^([a-z0-9]{3,10}):\/\/[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`),
	"assetCategory": regexp.MustCompile(`^(css|img|js)$`),
	"assetName":     regexp.MustCompile(`^[a-zA-Z0-9]{1,32}\.(css|js|png)$`),
}

func CheckShortURL(shortURL string) bool {
	return regexpMap["shortURL"].MatchString(shortURL)
}

func CheckLongURL(longURL string) bool {
	return regexpMap["longURL"].MatchString(longURL)
}

func CheckAssetCategory(category string) bool {
	return regexpMap["assetCategory"].MatchString(category)
}

func CheckAssetName(name string) bool {
	return regexpMap["assetName"].MatchString(name)
}
