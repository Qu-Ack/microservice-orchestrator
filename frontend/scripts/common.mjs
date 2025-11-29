export const API_URL = "http://localhost:8080"


export async function Request(url, options, data) {

	if (data) {
		options.body = JSON.stringify(data);
	}

	try {
		const response = await fetch(url, options);

		if (!response.ok) {
			const errorData = await response.json().catch(() => ({ message: 'Something went wrong' }));
			throw new Error(`HTTP error! status: ${response.status}, message: ${errorData.error || response.statusText}`);
		}

		return await response.json();
	} catch (err) {
		console.log(`error while making request: ${err}`)
		throw err;
	}
}

export function MakeNav() {
	const nav = document.createElement('nav')

	const login_btn = document.createElement('a')
	login_btn.id = "nav_login_link"
	login_btn.href = "login.html"
	login_btn.textContent = "Login"

	const register_btn = document.createElement('a')
	register_btn.id = "nav_register_link"
	register_btn.href = "register.html"
	register_btn.textContent = "Register"

	nav.appendChild(login_btn)
	nav.appendChild(register_btn)

	return nav
}

