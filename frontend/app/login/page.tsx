'use client'
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { axiosInstance } from "../fetcher/fetcher";
import LoginForm from "../UI/authentication/loginForm";
import { useAuthStore } from "../shared/store";

interface loginFormData {
    email: string;
    password: string;
}
function Login() {
    const router = useRouter();
    const successMessage = useAuthStore(state => state.successMessage);
    const setSuccessMessage = useAuthStore(state => state.setSuccessMessage);

    const [formData, setFormData] = useState<loginFormData>({
        email: '',
        password: '',

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
        try {
            const requestBody = JSON.stringify(formData);
            const response = await axiosInstance.post('/login', requestBody, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            router.push('/dashboard');


        } catch (error) {
            console.log(error);
        }
    };
    useEffect(() => {
        if (successMessage) {
            setTimeout(() => {
                setSuccessMessage(null);
            }, 3000);
        }
    });

    return (

        <>
            {successMessage && <div className="bg-green-500 text-white text-center p-3">{successMessage}</div>}
            <LoginForm formData={formData} handleChange={handleChange} handleSubmit={handleSubmit} />
        </>
    )
}

export default Login;