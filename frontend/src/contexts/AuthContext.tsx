import { createContext, useContext, useState } from "react";

interface AuthContextInterface {
	isAuthenticated: boolean;
	setAuthState: (state: boolean) => void;
}

const AuthContext = createContext<AuthContextInterface | null>(null);

export function AuthContextProvider({ children }: { children: React.ReactNode }) {
	const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);

	const setAuthState = (state: boolean) => {
		setIsAuthenticated(state);
	};

	return (
		<AuthContext.Provider
			value={{
				isAuthenticated,
				setAuthState,
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

	const { isAuthenticated, setAuthState } = ctx;
	return { isAuthenticated, setAuthState };
}

