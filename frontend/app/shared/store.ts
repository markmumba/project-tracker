import { create } from 'zustand';

interface AuthState {
    successMessage: string | null;
    setSuccessMessage: (message: string | null) => void;
}
interface User {
    id: number;
    name: string;
    email: string;
    role: string;
}
interface userDetails {
    user: User | null;
    setUserDetails: (data: User | null) => void;
    clearUserDetails: () => void;

}

export const useAuthStore = create<AuthState>((set) => ({
    successMessage: null,
    setSuccessMessage: (message) => set({ successMessage: message })
}));

export const useUserDetails = create<userDetails>((set) => ({
    user: null,
    setUserDetails: (user) => set({ user }),
    clearUserDetails: () => set({ user: null })
}));