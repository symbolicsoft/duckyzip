const getData = async (url) => {
	const response = await fetch(url, {
		method: `GET`,
		cache: `no-cache`,
		referrerPolicy: `no-referrer`
	});
	return response.json()
}

const verifyVRF = async () => {
	let shortID = document.getElementById(`shortID`).value
	const infoResponse = await getData(`https://ducky.zip/info/${shortID}`)
	if (infoResponse.status !== `OK`) {
		return { vrfResult: false, shortID: ``, payload: `` }
	}
	let payload = infoResponse.payload
	let vrfValue0 = infoResponse.vrfValue0
	let vrfProof0 = infoResponse.vrfProof0
	if (!/^([a-f0-9]{256})$/.test(`${vrfValue0}${vrfProof0}`)) {
		return { vrfResult: false, shortID: ``, payload: `` }
	}
	const contractResponse = await getData(`https://ducky.zip/contract/${vrfValue0}/${vrfProof0}`)
	if (contractResponse.status !== `OK`) {
		return { vrfResult: false, shortID: ``, payload: `` }
	}
	let vrfValue1 = contractResponse.vrfValue1
	let vrfProof1 = contractResponse.vrfProof1
	if (!/^([a-f0-9]{256})$/.test(`${vrfValue1}${vrfProof1}`)) {
		return { vrfResult: false, shortID: ``, payload: `` }
	}
	let vrfResult0 = VerifyShortIDProof(shortID, vrfValue0, vrfProof0)
	let vrfResult1 = VerifyPayloadProof(payload, vrfValue1, vrfProof1)
	let vrfResult = vrfResult0 && vrfResult1
	return { vrfResult, shortID, payload }
}

const go = new Go()
WebAssembly.instantiateStreaming(fetch(`vrf.wasm`), go.importObject).then((result) => {
	go.run(result.instance)
})

document.getElementById(`shortIDButton`).addEventListener(`click`, () => {
	verifyVRF().then(result => {
		if (result.vrfResult) {
			document.getElementById(`message`).innerText = `VRF proof verified on smart contract for short ID ${result.shortID} and payload ${result.payload}.`
		} else {
			document.getElementById(`message`).innerText = `VRF proof verification failed.`
		}
	})
})