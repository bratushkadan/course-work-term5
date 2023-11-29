import {create} from 'zustand';

export const useAuth = create<{
  token: string | null,
  setToken: (token: string | null) => void,
  clearToken: () => void,
}>(set => ({
  token: null,
  setToken: token => set({token}),
  clearToken: () => set({token: null}),
}))
