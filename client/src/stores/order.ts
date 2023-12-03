import { create } from 'zustand';
import { Order, OrderInfo } from '../api/types';

export const useOrder = create<{
  orderDetails: OrderInfo | null;
  setOrderDetails: (orderDetails: OrderInfo) => void;
  orders: Order[];
  setOrders: (orders: Order[]) => void;
}>((set) => ({
  orderDetails: null,
  setOrderDetails: (orderDetails) => set({ orderDetails }),
  orders: [],
  setOrders: (orders) => set({ orders }),
}));
