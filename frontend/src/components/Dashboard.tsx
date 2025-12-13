import { Link } from 'react-router';

const mockServices = [
	{
		id: 'frontend-web',
		name: 'frontend-web',
		namespace: 'production',
		replicas: 3,
		status: 'healthy',
		cpu: 45,
		memory: 67,
		uptime: '14d 6h',
	},
	{
		id: 'api-gateway',
		name: 'api-gateway',
		namespace: 'production',
		replicas: 5,
		status: 'healthy',
		cpu: 72,
		memory: 81,
		uptime: '21d 3h',
	},
	{
		id: 'auth-service',
		name: 'auth-service',
		namespace: 'production',
		replicas: 2,
		status: 'warning',
		cpu: 89,
		memory: 92,
		uptime: '8d 12h',
	},
	{
		id: 'database-proxy',
		name: 'database-proxy',
		namespace: 'production',
		replicas: 4,
		status: 'healthy',
		cpu: 34,
		memory: 56,
		uptime: '45d 9h',
	},
	{
		id: 'cache-redis',
		name: 'cache-redis',
		namespace: 'infrastructure',
		replicas: 3,
		status: 'healthy',
		cpu: 23,
		memory: 78,
		uptime: '32d 14h',
	},
	{
		id: 'message-queue',
		name: 'message-queue',
		namespace: 'infrastructure',
		replicas: 2,
		status: 'critical',
		cpu: 95,
		memory: 98,
		uptime: '2d 4h',
	},
];

export default function Services() {
	return (
		<div className="min-h-screen bg-[#c0c0c0] p-4">
			{/* Header */}
			<table width="100%" cellPadding="0" cellSpacing="0" border={0} className="bg-[#000080] mb-4">
				<tbody>
					<tr>
						<td className="p-3">
							<div className="flex items-center justify-between">
								<Link to="/" className="text-white text-xl">⚓ KubePirate</Link>
								<div className="flex gap-4 text-white text-sm">
									<span><b>Services</b></span>
									<Link to="/" className="text-white underline">Logout</Link>
								</div>
							</div>
						</td>
					</tr>
				</tbody>
			</table>

			<table width="100%" cellPadding="10" cellSpacing="0" border={1} className="bg-white border-black mb-4">
				<tbody>
					<tr className="bg-[#d3d3d3]">
						<td>
							<div className="flex items-center justify-between">
								<div>
									<h1 className="text-2xl">⛵ Your Services</h1>
									<small className="text-gray-600">Total: {mockServices.length} services</small>
								</div>
								<button className="bg-[#008080] text-white border-2 border-black px-4 py-1">
									+ Deploy New Service
								</button>
							</div>
						</td>
					</tr>
				</tbody>
			</table>

			{/* Search Box */}
			<table width="100%" cellPadding="8" cellSpacing="0" border={1} className="bg-white border-black mb-4">
				<tbody>
					<tr>
						<td>
							<b>Search:</b>{' '}
							<input
								type="text"
								placeholder="Filter services..."
								className="border-2 border-black px-2 py-1 ml-2"
								size={40}
							/>
						</td>
					</tr>
				</tbody>
			</table>

			<table width="100%" cellPadding="8" cellSpacing="0" border={1} className="bg-white border-black">
				<tbody>
					<tr className="bg-[#000080]">
						<th className="text-white text-left">Service Name</th>
						<th className="text-white text-left">Namespace</th>
						<th className="text-white text-center">Status</th>
						<th className="text-white text-center">Replicas</th>
						<th className="text-white text-center">CPU %</th>
						<th className="text-white text-center">Memory %</th>
						<th className="text-white text-center">Uptime</th>
						<th className="text-white text-center">Actions</th>
					</tr>
					{mockServices.map((service, index) => (
						<tr
							key={service.id}
							style={{
								backgroundColor: index % 2 === 0 ? '#ffffff' : '#f0f0f0'
							}}
						>							<td>
								<Link to={`/services/${service.id}`} className="text-blue-600 underline">
									<b>{service.name}</b>
								</Link>
							</td>
							<td>{service.namespace}</td>
							<td align="center">
								{service.status === 'healthy' && (
									<span className="text-green-700">✓ Healthy</span>
								)}
								{service.status === 'warning' && (
									<span className="text-orange-600">⚠ Warning</span>
								)}
								{service.status === 'critical' && (
									<span className="text-red-600">✗ Critical</span>
								)}
							</td>
							<td align="center">{service.replicas}</td>
							<td align="center">{service.cpu}%</td>
							<td align="center">{service.memory}%</td>
							<td align="center">{service.uptime}</td>
							<td align="center">
								<Link to={`/services/${service.id}`} className="text-blue-600 underline text-sm">
									View Details
								</Link>
							</td>
						</tr>
					))}
				</tbody>
			</table>

			{/* Legend */}
			<table width="100%" cellPadding="8" cellSpacing="0" border={1} className="bg-[#ffff99] border-black mt-4">
				<tbody>
					<tr>
						<td>
							<small>
								<b>Legend:</b>{' '}
								<span className="text-green-700">✓ Healthy</span> |
								<span className="text-orange-600"> ⚠ Warning</span> |
								<span className="text-red-600"> ✗ Critical</span>
							</small>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	);
}

