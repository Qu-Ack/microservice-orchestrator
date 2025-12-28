import { Link, Outlet, useLocation, useNavigate } from "react-router";
import { Home, Package, User, LogOut } from "lucide-react";
import { useAuth } from "@/contexts/AuthContext";
import { toast } from "sonner";
import {
	Sidebar,
	SidebarContent,
	SidebarGroup,
	SidebarGroupContent,
	SidebarGroupLabel,
	SidebarMenu,
	SidebarMenuButton,
	SidebarMenuItem,
	SidebarProvider,
	SidebarFooter,
} from "@/components/ui/sidebar";

export default function Dashboard() {
	const location = useLocation();
	const navigate = useNavigate();
	const { setAuthState } = useAuth();

	const navItems = [
		{ path: "/dashboard", icon: Home, label: "New Deployment" },
		{ path: "/dashboard/deployments", icon: Package, label: "Deployments" },
		{ path: "/dashboard/profile", icon: User, label: "Profile" },
	];

	const handleLogout = async () => {
		try {
			const resp = await fetch(`${import.meta.env.VITE_APP_API_URL}/v1/user/logout`, {
				method: "POST",
				credentials: "include",
			});

			if (resp.ok) {
				setAuthState(false);
				document.cookie = 'auth_id=; Max-Age=0; path=/;';
				toast.success("Successfully logged out");
				navigate("/");
			} else {
				toast.error("Logout failed");
			}
		} catch (err) {
			console.error("Logout error:", err);
			toast.error("Can't connect to the server");
		}
	};

	return (
		<SidebarProvider>
			<div className="flex w-full h-full">
				<Sidebar>
					<SidebarContent>
						<SidebarGroup>
							<SidebarGroupLabel>Dashboard</SidebarGroupLabel>
							<SidebarGroupContent>
								<SidebarMenu>
									{navItems.map((item) => {
										const Icon = item.icon;
										const isActive = location.pathname === item.path;
										return (
											<SidebarMenuItem key={item.path}>
												<SidebarMenuButton asChild isActive={isActive}>
													<Link to={item.path}>
														<Icon />
														<span>{item.label}</span>
													</Link>
												</SidebarMenuButton>
											</SidebarMenuItem>
										);
									})}
								</SidebarMenu>
							</SidebarGroupContent>
						</SidebarGroup>
					</SidebarContent>
					<SidebarFooter>
						<SidebarMenu>
							<SidebarMenuItem>
								<SidebarMenuButton onClick={handleLogout}>
									<LogOut />
									<span>Logout</span>
								</SidebarMenuButton>
							</SidebarMenuItem>
						</SidebarMenu>
					</SidebarFooter>
				</Sidebar>
				<main className="flex-1 overflow-auto">
					<Outlet />
				</main>
			</div>
		</SidebarProvider>
	);
}
