import { Button } from '@/components/ui/button';
import {
	Carousel,
	CarouselContent,
	CarouselItem,
	CarouselNext,
	CarouselPrevious,
} from '@/components/ui/carousel';

function App() {
	const steps = [
		{
			number: 1,
			title: "Add Docker Configuration",
			description: "Include a Dockerfile or docker-compose.yml in your project root.",
			icon: "🗺️"
		},
		{
			number: 2,
			title: "Push to GitHub",
			description: "Upload your code to a GitHub repository.",
			icon: "⚓"
		},
		{
			number: 3,
			title: "Deploy Your App",
			description: "Paste your repo link, choose a subdomain, and go live.",
			icon: "🏴‍☠️"
		}
	];

	return (
		<div className="flex mt-32 justify-center min-h-screen p-4">
			<div className="w-full max-w-2xl">
				<div className="text-center mb-12">
					<h1 className="text-5xl font-bold mb-4">
						The Last Hosting Platform You'll Ever Need
					</h1>
					<p className="text-xl text-gray-600 dark:text-gray-400 mb-2">
						Deploy your containerized apps in minutes, not hours.
					</p>
					<p className="text-lg text-gray-500 dark:text-gray-500">
						No complex configs. No DevOps headaches. Just pure deployment freedom.
					</p>
				</div>

				<h2 className="text-3xl font-bold text-center mb-8">
					Deploy in 3 Simple Steps
				</h2>

				<Carousel className="w-full">
					<CarouselContent>
						{steps.map((step) => (
							<CarouselItem key={step.number}>
								<div className="p-8 border rounded-lg">
									<div className="text-center">
										<div className="text-6xl mb-4">{step.icon}</div>
										<div className="text-sm font-semibold text-gray-500 mb-2">
											STEP {step.number}
										</div>
										<h3 className="text-2xl font-bold mb-4">
											{step.title}
										</h3>
										<p className="text-gray-600">
											{step.description}
										</p>
									</div>
								</div>
							</CarouselItem>
						))}
					</CarouselContent>
					<CarouselPrevious />
					<CarouselNext />
				</Carousel>

				<div className="text-center mt-8">
					<Button size="lg" variant="default" className="hover:cursor-pointer">
						Get Started
					</Button>
				</div>
			</div>
		</div>
	);
}

export default App;
