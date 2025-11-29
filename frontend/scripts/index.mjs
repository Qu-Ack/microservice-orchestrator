import {MakeNav} from "./common.mjs"

window.onload = function() {
	const app = document.querySelector("#app");
	const nav = MakeNav();

	app.appendChild(nav);
}



