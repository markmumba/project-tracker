'use client';
import { Suspense, useEffect, useState } from "react";
import UserCard from "../UI/dashboard/userCard";
import { axiosInstance } from "../fetcher/fetcher";
import { UserCardSkeleton } from "../UI/skeletons";
import { ProjectDetails, UserDetails } from "../shared/types";
import NoProject from "../UI/dashboard/noProject";
import Project from "../UI/dashboard/project";

// TODO : implement the get project details 



function Dashboard() {

  const [userDetails, setUserDetails] = useState<UserDetails | null>(null);
  const [projectDetails, setProjectDetails] = useState<ProjectDetails | null>(null);

  const getUserDetails = async () => {
    try {
      const response = await axiosInstance.get("/users", {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json'
        }
      });

      const data = response.data;
      setUserDetails(data);

    } catch (error: any) {
      console.log(error.response.data);
    }
  };

  const getProjectDetails = async () => {
    try {
      const response = await axiosInstance.get("/projects", {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json'
        }
      });

      const data = response.data;
      setProjectDetails(data);
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    getUserDetails();
    getProjectDetails();
  }, []);



  return (
    <div className="border p-4">
      <div className="flex flex-col md:flex-row justify-between">
        <div className="mb-4 md:mb-0 md:w-3/4 border p-4 flex-grow">
          {projectDetails && userDetails ?
            (<Project projectDetails={projectDetails} userDetails={userDetails} />)
            : (<NoProject userDetails={userDetails} />
            )}
        </div>
        <div className="md:w-1/4 border p-4">
          <Suspense fallback={<UserCardSkeleton />}>
            {userDetails && <UserCard
              userName={userDetails.name}
              projectName={projectDetails?.title}
              supervisorName={projectDetails?.lecturer_name}
              submissions={10}
            />
            }
          </Suspense>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
