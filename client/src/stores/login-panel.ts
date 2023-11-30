import { create } from 'zustand';

export const useLoginModal = create<{
  isOpen: boolean
  setIsOpen: (isOpen: boolean) => void
  email: string
  setEmail: (email: string) => void
  password: string
  setPassword: (password: string) => void
}>(set => ({
  isOpen: false,
  setIsOpen: (isOpen) => set(() => ({isOpen})),
  email: '',
  setEmail: (email) => set(() => ({email})),
  password: '',
  setPassword: (password) => set(() => ({password})),
}))
export const useSignUpModal = create<{
  isOpen: boolean
  setIsOpen: (isOpen: boolean) => void
  firstName: string
  setFirstName: (firstName: string) => void
  lastName: string
  setLastName: (lastName: string) => void
  email: string
  setEmail: (email: string) => void
  password: string
  setPassword: (password: string) => void
  phoneNumber: string
  setPhoneNumber: (phoneNumber: string) => void
}>(set => ({
  isOpen: false,
  setIsOpen: (isOpen) => set(() => ({isOpen})),
  firstName: '',
  setFirstName: (firstName) => set(() => ({firstName})),
  lastName: '',
  setLastName: (lastName) => set(() => ({lastName})),
  email: '',
  setEmail: (email) => set(() => ({email})),
  password: '',
  setPassword: (password) => set(() => ({password})),
  phoneNumber: '',
  setPhoneNumber: (phoneNumber) => set(() => ({phoneNumber})),
}))


