import ReactDOM from "react-dom/client";
import { createBrowserRouter, Outlet } from "react-router";
import { AuthContextProvider } from "./contexts/AuthContext"
import { ThemeProvider } from "@/contexts/ThemeProvider"
import { RouterProvider } from "react-router/dom";
import App from "./App"
import Nav from "./components/Nav";
import "./index.css"


function NavLayout() {
	return (
		<ThemeProvider>
			<div className="h-screen w-screen flex flex-col pl-32 pr-32 pt-12">
				<AuthContextProvider>
					<Nav></Nav>
					<main>
						<Outlet />
					</main>
				</AuthContextProvider>
			</div>
		</ThemeProvider>
	)
}


const router = createBrowserRouter([
	{
		path: "/",
		element: <NavLayout />,
		children: [
			{ index: true, element: <App /> }
		],
	},
]);

const root = document.getElementById("root");

if (!root) throw Error('no root element found bitch');

ReactDOM.createRoot(root).render(
	<RouterProvider router={router} />,
);

