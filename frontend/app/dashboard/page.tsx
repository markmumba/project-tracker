'use client';
import useSWR from 'swr';
import fetcher from "../fetcher/fetcher";
import { DashboardSkeleton, LecturerDashboardSkeleton, UserCardSkeleton } from "../UI/skeletons";
import { LecturerSubmissionDetails, ProjectDetails, SubmissionDetails, UserDetails } from "../shared/types";
import NoProject from "../UI/dashboard/student/noProject";
import Project from "../UI/dashboard/student/project";
import { useUserStore } from '../shared/store';
import StudentCard from '../UI/dashboard/student/studentCard';
import LecuturerCard from '../UI/dashboard/lecturer/lecturerCard';
import Submissions from '../UI/dashboard/lecturer/submissions';
import { useEffect } from 'react';

// TODO : can view due date of project 
// TODO : be alerted if the submission are not enough 
// TODO : middleware to check auth so no page can load without auth


function Dashboard() {
  const { data: userDetails, isLoading: userLoading, error: userError } = useSWR<UserDetails>('/users', fetcher);

  const setRole = useUserStore((state) => state.setRole);

  useEffect(() => {
    if (userDetails) {
      setRole(userDetails.role);
    }
  }, [userDetails, setRole]);

  const shouldFetch = userDetails && userDetails.role !== 'lecturer';

  const { data: projectDetails, error: projectError } = useSWR<ProjectDetails>(shouldFetch ? '/projects' : null, fetcher);
  const { data: submissions, error: submissionError } = useSWR<SubmissionDetails[]>(shouldFetch ? '/submissions/student' : null, fetcher);
  const { data: students, error: studentError } = useSWR<UserDetails[]>('/users/students', fetcher);
  const { data: lecturerSubmissions, error: lecturerSubmissionError } = useSWR<LecturerSubmissionDetails[]>('/submissions/lecturer', fetcher);
  console.log("this is the dashboard",lecturerSubmissions);

  console.log(lecturerSubmissions);

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

  if (studentError) {
    console.error(studentError);
  }
  if (lecturerSubmissionError) {
    console.error(lecturerSubmissionError);
  }

  const submissionCount = submissions ? submissions.length : 0;
  const studentCount = students ? students.length : 0;

  if (!userDetails) {
    return <div>No user data available</div>;
  }

  if (userDetails.role === 'lecturer') {
    return (
      <div className="">
        <div className="flex flex-col md:flex-row justify-between">
          <div className="mb-4 md:mb-0 md:w-3/4 p-4 flex-grow">
            <h1 className="text-2xl font-bold mb-4"> Welcome back Lecturer {userDetails.name}</h1>
            <h1 className="text-2xl font-bold mb-4"> Latest Submissions</h1>
            <Submissions lecturerSubmissions={lecturerSubmissions} />
          </div>
          <div className="md:w-1/4 p-4">
            <LecuturerCard userDetails={userDetails} studentCount={studentCount} />
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
