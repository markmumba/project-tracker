'use client';
import { useRouter } from 'next/navigation';
import useSWR from 'swr';
import React, { useState, useEffect } from 'react';
import { useSubmissionStore } from '@/app/shared/store';
import { CreateFeedbackFormData, SubmissionDetails } from '@/app/shared/types';
import fetcher, { axiosInstance } from '@/app/fetcher/fetcher';
import { formatDate } from '@/app/shared/helper';

function SubmissionDetail() {
    const router = useRouter();
    const selectedSubmissionId = useSubmissionStore((state) => state.selectedSubmissionId);
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (!selectedSubmissionId) {
            router.push('/dashboard');
        }
    }, [selectedSubmissionId, router]);

    const { data: submission, error, mutate } = useSWR<SubmissionDetails>(
        selectedSubmissionId ? `/submissions/${selectedSubmissionId}` : null,
        fetcher
    );

    console.log(submission);

    const [feedback, setFeedback] = useState<CreateFeedbackFormData>({
        submission_id: selectedSubmissionId,
        feedback_date: new Date().toISOString(),
        comment: '',
    });

    if (error) {
        return <div>Error loading submission details</div>;
    }

    if (!submission) {
        return <div>Loading...</div>;
    }

    async function handleFeedbackSubmit(e: React.FormEvent) {
        e.preventDefault();
        setIsLoading(true);
        try {
            const responseJson = JSON.stringify(feedback);
            console.log('Feedback:', responseJson);
            const response = await axiosInstance.post('/feedbacks', responseJson, {
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            // Fetch the current submission data
            const submissionResponse = await axiosInstance.get(`/submissions/${selectedSubmissionId}`, {
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            const currentSubmission = submissionResponse.data;

            // Update the 'reviewed' field on the current submission data
            const updatedSubmission = {
                ...currentSubmission,
                reviewed: true,
            };
            console.log(updatedSubmission.data)

            // Send the complete updated object in the PUT request
            await axiosInstance.put(`/submissions/${selectedSubmissionId}`, updatedSubmission, {
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            setFeedback({
                submission_id: selectedSubmissionId,
                feedback_date: new Date().toISOString(),
                comment: '',
            });

            // Refetch the submission details to reflect the updated reviewed status
            mutate();

            router.push('/dashboard');
        } catch (error) {
            console.error('Error submitting feedback:', error);
        } finally {
            setIsLoading(false);
        }
    }


    return (
        <div className="p-5 max-w-6xl">
            <h1 className="text-3xl font-bold mb-4">Submission Details</h1>
            <div className="bg-gray-100 p-4 rounded-lg mb-4">
                <p><strong>Description:</strong> {submission.description}</p>
                <p><strong>Submission Date:</strong> { formatDate(submission.submission_date)}</p>
                <p><strong>Project:</strong> {submission.project.title}</p>
                <p><strong>Student:</strong> {submission.student.name}</p>
                <a href={submission.document_path} className="text-blue-500 underline" target="_blank" rel="noopener noreferrer">
                    View Document
                </a>
            </div>
            <form onSubmit={handleFeedbackSubmit}>
                <div className="mb-4">
                    <label htmlFor="feedback" className="block text-lg font-medium text-gray-700">Feedback</label>
                    <textarea
                        id="feedback"
                        name="feedback"
                        value={feedback.comment}
                        onChange={(e) => setFeedback({ ...feedback, comment: e.target.value })}
                        rows={4}
                        className="mt-1 p-2 block w-full border border-gray-300 rounded-md shadow-sm focus:outline-none focus:border-blue-500 sm:text-sm"
                        required
                    />
                </div>
                <button
                    type="submit"
                    className="bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:bg-blue-700"
                    disabled={isLoading}
                >
                    {isLoading ? 'Submitting...' : 'Submit Feedback'}
                </button>
            </form>
        </div>
    );
}

export default SubmissionDetail;

