
import { UserCardProps } from "@/app/shared/types";
import avatar from "/public/images/user.png"


function UserCard({ userName, projectName, supervisorName, submissions }: UserCardProps) {
  return (
    <div className=" flex flex-col p-4 rounded-lg">
      <div className="w-40 h-40 mx-auto mb-4">
        <img src={avatar.src} alt="avatar" className="w-full h-full rounded-full object-cover" />
      </div>
      <div className="  bg-gray-300 rounded-xl p-6">
        <div className="mb-1 p-4 rounded-xl">
          <div className='text-lg font-bold '>Name</div>
          <div className='text-base'>{userName}</div>
        </div>

        <div className="mb-1 p-4 rounded-xl">
          <div className='text-lg font-bold '>Project Name</div>
          <div className='text-base'>{projectName}</div>
        </div>

        <div className="mb-1 p-4 rounded-xl">
          <div className='text-lg font-bold '>Supervisor Name</div>
          <div className='text-base'>{supervisorName}</div>
        </div>

        <div className='mb-1 p-4 rounded-xl'>
          <div className='text-lg font-bold'>Submissions</div>
          <div className='text-base'>{submissions}</div>
        </div>
      </div>
    </div>
  );
};

export default UserCard;
