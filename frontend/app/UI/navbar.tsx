
function Navbar() {

    return (
        <>
            <nav className="bg-gray-800 p-4">
                <div className="container mx-auto flex items-center justify-between">
                    <div className="flex items-center">
                        <a href="#" className="text-white text-2xl font-bold">Project Tracker</a>
                    </div>
                    <div className="hidden md:flex space-x-4">
                        <a href="#" className="text-gray-300 hover:text-white">Home</a>
                        <a href="#" className="text-gray-300 hover:text-white">About</a>
                        <a href="#" className="text-gray-300 hover:text-white">Services</a>
                        <a href="#" className="text-gray-300 hover:text-white">Contact</a>
                    </div>
                    <div className="hidden md:flex items-center space-x-2">
                        <input type="text" placeholder="Search" className="px-2 py-1 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-400" />
                        <button className="bg-indigo-500 text-white px-3 py-1 rounded">Search</button>
                    </div>
                    <div className="md:hidden">
                        <button id="mobile-menu-button" className="text-gray-300 hover:text-white focus:outline-none">
                            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
                            </svg>
                        </button>
                    </div>
                </div>
                <div id="mobile-menu" className="md:hidden mt-2 space-y-2 px-2 pb-3 hidden">
                    <a href="#" className="block text-gray-300 hover:text-white">Home</a>
                    <a href="#" className="block text-gray-300 hover:text-white">About</a>
                    <a href="#" className="block text-gray-300 hover:text-white">Services</a>
                    <a href="#" className="block text-gray-300 hover:text-white">Contact</a>
                    <div className="flex items-center space-x-2">
                        <input type="text" placeholder="Search" className="px-2 py-1 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-400" />
                        <button className="bg-indigo-500 text-white px-3 py-1 rounded">Search</button>
                    </div>
                </div>
            </nav>
        </>

    )
}
export default Navbar;