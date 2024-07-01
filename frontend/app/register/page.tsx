'use client'
import { useState } from "react";
import RegisterForm from "../UI/registerForm";


interface FormData {
    name: string;
    email: string;
    password: string;
}

function Register() {

    const [formData, setFormData] = useState<FormData>({
        name: '',
        email: '',
        password: ''
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log(formData);
        // Add your form submission logic here
    };
    return (
        <>
        <RegisterForm formData={formData} handleChange={handleChange} handleSubmit={handleSubmit} />
        </>
    )
}
export default Register;