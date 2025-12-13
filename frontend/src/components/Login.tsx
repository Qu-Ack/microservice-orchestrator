import { Link, useNavigate } from 'react-router';
import { useState } from 'react';
import { API_URL } from "../constants"

export default function Login() {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [error, setError] = useState('');
	const navigate = useNavigate();

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		setError('');

		try {
			const response = await fetch(`${API_URL}/v1/user/login`, {
				method: "POST",
				credentials: "include",
				body: JSON.stringify({ email: email, password: password })
			})

			if (!response.ok) {
				const error = await response.json();
				setError(error.error);
				return;
			}

			const data = await response.json();
			console.log(data);
			navigate("/dashboard")

		} catch (err) {
			setError('something went wrong. please try again!');
			console.log(err);
		}
	};

	return (
		<div className="min-h-screen bg-[#c0c0c0] p-4">

			<table width="100%" cellPadding="0" cellSpacing="0" border={0} className="bg-[#000080] mb-4">
				<tbody>
					<tr>
						<td className="p-3">
							<Link to="/" className="text-white text-xl">⚓ KubePirate</Link>
						</td>
					</tr>
				</tbody>
			</table>

			<center>
				<table width="500" cellPadding="15" cellSpacing="0" border={2} className="bg-white border-black mt-12">
					<tbody>
						<tr className="bg-[#000080]">
							<td>
								<center>
									<span className="text-white text-2xl">🏴‍☠️ Login to KubePirate</span>
								</center>
							</td>
						</tr>
						<tr>
							<td>
								<form onSubmit={handleSubmit} className="p-3">
									<table width="100%" cellPadding="8" cellSpacing="0" border={0}>
										<tbody>
											<tr>
												<td colSpan={2}>
													<p className="mb-4 font-bold text-center">
														Enter your credentials to access the control panel.
													</p>
												</td>
											</tr>
											<tr>
												<td width="30%" align="right">
													<b>Email:</b>
												</td>
												<td>
													<input
														type="email"
														value={email}
														onChange={(e) => setEmail(e.target.value)}
														className="w-full border-2 border-black px-2 py-1"
														size={30}
														required
													/>
												</td>
											</tr>
											<tr>
												<td align="right">
													<b>Password:</b>
												</td>
												<td>
													<input
														type="password"
														value={password}
														onChange={(e) => setPassword(e.target.value)}
														className="w-full border-2 border-black px-2 py-1"
														size={30}
														required
													/>
												</td>
											</tr>
											<tr>
												<td colSpan={2} align="center" className="pt-4">
													<button
														type="submit"
														className="bg-[#008080] text-white border-2 border-black px-8 py-2 text-lg"
													>
														Login
													</button>
												</td>
											</tr>
											<tr>
												<td colSpan={2} align="center" className="pt-4">
													<small className="text-red-900 font-bold">
														{error}
													</small>
												</td>
											</tr>
											<tr>
												<td colSpan={2} align="center" className="pt-4">
													<hr className="my-2 border-gray-400" />
													<small>
														Don't have an account? <Link to="/signup" className="text-blue-600 underline">Sign up here</Link>
													</small>
												</td>
											</tr>
										</tbody>
									</table>
								</form>
							</td>
						</tr>
					</tbody>
				</table>

				<div className="mt-4">
					<Link to="/" className="text-blue-600 underline">← Back to Home</Link>
				</div>
			</center>
		</div>
	);
}

