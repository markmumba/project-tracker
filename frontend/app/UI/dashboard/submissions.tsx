import { useSubmissionStore } from '@/app/shared/store';
import { LecturerSubmissionDetails } from '@/app/shared/types';
import { useRouter } from 'next/navigation';


function Submissions({ lecturerSubmissions }: { lecturerSubmissions?: LecturerSubmissionDetails[] | undefined | null }) {
    const router = useRouter();

    const setSelectedSubmissionId = useSubmissionStore((state) => state.setSelectedSubmissionId);

    const sortedSubmissions = lecturerSubmissions?.sort((a, b) => b.submission_id - a.submission_id);

    const truncateDescription = (description: string, maxLength: number) => {
        if (description.length > maxLength) {
            return description.slice(0, maxLength) + '...';
        }
        return description;
    };

    const handleSubmissionClick = (submissionId: number) => {
        setSelectedSubmissionId(submissionId);
        router.push('/submission/details');
    };

    return (
        <div className="relative">
            <div className="absolute left-2 top-0 bottom-0 border-l-2 border-gray-300"></div>
            {sortedSubmissions?.map((submission) => (
                <div
                    key={submission.submission_id}
                    className="relative pl-8 mb-4 cursor-pointer"
                    onClick={() => handleSubmissionClick(submission.submission_id)}
                >
                    <div className="flex items-center mb-2">
                        <div className="bg-blue-500 h-4 w-4 rounded-full border-2 border-white"></div>
                        <div className="ml-4 p-4 bg-sky-100 hover:bg-sky-200 rounded-lg flex-grow max-w-5xl">
                            <p className="mb-2">
                                <span className='text-lg font-bold'>Description:</span>
                                {` ${truncateDescription(submission.description, 40)}`}
                            </p>
                            <p className="mb-2">{`Submission Date: ${submission.submission_date}`}</p>
                            <h2>{`Project: ${submission.project_name}`}</h2>
                            <h3 className="text-medium text-gray-400">{`Student: ${submission.student_name}`}</h3>
                            <a
                                href={submission.document_path}
                                className="text-blue-500 underline"
                                target="_blank"
                                rel="noopener noreferrer"
                                onClick={(e) => e.stopPropagation()}
                            >
                                View Document
                            </a>
                        </div>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default Submissions;
