import { create } from 'zustand';
import { Product } from '../api/types';

export const useCatalog = create<{
  products: Product[];
  setProducts: (products: Product[]) => void;
}>((set) => ({
  products: [],
  setProducts: (products) => set({ products }),
}));
