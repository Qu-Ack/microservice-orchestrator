import {MakeNav, Request, API_URL} from "./common.mjs"



async function getService(svc_name){
	try {
		const response = await Request(`${API_URL}/v1/service/${svc_name}`, {method: "GET", credentials: "include"})
		console.log(response);
	} catch (err) {
		console.log(err.ToString())
	}
}

window.onload = () => {

	let params = new URLSearchParams(document.location.search);

	const app = document.querySelector("#app");
	let service_name = params.get("svc_name");
	console.log(service_name);

	getService(service_name);
}

