import { useAuth } from "@/contexts/AuthContext.tsx";
import { Link } from "react-router";
import { ModeToggle } from "./mode-toggle";
import { Button } from "@/components/ui/button";
import {
	NavigationMenu,
	NavigationMenuLink,
	NavigationMenuList,
} from "@/components/ui/navigation-menu";

export default function Nav() {
	const { isAuthenticated } = useAuth();

	if (!isAuthenticated) {
		return (
			<nav className="flex w-full justify-between items-center">
				<div>
					<Link
						to={'/'}
						className="font-bold transition text-xl hover:text-green-600"
					>
						Pirate
					</Link>
				</div>
				<NavigationMenu>
					<NavigationMenuList className="w-full flex items-center justify-between">
						<div className="flex items-center gap-4">
							<ModeToggle />
							<Button asChild variant="secondary">
								<Link to={'/signup'}>
									Sign Up
								</Link>
							</Button>
							<Button asChild>
								<Link to={'/login'}>
									Log In
								</Link>
							</Button>
						</div>
					</NavigationMenuList>
				</NavigationMenu>
			</nav>
		);
	} else {
		return (
			<nav className="flex w-full justify-between items-center">
				<div>
					<Link
						to={'/'}
						className="font-bold transition text-xl hover:text-green-600"
					>
						Pirate
					</Link>
				</div>
				<NavigationMenu>
					<NavigationMenuList className="w-full flex items-center justify-between">
						<div className="flex items-center gap-4">
							<ModeToggle />
							<NavigationMenuLink href={'/dashboard'}>
								Dashboard
							</NavigationMenuLink>
						</div>
					</NavigationMenuList>
				</NavigationMenu>
			</nav>
		);
	}
}
