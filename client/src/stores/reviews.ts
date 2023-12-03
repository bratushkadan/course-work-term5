import { create } from 'zustand';
import { Review } from '../api/types';

export const useReviews = create<{
  reviews: Review[];
  setReviews: (r: Review[]) => void;
}>((set) => ({
  reviews: [],
  setReviews: (reviews) => set({ reviews }),
}));
