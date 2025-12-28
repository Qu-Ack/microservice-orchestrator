import { createContext, useContext, useState, useEffect } from "react";
import { toast } from "sonner";

interface AuthContextInterface {
	isAuthenticated: boolean;
	setAuthState: (state: boolean) => void;
	checkAuth: () => Promise<void>;
}

const AuthContext = createContext<AuthContextInterface | null>(null);

export function AuthContextProvider({ children }: { children: React.ReactNode }) {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
	const [isLoading, setIsLoading] = useState(true);

	const checkAuth = async () => {
		try {
			const hasAuthCookie = document.cookie
				.split('; ')
				.some(cookie => cookie.startsWith('auth_id='));

			if (!hasAuthCookie) {
				setIsAuthenticated(false);
				setIsLoading(false);
				return;
			}

			const resp = await fetch(`${import.meta.env.VITE_APP_API_URL}/v1/user/me`, {
				credentials: "include",
			});

			if (resp.ok) {
				setIsAuthenticated(true);
			} else {
				toast.error("you need to log back in");
				setIsAuthenticated(false);
				// Optionally clear the invalid cookie
				document.cookie = 'auth_id=; Max-Age=0; path=/;';
			}
		} catch (err) {
			console.error("Auth check failed:", err);
			setIsAuthenticated(false);
		} finally {
			setIsLoading(false);
		}
	};
	useEffect(() => {
		checkAuth();
	}, []);

	const setAuthState = (state: boolean) => {
		setIsAuthenticated(state);
	};

	if (isLoading) {
		return <div>Loading...</div>;
	}

	return (
		<AuthContext.Provider
			value={{
				isAuthenticated,
				setAuthState,
				checkAuth,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}

export function useAuth() {
	const ctx = useContext(AuthContext);
	if (!ctx) {
		throw new Error("useAuth must be used within AuthContextProvider");
	}
	return ctx;
}
