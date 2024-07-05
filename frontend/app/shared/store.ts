import { create } from 'zustand';

interface AuthState {
    successMessage: string | null;
    setSuccessMessage: (message: string | null) => void;
}
interface AuthStatus {
    isAuthenticated: boolean;
    setAuthenticated: (value: boolean) => void;
}


export const useAuthStore = create<AuthState>((set) => ({
    successMessage: null,
    setSuccessMessage: (message) => set({ successMessage: message })
}));

export const useAuthStatus = create<AuthStatus>((set) => ({
    isAuthenticated: false,
    setAuthenticated: (value) => set({ isAuthenticated: value }),
}));