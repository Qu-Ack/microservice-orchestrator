import ReactDom from "react-dom/client"
import './index.css'
import App from './App.tsx'
import { createBrowserRouter } from "react-router"
import { RouterProvider } from "react-router/dom"

const router = createBrowserRouter([
	{
		path: "/",
		element: <App/>,
	},
	{
		path: "/health",
		element: <div> Healthy </div> 
	}
])

const root = document.getElementById('root')!;

ReactDom.createRoot(root).render(
	<RouterProvider router={router}/>
)
