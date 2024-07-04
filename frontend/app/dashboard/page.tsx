'use client';
import { useEffect, useState } from "react";
import SideNav from "../UI/dashboard/sidebar";
import UserCard from "../UI/dashboard/userCard";
import { axiosInstance } from "../fetcher/fetcher";
import avatar from "/public/images/user.png"

interface user {
  avatar: string; // Change the type of avatar to string
  userName: string;
  projectName: string;
  supervisorName: string;
  submissions: number;
}
interface userDetails {
  id: number;
  name: string;
  email: string;
  role: string;

}

function Dashboard() {

  const [userDetalis, setUserDetails] = useState<userDetails>();
  const [projectDetails, setProjectDetails] = useState(); 

  const getUserDetails = async () => {
    try {
      const response = await axiosInstance.get("/users", {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json'
        }
      });

      console.log(response);
      const data = response.data;
      setUserDetails(data);
      console.log(data);

    } catch (error: any) {
      if (error.response) {
        // Server responded with a status other than 2xx
        console.error("Error response data:", error.response.data);
        console.error("Error response status:", error.response.status);
        console.error("Error response headers:", error.response.headers);
      } else if (error.request) {
        // Request was made but no response was received
        console.error("Error request data:", error.request);
      } else {
        // Something else happened in making the request that triggered an error
        console.error("Error message:", error.message);
      }
      console.error("Axios config:", error.config);
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

      console.log(response);
      const data = response.data;
      setProjectDetails(data);
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  }

  const user: user = {
    avatar: avatar.src, // Assign the src property of the avatar to the avatar property of the user object
    userName: 'John Doe',
    projectName: 'Project Alpha',
    supervisorName: 'Dr. Smith',
    submissions: 5,
  };

  useEffect(() => {
    getUserDetails();
  }, []);
  return (
    <div className="border p-4">
      <div className="flex flex-col md:flex-row justify-between">
        <div className="mb-4 md:mb-0 md:w-3/4 border p-4 flex-grow">
          <p>
            1. should have latest feedback
            <br />
            2. should prompt user to create project if there is no project
          </p>
        </div>
        <div className="md:w-1/4 border  p-4">
          {userDetalis &&
            <UserCard
              avatar={user.avatar}
              userName={userDetalis.name}
              projectName={user.projectName}
              supervisorName={user.supervisorName}
              submissions={user.submissions}
            />
          }
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
