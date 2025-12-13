import ReactDom from "react-dom/client"
import './index.css'
import { createBrowserRouter } from "react-router"
import { RouterProvider } from "react-router/dom"
import Landing from "./components/Landing.tsx"
import Register from "./components/Register.tsx"
import Login from "./components/Login.tsx"
import Dashboard from "./components/Dashboard.tsx"
import Service from "./components/Service.tsx"

const router = createBrowserRouter([
	{
		path: "/",
		element: <Landing/>,
	},
	{
		path: "/health",
		element: <div> Healthy </div> 
	},
	{
		path: "/login",
		element: <Login/>
	},
	{
		path: "/signup",
		element: <Register/>
	},
	{
		path: "/dashboard",
		element: <Dashboard/>,
	},
	{
		path: "/service/:svc",
		element: <Service/>
	}
])

const root = document.getElementById('root')!;

ReactDom.createRoot(root).render(
	<RouterProvider router={router}/>
)
