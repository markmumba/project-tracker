'use client'
import { useState } from "react";
import RegisterForm from "../UI/registerForm";
import { axiosInstance } from "../fetcher/fetcher";
import { useRouter } from "next/router";
import { useAuthStore } from "../shared/store";


interface registerFormData {
    name: string;
    email: string;
    password: string;
    role_id: number;
}

function Register() {
    const router = useRouter();
    const setSuccessMessage = useAuthStore(state => state.setSuccessMessage);

    const [formData, setFormData] = useState<registerFormData>({
        name: '',
        email: '',
        password: '',
        role_id: 2
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log(formData);
        try {
            const requestBody = JSON.stringify(formData);
            const response = await axiosInstance.post('/register', requestBody, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            setSuccessMessage('Registration successful, please login');

            router.push('/login');


        } catch (error) {
            console.log(error);
        }
    };
    return (
        <>
            <RegisterForm formData={formData} handleChange={handleChange} handleSubmit={handleSubmit} />
        </>
    )
}
export default Register;