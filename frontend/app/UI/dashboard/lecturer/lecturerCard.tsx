
import { UserDetails } from "@/app/shared/types";
import avatar from "/public/images/user.png"
import Image from "next/image";


function LecturerCard({ userDetails,studentCount }:
    {
        userDetails?: UserDetails | null | undefined,
        studentCount?: number | null | undefined,

    }) {
    return (
        <div className=" flex flex-col p-4 rounded-lg">
            <div className="w-20 h-30 mx-auto mb-4">
                <Image src={avatar.src} alt="avatar" className="w-full h-full rounded-full object-cover" />
            </div>
            <div className="  bg-gray-100 rounded-xl p-6">
                <div className="mb-1 p-4 rounded-xl">
                    <div className='text-lg font-bold '>Name</div>
                    <div className='text-base'>{userDetails?.name} </div>
                </div>

                <div className="mb-1 p-4 rounded-xl">
                    <div className='text-lg font-bold '>Students Under Supervision</div>
                    <div className='text-base'>{studentCount}</div>
                </div>
            </div>
        </div>
    );
};

export default LecturerCard;