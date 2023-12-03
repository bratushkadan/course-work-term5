import { create } from 'zustand';
import { CartPosition } from '../api/types';

export const useCart = create<{
  cart: CartPosition[];
  setCart: (cart: CartPosition[]) => void;
}>((set) => ({
  cart: [],
  setCart: (cart) => set({ cart }),
}));
