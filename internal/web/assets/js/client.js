const postData = async (url, formData) => {
	const response = await fetch(url, {
		method: `POST`,
		cache: `no-cache`,
		referrerPolicy: `no-referrer`,
		body: formData
	});
	return response.json()
}

const shortenURLFormSubmit = (cb) => {
	const longURL = document.getElementById(`shortenInput`).value.trim()
	const captchaID = document.getElementById(`captchaID`).value
	const captchaResponse = document.getElementById(`captchaResponse`).value
	const formData = new FormData()
	formData.append(`longURL`, longURL)
	formData.append(`captchaID`, captchaID)
	formData.append(`captchaResponse`, captchaResponse)
	postData(`/shorten`, formData).then(response => {
		cb(response)
	})
}

const shortenURLHandleResponse = (response) => {
	switch (response.status) {
		case `OK`:
			const message = [
				`Short URL: https://ducky.zip/${response.shortURL}`,
				``,
				`VRF Proof committed to <a href="https://optimistic.etherscan.io/address/0x082ff59678c0c5781f164c29c5a8f90008d5b1c0">smart contract</a>.`,
				`Committed key: <span class="mono">${response.vrfValue0}${response.vrfProof0}</span>`,
				`Committed value: <span class="mono">${response.vrfValue1}${response.vrfProof1}</span>`,
				``,
				`DuckyZip's VRF public keys are hard-coded into the smart contract`,
				`as values <strong>VRFPK0</strong> and <strong>VRFPK1</strong>.` 
			].join('<br />')
			displayMessage(message, `good`)
			recycleCaptcha()
			document.getElementById(`shortenInput`).value = ``
			document.getElementById(`shortenInput`).focus()
			break
		default:
			displayMessage(`Error: ${response.message}`, `bad`)
			recycleCaptcha()
			break
	}
}

const displayMessage = (message, kind) => {
	const messageArea = document.getElementById(`messageArea`)
	switch (kind) {
		case `good`:
			messageArea.innerHTML = message
			messageArea.classList.remove(`bad`)
			messageArea.classList.add(`good`)
			break
		case `bad`:
			messageArea.innerText = message
			messageArea.classList.remove(`good`)
			messageArea.classList.add(`bad`)
			break
	}
}

const getNewCaptcha = async () => {
	const response = await fetch(`/captcha`, {
		method: `GET`,
		cache: `no-cache`,
		referrerPolicy: `no-referrer`,
	});
	return response.json()
}

const recycleCaptcha = () => {
	getNewCaptcha().then(response => {
		document.getElementById(`captchaSection`).style.opacity = 1
		switch (response.status) {
			case `OK`:
				document.getElementById(`captchaID`).value = response.captchaID
				document.getElementById(`captchaImg`).src = `data:image/jpeg;base64,${response.captchaImg}`
				document.getElementById(`captchaResponse`).value = ``
				break
			case `ERR`:
				displayMessage(response.message, `bad`)
				break
		}
	})	
}

document.getElementById(`shortenButton`).addEventListener(`click`, () => {
	shortenURLFormSubmit(shortenURLHandleResponse)
})

document.getElementById(`captchaResponse`).addEventListener(`keypress`, (event) => {
	if (event.key === `Enter`) {
		document.getElementById(`shortenButton`).click()
	}
})

recycleCaptcha()
document.getElementById(`shortenInput`).focus()