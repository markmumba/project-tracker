'use client'
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { axiosInstance } from "../fetcher/fetcher";
import LoginForm from "../UI/authentication/loginForm";
import { useAuthStore } from "../shared/store";
import { loginFormData } from "../shared/types";

// TODO : add the token to the window  local storage 



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
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            window.localStorage.setItem('token', response.data.token);
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