import { ProjectDetails, UserDetails } from "@/app/shared/types";



function Project({ userDetails, projectDetails }: { userDetails?: UserDetails | null | undefined , projectDetails: ProjectDetails }) {
    return (
        <div className="">
            <h1 className="text-2xl font-semibold text-gray-800">Welcome, {userDetails?.name}</h1>
            <div className="border rounded-md bg-sky-100 p-4 mt-4">
                <h2 className="text-2xl font-semibold text-gray-800">Project: {projectDetails.title}</h2>
                <p className=" text-lg text-gray-600">Supervisor: {projectDetails.lecturer_name}</p>
                <p className="text-gray-600">Description: {projectDetails.description}</p>
            </div>
        </div>
    )

}
export default Project;