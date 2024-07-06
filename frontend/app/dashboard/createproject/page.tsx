'use client'
import CreateProjectForm from "@/app/UI/createprojectform";
import fetcher, { axiosInstance } from "@/app/fetcher/fetcher";
import { CreateProjectFormData, UserDetails } from "@/app/shared/types";
import { useState } from "react";
import useSWR from "swr";

function CreateProjectPage() {

    const { data: projectDetails, error: projectError } = useSWR<UserDetails>('/users/lecturers', fetcher);

    const [formData, setFormData] = useState<CreateProjectFormData>({
        title: '',
        description: '',
        startDate: '',
        endDate: '',
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axiosInstance.post('/projects', formData, {
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json',
                },

            }); // Assuming your API endpoint for creating projects is '/api/projects'
            console.log('Project created successfully:', response.data);
            setFormData({
                title: '',
                description: '',
                startDate: '',
                endDate: '',
            });
        } catch (error) {
            console.error('Error creating project:', error);
        }
    };
    return (
        <div>
            <CreateProjectForm formData={formData} handleSubmit={handleSubmit} handleChange={handleChange} />
        </div>
    );
}

export default CreateProjectPage;