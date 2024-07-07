'use client';
import { SubmissionDetails } from '@/app/shared/types';
import Link from 'next/link';
import { useRouter } from 'next/navigation';


function Submissions({ submissions }: { submissions: SubmissionDetails[] | null | undefined }) {
    const router = useRouter();


    return (
        <div className="p-5">
            <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold">Submissions</h1>
                <Link href="/dashboard/submission/createsubmission">
                    <button
                        className="bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:bg-blue-700">
                        Create Submission
                    </button>
                </Link>
            </div>
            <div className="relative">
                <div className="absolute left-2 top-0 bottom-0 border-l-2 border-gray-300"></div>
                {submissions?.map((submission) => (
                    <div key={submission.submission_id} className="relative pl-8 mb-4">
                        <div className="flex items-center mb-2">
                            <div className="bg-blue-500 h-4 w-4 rounded-full border-2 border-white"></div>
                            <div className="ml-4 p-4 bg-white shadow-md rounded-lg">
                                <h2 className="text-xl font-semibold">{`Project ID: ${submission.project_id}`}</h2>
                                <p>{`Student ID: ${submission.student_id}`}</p>
                                <p>{`Submission Date: ${submission.submission_date}`}</p>
                                <p>{`Description: ${submission.description}`}</p>
                                <a
                                    href={submission.document_path}
                                    className="text-blue-500 underline"
                                    target="_blank"
                                    rel="noopener noreferrer"
                                >
                                    View Document
                                </a>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Submissions;
