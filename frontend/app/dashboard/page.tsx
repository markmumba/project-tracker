'use client';
import useSWR from 'swr';
import UserCard from "../UI/dashboard/studentCard";
import fetcher from "../fetcher/fetcher";
import { DashboardSkeleton, LecturerDashboardSkeleton } from "../UI/skeletons";
import { ProjectDetails, SubmissionDetails, UserDetails } from "../shared/types";
import NoProject from "../UI/dashboard/noProject";
import Project from "../UI/dashboard/project";
import { useUserStore } from '../shared/store';
import StudentCard from '../UI/dashboard/studentCard';
import LecuturerCard from '../UI/dashboard/lecturerCard';

// TODO : can view due date of project 
// TODO : be alerted if the submission are not enough 


function Dashboard() {
  const { data: userDetails, isLoading: userLoading, error: userError } = useSWR<UserDetails>('/users', fetcher, {
    revalidateOnFocus: false,
    revalidateOnReconnect: false
  });
  const setRole = useUserStore((state) => state.setRole);

  // Set role in Zustand store when userDetails are fetched
  if (userDetails) {
    setRole(userDetails.role);
  }

  const shouldFetch = userDetails && userDetails.role !== 'lecturer';

  const { data: projectDetails, error: projectError } = useSWR<ProjectDetails>(
    shouldFetch ? '/projects' : null,
    fetcher,
    {
      revalidateOnFocus: false,
      revalidateOnReconnect: false
    }
  );

  const { data: submissions, error: submissionError } = useSWR<SubmissionDetails[]>(
    shouldFetch ? '/submissions/all' : null,
    fetcher,
    {
      revalidateOnFocus: false,
      revalidateOnReconnect: false
    }
  );

  if (userLoading) {
    return <DashboardSkeleton />;
  }

  if (userError) {
    console.error(userError.response?.data);
    return <div>Error loading user data</div>;
  }

  if (projectError) {
    console.error(projectError);
  }

  if (submissionError) {
    console.error(submissionError);
  }

  const submissionCount = submissions ? submissions.length : 0;

  if (!userDetails) {
    return <div>No user data available</div>;
  }

  if (userDetails.role === 'lecturer') {
    return (
      <div className="border p-4">
        <div className="flex flex-col md:flex-row justify-between">
          <div className="mb-4 md:mb-0 md:w-3/4 border p-4 flex-grow">
        <h1>Feedbacks are here </h1>
          </div>
          <div className="md:w-1/4 border p-4">
            <LecuturerCard userDetails={userDetails} />
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="border p-4">
      <div className="flex flex-col md:flex-row justify-between">
        <div className="mb-4 md:mb-0 md:w-3/4 border p-4 flex-grow">
          {projectDetails ? (
            <Project projectDetails={projectDetails} userDetails={userDetails} />
          ) : (
            <NoProject userDetails={userDetails} />
          )}
        </div>
        <div className="md:w-1/4 border p-4">
          <StudentCard userDetails={userDetails} projectDetails={projectDetails} submissionCount={submissionCount} />
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
