import { createFileRoute } from '@tanstack/react-router';

const App = () => {
	return (
		<>
			<div className={'min-h-screen bg-purple-50 w-full p-4'}>
				<h1 className={'font-black text-4xl'}>Bastion</h1>
			</div>
		</>
	);
};

export const Route = createFileRoute('/')({ component: App });
