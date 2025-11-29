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

function logout() {
	deleteCookie("auth_id");
	window.location = "login.html";
}

export function MakeNav() {


	const nav = document.createElement('nav')

	const auth_id = getCookie("auth_id");

	if (auth_id != null) {
		const dashboard_btn = document.createElement('a');
		dashboard_btn.id = "nav_dash_link";
		dashboard_btn.href = "dashboard.html";
		dashboard_btn.textContent = "dashboard";

		const new_deployment_btn = document.createElement('a');
		new_deployment_btn.id = "nav_dash_link";
		new_deployment_btn.href = "new.html";
		new_deployment_btn.textContent = "new deployment";


		const logout_btn = document.createElement('button');
		logout_btn.id = "nav_logout_btn";
		logout_btn.textContent = "logout";
		logout_btn.onclick = logout;

		nav.appendChild(dashboard_btn);
		nav.appendChild(logout_btn);
		nav.appendChild(new_deployment_btn);
	} else {

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
	}	
	return nav
}



function deleteCookie(cookieName) {
    document.cookie = cookieName + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

export function getCookie(name) {
    var dc = document.cookie;
    var prefix = name + "=";
    var begin = dc.indexOf("; " + prefix);
    if (begin == -1) {
        begin = dc.indexOf(prefix);
        if (begin != 0) return null;
    }
    else
    {
        begin += 2;
        var end = document.cookie.indexOf(";", begin);
        if (end == -1) {
        end = dc.length;
        }
    }

    return decodeURI(dc.substring(begin + prefix.length, end));
} 
