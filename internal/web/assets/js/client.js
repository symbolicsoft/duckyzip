const postData = async (url, formData) => {
	const response = await fetch(url, {
		method: `POST`,
		cache: `no-cache`,
		referrerPolicy: `no-referrer`,
		body: formData
	});
	return response.json()
}

const linkFormSubmit = (cb) => {
	const payload = document.getElementById(`linkInput`).value.trim()
	const captchaID = document.getElementById(`captchaID`).value
	const captchaResponse = document.getElementById(`captchaResponse`).value
	const formData = new FormData()
	formData.append(`payload`, payload)
	formData.append(`captchaID`, captchaID)
	formData.append(`captchaResponse`, captchaResponse)
	postData(`/link`, formData).then(response => {
		cb(response)
	})
}

const linkHandleResponse = (response) => {
	switch (response.status) {
		case `OK`:
			const message = [
				`Short ID: https://ducky.zip/${response.shortID}`,
				``,
				`VRF Proof committed to <a href="https://optimistic.etherscan.io/address/0x082ff59678c0c5781f164c29c5a8f90008d5b1c0">smart contract</a>.`,
				`Committed key: <span class="mono">${response.vrfValue0}${response.vrfProof0}</span>`,
				`Committed value: <span class="mono">${response.vrfValue1}${response.vrfProof1}</span>`,
				``,
				`DuckyZip's VRF public keys are hard-coded into the smart contract`,
				`as values <strong>VRFPK0</strong> and <strong>VRFPK1</strong>.`,
				`We will soon make it easier to verify the above proof values.`
			].join('<br>')
			displayMessage(message, `good`)
			recycleCaptcha()
			document.getElementById(`linkInput`).value = ``
			document.getElementById(`linkInput`).focus()
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

document.getElementById(`linkButton`).addEventListener(`click`, () => {
	linkFormSubmit(linkHandleResponse)
})

document.getElementById(`captchaResponse`).addEventListener(`keypress`, (event) => {
	if (event.key === `Enter`) {
		document.getElementById(`linkButton`).click()
	}
})

recycleCaptcha()
document.getElementById(`linkInput`).focus()