import { Link } from "react-router-dom"

function NavBar() {
	return (
		<header>
			<nav>
				<ul className="flex items-center justify-end px-8 space-x-2 bg-blue-500">
					<li className="px-4 py-2">
						<Link className="hover:border-b-4 hover:border-blue-200" to={`/`}>Home</Link>
					</li>
					<li className="px-4 py-2">
						<Link className="hover:border-b-4 hover:border-blue-200" to={`about`}>About</Link>
					</li>
					<li className="px-4 py-2">
						<Link className="hover:border-b-4 hover:border-blue-200" to={`users-table`}>Users</Link>
					</li>
				</ul>
			</nav>
		</header>
	)
}

export default NavBar
