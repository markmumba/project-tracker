import { create } from 'zustand';

interface AuthState {
    successMessage: string | null;
    setSuccessMessage: (message: string | null) => void;
}
export const useAuthStore = create<AuthState>((set) => ({
    successMessage: null,
    setSuccessMessage: (message) => set({ successMessage: message })
}));

interface UserState {
    role: string | null;
    setRole: (role: string) => void;
    resetUser: () => void;
}

export const useUserStore = create<UserState>((set) => ({
    role: null,
    setRole: (role) => set({ role }),
    resetUser: () => set({ role: null }),
}));



interface SubmissionState {
    selectedSubmissionId: number | null;
    setSelectedSubmissionId: (id: number) => void;
}

export const useSubmissionStore = create<SubmissionState>((set) => ({
    selectedSubmissionId: null,
    setSelectedSubmissionId: (id: number) => set({ selectedSubmissionId: id }),
}));
