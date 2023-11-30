import React from 'react';
import ReactModal from 'react-modal';
import { BlockBtn, BlockInput } from '../generic';
import { useShallow } from 'zustand/react/shallow';
import { useSignUpModal } from '../../stores/login-panel';
import { api } from '../../api';
import { alertError } from '../../util/error';
import { useAuth, useMe } from '../../stores';

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
    phoneNumber,
    setPhoneNumber,
  } = useSignUpModal(
    useShallow((state) => ({
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
      phoneNumber: state.phoneNumber,
      setPhoneNumber: state.setPhoneNumber,
    }))
  );
  const setToken = useAuth((state) => state.setToken);
  const setMe = useMe((state) => state.setMe);

  const handleSubmit = async () => {
    try {
      const me = await api.createUser({
        first_name: firstName,
        last_name: lastName,
        email,
        password,
        phone_number: phoneNumber,
      });
      setMe(me);
      const { token } = await api.getToken({ email, password });
      setToken(token);
      setIsOpen(false)
    } catch (err) {
      alertError(err);
    }
  };

  return (
    <ReactModal
      isOpen={isOpen}
      shouldFocusAfterRender={true}
      shouldCloseOnOverlayClick={true}
      shouldCloseOnEsc={true}
      ariaHideApp={false}
      onRequestClose={() => setIsOpen(false)}
      contentLabel="modal"
    >
      <BlockInput
        name="first_name"
        type="text"
        placeholder="first name"
        value={firstName}
        onChange={(e) => setFirstName(e.currentTarget.value)}
      />
      <BlockInput
        name="last_name"
        type="text"
        placeholder="last name"
        value={lastName}
        onChange={(e) => setLastName(e.currentTarget.value)}
      />
      <BlockInput
        name="phone"
        type="phone"
        placeholder="phone number"
        value={phoneNumber}
        onChange={(e) => setPhoneNumber(e.currentTarget.value)}
      />
      <BlockInput
        name="email"
        type="email"
        placeholder="email"
        value={email}
        onChange={(e) => setEmail(e.currentTarget.value)}
      />
      <BlockInput
        name="password"
        type="password"
        placeholder="password"
        value={password}
        onChange={(e) => setPassword(e.currentTarget.value)}
      />
      <BlockBtn type="submit" onClick={handleSubmit}>
        Отправить
      </BlockBtn>
      <BlockBtn
        style={{
          position: 'absolute',
          top: '1rem',
          right: '1rem',
        }}
        onClick={() => setIsOpen(false)}
      >
        Закрыть окно
      </BlockBtn>
    </ReactModal>
  );
};
