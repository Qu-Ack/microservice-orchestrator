import {MakeNav, Request} from "./common.mjs"


function DisplayLoginForm() {
	const login_form = document.createElement("form");

	const email_ipt = document.createElement("input");
	email_ipt.type = "text";
	email_ipt.id = "login_email_ipt"
	email_ipt.name = "email"

	const password_ipt = document.createElement("input");
	password_ipt.type = "password";
	password_ipt.id = "login_password_ipt"
	password_ipt.name = "password"


	const submit_btn = document.createElement("button");
	submit_btn.type = "password";
	submit_btn.textContent = "Submit";
	submit_btn.id = "login_submit_btn"

	login_form.appendChild(email_ipt);
	login_form.appendChild(password_ipt);
	login_form.appendChild(submit_btn);
	return login_form;
}




window.onload = () => {
	const app = document.querySelector("#app");
	app.appendChild(MakeNav());
	const login_form = app.appendChild(DisplayLoginForm());


	login_form.onsubmit = async (e) => {
		e.preventDefault();
		const data = new FormData(e.currentTarget);

		const email = data.get("email");
		const password = data.get("password");

		try {
			const response = await Request("http://localhost:8080/v1/user/login", {method: "POST", credentials: "include"}, {email: email, password: password});
			console.log(response)
		} catch (err) {
			console.log(err.toString());
		}
	}
}
