'use client';
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { axiosInstance } from "../fetcher/fetcher";
import { useAuthStore } from "../shared/store";
import { registerFormData } from "../shared/types";
import RegisterForm from "../UI/authentication/registerForm";
import Spinner from "../UI/spinner";

function Register() {
    const router = useRouter();
    const setSuccessMessage = useAuthStore((state) => state.setSuccessMessage);

    const [formData, setFormData] = useState<registerFormData>({
        name: '',
        email: '',
        password: '',
        role_id: 2
    });
    const [loading, setLoading] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);
        try {
            const requestBody = JSON.stringify(formData);
            const response = await axiosInstance.post('/register', requestBody, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            setSuccessMessage('Registration successful! Please log in.');
            router.push('/login');
        } catch (error) {
            console.log(error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            {loading && <Spinner />} {/* Conditionally render the spinner */}
            <RegisterForm formData={formData} handleChange={handleChange} handleSubmit={handleSubmit} />
        </>
    )
}
export default Register;
