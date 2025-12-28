import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "@/components/ui/select";
import { toast } from "sonner";
import { useNavigate } from "react-router";
import { Loader2 } from "lucide-react";

export default function NewDeployment() {
	const [cloneUrl, setCloneUrl] = useState("");
	const [branch, setBranch] = useState("main");
	const [subdomain, setSubdomain] = useState("");
	const [type, setType] = useState("");
	const [isLoading, setIsLoading] = useState(false);
	const navigate = useNavigate();

	const handleSubmit = async () => {
		if (!cloneUrl || !branch || !subdomain || !type) {
			toast.error("Please fill in all fields");
			return;
		}

		setIsLoading(true);

		try {
			const resp = await fetch(`${import.meta.env.VITE_APP_API_URL}/v1/deploy`, {
				method: "POST",
				credentials: "include",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					clone_url: cloneUrl,
					branch: branch,
					subdomain: subdomain,
					type: type,
				}),
			});

			const data = await resp.json();

			if (!resp.ok) {
				toast.error(data.message || "Deployment failed");
				return;
			}

			toast.success("Deployment created successfully!");
			navigate("/dashboard/deployments");
		} catch (err) {
			console.error("Deployment error:", err);
			toast.error("Can't connect to the server");
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<div className="p-6 flex flex-col items-center mt-32">
			<div className="max-w-2xl w-full space-y-6">
				<div className="text-center space-y-3">
					<h1 className="text-4xl font-bold tracking-tight">Deploy in Seconds</h1>
					<p className="text-xl text-muted-foreground">
						From code to production in minutes. Connect your Git repository and let us handle the rest.
					</p>
					<div className="flex justify-center gap-8 pt-2">
						<div className="text-center">
							<div className="text-2xl font-bold">⚡</div>
							<div className="text-sm text-muted-foreground">Lightning Fast</div>
						</div>
						<div className="text-center">
							<div className="text-2xl font-bold">🔒</div>
							<div className="text-sm text-muted-foreground">Secure by Default</div>
						</div>
						<div className="text-center">
							<div className="text-2xl font-bold">🚀</div>
							<div className="text-sm text-muted-foreground">Auto Scaling</div>
						</div>
					</div>
				</div>

				<Card>
					<CardHeader>
						<CardTitle>Create New Deployment</CardTitle>
						<CardDescription>
							Deploy your application from a Git repository
						</CardDescription>
					</CardHeader>
					<CardContent>
						<div className="space-y-6">
							<div className="space-y-2">
								<Label htmlFor="cloneUrl">Git Clone URL</Label>
								<Input
									id="cloneUrl"
									type="url"
									placeholder="https://github.com/username/repo.git"
									value={cloneUrl}
									onChange={(e) => setCloneUrl(e.target.value)}
									disabled={isLoading}
								/>
								<p className="text-sm text-muted-foreground">
									The URL to clone your Git repository
								</p>
							</div>

							<div className="space-y-2">
								<Label htmlFor="branch">Branch</Label>
								<Input
									id="branch"
									type="text"
									placeholder="main"
									value={branch}
									onChange={(e) => setBranch(e.target.value)}
									disabled={isLoading}
								/>
								<p className="text-sm text-muted-foreground">
									The branch to deploy from
								</p>
							</div>

							<div className="space-y-2">
								<Label htmlFor="subdomain">Subdomain</Label>
								<Input
									id="subdomain"
									type="text"
									placeholder="my-app"
									value={subdomain}
									onChange={(e) => setSubdomain(e.target.value)}
									disabled={isLoading}
								/>
								<p className="text-sm text-muted-foreground">
									Your app will be available at: {subdomain || "my-app"}.yourdomain.com
								</p>
							</div>

							<div className="space-y-2">
								<Label htmlFor="type">Deployment Type</Label>
								<Select value={type} onValueChange={setType} disabled={isLoading}>
									<SelectTrigger id="type">
										<SelectValue placeholder="Select deployment type" />
									</SelectTrigger>
									<SelectContent>
										<SelectItem value="dockerfile">Dockerfile</SelectItem>
										<SelectItem value="docker_compose">Docker Compose</SelectItem>
									</SelectContent>
								</Select>
								<p className="text-sm text-muted-foreground">
									Choose how your application should be deployed
								</p>
							</div>

							<Button onClick={handleSubmit} className="w-full" disabled={isLoading}>
								{isLoading ? (
									<>
										<Loader2 className="mr-2 h-4 w-4 animate-spin" />
										Creating Deployment...
									</>
								) : (
									"Create Deployment"
								)}
							</Button>
						</div>
					</CardContent>
				</Card>
			</div>
		</div>
	);
}
