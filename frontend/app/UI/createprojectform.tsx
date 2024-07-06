
import { CreateProjectFormData } from '../shared/types';



function CreateProjectForm({ formData, handleSubmit, handleChange }: { formData: CreateProjectFormData, handleSubmit: (e: React.FormEvent<HTMLFormElement>) => void, handleChange: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void }) {


    return (
        <div className="max-w-md p-10 overflow-hidden md:max-w-5xl">
            <h2 className="text-3xl font-bold mb-4 text-center">Create Project</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
                <label className="block">
                    <span className="text-gray-700">Title:</span>
                    <input
                        type="text"
                        name="title"
                        value={formData.title}
                        onChange={handleChange}
                        required
                        className="mt-1 block w-full px-3 py-4  bg-gray-100 rounded-lg focus:outline-none  sm:text-sm"
                    />
                </label>

                <label className="block">
                    <span className="text-gray-700">Description:</span>
                    <textarea
                        name="description"
                        value={formData.description}
                        onChange={handleChange}
                        required
                        className="mt-1 block w-full px-3 py-6 bg-gray-100 rounded-lg focus:outline-none  sm:text-sm"
                    />
                </label>

                <label className="block">
                    <span className="text-gray-700">Start Date:</span>
                    <input
                        type="date"
                        name="startDate"
                        value={formData.startDate}
                        onChange={handleChange}
                        required
                        className="mt-1 block w-full px-3 py-4  bg-gray-100 rounded-lg focus:outline-none  sm:text-sm"
                    />
                </label>

                <label className="block">
                    <span className="text-gray-700">End Date:</span>
                    <input
                        type="date"
                        name="endDate"
                        value={formData.endDate}
                        onChange={handleChange}
                        required
                        className="mt-1 block w-full px-3 py-4  bg-gray-100 rounded-lg focus:outline-none  sm:text-sm"
                    />
                </label>

                <button
                    type="submit"
                    className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:bg-blue-700"
                >
                    Create Project
                </button>
            </form>
        </div>

    );
};

export default CreateProjectForm;
