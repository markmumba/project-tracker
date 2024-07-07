import SubmissionForm from "@/app/UI/Submission/createsubmission";
import fetcher, { axiosInstance } from "@/app/fetcher/fetcher";
import { CreateSubmissionFormData, ProjectDetails } from "@/app/shared/types";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import useSWR from "swr";

function Submission() {
    const router = useRouter();

    const { data: project, error: projectError } = useSWR<ProjectDetails>("/projects", fetcher);

    if (projectError) {
        console.log(projectError.response?.data);
    }
    console.log(project);

    const [formData, setFormData] = useState<CreateSubmissionFormData>({
        project_id: 0,
        student_id: 0,
        submission_date: '',
        document_path: '',
        description: '',
    });

    useEffect(() => {
        const currentDate = new Date().toISOString().split('T')[0];
        if (project) {
            setFormData({
                project_id: project.id,
                student_id: project.student_id, // Assuming project has a student_id field
                submission_date: currentDate,
                document_path: '',
                description: '',
            });
        } else {
            setFormData(prev => ({
                ...prev,
                submission_date: currentDate,
            }));
        }
    }, [project]);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const requestBody = JSON.stringify(formData);
            const response = await axiosInstance.post('/submissions', requestBody, {
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            console.log('Project created successfully:', response.data);
            setFormData({
                project_id: 0,
                student_id: 0,
                submission_date: '',
                document_path: '',
                description: '',
            });
            router.push('/dashboard');
        } catch (error) {
            console.error('Error creating project:', error);
        }
    };

    return (
        <>
            <SubmissionForm formData={formData} handleSubmit={handleSubmit} handleChange={handleChange} />
        </>
    );
}

export default Submission;
