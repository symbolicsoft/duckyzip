<!---
# SPDX-FileCopyrightText: © 2019-2022 Nadim Kobeissi <nadim@symbolic.software>
# SPDX-License-Identifier: CC-BY-SA-4.0
-->

# DuckyZip

<img src="https://ducky.zip/assets/img/logo.png" alt="" align="left" height="256" style="margin:10px" />

URL shorteners are a common online service that allows the shortening of a long URL (often a Google Maps URL or similar) into a much shorter one, to use for example on social media or in QR codes. However, URL shorteners are free to behave dishonestly: they can, for instance, map a short URL into a long URL honestly for one party, while redirecting some other party into a different malicious long URL for the same short URL.

DuckyZip is the first provably honest URL shortening service which cannot selectively provide different "long URLs" to different parties undetected. DuckyZip uses a combination of Verifiable Random Function (VRF) constructions and a smart contract in order to provide a URL shortening service with strong security guarantees: despite the transparency of the smart contract log, observers cannot feasibly create a mapping of all short URLs to long URLs that is faster than classical enumeration.

[Read Paper](https://eprint.iacr.org/2023/1069) - [Try It](https://ducky.zip)

**This is a proof-of-concept implementation that is not yet fully ready and that is under active development.**

## What Works

- [x] Web Server
- [x] Shorten URLs
- [x] Database
- [x] Sanitization
- [X] Captcha
- [ ] Smart Contract
- [X] VRF
- [ ] Improve UX
- [ ] Mobile UX
- [ ] Verification via Third Party Code

Importantly, **the VRF and smart contract features are still not implemented!**