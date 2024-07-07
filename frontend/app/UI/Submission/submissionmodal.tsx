// components/SubmissionModal.tsx

import { SubmissionDetails } from '@/app/shared/types';



function SubmissionModal({ submission, onClose }: { submission: SubmissionDetails, onClose: () => void }) {
    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
            <div className="absolute inset-0 bg-black opacity-50"></div>
            <div className="relative bg-sky-100 p-8 max-w-lg rounded-lg shadow-lg z-10">
                <h2 className="text-2xl font-bold mb-4">Submission Details</h2>
                <p className="mb-2"><span className='text-lg font-bold'>Description:</span> {submission.description}</p>
                <p className="mb-2">Submission Date: {submission.submission_date}</p>
                
                <div className="mt-4 flex justify-end">
                    <button
                        className="bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:bg-blue-700"
                        onClick={onClose}
                    >
                        Close
                    </button>
                </div>
            </div>
        </div>
    );
};

export default SubmissionModal;
