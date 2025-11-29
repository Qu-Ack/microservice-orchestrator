import {MakeNav, Request} from "./common.mjs"


function DisplayLoginForm() {
	const login_form = document.createElement("form");

	const email_ipt = document.createElement("input");
	email_ipt.type = "text";
	email_ipt.id = "login_email_ipt"

	const password_ipt = document.createElement("input");
	password_ipt.type = "password";
	password_ipt.id = "login_password_ipt"


	const submit_btn = document.createElement("button");
	submit_btn.type = "password";
	submit_btn.textContent = "Submit";
	submit_btn.id = "login_submit_btn"

	login_form.appendChild(email_ipt);
	login_form.appendChild(password_ipt);
	login_form.appendChild(submit_btn);
	return login_form;
}


function DisplayRegisterForm() {
	const register_form = document.createElement("form");

	const email_ipt = document.createElement("input");
	email_ipt.type = "text";
	email_ipt.id = "register_email_ipt"

	const password_ipt = document.createElement("input");
	password_ipt.type = "password";
	password_ipt.id = "register_password_ipt"

	const cnf_password_ipt = document.createElement("input");
	cnf_password_ipt.type = "password";
	cnf_password_ipt.id = "register_cnf_password_ipt"

	const submit_btn = document.createElement("button");
	submit_btn.type = "password";
	submit_btn.textContent = "Submit";
	submit_btn.id = "register_submit_btn"

	register_form.appendChild(email_ipt);
	register_form.appendChild(password_ipt);
	register_form.appendChild(submit_btn);
	return register_form;
}


window.onload = () => {
	const app = document.querySelector("#app");
	app.appendChild(MakeNav());
	const login_form = app.appendChild(DisplayLoginForm());


	login_form.onsubmit = (e) => {
		e.preventDefault();
		console.log("form was submitted");
	}
}
