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
}

export const useUserStore = create<UserState>((set) => ({
    role: null,
    setRole: (role) => set({ role }),
}));
