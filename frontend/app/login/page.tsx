import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { axiosInstance } from "../fetcher/fetcher";
import LoginForm from "../UI/loginForm";
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
        console.log(formData);
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
            <LoginForm formData={formData} handleChange={handleChange} handleSubmit={handleSubmit} successMessage={successMessage} />
        </>
    )
}

export default Login;