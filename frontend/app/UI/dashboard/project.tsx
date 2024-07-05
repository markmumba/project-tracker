import { ProjectDetails, UserDetails } from "@/app/shared/types";

//TODO : adding the lastest feedback 


function Project({ userDetails, projectDetails }: { userDetails: UserDetails, projectDetails: ProjectDetails }) {
    return (
        <>
            <h1 className="text-2xl font-semibold text-gray-800">Welcome, {userDetails.name}</h1>
            <div className="border p-4 mt-4">
                <h2 className="text-xl font-semibold text-gray-800">Project: {projectDetails.title}</h2>
                <p className="text-gray-600">Supervisor: {projectDetails.lecturer_name}</p>
                <p className="text-gray-600">Description: {projectDetails.description}</p>
            </div>
        </>
    )

}
export default Project;