import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Link, useNavigate } from "react-router";
import { toast } from "sonner"

export default function Signup() {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [confirmPassword, setConfirmPassword] = useState("");
	const [isLoading, setIsLoading] = useState(false);
	const navigate = useNavigate();

	const handleSubmit = async () => {
		if (password !== confirmPassword) {
			toast.error("confirm password and password don't match");
			return;
		}

		setIsLoading(true);

		try {
			const resp = await fetch(`${import.meta.env.VITE_APP_API_URL}/v1/user/register`, {
				credentials: "include",
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					email: email,
					password: password
				}),
			});

			const data = await resp.json();

			if (!resp.ok) {
				toast.error(data.message || "Something went wrong. Please try again.")
				return;
			}

			toast("You account has been created");
			navigate("/login");

		} catch (err) {
			console.error(err);
			toast.error("Can't connect to the server.");
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<div className="flex justify-center p-4 mt-64">
			<Card className="w-full max-w-md">
				<CardHeader>
					<CardTitle className="text-2xl">Sign Up</CardTitle>
					<CardDescription>
						Create an account to start deploying your apps
					</CardDescription>
				</CardHeader>
				<CardContent className="space-y-4">
					<div className="space-y-2">
						<Label htmlFor="email">Email</Label>
						<Input
							id="email"
							type="email"
							placeholder="you@example.com"
							value={email}
							onChange={(e) => setEmail(e.target.value)}
							disabled={isLoading}
						/>
					</div>
					<div className="space-y-2">
						<Label htmlFor="password">Password</Label>
						<Input
							id="password"
							type="password"
							placeholder="••••••••"
							value={password}
							onChange={(e) => setPassword(e.target.value)}
							disabled={isLoading}
						/>
					</div>
					<div className="space-y-2">
						<Label htmlFor="confirmPassword">Confirm Password</Label>
						<Input
							id="confirmPassword"
							type="password"
							placeholder="••••••••"
							value={confirmPassword}
							onChange={(e) => setConfirmPassword(e.target.value)}
							disabled={isLoading}
						/>
					</div>
				</CardContent>
				<CardFooter className="flex flex-col space-y-4">
					<Button className="w-full" onClick={handleSubmit} disabled={isLoading}>
						{isLoading ? "Creating account..." : "Sign Up"}
					</Button>
					<p className="text-sm text-center text-muted-foreground">
						Already have an account?{" "}
						<Link to="/login" className="underline hover:text-primary">
							Log in
						</Link>
					</p>
				</CardFooter>
			</Card>
		</div>
	);
}
