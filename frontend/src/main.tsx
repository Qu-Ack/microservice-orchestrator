import ReactDOM from "react-dom/client";
import { createBrowserRouter, Outlet } from "react-router";
import { AuthContextProvider } from "./contexts/AuthContext"
import { ThemeProvider } from "@/contexts/ThemeProvider"
import { RouterProvider } from "react-router/dom";
import App from "./App"
import { Toaster } from "./components/ui/sonner";
import Login from "./components/Login";
import NewDeployment from "./components/NewDeployment"
import Dashboard from "./components/Dashboard";
import Signup from "./components/Signup";
import Nav from "./components/Nav";
import "./index.css"

function NavLayout() {
	return (
		<div className="h-screen w-screen flex flex-col pl-32 pr-32 pt-12">
			<Nav></Nav>
			<main className="w-full h-full">
				<Outlet />
			</main>
		</div>
	)
}

function RootLayout() {
	return (
		<ThemeProvider>
			<AuthContextProvider>
				<Outlet />
				<Toaster />
			</AuthContextProvider>
		</ThemeProvider>
	)
}

const router = createBrowserRouter([
	{
		path: "/",
		element: <RootLayout />,
		children: [
			{
				path: "/",
				element: <NavLayout />,
				children: [
					{ index: true, element: <App /> },
					{ path: "/login", element: <Login /> },
					{ path: "/signup", element: <Signup /> },
				],
			},
			{
				path: "/dashboard",
				element: <Dashboard />,
				children: [
					{ index: true, element: <NewDeployment /> },
					{ path: "deployments", element: <div>Deployments </div> },
					{ path: "profile", element: <div> Profile</div> },
				]
			}
		],
	},
]);

const root = document.getElementById("root");
if (!root) throw Error('no root element found bitch');
ReactDOM.createRoot(root).render(
	<RouterProvider router={router} />,
);
