import {Request, MakeNav, API_URL} from "./common.mjs"

async function getServices() { 
	try {
		const response = await Request(`${API_URL}/v1/deploy`, {method: "GET", credentials: "include"});
		return response;
	} catch (err) {
		console.log(err)
	}
}

function MakeServicesList(services) {

	if (services.length == 0) {

		const no_svc = document.createElement('div');
		no_svc.id = "dash_svc_list_div";
		return no_svc;
	} else {
		
		const service_list = document.createElement('ul');
		service_list.id = "dash_svc_list";

		services.map(service => {
			const svc = document.createElement('li');
			svc.id = "dash_svc_li";

			const svc_link = document.createElement('a');


			svc_link.id = "dash_svc_link";
			svc_link.href = `service.html?svc_name=${encodeURIComponent(service)}`;
			svc_link.textContent = service;

			svc.appendChild(svc_link);

			service_list.appendChild(svc);
		})

		return service_list;
	}


}


window.onload = async () => {
	const app = document.querySelector("#app");

	app.appendChild(MakeNav());

	const svcs = await getServices();
	console.log(`svcs are : ${svcs.services}`);
	app.appendChild(MakeServicesList(svcs.services));
}





