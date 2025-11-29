import {MakeNav, Request, API_URL} from "./common.mjs"


function DisplayRegisterForm() {
	const register_form = document.createElement("form");

	const email_ipt = document.createElement("input");
	email_ipt.type = "text";
	email_ipt.id = "register_email_ipt"
	email_ipt.name = "email"

	const password_ipt = document.createElement("input");
	password_ipt.type = "password";
	password_ipt.id = "register_password_ipt"
	password_ipt.name = "password"


	const cnf_password_ipt = document.createElement("input");
	cnf_password_ipt.type = "password";
	cnf_password_ipt.id = "register_cnf_password_ipt"
	cnf_password_ipt.name = "confirm_password"

	const submit_btn = document.createElement("button");
	submit_btn.type = "password";
	submit_btn.textContent = "Submit";
	submit_btn.id = "register_submit_btn"

	register_form.appendChild(email_ipt);
	register_form.appendChild(password_ipt);
	register_form.appendChild(cnf_password_ipt);
	register_form.appendChild(submit_btn);
	return register_form;
}


window.onload = () => {
	const app = document.querySelector("#app");
	app.appendChild(MakeNav());
	const register_form = app.appendChild(DisplayRegisterForm());


	register_form.onsubmit = async (e) => {
		e.preventDefault();

		const data = new FormData(e.currentTarget);

		const email = data.get("email");
		const password = data.get("password");
		const confirm_password = data.get("confirm_password");
	
		if (password != confirm_password) {
			console.log("cnf pwd and pwd don't match")
			return;
		}

		try {
			const response = await Request(`${API_URL}/v1/user/register`, 
				{method: "POST", credentials: "include"},
				{email: email, password: password},
			)
			console.log(response);
		} catch (err) {
			console.log(err.toString());
		}

	}
}
