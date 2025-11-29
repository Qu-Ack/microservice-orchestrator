import {MakeNav, Request, API_URL} from "./common.mjs"

function MakeNewForm() {
	const new_form = document.createElement('form');

	const repo_ipt = document.createElement('input');
	repo_ipt.id = "new_form_repo_ipt";
	repo_ipt.type = "text";
	repo_ipt.name = "clone_url"
	repo_ipt.placeholder = "repo";

	const branch_ipt = document.createElement('input');
	branch_ipt.id = "new_form_branch_ipt";
	branch_ipt.type = "text";
	branch_ipt.name = "branch"
	branch_ipt.placeholder = "branch";

	const subdomain_ipt = document.createElement('input');
	subdomain_ipt.id = "new_form_subdomain_ipt";
	subdomain_ipt.type = "text";
	subdomain_ipt.name = "subdomain"
	subdomain_ipt.placeholder = "subdomain";

	const radio_dc = document.createElement('input');
	radio_dc.type = "radio";
	radio_dc.name = "type";
	radio_dc.value = "docker_compose";
	radio_dc.id = "new_form_radio";

	const radio_dc_label = document.createElement('label');
	radio_dc_label.htmlFor = "radio_dc";
	radio_dc_label.textContent = "Docker Compose";

	const radio_df = document.createElement('input');
	radio_df.type = "radio";
	radio_df.name = "type";
	radio_df.value = "dockerfile";
	radio_df.id = "new_form_radio";
	radio_df.checked = true;    

	const radio_df_label = document.createElement('label');
	radio_df_label.htmlFor = "radio_df";
	radio_df_label.textContent = "Dockerfile";

	const submit_btn = document.createElement('button');
	submit_btn.id = "new_form_btn";
	submit_btn.type = "submit";
	submit_btn.textContent = "submit";

	new_form.appendChild(repo_ipt);
	new_form.appendChild(branch_ipt);
	new_form.appendChild(subdomain_ipt);

	new_form.appendChild(radio_dc);
	new_form.appendChild(radio_dc_label);

	new_form.appendChild(radio_df);
	new_form.appendChild(radio_df_label);

	new_form.appendChild(submit_btn);

	return new_form;
}



window.onload = () => {
	const app = document.querySelector("#app");

	app.appendChild(MakeNav());
	const newForm = app.appendChild(MakeNewForm());

	newForm.onsubmit = async (e) => {
		e.preventDefault();
		const data = new FormData(e.currentTarget);

		const clone_url = data.get("clone_url");
		const branch = data.get("branch");
		const subdomain = data.get("subdomain");
		const type = data.get("type");

		const body = {
			clone_url, 
			branch,
			subdomain,
			type,
		};

		try {
			const response = await Request(`${API_URL}/v1/deploy`, {method: "POST", credentials: "include"}, body) 
			console.log(response);
		} catch (err) {
			console.log(err)
		}
	}

}

