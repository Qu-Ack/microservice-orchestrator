import {Request, MakeNav, API_URL} from "./common.mjs"

window.onload = () => {
	const app = document.querySelector("#app");

	app.appendChild(MakeNav());


	async function getServices() { 
		try {
			const response = await Request(`${API_URL}/v1/deploy`, {method: "GET", credentials: "include"});
			console.log(response);
		} catch (err) {
			console.log(err)
		}
	}

	getServices();
}





