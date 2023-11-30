import React from 'react';
import ReactModal from 'react-modal';
import { useShallow } from 'zustand/react/shallow';

import { BlockBtn, BlockInput } from '../generic';
import { useAuth, useLoginModal } from '../../stores';
import { api } from '../../api';
import {alertError} from '../../util/error';

export const Login: React.FC = () => {
  const { isOpen, setIsOpen, email, setEmail, password, setPassword } = useLoginModal(
    useShallow((state) => ({
      isOpen: state.isOpen,
      setIsOpen: state.setIsOpen,
      email: state.email,
      setEmail: state.setEmail,
      password: state.password,
      setPassword: state.setPassword,
    }))
  );

  const setToken = useAuth(state => state.setToken)

  const handleSubmit = async () => {
    try {
      const data = await api.getToken({ email, password });
      setToken(data.token)
      setIsOpen(false)
    } catch (err) {
      alertError(err)
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
