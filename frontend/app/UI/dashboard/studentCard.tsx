
import { ProjectDetails, UserDetails } from "@/app/shared/types";
import avatar from "/public/images/user.png"


function StudentCard({ userDetails, projectDetails, submissionCount }:
  {
    userDetails?: UserDetails | null | undefined,
    projectDetails?: ProjectDetails | null | undefined,
    submissionCount?: number | undefined
  }) {
  return (
    <div className=" flex flex-col p-4 rounded-lg">
      <div className="w-30 h-20 mx-auto mb-4">
        <img src={avatar.src} alt="avatar" className="w-full h-full rounded-full object-cover" />
      </div>
      <div className="  bg-gray-100 rounded-xl p-6">
        <div className="mb-1 p-4 rounded-xl">
          <div className='text-lg font-bold '>Name</div>
          <div className='text-base'>{userDetails?.name} </div>
        </div>

        <div className="mb-1 p-4 rounded-xl">
          <div className='text-lg font-bold '>Project Title</div>
          <div className='text-base'>{projectDetails?.title}</div>
        </div>

        <div className="mb-1 p-4 rounded-xl">
          <div className='text-lg font-bold '>Supervisor Name</div>
          <div className='text-base'>{projectDetails?.lecturer_name}</div>
        </div>

        <div className='mb-1 p-4 rounded-xl'>
          <div className='text-lg font-bold'>Submissions</div>
          <div className='text-base'>{submissionCount}</div>
        </div>
      </div>
    </div>
  );
};

export default StudentCard;