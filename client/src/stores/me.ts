import { create } from 'zustand';
import { User } from '../api/types';

export const useMe = create<{
  me: User | null;
  setMe: (user: User) => void;
}>((set) => ({
  me: null,
  setMe: (user) => set({ me: user }),
}));
