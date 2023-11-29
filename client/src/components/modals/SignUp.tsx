import React from 'react'
import ReactModal from 'react-modal'
import {BlockBtn, BlockInput} from '../generic'
import {useShallow} from 'zustand/react/shallow'
import {useSignUpModal} from '../../stores/login-panel'

export const SignUp: React.FC<{}> = () => {
  const {
    isOpen,
    setIsOpen,
    email,
    setEmail,
    firstName,
    setFirstName,
    lastName,
    setLastName,
    password,
    setPassword,
  } = useSignUpModal(useShallow(state => ({
    isOpen: state.isOpen,
    setIsOpen: state.setIsOpen,
    firstName: state.firstName,
    setFirstName: state.setFirstName,
    lastName: state.lastName,
    setLastName: state.setLastName,
    email: state.email,
    setEmail: state.setEmail,
    password: state.password,
    setPassword: state.setPassword,
  })))

  return <ReactModal
    isOpen={isOpen}
    shouldFocusAfterRender={true}
    shouldCloseOnOverlayClick={true}
    shouldCloseOnEsc={true}
    ariaHideApp={false}
    onRequestClose={() => setIsOpen(false)}
    contentLabel='modal'
  >
    <BlockInput 
      name="first_name"
      type="text"
      placeholder="first name"
      value={firstName}
      onChange={e => setFirstName(e.currentTarget.value)}/>
    <BlockInput 
      name="last_name"
      type="text"
      placeholder="last name"
      value={lastName}
      onChange={e => setLastName(e.currentTarget.value)}/>
    <BlockInput 
      name="email"
      type="email"
      placeholder="email"
      value={email}
      onChange={e => setEmail(e.currentTarget.value)}/>
    <BlockInput 
      name="password"
      type="password"
      placeholder="password"
      value={password}
      onChange={e => setPassword(e.currentTarget.value)}
    />
    <BlockBtn type="submit">Отправить</BlockBtn>
    <BlockBtn
      style={{
        position: 'absolute',
        top: '1rem',
        right: '1rem'
      }}
      onClick={() => setIsOpen(false)}
    >
      Закрыть окно
    </BlockBtn>
  </ReactModal>
}
