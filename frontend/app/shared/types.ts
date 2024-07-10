export interface ProjectDetails {
    project_id: string;
    student_id: string;
    lecturer_id: string;
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

export interface CreateProjectFormData {
    title: string;
    lecturer_id: number;
    description: string;
    start_date: string;
    end_date: string;
}

export interface SubmissionDetails {
    submission_id: number;
    project_id: number;
    student_id: number;
    submission_date: string;
    document_path: string;
    description: string;
    reviewed: boolean;
}
export interface CreateSubmissionFormData {
    project_id: string;
    student_id: string;
    submission_date: string;
    document_path: string;
    description: string;
    reviewed: boolean;
}

export interface LecturerSubmissionDetails {
    submission_id: number;
    submission_date: string;
    document_path: string;
    description: string;
    project_name: string;
    project_id: number;
    student_name: string;
    student_id: number;
    reviewed: boolean;
}
export interface CreateFeedbackFormData {
    submission_id: number | null | undefined;
    feedback_date: string;
    comments: string;
}

export interface FeedbackDetails {
    id: number;
    submission_id: number;
    feedback_date: string;
    comments: string;
    submission: SubmissionDetails;
}