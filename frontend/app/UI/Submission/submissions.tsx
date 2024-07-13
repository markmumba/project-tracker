// components/Submissions.tsx

import { useState } from 'react';
import { ProjectDetails, SubmissionDetails } from '@/app/shared/types';
import Link from 'next/link';
import SubmissionModal from './submissionmodal';
import { truncateDescription } from '@/app/shared/helper';

function Submissions({ submissions, project }: { submissions: SubmissionDetails[] | null | undefined, project: ProjectDetails | null | undefined }) {
    const [selectedSubmission, setSelectedSubmission] = useState<SubmissionDetails | null>(null);

    // Filter and sort submissions
    const unreviewed = submissions?.filter(submission => submission.reviewed === false) ?? [];
    const sortedSubmissions = unreviewed.sort((a, b) => b.submission_id - a.submission_id);

    const handleSubmissionClick = (submission: SubmissionDetails) => {
        setSelectedSubmission(submission);
    };

    const closeModal = () => {
        setSelectedSubmission(null);
    };

    return (
        <div className="p-5">
            <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold">Unreviewed Submissions</h1>
                <Link href="/dashboard/student/submission/createsubmission">
                    <button
                        className="bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:bg-blue-700"
                    >
                        Create Submission
                    </button>
                </Link>
            </div>
            <div className="relative">
                <div className="absolute left-2 top-0 bottom-0 border-l-2 border-gray-300"></div>
                {sortedSubmissions.length > 0 ? (
                    sortedSubmissions.map((submission) => (
                        <div
                            key={submission.submission_id}
                            className="relative pl-8 mb-4 cursor-pointer"
                            onClick={() => handleSubmissionClick(submission)}
                        >
                            <div className="flex items-center mb-2">
                                <div className="bg-blue-500 h-4 w-4 rounded-full border-2 border-white"></div>
                                <div className="ml-4 p-4 bg-blue-500/20 backdrop-blur-xl hover:bg-blue-500 hover:text-white group rounded-lg flex-grow max-w-5xl">
                                    <p className="mb-2">
                                        <span className='text-lg font-bold'>Description</span>
                                        {`: ${truncateDescription(submission.description, 40)}`}
                                    </p>
                                    <p className="mb-2">{`Submission Date: ${submission.submission_date}`}</p>
                                    <h2 className="">{`Project: ${project?.title}`}</h2>
                                    <Link href={submission.document_path} className="text-blue-500 group-hover:text-white underline">
                                        View Document
                                    </Link>
                                </div>
                            </div>
                        </div>
                    ))
                ) : (
                    <p className="text-center text-gray-500 mt-4">No unreviewed submissions found.</p>
                )}
            </div>
            {selectedSubmission && (
                <SubmissionModal submission={selectedSubmission} onClose={closeModal} />
            )}
        </div>
    );
}

export default Submissions;

