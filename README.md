<!---
# SPDX-FileCopyrightText: © 2019-2022 Nadim Kobeissi <nadim@symbolic.software>
# SPDX-License-Identifier: CC-BY-SA-4.0
-->

# DuckyZip

<img src="https://ducky.zip/assets/img/logo.png" alt="" align="left" height="256" style="margin:10px" />

DuckyZip is a provably honest global linking service which links short memorable identifiers to arbitrarily large payloads (URLs, text, documents, archives, etc.) without being able to undetectably provide different payloads for the same short identifier to different parties. DuckyZip uses a combination of Verifiable Random Function (VRF)-based zero knowledge proofs and a smart contract in order to provide strong security guarantees: despite the transparency of the smart contract log, observers cannot feasibly create a mapping of all short identifiers to payloads that is faster than O(n) classical enumeration.

[Read Paper](https://eprint.iacr.org/2023/1069) - [Try It](https://ducky.zip)

**This is a proof-of-concept implementation that is under active development.**