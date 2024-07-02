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

function Dashboard() {

  const getUserDetails = async () => {
    try {
      const response = await axiosInstance.get("/users", {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json'
        }
      });
      const data = response.data;
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
          <UserCard
            avatar={user.avatar}
            userName={user.userName}
            projectName={user.projectName}
            supervisorName={user.supervisorName}
            submissions={user.submissions}
          />
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
