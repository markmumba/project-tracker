export interface ProjectDetails {
    project_id: number;
    student_id: number;
    lecturer_id: number;
    lecturer_name: string;
    title: string;
    description: string;
    start_date: string; // Assuming the dates are in ISO string format
    end_date: string;
}

export interface UserDetails {
    id: number;
    name: string;
    email: string;
    role: string;
}
export interface loginFormData {
    email: string;
    password: string;
}
export interface registerFormData {
    name: string;
    email: string;
    password: string;
    role_id: number;
}
export interface UserCardProps {
    userName: string | null | undefined;
    projectName: string | null | undefined;
    supervisorName: string | null | undefined;
    submissions: number | null | undefined;
}