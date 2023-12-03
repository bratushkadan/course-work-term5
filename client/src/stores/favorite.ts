import { create } from 'zustand';
import {Favorite} from '../api/types';

export const useFavorite = create<{
  favorites: Favorite[],
  setFavorites: (favorites: Favorite[]) => void
}>(set => ({
  favorites: [],
  setFavorites: favorites => set({favorites})
}));
