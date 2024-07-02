import Image from "next/image";
import registerImage from "/public/images/Sandy_Bus-05_Single-08.jpg"
import Link from "next/link";

function RegisterForm({ formData, handleChange, handleSubmit }:
    {
        formData: { name: string; email: string; password: string; };
        handleChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
        handleSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
    }) {
    return (
        <>
            <div className="flex flex-col md:flex-row h-screen ">
                <div className="md:w-1/2 bg-gray-200 flex items-center justify-center">
                    <div className="max-w-md w-full p-8">
                        <h2 className="text-4xl font-bold text-center">Student</h2>
                        <h2 className="text-3xl font-bold mb-6 text-center">Register</h2>
                        <form onSubmit={handleSubmit}>
                            <div className="mb-4">
                                <label className="block text-gray-700 mb-2" htmlFor="name">
                                    Name
                                </label>
                                <input
                                    type="text"
                                    id="name"
                                    name="name"
                                    value={formData.name}
                                    onChange={handleChange}
                                    className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-400"
                                    required
                                />
                            </div>
                            <div className="mb-4">
                                <label className="block text-gray-700 mb-2" htmlFor="email">
                                    Email
                                </label>
                                <input
                                    type="email"
                                    id="email"
                                    name="email"
                                    value={formData.email}
                                    onChange={handleChange}
                                    className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-400"
                                    required
                                />
                            </div>
                            <div className="mb-6">
                                <label className="block text-gray-700 mb-2" htmlFor="password">
                                    Password
                                </label>
                                <input
                                    type="password"
                                    id="password"
                                    name="password"
                                    value={formData.password}
                                    onChange={handleChange}
                                    className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-400"
                                    required
                                />
                            </div>
                            <button
                                type="submit"
                                className="w-full mb-4 bg-indigo-500 text-white py-3 px-4 rounded hover:bg-indigo-600 focus:outline-none focus:bg-indigo-600"
                            >
                                Register
                            </button>
                        </form>
                        <Link href="/login">
                        <p className="text-blue-500">Already have an account. Go to Login</p>
                        </Link>
                    </div>
                </div>
                <div className="md:w-1/2 bg-white flex items-center justify-center">
                    <div className="w-full h-full">
                        <Image
                            src={registerImage}
                            alt="Register Image"
                            className="h-full w-full object-cover"
                        />
                    </div>
                </div>
            </div>
        </>
    )
}
export default RegisterForm;