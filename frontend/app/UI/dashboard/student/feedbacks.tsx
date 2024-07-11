import { useState } from 'react';
import { formatFeedbackDate, truncateDescription } from '@/app/shared/helper';
import { FeedbackDetails } from '@/app/shared/types';
import Link from 'next/link';
import FeedbackModal from './feedbackmodal';

function Feedbacks({ feedbackDetails }: { feedbackDetails: FeedbackDetails[] | undefined | null }) {
    const [showRecent, setShowRecent] = useState(true);

    const [selectedFeedback, setSelectedFeedback] = useState<FeedbackDetails | null>(null);
    const handleFeedbackClick = (feedback: FeedbackDetails) => {
        setSelectedFeedback(feedback);
    };

    const closeModal = () => {
        setSelectedFeedback(null);
    };

    // Sort feedback by date in descending order
    const sortedFeedback = feedbackDetails?.slice().sort((a, b) => new Date(b.feedback_date).getTime() - new Date(a.feedback_date).getTime());

    // Filter feedback based on the toggle state

    return (
        <div className="pt-10">
            <h1 className="text-3xl font-bold">Feedback</h1>
            <div className="flex justify-end mb-4">
                <button
                    className="bg-blue-300  py-2 px-4 rounded-full hover:bg-blue-400 focus:outline-none focus:bg-blue-400"
                    onClick={() => setShowRecent(!showRecent)}
                >View all feedback</button>
            </div>

            {sortedFeedback?.map((feedback) => (
                <div
                    key={feedback.feedback_id}
                    className="relative pl-8 mb-4 cursor-pointer"
                    onClick={() => handleFeedbackClick(feedback)}
                >
                    <div className="flex items-center mb-2">
                        <div className=" bg-blue-500 h-4 w-4 rounded-full border-2 border-white"></div>
                        <div className="ml-4 p-4 bg-gray-100 hover:bg-blue-500 hover:text-white group rounded-lg flex-grow max-w-5xl">
                            <p className="mb-2">
                                <span className='text-lg font-bold'>Feedback</span>
                                {`: ${truncateDescription(feedback.comments, 40)}`}
                            </p>
                            <p className="mb-2">{`Feedback Date: ${formatFeedbackDate(feedback.feedback_date)}`}</p>
                            <h2 className=""><span className="text-lg font-bold"> In reference to submission:</span>{truncateDescription(feedback.description, 40)}</h2>
                            <Link href={feedback.document_path} className="text-blue-500 group-hover:text-white underline">
                                View Document
                            </Link>
                        </div>
                    </div>
                </div>
            ))}
            {selectedFeedback && (
                <FeedbackModal feedback={selectedFeedback} onClose={closeModal} />
            )}
        </div>
    );
}

export default Feedbacks;
