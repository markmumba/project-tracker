
import {create} from 'zustand';

interface AuthState {
    successMessage :string | null ;
    setSuccessMessage: (message: string | null ) => void;
}

export const useAuthStore = create<AuthState>((set)=> ({
    successMessage: null,
    setSuccessMessage:(message) => set({successMessage: message})
}));

