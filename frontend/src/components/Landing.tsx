import {useEffect, useState} from "react"
import { Link } from "react-router";

function getCookie(name: string): string | null {
	const dc: string = document.cookie;
	const prefix: string = `${name}=`;

	let begin: number = dc.indexOf(`; ${prefix}`);

	if (begin === -1) {
		begin = dc.indexOf(prefix);
		if (begin !== 0) return null;
	} else {
		begin += 2;
	}

	let end: number = dc.indexOf(";", begin);
	if (end === -1) {
		end = dc.length;
	}

	return decodeURIComponent(
		dc.substring(begin + prefix.length, end)
	);
}


export default function Landing() {

	const [authKey, setAuthKey] = useState("")

	useEffect(() => {
		const auth_key = getCookie("auth_id");
		if (auth_key != null) {
			setAuthKey(auth_key);
		}
	}, [])

	return (
		<div className="min-h-screen bg-[#c0c0c0] p-4">
			<table
				width="100%"
				cellPadding="0"
				cellSpacing="0"
				border={0}
				className="bg-[#000080] mb-4"
			>
				<tbody>
					<tr>
						<td className="p-3">
							<div className="flex items-center justify-between">
								<span className="text-white text-xl">
									⚓ KubePirate - Microservice Orchestration
								</span>
								{authKey == "" ? <div className="flex gap-2">
									<Link to="/login">
										<button className="bg-[#c0c0c0] border-2 border-white px-4 py-1 text-sm">
											Login
										</button>
									</Link>
									<Link to="/signup">
										<button className="bg-[#c0c0c0] border-2 border-white px-4 py-1 text-sm">
											Sign Up
										</button>
									</Link>
								</div> : <div className="flex gap-2">
									<Link to="/dashboard">
										<button className="bg-[#c0c0c0] border-2 border-white px-4 py-1 text-sm">
											Dashboard
										</button>
									</Link>
								</div> }
							</div>
						</td>
					</tr>
				</tbody>
			</table>

			<table
				width="100%"
				cellPadding="10"
				cellSpacing="0"
				border={1}
				className="bg-white border-black mb-4"
			>
				<tbody>
					<tr>
						<td>
							<center>
								<h1 className="text-4xl mb-4 mt-4">
									🏴‍☠️ Welcome to KubePirate
								</h1>
								<p className="text-lg mb-6">
									<i>
										The Original Kubernetes Orchestration Tool
									</i>
								</p>
								<p className="mb-8 max-w-2xl">
									Manage your microservices like a captain
									commanding his fleet. Deploy, monitor, and
									scale your containers across the digital seas.
								</p>

								<table
									cellPadding="5"
									cellSpacing="0"
									border={0}
									className="mb-4"
								>
									<tbody>
										<tr>
											<td>
												<Link to="/signup">
													<button className="bg-[#008080] text-white border-2 border-black px-8 py-2 text-lg">
														Get Started Now
													</button>
												</Link>
											</td>
											<td width="20"></td>
											<td>
												<Link to="/services">
													<button className="bg-white border-2 border-black px-8 py-2 text-lg">
														View Demo
													</button>
												</Link>
											</td>
										</tr>
									</tbody>
								</table>
							</center>
						</td>
					</tr>
				</tbody>
			</table>

			<table
				width="100%"
				cellPadding="10"
				cellSpacing="0"
				border={1}
				className="bg-white border-black mb-4"
			>
				<tbody>
					<tr className="bg-[#000080]">
						<td colSpan={3}>
							<center>
								<span className="text-white text-xl">
									Features
								</span>
							</center>
						</td>
					</tr>
					<tr>
						<td
							width="33%"
							valign="top"
							className="border-r border-black"
						>
							<center>
								<h3 className="text-lg mb-3">
									⛵ Deploy Services
								</h3>
							</center>
							<p className="text-sm">
								Launch your containerized applications with
								ease. Scale replicas up or down based on demand.
							</p>
						</td>
						<td
							width="33%"
							valign="top"
							className="border-r border-black"
						>
							<center>
								<h3 className="text-lg mb-3">
									📊 Monitor Performance
								</h3>
							</center>
							<p className="text-sm">
								Track CPU, memory, and network usage in
								real-time. Get instant alerts when something
								goes wrong.
							</p>
						</td>
						<td width="33%" valign="top">
							<center>
								<h3 className="text-lg mb-3">🎯 Manage Pods</h3>
							</center>
							<p className="text-sm">
								View detailed statistics for each pod. Check
								logs, restart containers, and troubleshoot
								issues.
							</p>
						</td>
					</tr>
				</tbody>
			</table>

			{/* Stats Table */}
			<table
				width="100%"
				cellPadding="10"
				cellSpacing="0"
				border={1}
				className="bg-[#ffff99] border-black mb-4"
			>
				<tbody>
					<tr>
						<td>
							<center>
								<b>⚠️ System Status:</b> All systems operational
								|<b> Active Users:</b> 1,337 |
								<b> Services Running:</b> 12,845
							</center>
						</td>
					</tr>
				</tbody>
			</table>

			<table
				width="100%"
				cellPadding="5"
				cellSpacing="0"
				border={1}
				className="bg-[#d3d3d3] border-black"
			>
				<tbody>
					<tr>
						<td>
							<center>
								<small>
									© 2025 KubePirate |{" "}
									<a
										href="#"
										className="text-blue-600 underline"
									>
										About
									</a>{" "}
									|
									<a
										href="#"
										className="text-blue-600 underline"
									>
										{" "}
										Contact
									</a>{" "}
									|
									<a
										href="#"
										className="text-blue-600 underline"
									>
										{" "}
										Terms
									</a>
								</small>
							</center>
						</td>
					</tr>
					<tr>
						<td>
							<center>
								<small className="text-gray-600">
									Best viewed in Netscape Navigator 4.0 or
									higher
								</small>
							</center>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	);
}
