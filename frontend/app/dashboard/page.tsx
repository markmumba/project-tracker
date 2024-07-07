'use client';
import { Suspense } from "react";
import useSWR from 'swr';
import UserCard from "../UI/dashboard/userCard";
import fetcher from "../fetcher/fetcher";
import { DashboardSkeleton } from "../UI/skeletons";
import { ProjectDetails, SubmissionDetails, UserDetails } from "../shared/types";
import NoProject from "../UI/dashboard/noProject";
import Project from "../UI/dashboard/project";


function Dashboard() {

  const { data: userDetails, isLoading: userLoading, error: userError } = useSWR<UserDetails>('/users', fetcher);
  const { data: projectDetails, error: projectError } = useSWR<ProjectDetails>('/projects', fetcher);
  const { data: submissions, error: submissionError } = useSWR('/submissions/all', fetcher);

  if (userLoading) {
    return <DashboardSkeleton />;
  }
  if (userError) {
    console.log(userError.response?.data);
  }

  if (projectError) {
    console.error(projectError);
  }
  if (submissionError) {
    console.error(submissionError);
  }
  //count submissions 
  function countSubmissions(submissions: SubmissionDetails[] | undefined | null) {
    return submissions?.length;
  }
  var submissionCount = countSubmissions(submissions);

  return (
    <div className="border p-4">
      <div className="flex flex-col md:flex-row justify-between">
        <div className="mb-4 md:mb-0 md:w-3/4 border p-4 flex-grow">
          {projectDetails ?
            (
              <Project projectDetails={projectDetails} userDetails={userDetails} />

            )
            : (
              <NoProject userDetails={userDetails} />
            )}
        </div>
        <div className="md:w-1/4 border p-4">
          <UserCard userDetails={userDetails} projectDetails={projectDetails} submissionCount={submissionCount} />
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
