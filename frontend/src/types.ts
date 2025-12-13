export type ServiceList = {
	Services: Partial<Service>[];
}

export type Service = {
	Name: string;	
	Replicas: number;
	Age: string;
	StartedAt: string;
	Labels: string;
	Pods: Pod[];
	Stats: Stats;
}

export type Pod = {
	Name: string;
	Age: string;
	Port: number;
	Labels: string;
	StartedAt: string;
	Ready: string;
	ContainersReady: string;
	Node: string;
	Stats: Stats;
}

export type Stats = {
	CPU: number;
	Memory: number; 
}


